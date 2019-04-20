package cmd

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/haad/worktracker/model/entry"
	"github.com/haad/worktracker/wtime"
)

func init() {
	var entName string
	var entDesc string
	var entDura string
	var entStart string
	var entBillable bool
	var entTags string
	var entProjectID uint

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
		Short:   "Create Entry for given project",
		Long:    `Log some work done for a given project`,
		Run: func(cmd *cobra.Command, at []string) {
			entry.EntCreate(entName, entDesc, entDura, entStart, entProjectID, entBillable, entTags)
		},
	}

	entCreateCmd.Flags().StringVarP(&entName, "name", "n", "", "Entry short name")
	entCreateCmd.Flags().StringVarP(&entDesc, "desc", "D", "", "Entry description.")
	entCreateCmd.Flags().UintVarP(&entProjectID, "id", "i", 0, "ID of project to delete")
	entCreateCmd.Flags().StringVarP(&entDura, "duration", "u", "", "Duration of existing entry, valid units are s/m/h")
	entCreateCmd.Flags().StringVarP(&entStart, "start", "s", "", "Start date of work in format DD/MM/YYYY")
	entCreateCmd.Flags().BoolVarP(&entBillable, "billable", "B", true, "Is entry billable.")
	entCreateCmd.Flags().StringVarP(&entTags, "tags", "t", "", "Comma separated list of tags.")
	entCreateCmd.MarkFlagRequired("name")
	entCreateCmd.MarkFlagRequired("duration")

	var entDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete Entry",
		Long:  `Delete Entry`,
		Run: func(cmd *cobra.Command, args []string) {
			entry.EntDelete(entID)
		},
	}
	entDelCmd.Flags().UintVarP(&entID, "id", "i", 0, "ID of entry to delete")
	entDelCmd.MarkFlagRequired("id")

	var entListCmd = &cobra.Command{
		Use:   "list",
		Short: "List Entries",
		Long:  `List created Entries`,
		Run: func(cmd *cobra.Command, args []string) {
			entList(entProjectID, entStart)
		},
	}
	entListCmd.Flags().UintVarP(&entProjectID, "id", "i", 0, "ID of project to delete")
	entListCmd.Flags().StringVarP(&entStart, "date", "d", "", `Date string in following format @<>= MM/YYYY,
    to select entries with start date before use --date <10/2018,
    to select entries with start date after use --date >10/2018,
    to select entries done in a given month use --date =10/2018,
    for selecting current month there is a shortcut by using --date @.`)

	rootCmd.AddCommand(entCmd)
	entCmd.AddCommand(entCreateCmd)
	entCmd.AddCommand(entDelCmd)
	entCmd.AddCommand(entListCmd)
}

func entList(projectID uint, startDate string) {
	var entries []entry.EntryInt
	var timeSum int64

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Entry Name", "Start Date", "Duration", "Desc", "Project Name", "Customer Name"})
	//	table.AddTitle("Entries List")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	entries = entry.EntList(projectID, startDate)

	for _, e := range entries {
		table.Append([]string{strconv.FormatUint(uint64(e.GetID()), 10), e.GetName(), e.GetSDateString(), e.GetDurationString(), e.GetDesc(),
			e.GetProjectName(), e.GetCustomerName()})
		timeSum += e.GetDuration()
	}
	table.SetFooter([]string{"", "", "", "", "Worked hours", "", wtime.GetDurantionString(timeSum)})

	table.Render()
}
