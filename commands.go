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

func (c *Commands) Spawn() *EntityCommand {
	ec := NewEntityCommand(c.world.NextEntity())
	c.queue.Push(ec)
	return ec
}
