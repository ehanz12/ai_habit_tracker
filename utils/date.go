package utils

import (
	"fmt"
	"regexp"
)

var timeRegex = regexp.MustCompile(`^([01]\d|2[0-3]):([0-5]\d):([0-5]\d)$`)

func NormalizeAndValidateTime(timeStr string) (string, error) {
	// auto convert HH:MM → HH:MM:SS
	if len(timeStr) == 5 {
		timeStr += ":00"
	}

	if !timeRegex.MatchString(timeStr) {
		return "", fmt.Errorf("format waktu harus HH:MM atau HH:MM:SS")
	}

	return timeStr, nil
}