package ecs

import "sync"

type Scheduler struct {
	stage Stage
	wg    sync.WaitGroup
}

func NewScheduler(stage Stage) *Scheduler {
	return &Scheduler{
		stage: stage,
	}
}

func (s *Scheduler) RunSystems(world *World, events EventMap) {
	if s.stage.Threading() == ThreadingParallel {
		s.wg.Add(len(s.stage.Systems()))
		queue := NewQueue()
		for _, system := range s.stage.Systems() {
			go func(system System) {
				system(NewSystemContext(world, NewCommands(queue, world), events))
				s.wg.Done()
			}(system)
		}
		queue.Apply(world)
		s.wg.Wait()
	} else {
		queue := NewQueue()
		for _, system := range s.stage.Systems() {
			system(NewSystemContext(world, NewCommands(queue, world), events))
		}
		queue.Apply(world)
	}

}
