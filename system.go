package ecs

type StartUpSystem func(commands *Commands)

type System func(ctx *SystemContext)

type SystemContext struct {
	World    *World
	Commands *Commands
	events   EventMap
}

func NewSystemContext(w *World, c *Commands, events EventMap) *SystemContext {
	return &SystemContext{
		World:    w,
		Commands: c,
		events:   events,
	}
}

func (c *SystemContext) EventWriter(event Event) *EventWriter {
	return c.events[event]
}

func (c *SystemContext) EventReader(event Event) *EventReader {
	return NewEventReader(c.events[event])
}
