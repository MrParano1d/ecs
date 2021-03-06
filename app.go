package ecs

type App struct {
	world *World

	running bool

	stages *Stages

	events EventMap
}

func NewApp() *App {
	events := EventMap{}

	stages := NewStages()
	a := &App{
		world:   NewWorld(events),
		running: true,
		stages:  stages,
		events:  events,
	}

	a.AddEvent(
		func(eventMap EventMap) {
			AddEvent[AppExitEvent](eventMap)
		},
	)

	return a
}

func (a *App) Stages() *Stages {
	return a.stages
}

func (a *App) AddStage(stage Stage) *App {
	a.stages.Add(stage)
	return a
}

func (a *App) AddStageBefore(beforeStageName string, stage Stage) *App {
	a.stages.AddBefore(beforeStageName, stage)
	return a
}

func (a *App) AddStageAfter(afterStageName string, stage Stage) *App {
	a.stages.AddAfter(afterStageName, stage)
	return a
}

func (a *App) AddStartUpSystem(fn ...StartUpSystem) *App {
	return a.AddStartUpSystemToStage(StageUpdate, fn...)
}

func (a *App) AddStartUpSystemToStage(stageName string, fn ...StartUpSystem) *App {
	a.stages.GetStage(stageName).AddStartUpSystem(fn...)
	return a
}

func (a *App) AddSystem(system ...System) *App {
	return a.AddSystemToStage(StageUpdate, system...)
}

func (a *App) AddSystemToStage(stageName string, system ...System) *App {
	a.stages.GetStage(stageName).AddSystem(system...)
	return a
}

func (a *App) AddPlugin(plugin ...Plugin) *App {
	for _, p := range plugin {
		p.Build(a)
	}
	return a
}

func (a *App) AddEvent(cb EventInvoker) *App {
	cb(a.events)
	return a
}

func (a *App) Events() EventMap {
	return a.events
}

func (a *App) Cancel() {
	a.running = false
}

func (a *App) FlushEvents() {
	for _, w := range a.events {
		w.Flush()
	}
}

func (a *App) SetupSystems(stageFilters ...StageFilterOption) {
	for _, stage := range a.stages.GetOrderedStages(stageFilters...) {
		queue := NewQueue()
		for _, fn := range stage.StartUpSystems() {
			fn(NewCommands(&queue, a.world))
		}
		queue.Apply(a.world)

		reader := NewEventReader(a.Events()[AppExitEvent{}])
		if reader.Next() {
			a.Cancel()
		}
	}
}

func (a *App) World() *World {
	return a.world
}

func (a *App) RunSystems(stageFilters ...StageFilterOption) {
	for _, stage := range a.stages.GetOrderedStages(stageFilters...) {
		scheduler := NewScheduler(stage)
		scheduler.RunSystems(a.world, a.events)
	}

	reader := NewEventReader(a.Events()[AppExitEvent{}])
	if reader.Next() {
		a.Cancel()
	}

	a.FlushEvents()
}

func (a *App) Run() error {

	a.SetupSystems()

	for a.running {
		a.RunSystems()
	}

	return nil
}
