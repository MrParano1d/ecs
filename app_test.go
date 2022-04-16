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

type TestEvent struct {
	Data string
}

func TestApp_AddEvent(t *testing.T) {
	app := ecs.NewApp()
	app.AddEvent(func(eventMap ecs.EventMap) {
		ecs.AddEvent[TestEvent](eventMap)
	})
	app.AddSystem(func(ctx *ecs.SystemContext) {
		ctx.EventWriter(TestEvent{}).Send(TestEvent{Data: "test0"})
	})
	app.AddSystem(func(ctx *ecs.SystemContext) {
		ctx.EventWriter(TestEvent{}).Send(TestEvent{Data: "test1"})
	})
	app.AddSystem(func(ctx *ecs.SystemContext) {
		i := 0
		reader := ctx.EventReader(TestEvent{})
		for reader.Next() {
			assert.Equal(t, fmt.Sprintf("test%d", i), reader.Read().(TestEvent).Data)
			i++
		}
		assert.Equal(t, 2, i)
	})
	app.AddSystem(func(ctx *ecs.SystemContext) {
		i := 0
		reader := ctx.EventReader(TestEvent{})
		for reader.Next() {
			assert.Equal(t, fmt.Sprintf("test%d", i), reader.Read().(TestEvent).Data)
			i++
		}
		assert.Equal(t, 2, i)
		app.Cancel()
	})

	if err := app.Run(); err != nil {
		t.Fatal(err)
		return
	}
}

type AppRenderStage struct {
	systems       []ecs.System
	startUpSystem []ecs.StartUpSystem
}

func NewAppRenderStage() *AppRenderStage {
	return &AppRenderStage{
		systems:       []ecs.System{},
		startUpSystem: []ecs.StartUpSystem{},
	}
}

func (a *AppRenderStage) AddStartUpSystem(fn ...ecs.StartUpSystem) {
	a.startUpSystem = append(a.startUpSystem, fn...)
}

func (a *AppRenderStage) AddSystem(system ...ecs.System) {
	a.systems = append(a.systems, system...)
}

func (a *AppRenderStage) StartUpSystems() []ecs.StartUpSystem {
	return a.startUpSystem
}

func (a *AppRenderStage) Systems() []ecs.System {
	return a.systems
}

func TestApp_AddSystemToStage(t *testing.T) {

	updateCalls := 0
	renderCalls := 0

	app := ecs.NewApp()
	app.AddStageAfter(
		ecs.StageUpdate, "render", NewAppRenderStage(),
	).AddSystem(func(ctx *ecs.SystemContext) {
		updateCalls++
	}).AddSystemToStage("render", func(ctx *ecs.SystemContext) {
		renderCalls++
		app.Cancel()
	})

	if err := app.Run(); err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, 1, updateCalls)
	assert.Equal(t, 1, renderCalls)
}
