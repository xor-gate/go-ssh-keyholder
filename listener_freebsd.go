// +build freebsd
package main

import "golang.org/x/sys/unix"

func listenerGetSocketUidGid(fd int) (uint32, uint32, error) {
	xucred, err := unix.GetsockoptXucred(fd, unix.SOL_LOCAL, unix.LOCAL_PEERCRED)
	if err != nil {
		return 0,0,err
	}
	return xucred.Uid, xucred.Groups[0], nil
}
