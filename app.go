package ecs

type App struct {
	world *World

	running bool

	startUpSystems []StartUpSystem
	systems        []System
}

func NewApp() *App {
	return &App{
		world:          NewWorld(),
		startUpSystems: []StartUpSystem{},
		systems:        []System{},
		running:        true,
	}
}

func (a *App) AddStartUpSystem(fn ...StartUpSystem) {
	a.startUpSystems = append(a.startUpSystems, fn...)
}

func (a *App) AddSystem(system ...System) {
	a.systems = append(a.systems, system...)
}

func (a *App) Cancel() {
	a.running = false
}

func (a *App) Run() error {

	queue := NewQueue()
	for _, fn := range a.startUpSystems {
		fn(NewCommands(queue, a.world))
	}
	queue.Apply(a.world)

	for a.running {
		for _, s := range a.systems {
			s(NewSystemContext(a.world))
		}
	}

	return nil
}
