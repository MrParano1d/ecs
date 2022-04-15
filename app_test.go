package ecs_test

import (
	"fmt"
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

type HelloPlugin struct {
	T *testing.T
}

var _ ecs.Plugin = &HelloPlugin{}

func (p *HelloPlugin) Build(app *ecs.App) {
	app.AddStartUpSystem(func(commands *ecs.Commands) {
		commands.Spawn().Insert(&NameComponent{Name: "world"})
	})
	app.AddSystem(func(ctx *ecs.SystemContext) {
		q := ecs.NewQuery[*NameComponent](ctx.World)
		for _, c := range q.Find(nil) {
			assert.Equal(p.T, "hello, world", fmt.Sprintf("hello, %s", c.Name))
		}
		app.Cancel()
	})
}

func TestApp_AddPlugin(t *testing.T) {
	app := ecs.NewApp()
	app.AddPlugin(&HelloPlugin{T: t})
	if err := app.Run(); err != nil {
		t.Fatalf("failed to start app: %v", err)
	}
}
