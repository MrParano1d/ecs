package ecs_test

import (
	"fmt"
	"github.com/mrparano1d/ecs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := ecs.NewApp()
	app.AddPlugin(NewTestPlugin())
	app.AddStartUpSystem(func(commands ecs.Commands) {
		commands.Spawn().Insert(&NameComponent{Name: "test"})
	})
	app.AddSystem(func(ctx ecs.SystemContext) {
		q := ecs.NewQuery(ctx.World)
		for _, e := range q.Find(ecs.NewFilter(ecs.WithComponentFilter(&NameComponent{}))) {
			assert.Equal(t, "test", ecs.GetComponent[*NameComponent](ctx.World.Entities(), e).Name)
		}
	})
	app.AddSystem(func(ctx ecs.SystemContext) {
		q := ecs.NewQuery(ctx.World)
		assert.Equal(t, 0, len(q.Find(ecs.NewFilter(ecs.WithComponentFilter(&NameComponent{}, &PositionComponent{})))))
		ctx.EventWriter(ecs.AppExitEvent{}).Send(ecs.AppExitEvent{})
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
	app.AddStartUpSystem(func(commands ecs.Commands) {
		commands.Spawn().Insert(&NameComponent{Name: "world"})
	})
	app.AddSystem(func(ctx ecs.SystemContext) {
		q := ecs.NewQuery(ctx.World)
		for _, e := range q.Find(ecs.NewFilter(ecs.WithComponentFilter(&NameComponent{}))) {
			assert.Equal(p.T, "hello, world", fmt.Sprintf("hello, %s", ecs.GetComponent[*NameComponent](ctx.World.Entities(), e).Name))
		}
		ctx.EventWriter(ecs.AppExitEvent{}).Send(ecs.AppExitEvent{})
	})
}

func TestApp_AddPlugin(t *testing.T) {
	app := ecs.NewApp()
	app.AddPlugin(NewTestPlugin())
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
	app.AddPlugin(NewTestPlugin())
	app.AddEvent(func(eventMap ecs.EventMap) {
		ecs.AddEvent[TestEvent](eventMap)
	})
	app.AddSystem(func(ctx ecs.SystemContext) {
		ctx.EventWriter(TestEvent{}).Send(TestEvent{Data: "test0"})
	})
	app.AddSystem(func(ctx ecs.SystemContext) {
		ctx.EventWriter(TestEvent{}).Send(TestEvent{Data: "test1"})
	})
	app.AddSystem(func(ctx ecs.SystemContext) {
		i := 0
		reader := ctx.EventReader(TestEvent{})
		for reader.Next() {
			assert.Equal(t, fmt.Sprintf("test%d", i), reader.Read().(TestEvent).Data)
			i++
		}
		assert.Equal(t, 2, i)
	})
	app.AddSystem(func(ctx ecs.SystemContext) {
		i := 0
		reader := ctx.EventReader(TestEvent{})
		for reader.Next() {
			assert.Equal(t, fmt.Sprintf("test%d", i), reader.Read().(TestEvent).Data)
			i++
		}
		assert.Equal(t, 2, i)

		ctx.EventWriter(ecs.AppExitEvent{}).Send(ecs.AppExitEvent{})
	})

	if err := app.Run(); err != nil {
		t.Fatal(err)
		return
	}
}

type AppRenderStage struct {
	ecs.Stage
}

func NewAppRenderStage() *AppRenderStage {
	return &AppRenderStage{
		Stage: ecs.NewDefaultStage(),
	}
}

func (s AppRenderStage) Name() string {
	return "render"
}

func TestApp_AddSystemToStage(t *testing.T) {

	updateCalls := 0
	renderCalls := 0

	app := ecs.NewApp()
	app.AddPlugin(NewTestPlugin())
	app.AddStageAfter(
		ecs.StageUpdate, NewAppRenderStage(),
	).AddSystem(func(ctx ecs.SystemContext) {
		updateCalls++
	}).AddSystemToStage("render", func(ctx ecs.SystemContext) {
		renderCalls++
		ctx.EventWriter(ecs.AppExitEvent{}).Send(ecs.AppExitEvent{})
	})

	if err := app.Run(); err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, 1, updateCalls)
	assert.Equal(t, 1, renderCalls)
}
