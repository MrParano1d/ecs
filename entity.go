package ecs

type Entity uint64

type EntityMap map[Entity]ComponentMap

func (em EntityMap) assertEntity(entity Entity) {
	if _, ok := em[entity]; !ok {
		em[entity] = ComponentMap{}
	}
}

func (em EntityMap) AddComponents(entity Entity, components ...Component) {
	em.assertEntity(entity)

	for _, c := range components {
		em[entity][c.Type()] = c
	}
}

func (em EntityMap) RemoveComponents(entity Entity, components ...Component) {
	em.assertEntity(entity)
	for _, c := range components {
		delete(em[entity], c.Type())
	}
}

func (em EntityMap) Components(entity Entity) map[ComponentType]Component {
	em.assertEntity(entity)
	return em[entity]
}

func (em EntityMap) EntitiesByComponentTypes(components ...Component) []Entity {
	var entities []Entity
	for e, c := range em {
		hasWithTypes := false
		for _, wt := range components {
			if _, ok := c[wt.Type()]; !ok {
				hasWithTypes = false
				break
			} else {
				hasWithTypes = true
			}
		}
		if hasWithTypes {
			entities = append(entities, e)
		}
	}
	return entities
}

func GetComponent[C Component](entities EntityMap, entity Entity) C {
	var component C

	if c, ok := entities[entity].Get(component.Type()); ok {
		return c.(C)
	}

	return component
}
