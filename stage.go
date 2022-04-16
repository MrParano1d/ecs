package ecs

import "fmt"

type StageOrder []string

func (sc StageOrder) findStageIndex(stageName string) int {
	for i, s := range sc {
		if s == stageName {
			return i
		}
	}
	return -1
}

func (sc *StageOrder) Insert(stageName string) {
	if sc.findStageIndex(stageName) != -1 {
		return
	}
	*sc = append(*sc, stageName)
}

func (sc *StageOrder) InsertAfter(stageName string, newStageName string) {
	if sc.findStageIndex(newStageName) != -1 {
		return
	}
	stageIndex := sc.findStageIndex(stageName)
	if stageIndex == -1 || stageIndex >= len(*sc) {
		sc.Insert(newStageName)
		return
	}
	stageIndex++
	*sc = append(*sc, "")
	copy((*sc)[stageIndex+1:], (*sc)[stageIndex:])
	(*sc)[stageIndex] = newStageName
}

func (sc *StageOrder) InsertBefore(stageName string, newStageName string) {
	if sc.findStageIndex(newStageName) != -1 {
		return
	}
	stageIndex := sc.findStageIndex(stageName)
	if stageIndex == -1 || stageIndex >= len(*sc) {
		sc.Insert(newStageName)
		return
	}
	*sc = append(*sc, "")
	copy((*sc)[stageIndex+1:], (*sc)[stageIndex:])
	(*sc)[stageIndex] = newStageName
}

type StageMap map[string]Stage

func (sm StageMap) Get(stageName string) (Stage, bool) {
	if s, ok := sm[stageName]; ok {
		return s, true
	}
	return nil, false
}

func (sm StageMap) Set(stageName string, stage Stage) {
	sm[stageName] = stage
}

type Stages struct {
	registry StageMap
	order    StageOrder
}

func NewStages() *Stages {
	return &Stages{
		registry: StageMap{},
		order:    StageOrder{},
	}
}

func (s *Stages) Add(stageName string, stage Stage) {
	s.registry.Set(stageName, stage)
	s.order.Insert(stageName)
}

func (s *Stages) AddAfter(afterStageName string, stageName string, stage Stage) {
	s.registry.Set(stageName, stage)
	s.order.InsertAfter(afterStageName, stageName)
}

func (s *Stages) AddBefore(beforeStageName string, stageName string, stage Stage) {
	s.registry.Set(stageName, stage)
	s.order.InsertBefore(beforeStageName, stageName)
}

func (s *Stages) GetStage(stageName string) Stage {
	if stage, ok := s.registry.Get(stageName); ok {
		return stage
	}
	panic(fmt.Errorf("unknown stage: %s", stageName))
}

func (s *Stages) GetOrderedStages() []Stage {
	stages := make([]Stage, len(s.order))

	for i, stageName := range s.order {
		if stage, ok := s.registry.Get(stageName); !ok {
			panic(fmt.Errorf("mismatched stage order and registry length: registry=%d, order=%d", len(s.registry), len(s.order)))
		} else {
			stages[i] = stage
		}
	}
	return stages
}

const (
	StageUpdate = "update"
)

type Stage interface {
	AddStartUpSystem(fn ...StartUpSystem)
	AddSystem(system ...System)
	StartUpSystems() []StartUpSystem
	Systems() []System
}

type UpdateStage struct {
	startUpSystems []StartUpSystem
	systems        []System
}

func NewUpdateStage() *UpdateStage {
	return &UpdateStage{
		startUpSystems: []StartUpSystem{},
		systems:        []System{},
	}
}

func (s *UpdateStage) AddStartUpSystem(fn ...StartUpSystem) {
	s.startUpSystems = append(s.startUpSystems, fn...)
}

func (s *UpdateStage) AddSystem(system ...System) {
	s.systems = append(s.systems, system...)
}

func (s *UpdateStage) StartUpSystems() []StartUpSystem {
	return s.startUpSystems
}

func (s *UpdateStage) Systems() []System {
	return s.systems
}
