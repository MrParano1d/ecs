package ecs

type World struct {
	entities   EntityMap
	nextEntity Entity
}

func NewWorld() *World {
	return &World{
		entities:   EntityMap{},
		nextEntity: 1,
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
