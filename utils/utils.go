package utils

import (
	"fmt"
	"time"
)

func FormatDuration(d time.Duration) string {
	jam := int(d.Hours())
	menit := int(d.Minutes()) % 60
	detik := int(d.Seconds()) % 60

	result := ""
	if jam > 0 {
		result += fmt.Sprintf("%d jam ", jam)
	}
	if menit > 0 {
		result += fmt.Sprintf("%d menit ", menit)
	}
	if detik > 0 || result == "" {
		result += fmt.Sprintf("%d detik", detik)
	}

	return result
}
