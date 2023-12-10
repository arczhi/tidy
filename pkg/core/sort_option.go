package core

import (
	"time"
)

func NewSortOptions() *SortOptions {
	return &SortOptions{}
}

type SortOptions struct {
	byTimeSpan time.Duration
	byFileType bool
}

func (s *SortOptions) ByTimeSpan() time.Duration {
	return s.byTimeSpan
}

func (s *SortOptions) ByFileType() bool {
	return s.byFileType
}

type SortOpt func(options *SortOptions)

func WithTimeSpan(timeSpan time.Duration) SortOpt {
	return func(op *SortOptions) {
		op.byTimeSpan = timeSpan
	}
}

func WithFileType() SortOpt {
	return func(op *SortOptions) {
		op.byFileType = true
	}
}
