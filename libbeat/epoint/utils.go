package epoint

import "time"

func GAtCurrentDate() string {
	t := time.Now()
	return t.Format("2006-01-02")
}
