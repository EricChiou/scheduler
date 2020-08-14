package scheduler

import (
	"time"
)

// timeSchedule , typ: 0 = second, 1 = minute, 2 = hour
type timeSchedule struct {
	baseSchedule
	interval int
}

func (s *timeSchedule) From(timeStr string) *timeSchedule {
	defer func() {
		if err := recover(); err != nil {
			s.runTime.hour = 0
			s.runTime.minute = 0
			s.runTime.second = 0
		}
	}()

	s.runTime = parseTime(timeStr)
	return s
}

func (s *timeSchedule) Do(task func()) int {
	s.task = task
	s.run = true
	scheduleList = append(scheduleList, s)

	go s.runSchedule()

	return len(scheduleList) - 1
}

func (s *timeSchedule) calNextRunTime() {
	now := time.Now()
	fromTime := time.Date(now.Year(), now.Month(), now.Day(), s.runTime.hour, s.runTime.minute, s.runTime.second, 0, time.Local)

	if fromTime.Unix() > now.Unix() {
		s.nextTime = fromTime

	} else {
		// typ: 0 = second, 1 = minute, 2 = hour
		if s.typ == 0 {
			s.nextTime = time.Now().Add(time.Duration(s.interval) * time.Second)

		} else if s.typ == 1 {
			s.nextTime = time.Now().Add(time.Duration(s.interval*60) * time.Second)

		} else if s.typ == 2 {
			s.nextTime = time.Now().Add(time.Duration(s.interval*3600) * time.Second)
		}
	}
}

func (s *timeSchedule) runSchedule() {
	s.calNextRunTime()
	sleep := s.nextTime.Unix() - time.Now().Unix()
	time.Sleep(time.Duration(sleep) * time.Second)

	if s.run {
		s.lastTime = time.Now()
		s.task()
	}
	s.runSchedule()
}
