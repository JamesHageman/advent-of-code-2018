package day4

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

func scanTime(line string) (time.Time, error) {
	var year, month, day, hour, minute int
	_, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d", &year, &month, &day, &hour, &minute)
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC), err
}

type logType int

const (
	beginShift logType = iota
	fallAsleep
	wakeUp
)

type log struct {
	t       time.Time
	lt      logType
	guardID uint
}

func scanLogs(in io.Reader) []log {
	lineScanner := bufio.NewScanner(in)

	logs := make([]log, 0)

	for lineScanner.Scan() {
		var l log
		var guardID uint

		line := strings.Split(lineScanner.Text(), "] ")
		timeStr, log := line[0], line[1]

		scannedTime, err := scanTime(timeStr)
		if err != nil {
			break
		}

		l.t = scannedTime

		_, guardErr := fmt.Sscanf(log, "Guard #%d begins shift", &guardID)
		_, sleepErr := fmt.Sscanf(log, "falls asleep")
		_, wakeErr := fmt.Sscanf(log, "wakes up")
		switch {
		case guardErr == nil:
			l.lt = beginShift
			l.guardID = guardID
		case sleepErr == nil:
			l.lt = fallAsleep
		case wakeErr == nil:
			l.lt = wakeUp
		}
		logs = append(logs, l)
	}

	return logs
}

func buildSchedule(logs []log) (schedule map[uint][][2]time.Time) {
	schedule = make(map[uint][][2]time.Time)

	sortLogs(logs)

	var sleepTime, wakeTime time.Time
	var guardID uint

	for _, l := range logs {
		switch l.lt {
		case beginShift:
			guardID = l.guardID
		case fallAsleep:
			sleepTime = l.t
		case wakeUp:
			wakeTime = l.t
			times := schedule[guardID]
			times = append(times, [2]time.Time{sleepTime, wakeTime})
			schedule[guardID] = times
		}
	}

	return schedule
}

func sortLogs(logs []log) {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].t.Unix() < logs[j].t.Unix()
	})
}

func getSleepTime(schedule map[uint][][2]time.Time) (sleepTimes map[uint]int) {
	sleepTimes = make(map[uint]int)
	for guardID, times := range schedule {
		for _, nap := range times {
			sleepTimes[guardID] += (nap[1].Minute() - nap[0].Minute())
		}
	}
	return sleepTimes
}

func calcSleepMinutes(sleeps [][2]time.Time) (minutes [60]uint) {
	for _, nap := range sleeps {
		for minute := nap[0].Minute(); minute < nap[1].Minute(); minute++ {
			minutes[minute]++
		}
	}
	return minutes
}

func calcMostPopularMinute(minutes [60]uint) int {
	maxMinute := 0

	for minute, count := range minutes {
		if count > minutes[maxMinute] {
			maxMinute = minute
		}
	}
	return maxMinute
}

func Part1(in io.Reader) uint {
	logs := scanLogs(in)
	schedule := buildSchedule(logs)
	sleepTimes := getSleepTime(schedule)

	var sleepiestGuardID uint
	var maxSleepTime int
	for guardID, sleepTime := range sleepTimes {
		if sleepTime > maxSleepTime {
			sleepiestGuardID = guardID
			maxSleepTime = sleepTime
		}
	}

	return sleepiestGuardID *
		uint(calcMostPopularMinute(calcSleepMinutes(schedule[sleepiestGuardID])))
}

func Part2(in io.Reader) uint {
	logs := scanLogs(in)
	schedule := buildSchedule(logs)

	var maxOnMinute uint
	var maxPerMinute uint
	var maxGuardID uint

	for guardID, guardSchedule := range schedule {
		minutes := calcSleepMinutes(guardSchedule)
		maxMinute := calcMostPopularMinute(minutes)
		if minutes[maxMinute] > maxPerMinute {
			maxOnMinute = uint(maxMinute)
			maxPerMinute = minutes[maxMinute]
			maxGuardID = guardID
		}
	}

	return maxGuardID * maxOnMinute
}
