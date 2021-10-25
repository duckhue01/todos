package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

type Task struct {
	Des    []string
	IsDone bool
	Tag    string
	Id int
}
type Seed struct {
	Des     string
	Section string
}

var todos map[string][]Task
var schedule map[string]Seed

func init() {
	rootCmd.AddCommand(listCmd)

	todosRaw, err1 := ioutil.ReadFile("todos.json")
	scheduleRaw, err2 := ioutil.ReadFile("schedule.json")
	if err1 != nil && err2 != nil {
		panic("some thing went wrong !!!")
	}

	json.Unmarshal(todosRaw, &todos)
	json.Unmarshal(scheduleRaw, &schedule)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "print list of tasks",
	Args:  cobra.ExactArgs(1),
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case args[0] == "task":
			taskHandler()

		case args[0] == "stats":
			fmt.Println("this is stats")
		case args[0] == "schedule":
			fmt.Println("this is schedule")
		}

	},
}

func taskHandler() {

	// fmt.Println(schedule)
	for i := 0; i < len(todos["25-10-2021"]); i++ {

		fmt.Printf("#%s: ", todos["25-10-2021"][i].Tag)

		for a := 0; a < len(todos["25-10-2021"][i].Des); a++ {
			fmt.Println(todos["25-10-2021"][i].Des[a])

		}

	}

	// fmt.Println(todos)

}

// func statsHandler() {

// }

// func scheduleHandle() {

// }
