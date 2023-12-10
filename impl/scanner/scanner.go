package scanner

import (
	"io/fs"
	"os"
)

type Scanner struct {
	path string
}

// func New() *Scanner {
// 	return &Scanner{}
// }

func (s *Scanner) SetUp(path string) error {
	s.path = path
	return nil
}

func (s *Scanner) Scan() (*[]fs.DirEntry, error) {
	dirEntries, err := os.ReadDir(s.path)
	if err != nil {
		return nil, err
	}
	return &dirEntries, nil
}
