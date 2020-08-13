package scheduler

import (
	"strconv"
	"strings"
	"time"
)

type schedule struct {
	interval uint
	task     func()
	lastTime time.Time
	nextTime time.Time
}

func (s *schedule) Do(task func()) {
	s.task = task
}

type runTime struct {
	hour   int
	minute int
	second int
}

// timeSchedule , unit type: 0 = second, 1 = minute, 2 = hour
type timeSchedule struct {
	schedule
	unit int
	from runTime
}

func (s *timeSchedule) From(timeStr string) {
	defer func() {
		if err := recover(); err != nil {
			s.from.hour = time.Now().Hour()
			s.from.minute = time.Now().Minute()
			s.from.second = time.Now().Second()
		}
	}()

	s.from = parseTime(timeStr)
}

type dateSchedule struct {
	schedule
	at runTime
}

func (s *dateSchedule) At(timeStr string) {
	defer func() {
		if err := recover(); err != nil {
			s.at.hour = time.Now().Hour()
			s.at.minute = time.Now().Minute()
			s.at.second = time.Now().Second()
		}
	}()

	s.at = parseTime(timeStr)
}

func parseTime(timeStr string) runTime {
	runTime := runTime{}

	if len(timeStr) <= 0 {
		panic("time format parse error")
	}

	timeSplit := strings.Split(timeStr, ":")
	for i := 0; i < 3; i++ {
		if num, err := strconv.ParseInt(timeSplit[i], 10, 64); err == nil {
			if i == 0 {
				runTime.hour = int(num)

			} else if i == 1 {
				runTime.minute = int(num)

			} else if i == 2 {
				runTime.second = int(num)
			}
		} else {
			panic("time format parse error")
		}
	}

	return runTime
}
