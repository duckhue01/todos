package cmd

import (
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
		switch {
		case args[0] == "list":
		}

	},
}
