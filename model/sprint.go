package model

import "time"

type Sprint struct {
	Id     int64
	Number int64
	Start  time.Time
	Days   [10]int
}
