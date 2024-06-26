package date

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetData() string {
	date := time.Now().UTC()
	return date.Format(apiDateLayout)
}

