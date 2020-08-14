package scheduler

import "time"

var scheduleList []schedule

func Every() *scheduler {
	return &scheduler{}
}

func Start(index int) {
	if index > 0 && len(scheduleList) > index && scheduleList[index] != nil {
		scheduleList[index].start()
	}
}

func Stop(index int) {
	if index > 0 && len(scheduleList) > index && scheduleList[index] != nil {
		scheduleList[index].stop()
	}
}

type scheduler struct {
}

func (s *scheduler) Second(interval int) *timeSchedule {
	return initTimeSchedule(interval, 0)
}

func (s *scheduler) Minute(interval int) *timeSchedule {
	return initTimeSchedule(interval, 1)
}

func (s *scheduler) Hour(interval int) *timeSchedule {
	return initTimeSchedule(interval, 2)
}

func (s *scheduler) Day(day int) *dateSchedule {
	return initDateSchedule(day, 0)
}

func (s *scheduler) Week(day int) *dateSchedule {
	if day < 0 || 6 < day {
		day = int(time.Now().Weekday())
	}
	return initDateSchedule(day, 1)
}

func (s *scheduler) Month(day int) *dateSchedule {
	if day < 1 || 31 < day {
		day = time.Now().Day()
	}
	return initDateSchedule(day, 2)
}

func initTimeSchedule(interval int, typ int) *timeSchedule {
	if interval < 1 {
		interval = 1
	}
	var timeSchedule timeSchedule
	timeSchedule.interval = interval
	timeSchedule.typ = typ
	timeSchedule.runTime = runTime{}

	return &timeSchedule
}

func initDateSchedule(day int, typ int) *dateSchedule {
	var dateSchedule dateSchedule
	dateSchedule.typ = typ
	dateSchedule.day = day
	dateSchedule.runTime = runTime{
		hour:   time.Now().Hour(),
		minute: time.Now().Minute(),
		second: time.Now().Second(),
	}

	return &dateSchedule
}
