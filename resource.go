package ecs

import "fmt"

type ResourceMap map[any]Resource

func InitResource[R any](rm ResourceMap, world *World) error {
	var resourceType R
	if r, ok := any(resourceType).(FromWorldResource); ok {
		AddResource[R](rm, r.FromWorld(world))
		return nil
	}
	return fmt.Errorf("resouce doesn't implement FromWorldResource, instead: %T", resourceType)
}

func AddResource[R any](rm ResourceMap, r Resource) {
	var resourceType R
	rm[resourceType] = r
}

func GetResource[R any](rm ResourceMap) R {
	var resourceType R
	if r, ok := rm[resourceType]; ok {
		return r.(R)
	}
	return resourceType
}

func RemoveResource[R any](rm ResourceMap) {
	var resourceType R
	delete(rm, resourceType)
}

type Resource interface {
}

type FromWorldResource interface {
	FromWorld(world *World) Resource
}
