package ecs

import "reflect"

type ComponentType reflect.Type

type ComponentMap map[ComponentType]Component

func (cm ComponentMap) Set(component Component) {
	cm[component.Type()] = component
}

func (cm ComponentMap) Get(componentType ComponentType) (Component, bool) {
	c, ok := cm[componentType]
	return c, ok
}

type Component interface {
	Type() ComponentType
}
