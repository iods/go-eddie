package time

import (
	"strconv"
	"strings"
	"time"
)

func UpdateTime(t string) time.Time {
	n := time.Now()
	s := strings.Split(t, ":")

	h, err := strconv.Atoi(s[0])
	m, err := strconv.Atoi(s[1])
	if err != nil {
		panic(err)
	}

	d := time.Date(n.Year(), n.Month(), n.Day(), h, m, n.Second(), 0, time.Local)

	return d
}