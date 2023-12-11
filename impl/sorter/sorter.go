package sorter

import (
	"io/fs"
	"sync"

	"github.com/arczhi/tidy/pkg/constants"
	"github.com/arczhi/tidy/pkg/core"
	"github.com/arczhi/tidy/pkg/tool"
)

type Sorter struct {
	option        *core.SortOptions
	entries       *[]fs.DirEntry
	sortedEntries map[string][]*fs.DirEntry
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
	s.sortedEntries = make(map[string][]*fs.DirEntry, len(*dirEntries))
	s.setDirEntries(dirEntries)
	s.setOption(opt)
	return
}

func (s *Sorter) setDirEntries(dirEntries *[]fs.DirEntry) {
	s.entries = dirEntries
}

func (s *Sorter) setOption(option *core.SortOptions) {
	s.option = option
}

func (s *Sorter) Sort() (map[string][]*fs.DirEntry, error) {

	//TODO 支持手动选择分类顺序
	if s.option.SortByTimeSpan() {
		s.sort(constants.SORT_BY_TIME, s.entries)
	}
	if s.option.SortByFileType() {
		s.sort(constants.SORT_BY_FILE_TYPE, s.entries)
	}

	return s.sortedEntries, nil
}

func (s *Sorter) sort(sortType string, entries *[]fs.DirEntry) *Sorter {
	var wg sync.WaitGroup

	for _, entry := range *entries {
		wg.Add(1)
		go func(entry fs.DirEntry) {
			defer wg.Done()

			if tool.IsBinary(entry.Name()) && tool.IsBinaryInPath(core.PathParam) {
				return
			}

			fileInfo, _ := entry.Info()
			if fileInfo.IsDir() {
				//TODO 处理文件夹的情况
				return
			}

			var sortKey string
			switch sortType {
			case constants.SORT_BY_TIME:
				sortKey = fileInfo.ModTime().Format(s.option.TimeFormat())
			case constants.SORT_BY_FILE_TYPE:
				sortKey = tool.GetFileType(fileInfo.Name())
			}

			s.locker.Lock()
			defer s.locker.Unlock()
			if _, exist := s.sortedEntries[sortKey]; !exist {
				s.sortedEntries[sortKey] = []*fs.DirEntry{}
			}
			s.sortedEntries[sortKey] = append(s.sortedEntries[sortKey], &entry)

		}(entry)
	}
	wg.Wait()

	return s
}
