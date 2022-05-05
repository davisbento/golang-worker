package utils

import "time"

func GetDateIso() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}
