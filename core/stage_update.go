package core

import "github.com/mrparano1d/ecs"

const (
	StageUpdate = "update"
)

type UpdateStage struct {
	ecs.Stage
}

var _ ecs.Stage = &UpdateStage{}

func NewUpdateStage() *UpdateStage {
	return &UpdateStage{
		Stage: ecs.NewDefaultStage(),
	}
}

func (p *UpdateStage) Name() string {
	return StageUpdate
}
