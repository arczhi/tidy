package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"github.com/arczhi/tidy/impl"
	"github.com/arczhi/tidy/pkg/constants"
	"github.com/arczhi/tidy/pkg/core"
)

var (
	tidy core.Tidy
	err  error
)

func main() {

	startAt := time.Now()
	core.ParseParam()

	path := core.PathParam
	if core.SortTypeParam == constants.SORT_BY_FILE_TYPE {
		tidy, err = impl.New(path, core.WithFileType())
		if err != nil {
			log.Println(err)
			return
		}
	} else if core.SortTypeParam == constants.SORT_BY_TIME {
		tidy, err = impl.New(path, core.WithTimeSpan(core.TimeSpanParam))
		if err != nil {
			log.Println(err)
			return
		}
	}

	err = tidy.Exec()
	if err != nil {
		log.Println(err)
		log.Println(string(debug.Stack()))
		return
	}

	fmt.Println("🧹", "tidy finished")
	fmt.Println("⏱️", " cost", time.Now().Sub(startAt).Milliseconds(), "ms")
}
