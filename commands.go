package ecs

type Command interface {
	Write(world *World)
}

type Commands struct {
	queue *Queue
	world *World
}

func NewCommands(queue *Queue, world *World) *Commands {
	return &Commands{
		queue: queue,
		world: world,
	}
}

func (c *Commands) Spawn(components ...Component) *EntityCommand {
	ec := NewEntityCommand(c.world.NextEntity())
	for _, c := range components {
		ec.Insert(c)
	}
	c.queue.Push(ec)
	return ec
}

func (c *Commands) InvokeResource(cb ResourceInvoker) *ResourceCommand {
	cmd := NewResourceCommand(cb)
	c.queue.Push(cmd)
	return cmd
}

func (c *Commands) Cancel() *CancelCommand {
	cmd := NewCancelCommand(true)
	c.queue.Push(cmd)
	return cmd
}
