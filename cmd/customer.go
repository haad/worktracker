package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xlab/tablewriter"

	"github.com/haad/worktracker/sql"
)

func init() {

	var customerName string
	var customerContactName string
	var customerEmail string
	var rate uint

	var custCmd = &cobra.Command{
		Use:   "customer",
		Short: "Manipulate worktracker customers",
		Long:  `Create/Update/List/Delete customers`,
	}

	var custCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create customer with given name",
		Long:  `Create new customers`,
		Run: func(cmd *cobra.Command, args []string) {
			CustomerCreate(customerName, rate, customerContactName, customerEmail)
		},
	}
	custCreateCmd.Flags().UintVarP(&rate, "rate", "r", 0, "Default rate for a given customer")
	custCreateCmd.Flags().StringVarP(&customerName, "name", "n", "", "customer name to work with")
	custCreateCmd.Flags().StringVarP(&customerContactName, "contact-name", "C", "", "customer contact name")
	custCreateCmd.Flags().StringVarP(&customerEmail, "email", "e", "", "customer contact email")
	custCreateCmd.MarkFlagRequired("name")
	custCreateCmd.MarkFlagRequired("rate")

	var custDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete customer with given name",
		Long:  `Delete existing customers`,
		Run: func(cmd *cobra.Command, args []string) {
			CustomerDelete(customerName)
		},
	}
	custDeleteCmd.Flags().StringVarP(&customerName, "name", "n", "", "customer name to work with")
	custDeleteCmd.MarkFlagRequired("name")

	var custListCmd = &cobra.Command{
		Use:   "list",
		Short: "List customers",
		Long:  `List created customers`,
		Run: func(cmd *cobra.Command, args []string) {
			CustomerList()
		},
	}

	rootCmd.AddCommand(custCmd)
	custCmd.AddCommand(custCreateCmd)
	custCmd.AddCommand(custDeleteCmd)
	custCmd.AddCommand(custListCmd)
}

func CustomerCreate(name string, rate uint, contact string, email string) {
	fmt.Println("Creating customer:", name, "with default rate:", rate)
	sql.DBc.Create(&sql.Customer{Name: name, Rate: rate, ContactEmail: email, ContactName: contact})
}

func CustomerDelete(name string) {
	var customer sql.Customer
	var projects []sql.Project

	sql.DBc.Where("name = ?", name).First(&customer)

	fmt.Println("Deleting projects:")
	sql.DBc.Where("customer_id = ?", customer.ID).Find(&projects)
	for _, project := range projects {
		sql.DBc.Unscoped().Delete(&project)
	}

	fmt.Println("Deleting customer:", name)
	sql.DBc.Where("name = ?", name).Delete(&customer)
	sql.DBc.Unscoped().Delete(&customer)
}

func customerGetProjects(id uint) string {
	var projects []sql.Project
	var p string

	sql.DBc.Where("customer_id = ?", id).Find(&projects)
	fmt.Println("projects:", projects)
	for _, project := range projects {
		p += project.Name + " "
	}
	return p
}

func CustomerList() {

	var customers []sql.Customer

	table := tablewriter.CreateTable()
	table.AddHeaders("Customer Name", "Rate", "Contact Name", "Contact Email")

	sql.DBc.Find(&customers)
	fmt.Println("List existing customers: ")

	for _, customer := range customers {

		table.AddRow(customer.Name, customer.Rate, customer.ContactName, customer.ContactEmail)
	}
	fmt.Println(table.Render())

}
