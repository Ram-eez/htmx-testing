package utils

import "time"

func Sleep(seconds int) string {
	time.Sleep(time.Duration(seconds) * time.Second)
	return ""
}
