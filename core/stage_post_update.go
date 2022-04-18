package core

import "github.com/mrparano1d/ecs"

const (
	StagePostUpdate = "post_update"
)

type PostUpdateStage struct {
	ecs.Stage
}

var _ ecs.Stage = &PostUpdateStage{}

func NewPostUpdateStage() *PostUpdateStage {
	return &PostUpdateStage{
		Stage: ecs.NewDefaultStage(),
	}
}

func (p *PostUpdateStage) Name() string {
	return StagePostUpdate
}
