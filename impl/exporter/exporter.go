package exporter

import (
	"io"
	"io/fs"
	"os"
	"sync"
	"time"

	"github.com/arczhi/tidy/pkg/constants"
	"github.com/arczhi/tidy/pkg/core"
	"github.com/arczhi/tidy/pkg/tool"

	"github.com/arczhi/tidy/pkg/tool/file_time"
)

type Exporter struct {
	setting *core.Setting
}

func (e *Exporter) SetUp(setting *core.Setting) {
	e.setting = setting
}

func (e *Exporter) Export() error {
	var (
		wg                                    sync.WaitGroup
		errorChan                             = make(chan error, constants.ERROR_CHAN_NUM)
		subDirectory                          string
		createdTime, accessTime, modifiedTime time.Time
	)

	// SortedEntries key
	// sort_by_time, for example: "2006-01-02 15.04.05" / "2006-01-02 15.04" / "2006-01-02 15"
	// sort_by_file_type, for example: "docx" / "doc" / "pptx"
	for key, entries := range e.setting.SortedEntries {
		// fmt.Println("测试", key)
		for _, entry := range entries {
			wg.Add(1)
			go func(key string, entry fs.DirEntry) {
				defer wg.Done()

				path := e.setting.ImportPath + "/" + entry.Name()
				file, err := os.Open(path)
				if err != nil {
					errorChan <- err
					return
				}
				defer file.Close()

				fileInfo, err := entry.Info()
				if err != nil {
					errorChan <- err
					return
				}

				if e.setting.SortOptions.SortByTimeSpan() {
					subDirectory = fileInfo.ModTime().Format(e.setting.SortOptions.TimeFormat())
				} else if e.setting.SortOptions.SortByFileType() {
					subDirectory = key
				}

				newFilePath := tool.NewFilePath(e.setting.ImportPath, core.DirectoryNameParam, subDirectory, tool.GetFileName(entry.Name()))
				newFile, err := os.Create(newFilePath)
				if err != nil {
					errorChan <- err
					return
				}
				defer newFile.Close()
				data, err := io.ReadAll(file)
				if err != nil {
					errorChan <- err
					return
				}
				newFile.Write(data)

				if tool.OsType == constants.OS_WINDOWS {
					createdTime, accessTime, modifiedTime = file_time.GetFileTime(e.setting.ImportPath + "/" + entry.Name())
				} else if tool.OsType == constants.OS_LINUX || tool.OsType == constants.OS_DARWIN {
					createdTime, accessTime, modifiedTime = file_time.GetFileTime(e.setting.ImportPath + "/" + entry.Name())
					// createdTime, accessTime, modifiedTime = fileInfo.ModTime(), fileInfo.ModTime(), fileInfo.ModTime()
				}
				err = file_time.SetFileTime(newFilePath, createdTime, accessTime, modifiedTime)
				if err != nil {
					errorChan <- err
					return
				}

			}(key, *entry)
		}
	}

	wg.Wait()

	select {
	case err := <-errorChan:
		e.clean()
		return err
	default:
	}

	return nil
}

func (e *Exporter) clean() {
	os.Remove(e.setting.ImportPath + "/" + constants.NEW_DIRECTORY_NAME)
}
