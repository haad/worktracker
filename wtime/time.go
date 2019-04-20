package wtime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DateForm = "01/2006"

// GetDurantionString returns duration converted to a string with s/m/h
func GetDurantionString(duration int64) string {
	d, err := time.ParseDuration(strconv.FormatInt(duration, 10) + "s")

	if err != nil {
		return ""
	}

	return d.String()
}

// GetStartEndMonth gets first second of a month described in startD and first second of next month.
// XXX Error Handling
func GetStartEndMonth(startD time.Time) (int64, int64) {
	var eyear int
	var emonth time.Month

	emonth = startD.Month() + 1
	eyear = startD.Year()

	if startD.Month() == 12 {
		emonth = time.January
		eyear = startD.Year() + 1
	}

	start := time.Date(startD.Year(), startD.Month(), 0, 0, 0, 0, 1, time.UTC).Unix()
	end := time.Date(eyear, emonth, 0, 0, 0, 0, 0, time.UTC).Unix()

	return start, end
}

// XXX Error Handling
func CompareStartDate(startDate string, entryDate int64) bool {
	var sd time.Time
	var err error

	if strings.TrimLeft(startDate, "<>@=") == "" {
		sd = time.Now()
	} else {
		sd, err = time.Parse(DateForm, strings.TrimLeft(startDate, "<>@="))
		if err != nil {
			panic(err)
		}
	}

	switch string(startDate[0]) {
	case ">":
		if entryDate >= sd.Unix() {
			return true
		}
	case "<":
		if entryDate <= sd.Unix() {
			return true
		}
	case "=":
		start, end := GetStartEndMonth(sd)
		if entryDate >= start && entryDate < end {
			return true
		}
	case "@":
		// @ means to print only entries related to current month
		start, end := GetStartEndMonth(sd)
		if entryDate >= start && entryDate < end {
			return true
		}
	default:
		return false
	}

	return false
}

// DehumanizeDuration converts string in format 1w1d1h1m into number of seconds
func DehumanizeDuration(dura string) (int64, error) {
	var num string
	var base int64
	var duration int64

	for i := 0; i < len(dura); i++ {
		// Check if current char is a number or string
		if _, err := strconv.Atoi(string(dura[i])); err == nil {
			num += string(dura[i])
		} else {
			// Convert existing num to number as we got our first char
			if d, err := strconv.Atoi(num); err == nil {
				base = int64(d)
				num = ""
			}

			switch string(dura[i]) {
			case "w":
				duration += (base * 604800)
			case "d":
				duration += (base * 86400)
			case "h":
				duration += (base * 3600)
			case "m":
				duration += (base * 60)
			default:
				return 0, fmt.Errorf("wrong char only w,d,h,m are accepted: %s", string(dura[i]))
			}

			base = 0
		}
	}

	return duration, nil
}
