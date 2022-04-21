package ecs_test

import (
	"github.com/mrparano1d/ecs"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type NameComponent struct {
	Name string
}

var _ ecs.Component = &NameComponent{}

func (nc *NameComponent) Type() ecs.ComponentType {
	return reflect.TypeOf(nc)
}

type PositionComponent struct {
	X, Y, Z float32
}

var _ ecs.Component = &PositionComponent{}

func (pc *PositionComponent) Type() ecs.ComponentType {
	return reflect.TypeOf(pc)
}

func TestNewQuery(t *testing.T) {
	world := ecs.NewWorld(ecs.EventMap{})
	world.Entities().AddComponents(world.NextEntity(), &NameComponent{Name: "test"})
	query := ecs.NewQuery(world)
	assert.Equal(t, 1, len(query.Find(ecs.NewFilter(ecs.WithComponentFilter(&NameComponent{})))))
	assert.Equal(t, 0, len(query.Find(ecs.NewFilter(ecs.WithComponentFilter(&PositionComponent{})))))

	world.Entities().AddComponents(world.NextEntity(), &NameComponent{Name: "test"}, &PositionComponent{0, 0, 0})
	query = ecs.NewQuery(world)
	assert.Equal(t, 1, len(query.Find(ecs.NewFilter(ecs.WithComponentFilter(&NameComponent{}, &PositionComponent{})))))
}

func BenchmarkQuery_Find(b *testing.B) {
	world := ecs.NewWorld(ecs.EventMap{})
	world.Entities().AddComponents(world.NextEntity(), &NameComponent{Name: "test"})
	for n := 0; n < b.N; n++ {
		query := ecs.NewQuery(world)
		query.Find(ecs.NewFilter(ecs.WithComponentFilter(&NameComponent{}, &PositionComponent{})))
	}
}
