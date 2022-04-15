package ecs

type Plugin interface {
	Build(app *App)
}
