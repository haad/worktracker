package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var EntName string
var EntDesc string
var EntDura uint64
var EntStart uint64
var EntEnd uint64
var EntType uint16
var EntBillable bool
var EntTags string

var EntProject string

func init() {
	rootCmd.AddCommand(EntCmd)
	EntCmd.AddCommand(EntCmdLog)

	EntCmdLog.Flags().StringVarP(&EntName, "name", "n", "", "Entry short name")
	EntCmdLog.Flags().StringVarP(&EntDesc, "desc", "D", "", "Entry description.")
	EntCmdLog.Flags().Uint64VarP(&EntDura, "duration", "u", 0, "Duration of existing entry.")
	EntCmdLog.Flags().Uint16VarP(&EntType, "type", "T", 1, "Entry type")
	EntCmdLog.Flags().BoolVarP(&EntBillable, "billable", "B", true, "IS entry billable.")
	EntCmdLog.Flags().StringVarP(&EntTags, "tags", "t", "", "Comma separated list of tags.")
	EntCmdLog.MarkFlagRequired("name")
	EntCmdLog.MarkFlagRequired("duration")
}

var EntCmd = &cobra.Command{
	Use:   "entry",
	Short: "Manipulate worktracker entries",
	Long:  `Entry is basic unit of time in worktracker `,
}

var EntCmdLog = &cobra.Command{
	Use:   "log",
	Short: "Create customer with given name",
	Long:  `Log some work done for a given project`,
	Run: func(cmd *cobra.Command, at []string) {
		EntLog(EntName, EntDesc, EntDura, EntType, EntBillable, EntTags)
	},
}

func EntLog(name string, desc string, dura uint64, etype uint16, billable bool, tags string) {
	fmt.Println("Creating entry:", name, "with desc:", desc, "duration was:", dura,
		"Entry is type:", etype, "and is billable:", billable, "with tags: ", tags)
}
