package ecs_test

import (
	"github.com/mrparano1d/ecs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime_Delta(t *testing.T) {
	gameTime := ecs.NewTime()

	time.Sleep(1 * time.Microsecond)

	assert.Greater(t, gameTime.Delta(), float64(0))
}

func TestNewTime(t *testing.T) {
	loops := 0
	app := ecs.NewApp()
	app.AddSystem(func(ctx *ecs.SystemContext) {
		time.Sleep(1 * time.Microsecond)
	})
	app.AddSystem(func(ctx *ecs.SystemContext) {
		assert.Greater(t, ctx.Time().Delta(), float64(0))
	})
	app.AddSystem(func(ctx *ecs.SystemContext) {
		if loops == 2 {
			app.Cancel()
		}
		loops++
	})
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
