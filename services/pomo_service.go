package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path"
	"path/filepath"
	"strconv"
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

var base = "/Users/duckhue01/code/side/todos"


var set Set
var setRaw, _ = ioutil.ReadFile(filepath.Join(base, "pomo.json"))

func StarPomotHandler(isMusic bool) {
	json.Unmarshal(setRaw, &set)
	var state string
	var inst string
	start := time.Now().Format(time.RFC1123)
	round := 1
	if isMusic {
		go func() {
			for {
				rand.Seed(time.Now().UnixNano())
				utils.RunMP3(path.Join(base, fmt.Sprintf("musics/%d.mp3", rand.Intn(8))))
			}
		}()
	}
	for {
		quit := make(chan bool)
		tm.Clear()
		state = "focus"
		clock(start, state, round)
		fmt.Println("press enter to continuos...")
		go func() {
			select {
			case <-quit:
				return
			default:
				for {
					utils.RunMP3(path.Join(base, "alarm.mp3"))
				}
			}
		}()
		fmt.Scanln(&inst)
		quit <- true
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

	timeF := func(m, s int) string {
		mf, sf := strconv.Itoa(m), strconv.Itoa(s)
		if m < 10 {
			mf = "0" + mf
		}
		if s < 10 {
			sf = "0" + sf
		}

		return fmt.Sprintf("%s:%s", mf, sf)
	}

	if state == "focus" {
		for i := 0; i <= (set.Pomo * 60); i++ {
			m := (i / 60) % 60
			s := i % 60
			tm.MoveCursor(1, 5)
			tm.Println(tm.Color("===============================================", tm.RED))
			tm.Println("start time: ", start)
			tm.Println("\n", (tm.Color(timeF(m, s), tm.RED)), "\n")
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
			tm.MoveCursor(1, 5)
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
