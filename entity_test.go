package ecs_test

import (
	"github.com/mrparano1d/ecs"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type TestComponent struct {
	a int
}

var tcComponentType = reflect.TypeOf(&TestComponent{})

func (tc *TestComponent) Type() ecs.ComponentType {
	return tcComponentType
}

var _ ecs.Component = &TestComponent{}

type TestComponent2 struct {
	b int
}

var tc2ComponentType = reflect.TypeOf(&TestComponent2{})

func (tc *TestComponent2) Type() ecs.ComponentType {
	return tc2ComponentType
}

var _ ecs.Component = &TestComponent2{}

type TestComponent3 struct {
	b int
}

var tc3ComponentType = reflect.TypeOf(&TestComponent3{})

func (tc *TestComponent3) Type() ecs.ComponentType {
	return tc3ComponentType
}

var _ ecs.Component = &TestComponent3{}

func TestEntityMap_AddComponents(t *testing.T) {
	em := ecs.EntityMap{}
	em.AddComponents(1, &TestComponent{a: 0})
	em.AddComponents(1, &TestComponent{a: 0})

	assert.Equal(t, 1, len(em.Components(1)), "entity 1 should contain 1 component")
}

func TestEntityMap_ComponentsByType(t *testing.T) {
	em := ecs.EntityMap{}
	em.AddComponents(1, &TestComponent{a: 0}, &TestComponent2{b: 0})

	assert.Equal(t, 1, len(em.ComponentsByType(&TestComponent{})))
	assert.Equal(t, 1, len(em.ComponentsByType(&TestComponent{}, &TestComponent2{})))
	assert.Equal(t, 0, len(em.ComponentsByType(&TestComponent{}, &TestComponent2{}, &TestComponent3{})))
}
