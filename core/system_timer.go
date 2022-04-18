package core

import (
	"github.com/mrparano1d/ecs"
)

func TimerStartUp() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			ecs.AddResource[*Time](resourceMap, NewTime())
		})
	}
}

func TimerSystem() ecs.System {
	return func(ctx ecs.SystemContext) {
		ctx.Commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			t := ecs.GetResource[*Time](resourceMap)
			t.Update()
		})
	}
}
