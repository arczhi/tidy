//go:build windows
// +build windows

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
	winFileAttr := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	return SecondToTime(winFileAttr.CreationTime.Nanoseconds() / 1e9), SecondToTime(winFileAttr.LastAccessTime.Nanoseconds() / 1e9), SecondToTime(winFileAttr.LastWriteTime.Nanoseconds() / 1e9)
}

func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func SetFileTime(path string, createdTime time.Time, accessTime time.Time, modifiedTime time.Time) error {

	pathPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil
	}
	handle, err := syscall.CreateFile(pathPtr, syscall.FILE_WRITE_ATTRIBUTES, syscall.FILE_SHARE_WRITE, nil, syscall.OPEN_EXISTING, syscall.FILE_FLAG_BACKUP_SEMANTICS, 0)
	if err != nil {
		return nil
	}
	defer syscall.Close(handle)
	c := syscall.NsecToFiletime(syscall.TimespecToNsec(syscall.NsecToTimespec(createdTime.UnixNano())))
	a := syscall.NsecToFiletime(syscall.TimespecToNsec(syscall.NsecToTimespec(accessTime.UnixNano())))
	m := syscall.NsecToFiletime(syscall.TimespecToNsec(syscall.NsecToTimespec(modifiedTime.UnixNano())))
	syscall.SetFileTime(handle, &c, &a, &m)

	return nil
}
