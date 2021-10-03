// +build linux
package main

import "golang.org/x/sys/unix"

func listenerGetSocketUidGid(fd int) (uint32, uint32, error) {
	ucred, err := unix.GetsockoptUcred(fd, unix.SOL_SOCKET, unix.SO_PEERCRED)
	if err != nil {
		return 0,0,err
	}
	return ucred.Uid, ucred.Gid, nil
}
