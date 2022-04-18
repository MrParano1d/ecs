package core

import (
	"fmt"
	"github.com/mrparano1d/ecs"
	"go.uber.org/zap"
)

const (
	StageFirst = "first"
)

type FirstStage struct {
	ecs.Stage
}

var _ ecs.Stage = &FirstStage{}

func NewFirstStage(environment string) *FirstStage {
	s := &FirstStage{
		Stage: ecs.NewDefaultStage(),
	}

	s.AddStartUpSystem(TimerStartUp())
	s.AddSystem(TimerSystem())

	s.AddStartUpSystem(func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			if environment == EnvDebug {
				ecs.AddResource[*ConfigRes](resourceMap, NewDebugConfig())
				logger, err := zap.NewDevelopment()
				if err != nil {
					panic(fmt.Errorf("failed to init logger: %v", err))
				}
				ecs.AddResource[*zap.Logger](resourceMap, logger)
			} else {
				logger, err := zap.NewProduction()
				if err != nil {
					panic(fmt.Errorf("failed to init logger: %v", err))
				}
				ecs.AddResource[*ConfigRes](resourceMap, NewConfig())
				ecs.AddResource[*zap.Logger](resourceMap, logger)
			}
		})
	})

	return s
}

func (p *FirstStage) Name() string {
	return StageFirst
}
