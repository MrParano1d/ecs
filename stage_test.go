package ecs_test

import (
	"github.com/mrparano1d/ecs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStageOrder_InsertAfter(t *testing.T) {

	emptyStageOrder := ecs.StageOrder{}

	emptyStageOrder.InsertAfter("update", "render")

	assert.Equal(t, ecs.StageOrder{"render"}, emptyStageOrder)

	singleStageOrder := ecs.StageOrder{"update"}

	singleStageOrder.InsertAfter("update", "render")

	assert.Equal(t, ecs.StageOrder{"update", "render"}, singleStageOrder)

	multipleStageOrder := ecs.StageOrder{"update", "render"}
	multipleStageOrder.InsertAfter("update", "debug_update")
	multipleStageOrder.InsertAfter("render", "debug_render")

	assert.Equal(t, ecs.StageOrder{"update", "debug_update", "render", "debug_render"}, multipleStageOrder)

	duplicateStageOrder := ecs.StageOrder{"update"}
	duplicateStageOrder.InsertAfter("update", "update")
	assert.Equal(t, ecs.StageOrder{"update"}, duplicateStageOrder)
}

func TestStageOrder_InsertBefore(t *testing.T) {
	emptyStageOrder := ecs.StageOrder{}

	emptyStageOrder.InsertBefore("update", "render")

	assert.Equal(t, ecs.StageOrder{"render"}, emptyStageOrder)

	singleStageOrder := ecs.StageOrder{"update"}

	singleStageOrder.InsertBefore("update", "render")

	assert.Equal(t, ecs.StageOrder{"render", "update"}, singleStageOrder)

	multipleStageOrder := ecs.StageOrder{"update", "render"}
	multipleStageOrder.InsertBefore("update", "debug_update")
	multipleStageOrder.InsertBefore("render", "debug_render")

	assert.Equal(t, ecs.StageOrder{"debug_update", "update", "debug_render", "render"}, multipleStageOrder)

	duplicateStageOrder := ecs.StageOrder{"update"}
	duplicateStageOrder.InsertBefore("update", "update")
	assert.Equal(t, ecs.StageOrder{"update"}, duplicateStageOrder)
}

type RenderStage struct {
	ecs.Stage
}

func (s *RenderStage) Name() string {
	return "render"
}

type DebugStage struct {
	ecs.Stage
}

func (s *DebugStage) Name() string {
	return "debug"
}

func TestNewStages(t *testing.T) {
	stages := ecs.NewStages()

	renderStage := &RenderStage{}
	stages.Add(renderStage)
	updateStage := &UpdateStage{}
	stages.AddBefore("render", updateStage)
	debugStage := &DebugStage{}
	stages.AddAfter("render", debugStage)

	expectedStageOrder := []ecs.Stage{updateStage, renderStage, debugStage}
	assert.Equal(t, expectedStageOrder, stages.GetOrderedStages())
}

type TestLabelStage struct {
	ecs.Stage
}

func NewTestLabelStage() *TestLabelStage {
	return &TestLabelStage{
		Stage: ecs.NewDefaultStage(ecs.WithStageLabel(ecs.LabelRender)),
	}
}

func (s *TestLabelStage) Name() string {
	return "test"
}

func TestWithStageLabel(t *testing.T) {

	s := NewTestLabelStage()

	stages := ecs.NewStages()
	stages.Add(s)

	assert.Equal(t, 1, len(stages.GetOrderedStages()))
	assert.Equal(t, 1, len(stages.GetOrderedStages(ecs.WithStageLabelFilter(ecs.LabelRender))))
	assert.Equal(t, 0, len(stages.GetOrderedStages(ecs.WithStageLabelFilter(ecs.LabelUpdate))))
	assert.Equal(t, 1, len(stages.GetOrderedStages(ecs.WithStageLabelFilter(ecs.LabelUpdate, ecs.LabelRender))))
}
