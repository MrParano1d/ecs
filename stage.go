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

func (s *Stages) Add(stage Stage) {
	s.registry.Set(stage.Name(), stage)
	s.order.Insert(stage.Name())
}

func (s *Stages) AddAfter(afterStageName string, stage Stage) {
	s.registry.Set(stage.Name(), stage)
	s.order.InsertAfter(afterStageName, stage.Name())
}

func (s *Stages) AddBefore(beforeStageName string, stage Stage) {
	s.registry.Set(stage.Name(), stage)
	s.order.InsertBefore(beforeStageName, stage.Name())
}

func (s *Stages) GetStage(stageName string) Stage {
	if stage, ok := s.registry.Get(stageName); ok {
		return stage
	}
	panic(fmt.Errorf("unknown stage: %s", stageName))
}

func (s *Stages) GetOrderedStages(filters ...StageFilterOption) []Stage {
	stages := make([]Stage, len(s.order))

	filter := NewStageFilter(filters...)

	filterLabelsLen := len(filter.labels)

	for i, stageName := range s.order {
		if stage, ok := s.registry.Get(stageName); !ok {
			panic(
				fmt.Errorf(
					"mismatched stage order and registry length: registry=%d, order=%d", len(s.registry), len(s.order),
				),
			)
		} else {
			if filterLabelsLen > 0 {
				for _, l := range filter.labels {
					if l == stage.Label() {
						stages[i] = stage
						break
					}
				}
			} else {
				stages[i] = stage
			}
		}
	}

	// if a filter was applied the stage slice could contain nil elements, so we need to remove them
	if filter.HasFilters() {
		for i := 0; i < len(stages); {
			if stages[i] != nil {
				i++
				continue
			}

			if i < len(stages)-1 {
				copy(stages[i:], stages[i+1:])
			}

			stages[len(stages)-1] = nil
			stages = stages[:len(stages)-1]
		}
	}

	return stages
}

const (
	StageUpdate = "update"
	LabelNone   = "none"
	LabelUpdate = "update"
	LabelRender = "render"
)

const (
	ThreadingParallel = true
	ThreadingSingle   = false
)

type Stage interface {
	Name() string
	AddStartUpSystem(fn ...StartUpSystem)
	AddSystem(system ...System)
	StartUpSystems() []StartUpSystem
	Systems() []System
	Threading() bool
	SetLabel(label string)
	Label() string
}

type DefaultStage struct {
	startUpSystems []StartUpSystem
	systems        []System
	label          string
}

type StageOption func(stage Stage)

func WithStageLabel(label string) StageOption {
	return func(stage Stage) {
		stage.SetLabel(label)
	}
}

func NewDefaultStage(opts ...StageOption) *DefaultStage {
	d := &DefaultStage{
		startUpSystems: []StartUpSystem{},
		systems:        []System{},
		label:          LabelNone,
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func (s *DefaultStage) AddStartUpSystem(fn ...StartUpSystem) {
	s.startUpSystems = append(s.startUpSystems, fn...)
}

func (s *DefaultStage) AddSystem(system ...System) {
	s.systems = append(s.systems, system...)
}

func (s *DefaultStage) StartUpSystems() []StartUpSystem {
	return s.startUpSystems
}

func (s *DefaultStage) Systems() []System {
	return s.systems
}

func (s *DefaultStage) Threading() bool {
	return ThreadingSingle
}

func (s *DefaultStage) SetLabel(label string) {
	s.label = label
}

func (s *DefaultStage) Label() string {
	return s.label
}

func (s *DefaultStage) Name() string {
	panic("implement me")
}
