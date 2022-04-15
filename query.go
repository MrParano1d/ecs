package ecs

type Query[C Component] struct {
	world *World
}

func NewQuery[C Component](w *World) *Query[C] {
	return &Query[C]{
		world: w,
	}
}

type Filter struct {
	with []Component
}

type FilterOption func(f *Filter)

func NewFilter(opts ...FilterOption) *Filter {
	f := &Filter{
		with: []Component{},
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func WithComponentFilter(component Component) FilterOption {
	return func(f *Filter) {
		f.with = append(f.with, component)
	}
}

func (q *Query[C]) Find(filter *Filter) []C {
	var comp C
	var cs []C

	if filter == nil {
		filter = NewFilter()
	}

	for _, c := range q.world.Entities().ComponentsByType(comp, filter.with...) {
		cs = append(cs, c.(C))
	}

	return cs
}
