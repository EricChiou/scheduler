package scheduler

import (
	"time"
)

type schedule interface {
	Do(task func()) int
	start()
	stop()
	calNextRunTime()
	runSchedule()
}

type runTime struct {
	hour   int
	minute int
	second int
}

type baseSchedule struct {
	task     func()
	typ      int
	lastTime time.Time
	nextTime time.Time
	runTime  runTime
	run      bool
}

func (s *baseSchedule) start() {
	s.run = true
}

func (s *baseSchedule) stop() {
	s.run = false
}
