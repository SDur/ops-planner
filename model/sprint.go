package model

import "time"

type Sprint struct {
	Id    int64     `json:"id"`
	Nr    int64     `json:"nr"`
	Start time.Time `json:"start"`
	Days  [10]int64 `json:"days"`
}
