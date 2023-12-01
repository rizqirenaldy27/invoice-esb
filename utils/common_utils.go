package utils

import "time"

func ParseDate(dateString string) *time.Time {
	layout := "02/01/2006" // format: dd/mm/yyyy
	parsedDate, _ := time.Parse(layout, dateString)
	return &parsedDate
}
