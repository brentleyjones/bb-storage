// +build darwin

package filesystem

import (
	"golang.org/x/sys/unix"
)

func (d *localDirectory) linkat(srcDirfd int, src string, dstDirfd int, dst string, flags int) error {
	return unix.Clonefileat(srcDirfd, src, dstDirfd, dst, flags)
}
