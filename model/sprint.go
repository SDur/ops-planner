package model

import "time"

type Sprint struct {
	Id    int64
	Nr    int64
	Start time.Time
	Days  [10]int64
}
