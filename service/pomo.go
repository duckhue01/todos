package service

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(fmt.Errorf("fail to read pomo config: %w", err))
	}
	startTime := time.Now().Format(time.Kitchen)
	if needMusic {
		go func() {
			for {
				RunMP3(pomo.db.ReadChillMusic())
			}
		}()
	}

	view.StartPomo(pomoConfig, "focus", startTime)

}

func (pomo *Pomo) SetPomo(key string, value time.Duration) {
	pomoConfig, err := pomo.db.ReadPomoConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("fail to read pomo config: %w", err))
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
		fmt.Println(fmt.Errorf("fail to marshal pomo config: %w", err))
	}

	err = pomo.db.WritePomoConfig(raw)
	if err != nil {
		fmt.Println(fmt.Errorf("fail to write pomo config: %w", err))
	}

	fmt.Printf("%+v", pomoConfig)

}

func (pomo *Pomo) InfoPomo() {
	pomoConfig, err := pomo.db.ReadPomoConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("fail to read pomo config: %w", err))
	}
	fmt.Printf("%+v", pomoConfig)

}
