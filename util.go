package scheduler

import (
	"strconv"
	"strings"
)

func parseTime(timeStr string) runTime {
	runTime := runTime{}

	if len(timeStr) <= 0 {
		panic("time format parse error")
	}

	timeSplit := strings.Split(timeStr, ":")
	if len(timeSplit) > 3 {
		panic("time format parse error")
	}

	for i := 0; i < 3; i++ {
		if num, err := strconv.ParseInt(timeSplit[i], 10, 64); err == nil {
			if i == 0 {
				if int(num) < 0 || 23 < int(num) {
					panic("time format parse error")
				}
				runTime.hour = int(num)

			} else if i == 1 {
				if int(num) < 0 || 59 < int(num) {
					panic("time format parse error")
				}
				runTime.minute = int(num)

			} else if i == 2 {
				if int(num) < 0 || 59 < int(num) {
					panic("time format parse error")
				}
				runTime.second = int(num)
			}
		} else {
			panic("time format parse error")
		}
	}

	return runTime
}
