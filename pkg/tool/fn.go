package tool

import (
	"errors"
	"io/fs"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/arczhi/tidy/pkg/constants"
)

var (
	OsType      = runtime.GOOS
	currentPath string
)

func CheckPath(path string) (string, error) {
	if len(path) > 0 {
		_, err := os.Stat(path)
		if err != nil {
			return "", err
		}
		return path, nil
	}
	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return currentPath, nil
}

func GetFileType(path string) string {
	list := strings.Split(path, ".")
	return list[len(list)-1]
}

func GetFileName(path string) string {
	if OsType == "linux" || OsType == "darwin" {
		list := strings.Split(path, "/")
		return list[len(list)-1]
	}
	if OsType == "windows" {
		list := strings.Split(path, "\\")
		return list[len(list)-1]
	}
	return ""
}

func NewFilePath(currentPath, newDirectory, subDirectory, fileName string) string {
	var dirPath string
	if OsType == "linux" || OsType == "darwin" {
		dirPath = currentPath + "/" + newDirectory + "/" + subDirectory + "/"

	}
	if OsType == "windows" {
		dirPath = currentPath + "\\" + newDirectory + "\\" + subDirectory + "\\"
	}
	os.MkdirAll(dirPath, os.ModePerm)
	return dirPath + fileName
}

func IntToString(num int) string {
	return strconv.Itoa(num)
}

func ExtendDirEntries(entries *[][]fs.DirEntry, index int) {
	if len(*entries) < index+1 {
		length := index + 1 - len(*entries)
		for i := 0; i < length; i++ {
			*entries = append(*entries, []fs.DirEntry{})
		}
	}
}

func GetDirEntryNum(entries map[string][]*fs.DirEntry) int {
	var num int
	for _, entryList := range entries {
		num += len(entryList)
	}
	return num
}

func ParseTimeSpan(timeSpan string) (time.Duration, error) {
	var (
		numStr      string
		timeUnitStr string
	)
	for _, r := range timeSpan {
		str := string(r)
		_, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			numStr += str
		}
		if err != nil {
			timeUnitStr += str
		}
	}
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return time.Second, err
	}

	var timeUnit time.Duration
	switch timeUnitStr {
	case "d":
		timeUnit = 24 * time.Hour
	case "h":
		timeUnit = time.Hour
	case "min":
		timeUnit = time.Minute
	case "s":
		timeUnit = time.Second
	default:
		return time.Second, errors.New("invalid time unit")
	}

	return time.Duration(num) * timeUnit, nil

}

func IsBinary(name string) bool {
	if name == constants.BINARY_NAME_IN_LINUX || name == constants.BINARY_NAME_IN_WINDOWS {
		return true
	}
	return false
}

func IsBinaryInPath(path string) bool {
	if currentPath == "" {
		currentPath, _ = os.Getwd()
	}
	if path == "" || path == currentPath {
		return true
	}
	return false
}
