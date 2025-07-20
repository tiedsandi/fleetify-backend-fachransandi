package helpers

import (
	"errors"
	"time"
)

func parseFlexibleTime(input string) (time.Time, error) {
	layouts := []string{"15:04", "15:04:05"}
	var t time.Time
	var err error

	for _, layout := range layouts {
		t, err = time.Parse(layout, input)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("format waktu tidak valid, gunakan format HH:mm atau HH:mm:ss")
}

func ValidateClockTimes(in, out string) error {
	clockIn, err := parseFlexibleTime(in)
	if err != nil {
		return errors.New("jam Masuk tidak valid: " + err.Error())
	}
	clockOut, err := parseFlexibleTime(out)
	if err != nil {
		return errors.New("jam Keluar tidak valid: " + err.Error())
	}
	if !clockIn.Before(clockOut) {
		return errors.New("jam Masuk harus lebih kecil dari jam Keluar")
	}
	return nil
}
