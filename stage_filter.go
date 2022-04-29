package ecs

type StageFilter struct {
	hasFilter bool
	labels    []string
}

type StageFilterOption func(f *StageFilter)

func WithStageLabelFilter(labels ...string) StageFilterOption {
	return func(f *StageFilter) {
		f.labels = append(f.labels, labels...)
	}
}

func NewStageFilter(opts ...StageFilterOption) *StageFilter {
	f := &StageFilter{
		labels:    []string{},
		hasFilter: false,
	}

	for _, opt := range opts {
		f.hasFilter = true
		opt(f)
	}

	return f
}

func (f *StageFilter) HasFilters() bool {
	return f.hasFilter
}
