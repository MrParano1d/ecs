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

type UpdateStage struct {
	ecs.Stage
}

type RenderStage struct {
	ecs.Stage
}

type DebugStage struct {
	ecs.Stage
}

func TestNewStages(t *testing.T) {
	stages := ecs.NewStages()

	renderStage := &RenderStage{}
	stages.Add("render", renderStage)
	updateStage := &UpdateStage{}
	stages.AddBefore("render", "update", updateStage)
	debugStage := &DebugStage{}
	stages.AddAfter("render", "debug", debugStage)

	expectedStageOrder := []ecs.Stage{updateStage, renderStage, debugStage}
	assert.Equal(t, expectedStageOrder, stages.GetOrderedStages())
}
