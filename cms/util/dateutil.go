package util

import (
	"fmt"
	"time"
)

/**
 * 取得当前日期时间字符串，为RFC1123格式
 */
func GetRFCDate() string {
	now := time.Now()
	utcNow := now.UTC()

	return fmt.Sprintf(utcNow.Format("Mon, 02 Jan 2006 15:04:05 GMT"))
}
