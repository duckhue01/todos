package utils

import (
	"fmt"
	"time"
)

func DateToString(time time.Time) string {
	day := time.Day()
	mon := time.Month()
	year := time.Year()

	return fmt.Sprintf("%02d-%02d-%04d", day, mon, year)

}

