package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	tm "github.com/buger/goterm"
	"github.com/duckhue01/todos/utils"
)

type Set struct {
	Pomo     int
	Short    int
	Long     int
	Interval int
}

var set Set
var setRaw, _ = ioutil.ReadFile(filepath.Join(base, "pomo.json"))

func StarPomotHandler(isMusic bool) {
	json.Unmarshal(setRaw, &set)
	var state string
	var inst string
	start := time.Now().Format(time.RFC1123)
	round := 1
	if isMusic {
		go utils.RunMP3("musics/1.mp3")
	}
	for {
		tm.Clear()
		state = "focus"
		clock(start, state, round)
		fmt.Println("press enter to continuos...")
		fmt.Scanln(&inst)
		if round%set.Interval == 0 {
			state = "long break"
		}
		state = "short break"
		tm.Clear()
		clock(start, state, round)
		fmt.Println("press enter to continuos...")
		fmt.Scanln(&inst)
		round++
	}
}

func SetPomoHandler(key string, value int) {
	json.Unmarshal(setRaw, &set)
	switch key {
	case "pomo":
		set.Pomo = value
	case "short":
		set.Short = value
	case "long":
		set.Long = value
	case "interval":
		set.Interval = value
	}
	raw, _ := json.Marshal(set)

	err := ioutil.WriteFile(filepath.Join(base, "pomo.json"), raw, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func clock(start string, state string, round int) {

	if state == "focus" {
		for i := 0; i <= (set.Pomo * 60); i++ {
			minute := (i / 60) % 60
			second := i % 60
			tm.MoveCursor(1, 10)
			tm.Println(tm.Color("===============================================", tm.RED))
			tm.Println("start time: ", start)
			fTime := fmt.Sprintf("%d:%d", minute, second)
			tm.Println("\n", (tm.Color(fTime, tm.RED)), "\n")

			tm.Println("round:", round, "- state:", tm.Color(state, tm.RED))
			tm.Println(tm.Color("===============================================", tm.RED))
			tm.Flush() // Call it every time at the end of rendering
			time.Sleep(time.Second)
		}
	} else {
		var bTime int
		if state == "short break" {
			bTime = set.Short * 60
		} else {
			bTime = set.Long * 60
		}

		for i := 0; i <= bTime; i++ {
			minute := (i / 60) % 60
			second := i % 60
			tm.MoveCursor(1, 10)
			tm.Println(tm.Color("===============================================", tm.GREEN))
			tm.Println("start time: ", start)
			fTime := fmt.Sprintf("%d:%d", minute, second)
			tm.Println("\n", (tm.Color(fTime, tm.GREEN)), "\n")

			tm.Println("round:", round, "- state: ", tm.Color(state, tm.GREEN))
			tm.Println(tm.Color("===============================================", tm.GREEN))
			tm.Flush() // Call it every time at the end of rendering
			time.Sleep(time.Second)
		}
	}

}
