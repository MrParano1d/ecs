package ecs

import "sync"

type EventMap map[any]*EventWriter

type EventInvoker func(eventMap EventMap)

func AddEvent[E any](em EventMap) {
	var eventType E
	em[eventType] = NewEventWriter()
}

type Event interface {
}

type EventWriter struct {
	mutex *sync.Mutex
	queue []Event
}

func NewEventWriter() *EventWriter {
	return &EventWriter{
		mutex: &sync.Mutex{},
		queue: []Event{},
	}
}

func (w *EventWriter) Flush() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.queue = []Event{}
}

func (w *EventWriter) Send(payload Event) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.queue = append(w.queue, payload)
}

func (w *EventWriter) Queue() []Event {
	return w.queue
}

type EventReader struct {
	mutex  *sync.Mutex
	writer *EventWriter
	index  int
}

func NewEventReader(writer *EventWriter) *EventReader {
	return &EventReader{
		mutex:  &sync.Mutex{},
		writer: writer,
		index:  0,
	}
}

func (r *EventReader) Next() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.index < len(r.writer.queue)
}

func (r *EventReader) Read() Event {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	defer func() {
		r.index++
	}()
	return r.writer.Queue()[r.index]
}
