package ecs

type Query struct {
	world *World
}

func NewQuery(w *World) *Query {
	return &Query{
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

func WithComponentFilter(components ...Component) FilterOption {
	return func(f *Filter) {
		f.with = append(f.with, components...)
	}
}

func (q *Query) Find(filter *Filter) []Entity {
	var entities []Entity

	if filter == nil {
		filter = NewFilter()
	}

	for _, e := range q.world.Entities().EntitiesByComponentTypes(filter.with...) {
		entities = append(entities, e)
	}

	return entities
}
