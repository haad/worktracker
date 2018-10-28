package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var entName string
	var entDesc string
	var entDura uint64
	var entBillable bool
	var entTags string

	var entProjectName string

	var entCmd = &cobra.Command{
		Use:   "entry",
		Short: "Manipulate worktracker entries",
		Long:  `Entry is basic unit of time in worktracker `,
	}

	var entCmdLog = &cobra.Command{
		Use:   "log",
		Short: "Create customer with given name",
		Long:  `Log some work done for a given project`,
		Run: func(cmd *cobra.Command, at []string) {
			EntLog(entName, entDesc, entDura, entProjectName, entBillable, entTags)
		},
	}

	rootCmd.AddCommand(entCmd)
	entCmd.AddCommand(entCmdLog)

	entCmdLog.Flags().StringVarP(&entName, "name", "n", "", "Entry short name")
	entCmdLog.Flags().StringVarP(&entDesc, "desc", "D", "", "Entry description.")
	entCmdLog.Flags().StringVarP(&entProjectName, "project", "P", "", "Project to which entry belongs")
	entCmdLog.Flags().Uint64VarP(&entDura, "duration", "u", 0, "Duration of existing entry.")
	entCmdLog.Flags().BoolVarP(&entBillable, "billable", "B", true, "IS entry billable.")
	entCmdLog.Flags().StringVarP(&entTags, "tags", "t", "", "Comma separated list of tags.")
	entCmdLog.MarkFlagRequired("name")
	entCmdLog.MarkFlagRequired("duration")
	entCmdLog.MarkFlagRequired("project")
}

func EntLog(name string, desc string, dura uint64, eproj string, billable bool, tags string) {
	fmt.Println("Creating entry:", name, "with desc:", desc, "duration was:", dura,
		"Entry belongs to :", eproj, "and is billable:", billable, "with tags: ", tags)
}
