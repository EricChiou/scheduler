package scheduler

import "time"

type scheduler struct {
	interval uint
}

func (s *scheduler) Second() timeSchedule {
	return initTimeSchedule(s.interval, 0)
}

func (s *scheduler) Minute() timeSchedule {
	return initTimeSchedule(s.interval, 1)
}

func (s *scheduler) Hour() timeSchedule {
	return initTimeSchedule(s.interval, 2)
}

func (s *scheduler) Day() dateSchedule {
	var dateSchedule dateSchedule
	dateSchedule.interval = s.interval

	return dateSchedule
}

func (s *scheduler) Week() dateSchedule {
	var dateSchedule dateSchedule

	return dateSchedule
}

func (s *scheduler) Month() dateSchedule {
	var dateSchedule dateSchedule

	return dateSchedule
}

func Every(interval uint) scheduler {
	return scheduler{interval: interval}
}

func initTimeSchedule(interval uint, unit int) timeSchedule {
	var timeSchedule timeSchedule
	timeSchedule.interval = interval
	timeSchedule.unit = unit
	timeSchedule.from = newRunTime()

	return timeSchedule
}

func initDateSchedule(interval uint) dateSchedule {
	var dateSchedule dateSchedule
	dateSchedule.interval = interval
	dateSchedule.at = newRunTime()

	return dateSchedule
}

func newRunTime() runTime {
	return runTime{
		hour:   time.Now().Hour(),
		minute: time.Now().Minute(),
		second: time.Now().Second(),
	}
}
