package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xlab/tablewriter"

	"github.com/haad/worktracker/model/project"
)

func init() {
	var projName string
	var projCustName string

	var projCmd = &cobra.Command{
		Use:   "project",
		Short: "Manipulate worktracker projects",
		Long:  `Create project under which we can track work`,
	}

	var projCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create project with given name",
		Long:  `Create project which belongs to one customer`,
		Run: func(cmd *cobra.Command, args []string) {
			ProjectCreate(projName, projCustName)
		},
	}
	projCreateCmd.Flags().StringVarP(&projName, "name", "n", "", "Project name")
	projCreateCmd.Flags().StringVarP(&projCustName, "customer", "c", "", "Project customer name, needs to be created before.")
	projCreateCmd.MarkFlagRequired("name")
	projCreateCmd.MarkFlagRequired("customer")

	var projDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete projects",
		Long:  `Delete created project`,
		Run: func(cmd *cobra.Command, args []string) {
			ProjectDelete(projName)
		},
	}
	projDelCmd.Flags().StringVarP(&projName, "name", "n", "", "Project name")
	projDelCmd.MarkFlagRequired("name")

	var projListCmd = &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Long:  `List created projects`,
		Run: func(cmd *cobra.Command, args []string) {
			ProjectList()
		},
	}

	rootCmd.AddCommand(projCmd)
	projCmd.AddCommand(projCreateCmd)
	projCmd.AddCommand(projListCmd)
}

func ProjectCreate(name string, customerName string) {
	project.ProjectCreate(name, customerName)
}

func ProjectDelete(name string) {
	project.ProjectDelete(name)
}

func ProjectList() {
	var projects []project.ProjectInt

	table := tablewriter.CreateTable()
	table.AddHeaders("Project Name", "Customer")

	projects = project.ProjectList()

	for _, p := range projects {
		table.AddRow(p.GetName(), p.GetCustomerName())
	}

	fmt.Println(table.Render())
}
