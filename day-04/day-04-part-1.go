package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	dateTime time.Time
	action string
}

func parseEntry(line string) Entry {
	dateRegex := regexp.MustCompile(`\d*-\d*-\d*`)
	dateStr := dateRegex.FindString(line)

	timeRegex := regexp.MustCompile(`\d*:\d*`)
	timeStr := timeRegex.FindString(line)

	dateTime, _ := time.Parse("2006-01-02T15:04", dateStr + "T" + timeStr)

	actionRegex := regexp.MustCompile(`(?:\]\s)(.*)`)
	action := actionRegex.FindStringSubmatch(line)[1]

	return Entry{
		dateTime: dateTime,
		action: action,
	}
}

func sortEntries(entries []Entry) {
	sort.Slice(entries, func(i, j int) bool {
		dateTimeOne := entries[i].dateTime
		dateTimeTwo := entries[j].dateTime

		if dateTimeOne.Equal(dateTimeTwo) && dateTimeOne.Hour() == dateTimeTwo.Hour() {
			return dateTimeOne.Minute() < dateTimeTwo.Minute()
		} else if dateTimeOne.Equal(dateTimeTwo) {
			return dateTimeOne.Hour() < dateTimeTwo.Hour()
		} else {
			return dateTimeOne.Before(dateTimeTwo)
		}
	})
}

func findGuardId(action string) int {
	idRegex := regexp.MustCompile(`(?:#)(\d*)`)
	guardIdStr := idRegex.FindStringSubmatch(action)[1]
	guardId, _ := strconv.Atoi(guardIdStr)
	return guardId
}

func findSleepiestGuardId(entries []Entry) int {
	minutesSlept := make(map[int]int)
	
	var currentGuardId int
	
	for i, entry := range entries {
		action := entry.action
		
		if strings.Contains(action, "begins") {
			currentGuardId = findGuardId(action)
		} else if strings.Contains(action, "asleep") {
			sleepTime := entry.dateTime
			wakeTime := entries[i+1].dateTime
			elapsedMinutes := int(wakeTime.Sub(sleepTime).Minutes())

			minutesSlept[currentGuardId] += elapsedMinutes
		}
	}

	sleepiestGuardId := -1

	for k, v := range minutesSlept {
		if sleepiestGuardId == -1 || v > minutesSlept[sleepiestGuardId] {
			sleepiestGuardId = k
		}
	}

	return sleepiestGuardId
}

func recordSleepSession(sleepLog map[int]int, sleepTime time.Time, wakeTime time.Time) {
	elapsedMinutes := int(wakeTime.Sub(sleepTime).Minutes())

	for i := 0; i < elapsedMinutes; i++ {
		sleepMinute := sleepTime.Add(time.Minute * time.Duration(i))
		sleepLog[sleepMinute.Minute()] += 1
	}
}

func findSleepiestMinute(entries []Entry, longestSleeperId int) int {
	var sleepLog = make(map[int]int)
	
	var currentGuardId int
	
	for i, entry := range entries {
		action := entry.action
		
		if strings.Contains(action, "begins") {
			currentGuardId = findGuardId(action)
		} else if strings.Contains(action, "asleep") {
			if currentGuardId == longestSleeperId {
				sleepTime := entry.dateTime
				wakeTime := entries[i+1].dateTime
				recordSleepSession(sleepLog, sleepTime, wakeTime)
			}
		}
	}

	sleepiestMinute := -1

	for k, v := range sleepLog {
		if sleepiestMinute == -1 || v > sleepLog[sleepiestMinute] {
			sleepiestMinute = k
		}
	}

	return sleepiestMinute
}

func main() {
	file, _ := os.Open("./day-04-input.txt")
	
	defer file.Close()

	scanner := bufio.NewScanner(file)

	entries := make([]Entry, 0)

	for scanner.Scan() {
		entry := parseEntry(scanner.Text())
		entries = append(entries, entry)
	}

	sortEntries(entries)

	sleepiestGuardId := findSleepiestGuardId(entries)
	sleepiestMinute := findSleepiestMinute(entries, sleepiestGuardId)

	fmt.Println("sleepiest guard ID * sleepiest minute:", sleepiestGuardId * sleepiestMinute)
}
