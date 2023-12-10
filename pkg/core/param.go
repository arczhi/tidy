package core

import (
	"flag"
)

var (
	PathParam          string
	DirectoryNameParam string
	SortTypeParam      string
	TimeSpanParam      string
)

func ParseParam() {
	flag.StringVar(&PathParam, "path", "", "please give your directory path")
	flag.StringVar(&DirectoryNameParam, "dir", "new_directory", "please give new a directory name")
	flag.StringVar(&SortTypeParam, "type", "time", "time: sort by time\tfile_type: sort by file_type")
	flag.StringVar(&TimeSpanParam, "time_span", "", "give a time span, for example: 1d,1h,30min,30s")
	flag.Parse()
}
