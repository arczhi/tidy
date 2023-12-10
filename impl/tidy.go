package impl

import (
	"errors"
	"os"

	exporter_impl "github.com/arczhi/tidy/impl/exporter"
	scanner_impl "github.com/arczhi/tidy/impl/scanner"
	sorter_impl "github.com/arczhi/tidy/impl/sorter"

	"github.com/arczhi/tidy/pkg/core"
	"github.com/arczhi/tidy/pkg/tool"
)

var scanner core.Scanner = &scanner_impl.Scanner{}
var sorter core.Sorter = &sorter_impl.Sorter{}
var exporter core.Exporter = &exporter_impl.Exporter{}

var (
	err         error
	importPath  string
	sortOptions *core.SortOptions
)

type tidy struct {
}

func New(path string, options ...core.SortOpt) (*tidy, error) {
	var t = &tidy{}
	path, err := tool.CheckPath(path)
	if err != nil {
		return nil, err
	}
	importPath = path
	dir, _ := os.Stat(importPath + "/" + core.DirectoryNameParam)
	if dir != nil {
		return nil, errors.New("duplicate directory existed!")
	}

	scanner.SetUp(path)
	entries, err := scanner.Scan()
	if err != nil {
		return nil, err
	}
	sorter.SetUp(entries, options...)

	t.initSortOptions(options...)
	return t, nil
}

func (t *tidy) initSortOptions(options ...core.SortOpt) {
	sortOptions = core.NewSortOptions()
	for _, option := range options {
		option(sortOptions)
	}
}

func (t *tidy) Exec() error {
	sortedEntries, err := sorter.Sort()
	if err != nil {
		return err
	}
	exporter.SetUp(&core.Setting{
		ImportPath:    importPath,
		SortedEntries: sortedEntries,
		SortOptions:   sortOptions,
	})
	return exporter.Export()
}
