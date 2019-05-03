package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

const DTF string = "2006-01-02" // Default Time Format
var timeFormat = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var ISO8601 = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} [+-]\d{4}$`)
var findISO8601 = regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} [+-]\d{4}`)
var RFC3339 = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[+-]\d{2}:\d{2}$`)

func timeZoneOffset() string {
	_, offset := time.Now().Zone()
	if offset > 0 {
		return fmt.Sprintf("+%04d", offset/60/60*100)
	}
	return fmt.Sprintf("-%04d", (-1*offset)/60/60*100)
}

func ISO8601_to_RFC3339(t string) (string, error) {
	if !ISO8601.MatchString(t) {
		return t, errors.New("time string is not ISO8601 format.")
	}
	return fmt.Sprintf("%sT%s%s:%s", t[0:10], t[11:19], t[20:23], t[23:25]), nil
}

func beforeMonth() (string, string) {
	y, m, _ := time.Now().Date()
	if m == 1 {
		y -= 1
		m = 12
	} else {
		m -= 1
	}
	since := time.Date(y, m, 1, 0, 0, 0, 0, time.Now().Location())
	return fmt.Sprintf(since.Format(DTF)), fmt.Sprintf(since.AddDate(0, 1, 0).Add(-time.Nanosecond).Format(DTF))
}

func thisMonth() (string, string) {
	y, m, _ := time.Now().Date()
	since := time.Date(y, m, 1, 0, 0, 0, 0, time.Now().Location())
	return fmt.Sprintf(since.Format(DTF)), fmt.Sprintf(since.AddDate(0, 1, 0).Add(-time.Nanosecond).Format(DTF))
}
