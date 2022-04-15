package ecs

type Queue struct {
	commands []Command
}

func NewQueue() *Queue {
	return &Queue{
		commands: []Command{},
	}
}

func (q *Queue) Push(c Command) {
	q.commands = append(q.commands, c)
}

func (q *Queue) Apply(w *World) {
	for _, c := range q.commands {
		c.Write(w)
	}
}
