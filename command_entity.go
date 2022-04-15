package ecs

type EntityCommand struct {
	entity  Entity
	inserts ComponentMap
}

func NewEntityCommand(entity Entity) *EntityCommand {
	return &EntityCommand{
		entity:  entity,
		inserts: ComponentMap{},
	}
}

func (e *EntityCommand) Insert(component Component) *EntityCommand {
	e.inserts.Set(component)
	return e
}

func (e *EntityCommand) Write(w *World) {
	var components []Component
	for _, c := range e.inserts {
		components = append(components, c)
	}
	w.Entities().AddComponents(e.entity, components...)
}
