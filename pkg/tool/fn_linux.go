// + build linux,darwin

package tool

func GetFileTime(path string) (createdTime time.Time, accessTime time.Time, modifiedTime time.Time) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}
	}
	// windows环境下 *syscall.Stat_t未定义
	linuxFileAttr := fileInfo.Sys().(*syscall.Stat_t)
	return SecondToTime(linuxFileAttr.Ctim.Sec), SecondToTime(linuxFileAttr.Atim.Sec), SecondToTime(linuxFileAttr.Mtim.Sec)
}
