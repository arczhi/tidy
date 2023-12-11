//go:build linux || darwin

// + build linux darwin

package file_time

import (
	"os"
	"syscall"
	"time"
)

func GetFileTime(path string) (createdTime time.Time, accessTime time.Time, modifiedTime time.Time) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}
	}
	linuxFileAttr := fileInfo.Sys().(*syscall.Stat_t)
	return SecondToTime(linuxFileAttr.Ctim.Sec),
		SecondToTime(linuxFileAttr.Atim.Sec),
		SecondToTime(linuxFileAttr.Mtim.Sec)
}

func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// created time is not supported to be set in linux
func SetFileTime(path string, createdTime time.Time, accessTime time.Time, modifiedTime time.Time) error {
	// func Chtimes(name string, atime time.Time, mtime time.Time) error
	return os.Chtimes(path, accessTime, modifiedTime)
}
