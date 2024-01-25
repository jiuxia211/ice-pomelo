package utils

import (
	"fmt"
	"time"
)

func GetTime() string {
	// 获取当前时间
	currentTime := time.Now()

	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	return fmt.Sprintf("%d%02d%02d_%02d%02d", year, month, day, hour, minute)
}
