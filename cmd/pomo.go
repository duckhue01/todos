package cmd

import (
	"fmt"
	"time"

	"github.com/duckhue01/todos/db"
	"github.com/duckhue01/todos/services"
	"github.com/spf13/cobra"
)

// pomoCmd represents the pomo command
var pomoCmd = &cobra.Command{
	Use:   "pomo",
	Short: "pomodoro",
	Long:  `pomodoro`,
	Run: func(cmd *cobra.Command, args []string) {
		needMusic, _ := cmd.Flags().GetBool("music")
		// todo: need to replace placeholder
		pomoService := services.NewPomo(db.New("some-path"))
		if len(args) > 0 {
			switch {
			case args[0] == "start":
				pomoService.StartPomo(needMusic)
			case args[0] == "set":
				key, _ := cmd.Flags().GetString("key")
				value, _ := cmd.Flags().GetInt("value")
				if key != "" && value > 0 {
					pomoService.SetPomo(key, time.Duration(value))
				}
			case args[0] == "info":
				pomoService.InfoPomo()
			case args[0] == "test":
			}

		} else {
			fmt.Println("operation is required. please use [start/set]")
		}

	},
}

func init() {
	rootCmd.AddCommand(pomoCmd)
	pomoCmd.Flags().BoolP("music", "m", true, "open music or not")
}
