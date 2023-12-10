// + build windows

package tool

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
	winFileAttr := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	return SecondToTime(winFileAttr.CreationTime.Nanoseconds() / 1e9), SecondToTime(winFileAttr.LastAccessTime.Nanoseconds() / 1e9), SecondToTime(winFileAttr.LastWriteTime.Nanoseconds() / 1e9)
}

func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}
