package ecs_test

import (
	"github.com/mrparano1d/ecs"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestOnceEvent struct {
	Data int
}

func TestEventWriter_SendOnce(t *testing.T) {
	counter := 0
	events := ecs.EventMap{}

	ecs.AddEvent[TestOnceEvent](events)

	events[TestOnceEvent{}].SendOnce(TestOnceEvent{})
	events[TestOnceEvent{}].SendOnce(TestOnceEvent{})

	reader := ecs.NewEventReader(events[TestOnceEvent{}])
	for reader.Next() {
		_ = reader.Read()
		counter++
	}

	assert.Equal(t, 1, counter)

	events[TestOnceEvent{}].Flush()

	counter = 0

	events[TestOnceEvent{}].SendOnce(TestOnceEvent{Data: 1})
	events[TestOnceEvent{}].SendOnce(TestOnceEvent{Data: 2})

	reader = ecs.NewEventReader(events[TestOnceEvent{}])
	for reader.Next() {
		_ = reader.Read()
		counter++
	}

	assert.Equal(t, 2, counter)
}
