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

	"github.com/duckhue01/todos/models"
	"github.com/duckhue01/todos/utils"
	"github.com/duckhue01/todos/view"
)

type Pomo struct {
	base string
}

func NewPomo(base string) *Pomo {
	return &Pomo{
		base: base,
	}
}

func (pomo *Pomo) StartPomoHanddler(needMusic bool) {
	var pomoConfig models.PomoConfig
	var setRaw, pomoConfigErr = ioutil.ReadFile(filepath.Join(pomo.base, "pomo.json"))
	if pomoConfigErr != nil {
		fmt.Println("can't read pomo config file")
	}

	json.Unmarshal(setRaw, &pomoConfig)

	startTime := time.Now().Format(time. Kitchen)
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

func (pomo *Pomo) SetPomoHandler(key string, value time.Duration) {
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

func (pomo *Pomo) InfoPomoHandler() {

}
