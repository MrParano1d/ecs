package ecs

type World struct {
	entities   EntityMap
	nextEntity Entity

	resources ResourceMap
}

func NewWorld() *World {
	return &World{
		entities:   EntityMap{},
		nextEntity: 1,
		resources:  ResourceMap{},
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
