package ecs_test

import (
	"github.com/mrparano1d/ecs"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ContextTestEvent struct {
	Data string
}

func TestNewSystemContext(t *testing.T) {
	w := ecs.NewWorld()
	q := ecs.NewQueue()
	ev := ecs.EventMap{}

	ecs.AddEvent[ContextTestEvent](ev)

	ctx := ecs.NewSystemContext(w, ecs.NewCommands(&q, w), ev)
	writer := ctx.EventWriter(ContextTestEvent{})
	writer.Send(ContextTestEvent{Data: "test"})

	reader := ctx.EventReader(ContextTestEvent{})

	for reader.Next() {
		assert.Equal(t, "test", reader.Read().(ContextTestEvent).Data)
	}
}
