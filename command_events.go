package ecs

type EventsCommand struct {
	cb EventInvoker
}

var _ Command = &EventsCommand{}

func NewEventsCommand(cb EventInvoker) *EventsCommand {
	return &EventsCommand{
		cb: cb,
	}
}

func (r *EventsCommand) Write(world *World) {
	r.cb(world.Events())
}
