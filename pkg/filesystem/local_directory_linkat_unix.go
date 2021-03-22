// +build freebsd linux

package filesystem

import (
	"golang.org/x/sys/unix"
)

func (d *localDirectory) linkat(srcDirfd int, src string, dstDirfd int, dst string, flags int) error {
	return unix.Linkat(srcDirfd, src, dstDirfd, dst, flags)
}
