package cmd

import (
	"fmt"

	"github.com/duckhue01/todos/services"
	"github.com/spf13/cobra"
)

// zoomCmd represents the zoom command
var zoomCmd = &cobra.Command{
	Use:   "zoom",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			switch {
			case args[0] == "start":

				username, _ := cmd.Flags().GetString("username")
				password, _ := cmd.Flags().GetString("password")
				auto, _ := cmd.Flags().GetBool("auto")
				services.StartZoomhandler(username, password, auto)
			case args[0] == "update":
				user, _ := cmd.Flags().GetString("username")
				pass, _ := cmd.Flags().GetString("password")
				services.UpdateZoomHandler(user, pass)
			case args[0] == "check":
				code, _ := cmd.Flags().GetString("code")
				if code != "" {
					services.CheckRoomHandler(code)
				}

				services.CheckTodayHandler()
			}

		} else {
			fmt.Println("enter what you want to do")
		}
	},
}

func init() {
	rootCmd.AddCommand(zoomCmd)
	zoomCmd.Flags().StringP("username", "u", "", "user name of fucking tlu account")
	zoomCmd.Flags().StringP("password", "p", "", "password")
	zoomCmd.Flags().BoolP("auto", "a", false, "auto mode")
	zoomCmd.Flags().StringP("code", "c", "", "check room code")
}
