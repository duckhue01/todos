package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	tm "github.com/buger/goterm"
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
		music()
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

func SetPomoHandler() {
	fmt.Println("set is called")
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

func music() {

}
