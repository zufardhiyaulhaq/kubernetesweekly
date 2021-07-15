package utils

import (
	"strings"
	"time"
)

func GetDate() string {
	timedate := time.Now()
	date := timedate.Format("01-02-2006")
	date = strings.ReplaceAll(date, "-", "/")

	return date
}
