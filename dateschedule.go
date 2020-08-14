package scheduler

import (
	"time"
)

// dateSchedule typ: 0 = Day, 1 = Week, 2 = Month
type dateSchedule struct {
	baseSchedule
	day int
}

func (s *dateSchedule) At(timeStr string) *dateSchedule {
	defer func() {
		if err := recover(); err != nil {
			s.runTime.hour = time.Now().Hour()
			s.runTime.minute = time.Now().Minute()
			s.runTime.second = time.Now().Second()
		}
	}()

	s.runTime = parseTime(timeStr)
	return s
}

func (s *dateSchedule) Do(task func()) int {
	s.task = task
	s.run = true
	scheduleList = append(scheduleList, s)

	go s.runSchedule()

	return len(scheduleList) - 1
}

func (s *dateSchedule) calNextRunTime() {
	now := time.Now()

	// typ: 0 = Day, 1 = Week, 2 = Month
	if s.typ == 0 {
		atTime := time.Date(now.Year(), now.Month(), now.Day(), s.runTime.hour, s.runTime.minute, s.runTime.second, 0, time.Local)
		if atTime.Unix() <= now.Unix() {
			s.nextTime = time.Date(now.Year(), now.Month(), now.Day()+s.day, s.runTime.hour, s.runTime.minute, s.runTime.second, 0, time.Local)
		} else {
			s.nextTime = atTime
		}

	} else if s.typ == 1 {
		for i := 0; i < 8; i++ {
			atTime := time.Date(now.Year(), now.Month(), now.Day()+i, s.runTime.hour, s.runTime.minute, s.runTime.second, 0, time.Local)
			if s.day == int(atTime.Weekday()) {
				if atTime.Unix() > now.Unix() {
					s.nextTime = atTime
					return
				}
			}
		}

	} else if s.typ == 2 {
		atTime := time.Date(now.Year(), now.Month(), s.day, s.runTime.hour, s.runTime.minute, s.runTime.second, 0, time.Local)
		if atTime.Unix() <= now.Unix() {
			s.nextTime = time.Date(now.Year(), now.Month()+1, s.day, s.runTime.hour, s.runTime.minute, s.runTime.second, 0, time.Local)
		} else {
			s.nextTime = atTime
		}
	}
}

func (s *dateSchedule) runSchedule() {
	s.calNextRunTime()
	sleep := s.nextTime.Unix() - time.Now().Unix()
	time.Sleep(time.Duration(sleep) * time.Second)

	if s.run {
		s.lastTime = time.Now()
		s.task()
	}
	s.runSchedule()
}
