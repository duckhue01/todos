
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
				services.SetPomoHandler()
			}

		} else {
			fmt.Println("st happened")
		}

	},
}

func init() {
	rootCmd.AddCommand(pomoCmd)
	pomoCmd.Flags().BoolP("music", "m", true, "open music or not")

}
