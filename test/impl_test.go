package test

import (
	"log"
	"os"
	"testing"

	"github.com/arczhi/tidy/impl"
	"github.com/arczhi/tidy/pkg/constants"
	"github.com/arczhi/tidy/pkg/core"
)

const PATH = "./test_directory"

func TestMain(m *testing.M) {
	os.MkdirAll(PATH, 0755)
	os.Create(PATH + "/file_01.txt")
	os.Create(PATH + "/file_02.txt")
}

func TestSortByTime(t *testing.T) {
	defer clean()
	tidy, err := impl.New(PATH, core.WithTimeSpan(constants.TIME_FORMAT_ACCURATE_TO_HOUR))
	if err != nil {
		t.Error(err)
		return
	}
	if err := tidy.Exec(); err != nil {
		t.Error(err)
		return
	}
}

func TestSortByFileType(t *testing.T) {
	defer clean()
	tidy, err := impl.New(PATH, core.WithFileType())
	if err != nil {
		t.Error(err)
		return
	}
	if err := tidy.Exec(); err != nil {
		t.Error(err)
		return
	}
}

func clean() {
	err := os.RemoveAll(PATH)
	if err != nil {
		log.Println(err)
	}
}
