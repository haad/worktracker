package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const major uint = 0
const minor uint = 0
const patch uint = 1

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of worktracker",
	Long:  `All software has versions. This is Worktrackers's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Work time tracker v%d.%d.%d -- HEAD\n", major, minor, patch)
	},
}
