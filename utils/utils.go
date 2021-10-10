package utils

import (
	"strconv"
	"time"
)

func GetDateAndHour() (string, string) {
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)

	tday := tomorrow.Day()
	tmonth := tomorrow.Month()
	tyear := tomorrow.Year()
	hr, min, _ := tomorrow.Clock()

	day := strconv.Itoa(tday) + "-" + strconv.Itoa(int(tmonth)) + "-" + strconv.Itoa(tyear)
	hour := strconv.Itoa(hr) + ":" + strconv.Itoa(int(min))

	return day, hour
}
