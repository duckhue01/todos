package cmd

import (
	"fmt"
	"github.com/duckhue01/todos/services"
	"github.com/spf13/cobra"
)

type Set struct {
	Pomo     int
	Short    int
	Long     int
	Interval int
}

// pomoCmd represents the pomo command
var pomoCmd = &cobra.Command{
	Use:   "pomo",
	Short: "pomodoro",
	Long:  `pomodoro`,
	Run: func(cmd *cobra.Command, args []string) {

		isMusic, _ := cmd.Flags().GetBool("music")

		if len(args) > 0 {
			switch {
			case args[0] == "start":
				services.StarPomotHandler(isMusic)
			case args[0] == "set":
				key, _ := cmd.Flags().GetString("key")
				value, _ := cmd.Flags().GetInt("value")
				if key != "" && value > 0 {
					services.SetPomoHandler(key, value)
				}
			}

		} else {
			fmt.Println("st happened")
		}

	},
}

func init() {
	rootCmd.AddCommand(pomoCmd)
	pomoCmd.Flags().BoolP("music", "m", true, "open music or not")
	pomoCmd.Flags().IntP("value", "v", 0, "value")
	pomoCmd.Flags().StringP("key", "k", "", "key")
}
