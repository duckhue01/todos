package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{

	Use:   "add",
	Short: "print list of tasks",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}
