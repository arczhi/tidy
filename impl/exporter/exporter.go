package exporter

import (
	"io"
	"io/fs"
	"os"
	"sync"
	"tidy/pkg/constants"
	"tidy/pkg/core"
	"tidy/pkg/tool"
)

type Exporter struct {
	setting *core.Setting
}

func (e *Exporter) SetUp(setting *core.Setting) {
	e.setting = setting
}

func (e *Exporter) Export() error {
	var (
		wg           sync.WaitGroup
		errorChan    = make(chan error, tool.GetDirEntryNum(e.setting.SortedEntries))
		subDirectory string
	)

	for index, entries := range *e.setting.SortedEntries {
		for _, entry := range entries {
			wg.Add(1)
			go func(index int, entry fs.DirEntry) {
				defer wg.Done()

				path := e.setting.ImportPath + "/" + entry.Name()
				file, err := os.Open(path)
				if err != nil {
					errorChan <- err
					return
				}
				defer file.Close()

				if e.setting.SortOptions.ByTimeSpan().Seconds() > 0 {

					firstFileInfo, err := (*e.setting.SortedEntries)[index][0].Info()
					if err != nil {
						errorChan <- err
						return
					}
					subDirectory = firstFileInfo.ModTime().Format(constants.TIME_FORMAT)
				} else if e.setting.SortOptions.ByFileType() {
					subDirectory = tool.GetFileType(path)
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

				createdTime, accessTime, modifiedTime := tool.GetFileTime(e.setting.ImportPath + "/" + entry.Name())
				err = tool.SetFileTime(newFilePath, createdTime, accessTime, modifiedTime)
				if err != nil {
					errorChan <- err
					return
				}

			}(index, entry)
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
