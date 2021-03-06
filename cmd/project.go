package cmd

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/haad/worktracker/model/project"
	"github.com/haad/worktracker/wtime"
)

func init() {
	var projName string
	var projCustName string
	var projEstimate string
	var projFinished bool

	var projID uint

	var projCmd = &cobra.Command{
		Use:     "project",
		Aliases: []string{"p"},
		Short:   "Manipulate worktracker projects",
		Long:    `Create project under which we can track work`,
	}

	var projCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create project with given name",
		Long:  `Create project which belongs to one customer`,
		Run: func(cmd *cobra.Command, args []string) {
			project.ProjectCreate(projName, projEstimate, projCustName)
		},
	}
	projCreateCmd.Flags().StringVarP(&projName, "name", "n", "", "Project name")
	projCreateCmd.Flags().StringVarP(&projCustName, "customer", "c", "", "Project customer name, needs to be created before.")
	projCreateCmd.Flags().StringVarP(&projEstimate, "estimate", "e", "", "Project hour estimate, valid units are s/m/h")
	projCreateCmd.MarkFlagRequired("name")
	projCreateCmd.MarkFlagRequired("customer")

	var projEditCmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit project with given name",
		Long:  `Edit project which belongs to one customer`,
		Run: func(cmd *cobra.Command, args []string) {
			project.ProjectEdit(projID, projName, projEstimate, projFinished)
		},
	}
	projEditCmd.Flags().UintVarP(&projID, "id", "i", 0, "ID of project to delete")
	projEditCmd.Flags().StringVarP(&projName, "name", "n", "", "Project name")
	projEditCmd.Flags().StringVarP(&projEstimate, "estimate", "e", "", "Project hour estimate, valid units are s/m/h")
	projEditCmd.Flags().BoolVarP(&projFinished, "finished", "f", false, "Is project finished ?")
	projEditCmd.MarkFlagRequired("id")

	var projDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete projects",
		Long:  `Delete created project`,
		Run: func(cmd *cobra.Command, args []string) {
			project.ProjectDelete(projID)
		},
	}
	projDelCmd.Flags().UintVarP(&projID, "id", "i", 0, "ID of project to delete")
	projDelCmd.MarkFlagRequired("id")

	var projListCmd = &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Long:  `List created projects`,
		Run: func(cmd *cobra.Command, args []string) {
			projectList(projCustName, projFinished)
		},
	}
	projListCmd.Flags().StringVarP(&projCustName, "customer", "c", "", "Project customer name.")
	projListCmd.Flags().BoolVarP(&projFinished, "finished", "F", false, "Include finished projects in a list.")

	rootCmd.AddCommand(projCmd)
	projCmd.AddCommand(projCreateCmd)
	projCmd.AddCommand(projEditCmd)
	projCmd.AddCommand(projDelCmd)
	projCmd.AddCommand(projListCmd)
}

func projectList(customerName string, includeFinished bool) {
	var projects []project.ProjectInt
	var timeSum int64

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Customer", "ID", "Project Name", "Finished", "Original Estimate", "Work Logged"})

	// table.AddTitle("Projects List")
	// table.SetAutoMergeCells(true)

	projects = project.ProjectList(customerName, includeFinished)

	for _, p := range projects {
		table.Append([]string{strconv.FormatUint(uint64(p.GetID()), 10), p.GetName(), p.GetCustomerName(),
			strconv.FormatBool(p.GetFinished()), p.GetEstimateString(), p.GetWorkLoggedString()})
		timeSum += p.GetWorkLogged()
	}

	table.SetFooter([]string{"", "", "", "Worked hours", "", wtime.GetDurantionString(timeSum)})

	table.Render()
}
