package core

import (
	"io/fs"
)

type Tidy interface {
	Exec() error
}

type Scanner interface {
	SetUp(path string) error
	Scan() (*[]fs.DirEntry, error)
}

type Sorter interface {
	SetUp(dirEntries *[]fs.DirEntry, options ...SortOpt)
	Sort() (map[string][]*fs.DirEntry, error)
}

type Exporter interface {
	SetUp(setting *Setting)
	Export() error
}
