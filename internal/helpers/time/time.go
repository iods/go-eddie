package time

import (
	"strconv"
	"strings"
	"time"
)

func Create(test string) time.Time {
	d := time.Now()
	t := strings.Split(test, ":")
	hr := t[0]
	min := t[1]
	h, err := strconv.Atoi(hr)
	m, err := strconv.Atoi(min)
	if err != nil {
		panic(err)
	}
	tm := time.Date(d.Year(), d.Month(), d.Day(), h, m, 0, 0, time.Local)
	return tm
}
