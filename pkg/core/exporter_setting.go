package core

import "io/fs"

type Setting struct {
	ImportPath    string
	SortedEntries *[][]fs.DirEntry
	SortOptions   *SortOptions
}
