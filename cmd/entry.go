package cmd

import (
	"fmt"
	//"time"

	"github.com/spf13/cobra"
	"github.com/xlab/tablewriter"

	"github.com/haad/worktracker/model/entry"
)

func init() {
	var entName string
	var entDesc string
	var entDura string
	var entBillable bool
	var entTags string
	var entProjectName string
	var entCustomerName string

	var entID uint

	var entCmd = &cobra.Command{
		Use:     "entry",
		Aliases: []string{"e"},
		Short:   "Manipulate worktracker entries",
		Long:    `Entry is basic unit of time in worktracker `,
	}

	var entCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"log", "l"},
		Short:   "Create customer with given name",
		Long:    `Log some work done for a given project`,
		Run: func(cmd *cobra.Command, at []string) {
			entry.EntCreate(entName, entDesc, entDura, entProjectName, entCustomerName, entBillable, entTags)
		},
	}

	entCreateCmd.Flags().StringVarP(&entName, "name", "n", "", "Entry short name")
	entCreateCmd.Flags().StringVarP(&entDesc, "desc", "D", "", "Entry description.")
	entCreateCmd.Flags().StringVarP(&entProjectName, "project", "P", "", "Project to which entry belongs")
	entCreateCmd.Flags().StringVarP(&entCustomerName, "customer", "c", "", "Customer to which entry belongs")
	entCreateCmd.Flags().StringVarP(&entDura, "duration", "u", "", "Duration of existing entry.")
	entCreateCmd.Flags().BoolVarP(&entBillable, "billable", "B", true, "Is entry billable.")
	entCreateCmd.Flags().StringVarP(&entTags, "tags", "t", "", "Comma separated list of tags.")
	entCreateCmd.MarkFlagRequired("name")
	entCreateCmd.MarkFlagRequired("duration")
	entCreateCmd.MarkFlagRequired("project")

	var entDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete Entects",
		Long:  `Delete created Entect`,
		Run: func(cmd *cobra.Command, args []string) {
			entry.EntDelete(entID)
		},
	}
	entDelCmd.Flags().UintVarP(&entID, "id", "i", 0, "ID of entry to delete")
	entDelCmd.MarkFlagRequired("id")

	var entListCmd = &cobra.Command{
		Use:   "list",
		Short: "List Entects",
		Long:  `List created Entects`,
		Run: func(cmd *cobra.Command, args []string) {
			EntList(entProjectName, entCustomerName)
		},
	}
	entListCmd.Flags().StringVarP(&entProjectName, "project", "P", "", "Project to which entry belongs")
	entListCmd.Flags().StringVarP(&entCustomerName, "customer", "c", "", "Customer to which entry belongs")

	rootCmd.AddCommand(entCmd)
	entCmd.AddCommand(entCreateCmd)
	entCmd.AddCommand(entDelCmd)
	entCmd.AddCommand(entListCmd)
}

func EntList(projectName string, customerName string) {
	var entries []entry.EntryInt

	table := tablewriter.CreateTable()
	table.AddHeaders("ID", "Entry Name", "Duration", "Desc", "Project Name", "Customer Name")
	table.AddTitle("Entries List")

	entries = entry.EntList(projectName, customerName)

	for _, e := range entries {
		table.AddRow(e.GetID(), e.GetName(), e.GetDurationString(), e.GetDesc(),
			e.GetProjectName(), e.GetCustomerName())
	}
	fmt.Println(table.Render())
}
