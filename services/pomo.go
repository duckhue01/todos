package services

import (
	"encoding/json"
	"time"

	"github.com/duckhue01/todos/db"
	"github.com/duckhue01/todos/view"
)

type Pomo struct {
	db db.DB
}

func NewPomo(db db.DB) *Pomo {
	return &Pomo{
		db: db,
	}
}

func (pomo *Pomo) StartPomo(needMusic bool) {
	pomoConfig, err := pomo.db.ReadPomoConfig()
	if err != nil {

	}
	startTime := time.Now().Format(time.Kitchen)
	if needMusic {
		go func() {
			for {
				RunMP3(pomo.db.ReadEpicMusic())
			}
		}()
	}

	view.StartPomo(pomoConfig, "focus", startTime)

}

func (pomo *Pomo) SetPomo(key string, value time.Duration) {
	pomoConfig, err := pomo.db.ReadPomoConfig()
	if err != nil {

	}

	switch key {
	case "pomo":
		pomoConfig.Pomo = value
	case "break":
		pomoConfig.Break = value
	case "interval":
		pomoConfig.Interval = int(value)
	}
	raw, err := json.Marshal(pomoConfig)
	if err != nil {

	}

	err = pomo.db.WritePomoConfig(raw)
	if err != nil {

	}

}

func (pomo *Pomo) InfoPomo() {

}
