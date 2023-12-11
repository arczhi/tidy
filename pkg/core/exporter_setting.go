package core

import "io/fs"

type Setting struct {
	ImportPath    string
	SortedEntries map[string][]*fs.DirEntry `comment:"map key: timestamp ,for example '2006-01-02 15.04.05'; file_type ,for example:'docx' "`
	SortOptions   *SortOptions
}
