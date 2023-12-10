package sorter

import (
	"hash/crc32"
	"io/fs"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/arczhi/tidy/pkg/core"
	"github.com/arczhi/tidy/pkg/tool"
)

type Sorter struct {
	option        *core.SortOptions
	sortedEntries *[][]fs.DirEntry
	locker        sync.Mutex
}

// func New() *Sorter {
// 	return &Sorter{}
// }

func (s *Sorter) SetUp(dirEntries *[]fs.DirEntry, options ...core.SortOpt) {
	opt := core.NewSortOptions()
	for _, option := range options {
		option(opt)
	}
	s.setDirEntries(dirEntries)
	s.setOption(opt)
	return
}

func (s *Sorter) setDirEntries(dirEntries *[]fs.DirEntry) {
	s.sortedEntries = &[][]fs.DirEntry{}
	(*s.sortedEntries) = append((*s.sortedEntries), *dirEntries)
}

func (s *Sorter) setOption(option *core.SortOptions) {
	s.option = option
}

func (s *Sorter) Sort() (*[][]fs.DirEntry, error) {

	//TODO 解决此处排序不稳定的问题
	s.sortByModTimeAsc()
	//TODO 支持手动选择分类顺序
	if s.option.ByTimeSpan().Seconds() > 0 {
		s.sortByTimeSpan()
	}
	if s.option.ByFileType() {
		s.sortByFileType()
	}

	return s.sortedEntries, nil
}

func (s *Sorter) sortByModTimeAsc() *Sorter {
	for _, entries := range *s.sortedEntries {
		if len(entries) == 0 {
			continue
		}
		sort.SliceStable(entries, func(i, j int) bool {
			f1, _ := (entries)[i].Info()
			f2, _ := (entries)[j].Info()
			return f1.ModTime().Before(f2.ModTime())
		})
	}
	return s
}

func (s *Sorter) sortByTimeSpan() *Sorter {
	var (
		sortedEntries    = &[][]fs.DirEntry{}
		wg               sync.WaitGroup
		firstFileModTime time.Time
		// lastFileModTime  time.Time
	)

	if len(*s.sortedEntries) > 0 {
		firstFileList := (*s.sortedEntries)[0]
		if len(firstFileList) > 0 {
			firstFile := firstFileList[0]
			firstFileInfo, _ := firstFile.Info()
			firstFileModTime = firstFileInfo.ModTime()
		}
		// lastFileList := (*s.sortedEntries)[len(*s.sortedEntries)-1]
		// if len(lastFileList) > 0 {
		// 	lastFile := lastFileList[len(lastFileList)-1]
		// 	lastFileInfo, _ := lastFile.Info()
		// 	lastFileModTime = lastFileInfo.ModTime()
		// }
	}

	for _, entries := range *s.sortedEntries {
		for index, entry := range entries {
			wg.Add(1)
			go func(index int, entry fs.DirEntry) {
				defer wg.Done()

				fileInfo, _ := entry.Info()
				if fileInfo.IsDir() {
					//TODO 处理文件夹的情况
					return
				}
				sortedIndex := int(math.Round(fileInfo.ModTime().Sub(firstFileModTime).Hours()) / s.option.ByTimeSpan().Hours())
				// fmt.Println("排序下标", sortedIndex)

				s.locker.Lock()
				defer s.locker.Unlock()
				tool.ExtendDirEntries(sortedEntries, sortedIndex)
				(*sortedEntries)[sortedIndex] = append((*sortedEntries)[sortedIndex], entry)

			}(index, entry)
		}
	}
	wg.Wait()

	s.sortedEntries = sortedEntries
	return s
}

func (s *Sorter) sortByFileType() *Sorter {
	var (
		sortedEntries = &[][]fs.DirEntry{}
		wg            sync.WaitGroup
	)

	for _, entries := range *s.sortedEntries {
		for index, entry := range entries {
			wg.Add(1)
			go func(index int, entry fs.DirEntry) {
				defer wg.Done()

				fileInfo, _ := entry.Info()
				if fileInfo.IsDir() {
					//TODO 处理文件夹的情况
					return
				}

				sortedIndex := int(
					math.Mod(
						float64(crc32.ChecksumIEEE([]byte(tool.GetFileType(fileInfo.Name())))),
						20),
				)
				s.locker.Lock()
				defer s.locker.Unlock()
				tool.ExtendDirEntries(sortedEntries, sortedIndex)
				(*sortedEntries)[sortedIndex] = append((*sortedEntries)[sortedIndex], entry)
			}(index, entry)
		}
	}
	wg.Wait()

	s.sortedEntries = sortedEntries
	return s
}
