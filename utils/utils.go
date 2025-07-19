package utils

import (
	"fmt"
	"time"
)

func FormatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	result := ""
	if hours > 0 {
		result += fmt.Sprintf("%d hour%s ", hours, plural(hours))
	}
	if minutes > 0 {
		result += fmt.Sprintf("%d minute%s ", minutes, plural(minutes))
	}
	if seconds > 0 || result == "" {
		result += fmt.Sprintf("%d second%s", seconds, plural(seconds))
	}

	return result
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
