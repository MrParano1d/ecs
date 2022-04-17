package ecs

func TimerStartUp() StartUpSystem {
	return func(commands *Commands) {
		commands.InvokeResource(func(resourceMap ResourceMap) {
			AddResource[*Time](resourceMap, NewTime())
		})
	}
}

func TimerSystem() System {
	return func(ctx SystemContext) {
		ctx.Commands.InvokeResource(func(resourceMap ResourceMap) {
			t := GetResource[*Time](resourceMap)
			t.Update()
		})
	}
}
