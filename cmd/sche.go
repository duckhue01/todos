package cmd

import (
	"github.com/duckhue01/todos/view"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "sche",
	Short: "print list of tasks",
	Args:  cobra.ExactArgs(1),
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		// scheService := services.NewScheService("/Users/duckhue01/code/side/todos")
		switch {
		case args[0] == "list":
			// scheService.GetSche()
			view.StartSche()
		}

	},
}
