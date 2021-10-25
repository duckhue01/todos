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
}

func init() {
	rootCmd.AddCommand(listCmd)
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

}

func statsHandler() {

}

func scheduleHandle() {

}