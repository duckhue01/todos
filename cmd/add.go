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

// colorBlue := "\033[34m"
// colorCyan := "\033[36m"

// colorReset := "\033[0m"
// colorRed := "\033[31m"
// colorGreen := "\033[32m"
// colorYellow := "\033[33m"
// colorPurple := "\033[35m"
// colorWhite := "\033[37m"
