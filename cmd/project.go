package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xlab/tablewriter"

	"github.com/haad/worktracker/sql"
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
	var customer sql.Customer
	sql.DBc.Where("name = ?", customerName).First(&customer)

	fmt.Println("Creating project:", name, "with default rate:", customer)
	sql.DBc.Create(&sql.Project{Name: name, CustRef: customer.ID})
}

func ProjectList() {
	var projects []sql.Project

	table := tablewriter.CreateTable()
	table.AddHeaders("Project Name", "Customer")

	sql.DBc.Set("gorm:auto_preload", true).Find(&projects)

	for _, project := range projects {
		table.AddRow(project.Name, project.Customer.Name)
	}

	fmt.Println(table.Render())
}
