package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ProjectName string
var ProjectOwnerName string

func init() {
	rootCmd.AddCommand(ProjCmd)
	ProjCmd.AddCommand(ProjCreateCmd)

	ProjCreateCmd.Flags().StringVarP(&ProjectName, "name", "n", "", "Project name")
	ProjCreateCmd.Flags().StringVarP(&ProjectOwnerName, "customer", "c", "", "Project customer name, needs to be created before.")
	ProjCreateCmd.MarkFlagRequired("name")
	ProjCreateCmd.MarkFlagRequired("customer")
}

var ProjCmd = &cobra.Command{
	Use:   "project",
	Short: "Manipulate worktracker projects",
	Long:  `Create project under which we can track work`,
}

var ProjCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create project with given name",
	Long:  `Create project which belongs to one customer`,
	Run: func(cmd *cobra.Command, args []string) {
		ProjectCreate(ProjectName, ProjectOwnerName)
	},
}

func ProjectCreate(name string, owner string) {
	fmt.Println("Creating customer:", name, "with default rate:", owner)
}
