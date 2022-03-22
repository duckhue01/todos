package models

import "time"

type PomoConfig struct {
	Pomo     time.Duration `json:"pomo"`
	Break    time.Duration `json:"break"`
	Interval int           `json:"interval"`
}

