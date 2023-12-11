package core

import "github.com/arczhi/tidy/pkg/constants"

func NewSortOptions() *SortOptions {
	return &SortOptions{}
}

type SortOptions struct {
	timeSpanOption
	fileTypeOption
}

type timeSpanOption struct {
	sortByTimeSpan bool
	timeSpan       string
}

type fileTypeOption struct {
	sortByFileType bool
}

func (s *timeSpanOption) SortByTimeSpan() bool {
	return s.sortByTimeSpan
}

func (s *timeSpanOption) TimeFormat() string {

	switch s.timeSpan {
	case "":
		return constants.TIME_FORMAT_ACCURATE_TO_HOUR
	case "year":
		return constants.TIME_FORMAT_ACCURATE_TO_YEAR
	case "month":
		return constants.TIME_FORMAT_ACCURATE_TO_MONTH
	case "day":
		return constants.TIME_FORMAT_ACCURATE_TO_DAY
	case "hour":
		return constants.TIME_FORMAT_ACCURATE_TO_HOUR
	case "minute":
		return constants.TIME_FORMAT_ACCURATE_TO_MINUTE
	case "second":
		return constants.TIME_FORMAT_ACCURATE_TO_SECOND
	default:
		return constants.TIME_FORMAT_ACCURATE_TO_HOUR
	}
}

func (s *fileTypeOption) SortByFileType() bool {
	return s.sortByFileType
}

type SortOpt func(options *SortOptions)

func WithTimeSpan(timeSpan string) SortOpt {
	return func(op *SortOptions) {
		op.sortByTimeSpan = true
		op.timeSpan = timeSpan
	}
}

func WithFileType() SortOpt {
	return func(op *SortOptions) {
		op.sortByFileType = true
	}
}
