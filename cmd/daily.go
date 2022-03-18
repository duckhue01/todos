package cmd

import (
	"fmt"
	"time"

	"github.com/duckhue01/todos/services"
	"github.com/duckhue01/todos/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dailyCmd)
	dailyCmd.Flags().StringP("date", "d", utils.TimeToString(time.Now()), "get todos with specific date. example: 10-10-2021")

}

var dailyCmd = &cobra.Command{
	Use:   "daily [list/add/delete/done]",
	Short: "daily",
	Long:  `daily`,
	Run: func(cmd *cobra.Command, args []string) {
		daily := services.NewDaily()
		date, err := cmd.Flags().GetString("date")

		if err != nil {
			fmt.Println(fmt.Errorf("%v", err))
		}

		switch args[0] {
		case "list":
			fmt.Println(daily.List(date))
		case "add":
		case "delete":
		case "done":
		default:
			fmt.Println("please use: daily [list/add/delete/done]")

		}

	},
}
