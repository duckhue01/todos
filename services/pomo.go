package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path"
	"path/filepath"
	"time"

	"github.com/duckhue01/todos/db"
	"github.com/duckhue01/todos/models"
	"github.com/duckhue01/todos/utils"
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

	pomoConfig, err := pomo.db.GetPomoConfig()
	if err != nil {

	}
	startTime := time.Now().Format(time.Kitchen)
	if needMusic {
		go func() {
			for {
				rand.Seed(time.Now().UnixNano())
				utils.RunMP3(path.Join(pomo.base, fmt.Sprintf("musics/%d.mp3", rand.Intn(8))))
			}
		}()
	}

	view.StartPomo(pomoConfig, "focus", startTime)

}

func (pomo *Pomo) SetPomo(key string, value time.Duration) {
	var pomoConfig models.PomoConfig
	var setRaw, pomoConfigErr = ioutil.ReadFile(filepath.Join(pomo.base, "pomo.json"))
	if pomoConfigErr != nil {
		fmt.Println("can't read pomo config file")
	}

	json.Unmarshal(setRaw, &pomoConfig)
	switch key {
	case "pomo":
		pomoConfig.Pomo = value
	case "break":
		pomoConfig.Break = value
	case "interval":
		pomoConfig.Interval = int(value)
	}
	raw, _ := json.Marshal(pomoConfig)

	err := ioutil.WriteFile(filepath.Join(pomo.base, "pomo.json"), raw, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (pomo *Pomo) InfoPomo() {

}
