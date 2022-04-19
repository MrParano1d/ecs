package ecs

type World struct {
	entities   EntityMap
	nextEntity Entity

	resources ResourceMap
	events    EventMap
}

func NewWorld(events EventMap) *World {
	return &World{
		entities:   EntityMap{},
		nextEntity: 1,
		resources:  ResourceMap{},
		events:     events,
	}
}

func (w *World) Entities() EntityMap {
	return w.entities
}

func (w *World) NextEntity() Entity {
	defer func() {
		w.nextEntity++
	}()
	return w.nextEntity
}

func (w *World) Resources() ResourceMap {
	return w.resources
}

func (w *World) Events() EventMap {
	return w.events
}
