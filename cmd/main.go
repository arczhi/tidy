package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"tidy/impl"
	"tidy/pkg/constants"
	"tidy/pkg/core"
	"tidy/pkg/tool"
	"time"
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
		var timeSpan time.Duration
		if core.TimeSpanParam == "" {
			timeSpan = 6 * time.Hour
		} else {
			timeSpan, err = tool.ParseTimeSpan(core.TimeSpanParam)
			if err != nil {
				log.Println(err)
				return
			}
		}
		tidy, err = impl.New(path, core.WithTimeSpan(timeSpan))
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
