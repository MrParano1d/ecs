package core_test

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/ecs/core"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime_Delta(t *testing.T) {
	gameTime := core.NewTime()

	time.Sleep(1 * time.Microsecond)

	assert.Greater(t, gameTime.Delta(), float64(0))
}

func TestNewTime(t *testing.T) {
	loops := 0
	app := ecs.NewApp()
	app.AddPlugin(core.NewPlugin(core.EnvDebug))
	app.AddSystem(func(ctx ecs.SystemContext) {
		time.Sleep(1 * time.Millisecond)
	})
	app.AddSystem(func(ctx ecs.SystemContext) {
		gameTime := ecs.GetResource[*core.Time](ctx.Resources)
		assert.Greater(t, gameTime.Delta(), float64(0))
	})
	app.AddSystem(func(ctx ecs.SystemContext) {
		if loops == 2 {
			ctx.EventWriter(ecs.AppExitEvent{}).Send(ecs.AppExitEvent{})
		}
		loops++
	})
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
