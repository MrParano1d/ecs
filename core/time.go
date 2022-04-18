package core

import "time"

type Time struct {
	startTime time.Time
	lastFrame time.Time
}

func NewTime() *Time {
	return &Time{
		startTime: time.Now(),
		lastFrame: time.Now(),
	}
}

func (t *Time) Startup() time.Time {
	return t.startTime
}

func (t *Time) TimeSinceStartUp() time.Duration {
	return time.Now().Sub(t.Startup())
}

func (t *Time) SecondsSinceStart() float64 {
	return time.Now().Sub(t.startTime).Seconds()
}

func (t *Time) Delta() float64 {
	delta := time.Now().Sub(t.lastFrame).Seconds()
	if delta == 0 {
		delta = 0.00001
	}
	return delta
}

func (t *Time) Update() {
	t.lastFrame = time.Now()
}
