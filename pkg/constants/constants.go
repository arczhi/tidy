package constants

type SortType = string

const (
	SORT_BY_TIME      SortType = "time"
	SORT_BY_FILE_TYPE SortType = "file_type"
)

const (
	NEW_DIRECTORY_NAME = "your_directory"
	ERROR_CHAN_NUM     = 10
)

type timeSpan = string

const (
	TIME_SPAN_YEAR   = "year"
	TIME_SPAN_MONTH  = "month"
	TIME_SPAN_DAY    = "day"
	TIME_SPAN_HOUR   = "hour"
	TIME_SPAN_MINUTE = "minute"
	TIME_SPAN_SECOND = "second"
)

const (
	TIME_FORMAT_ACCURATE_TO_YEAR   = "2006"
	TIME_FORMAT_ACCURATE_TO_MONTH  = "2006-01"
	TIME_FORMAT_ACCURATE_TO_DAY    = "2006-01-02"
	TIME_FORMAT_ACCURATE_TO_HOUR   = "2006-01-02 15"
	TIME_FORMAT_ACCURATE_TO_MINUTE = "2006-01-02 15.04"
	TIME_FORMAT_ACCURATE_TO_SECOND = "2006-01-02 15.04.05"
)

const (
	OS_WINDOWS = "windows"
	OS_LINUX   = "linux"
	OS_DARWIN  = "darwin"
)

const (
	BINARY_NAME_IN_WINDOWS = "tidy.exe"
	BINARY_NAME_IN_LINUX   = "tidy"
)
