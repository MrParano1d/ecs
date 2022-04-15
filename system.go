package ecs

type StartUpSystem func(commands *Commands)

type System func(ctx *SystemContext)

type SystemContext struct {
	World *World
}

func NewSystemContext(w *World) *SystemContext {
	return &SystemContext{
		World: w,
	}
}
