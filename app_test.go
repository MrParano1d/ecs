package ecs_test

import (
	"github.com/mrparano1d/ecs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := ecs.NewApp()
	app.AddStartUpSystem(func(commands *ecs.Commands) {
		commands.Spawn().Insert(&NameComponent{Name: "test"})
	})
	app.AddSystem(func(ctx *ecs.SystemContext) {
		q := ecs.NewQuery[*NameComponent](ctx.World)
		for _, c := range q.Find(nil) {
			assert.Equal(t, "test", c.Name)
		}
	})
	app.AddSystem(func(ctx *ecs.SystemContext) {
		q := ecs.NewQuery[*NameComponent](ctx.World)
		assert.Equal(t, 0, len(q.Find(ecs.NewFilter(ecs.WithComponentFilter(&PositionComponent{})))))
		app.Cancel()
	})
	if err := app.Run(); err != nil {
		t.Fatalf("failed to start app: %v", err)
	}
}
