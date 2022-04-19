package ecs_test

import (
	"github.com/mrparano1d/ecs"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestResource struct {
	Name string
}

type WorldResource struct {
	Name string
}

var _ ecs.FromWorldResource = &WorldResource{}

func (res *WorldResource) FromWorld(_ *ecs.World) ecs.Resource {
	return &WorldResource{Name: "test"}
}

func TestGetResource(t *testing.T) {
	rm := ecs.ResourceMap{}
	testResource := &TestResource{Name: "test"}
	ecs.AddResource[*TestResource](rm, testResource)

	assert.Equal(t, testResource, ecs.GetResource[*TestResource](rm))
	assert.Equal(t, "test", ecs.GetResource[*TestResource](rm).Name)
}

func TestRemoveResource(t *testing.T) {
	rm := ecs.ResourceMap{}
	testResource := &TestResource{Name: "test"}
	ecs.AddResource[*TestResource](rm, testResource)

	assert.Equal(t, 1, len(rm))

	ecs.RemoveResource[*TestResource](rm)

	assert.Equal(t, 0, len(rm))
}

func TestInitResource(t *testing.T) {
	world := ecs.NewWorld(ecs.EventMap{})

	rm := ecs.ResourceMap{}
	if err := ecs.InitResource[*WorldResource](rm, world); err != nil {
		t.Fatal(err)
		return
	}

	res := ecs.GetResource[*WorldResource](rm)

	assert.Equal(t, "test", res.Name)
}
