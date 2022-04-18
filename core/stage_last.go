package core

import "github.com/mrparano1d/ecs"

const (
	StageLast = "last"
)

type LastStage struct {
	ecs.Stage
}

var _ ecs.Stage = &LastStage{}

func NewLastStage() *LastStage {
	return &LastStage{
		Stage: ecs.NewDefaultStage(),
	}
}

func (p *LastStage) Name() string {
	return StageLast
}
