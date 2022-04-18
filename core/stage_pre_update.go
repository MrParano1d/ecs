package core

import "github.com/mrparano1d/ecs"

const (
	StagePreUpdate = "pre_update"
)

type PreUpdateStage struct {
	ecs.Stage
}

var _ ecs.Stage = &PreUpdateStage{}

func NewPreUpdateStage() *PreUpdateStage {
	return &PreUpdateStage{
		Stage: ecs.NewDefaultStage(),
	}
}

func (p *PreUpdateStage) Name() string {
	return StagePreUpdate
}
