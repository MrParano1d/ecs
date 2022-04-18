package ecs_test

import "github.com/mrparano1d/ecs"

type UpdateStage struct {
	ecs.Stage
}

func NewUpdateStage() *UpdateStage {
	s := &UpdateStage{
		Stage: ecs.NewDefaultStage(),
	}

	return s
}

func (s *UpdateStage) Name() string {
	return "update"
}

type TestPlugin struct {
}

var _ ecs.Plugin = &TestPlugin{}

func NewTestPlugin() *TestPlugin {
	return &TestPlugin{}
}

func (p *TestPlugin) Build(app *ecs.App) {
	app.AddStage(NewUpdateStage())
}
