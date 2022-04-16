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

func (em EntityMap) ComponentsByType(ct Component, withTypes ...Component) []Component {
	var components []Component
	withTypesLen := len(withTypes)
	for _, e := range em {
		if c, ok := e[ct.Type()]; !ok {
			continue
		} else {
			hasWithTypes := withTypesLen == 0
			for _, wt := range withTypes {
				if _, ok := e[wt.Type()]; !ok {
					hasWithTypes = false
					break
				} else {
					hasWithTypes = true
				}
			}
			if hasWithTypes {
				components = append(components, c)
			}
		}
	}
	return components
}
