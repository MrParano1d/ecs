package core

import "github.com/mrparano1d/ecs"

type Plugin struct {
	environment string
}

var _ ecs.Plugin = &Plugin{}

func NewPlugin(environment string) *Plugin {
	return &Plugin{
		environment: environment,
	}
}

func (p *Plugin) Build(app *ecs.App) {
	app.AddStage(NewFirstStage(p.environment))
	app.AddStageAfter(StageFirst, NewLastStage())
	app.AddStageAfter(StageFirst, NewPreUpdateStage())
	app.AddStageAfter(StagePreUpdate, NewUpdateStage())
	app.AddStageAfter(StageUpdate, NewPostUpdateStage())
}
