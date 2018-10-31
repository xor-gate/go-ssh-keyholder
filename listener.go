package main

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/sys/unix"
)

var ErrListenerSocketPathIsNotAUnixSocket = fmt.Errorf("listener socket path is not unix socket")

type Listener struct {
	cfg *Config
	fd  net.Listener
	kr  agent.Agent
}

func connPeerCredentials(conn *net.UnixConn) (uint32, uint32, error) {
	f, err := conn.File()
	if err != nil {
		return 0, 0, err
	}
	creds, err := unix.GetsockoptUcred(int(f.Fd()), unix.SOL_SOCKET, unix.SO_PEERCRED)
	f.Close()

	if err != nil {
		return 0, 0, err
	}
	return creds.Uid, creds.Gid, nil
}

func NewListener(cfg *Config) (*Listener, error) {
	fi, err := os.Stat(cfg.SocketPath)
	if !os.IsNotExist(err) {
		switch mode := fi.Mode(); {
		case mode&os.ModeSocket != 0:
			err = os.Remove(cfg.SocketPath)
			if err != nil {
				return nil, err
			}
		default:
			return nil, ErrListenerSocketPathIsNotAUnixSocket
		}
	}

	fd, err := net.Listen("unix", cfg.SocketPath)
	if err != nil {
		return nil, err
	}

	err = os.Chown(cfg.SocketPath, int(cfg.TrustedUserId), int(cfg.AllowedGroupId))
	if err != nil {
		fd.Close()
		return nil, err
	}

	err = os.Chmod(cfg.SocketPath, 0660)
	if err != nil {
		fd.Close()
		return nil, err
	}

	l := &Listener{
		cfg: cfg,
		fd:  fd,
		kr:  agent.NewKeyring(),
	}
	return l, nil
}

func (l *Listener) Serve() error {
	for {
		fd, err := l.fd.Accept()
		if err != nil {
			return err
		}

		go func() {
			l.serveClient(fd)
			fd.Close()
		}()
	}
}

func (l *Listener) serveClient(c net.Conn) {
	uc, ok := c.(*net.UnixConn)
	if !ok {
		return
	}

	uid, gid, err := connPeerCredentials(uc)
	if err != nil {
		return
	}

	if uid == l.cfg.TrustedUserId {
		agent.ServeAgent(l.kr, c)
	} else if gid == l.cfg.AllowedGroupId {
		agent.ServeAgent(&SignOnlyKeyring{Agent: l.kr}, c)
	} else {
		agent.ServeAgent(&NoopKeyring{}, c)
	}
}
