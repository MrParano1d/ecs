package ecs

type ResourceInvoker func(resourceMap ResourceMap)

type ResourceCommand struct {
	cb ResourceInvoker
}

var _ Command = &ResourceCommand{}

func NewResourceCommand(cb ResourceInvoker) *ResourceCommand {
	return &ResourceCommand{
		cb: cb,
	}
}

func (r *ResourceCommand) Write(world *World) {
	r.cb(world.Resources())
}
