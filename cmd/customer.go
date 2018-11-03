package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xlab/tablewriter"

	"github.com/haad/worktracker/model/customer"
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
	customer.CustomerCreate(name, rate, contact, email)
}

func CustomerDelete(name string) {
	customer.CustomerDelete(name)
}

func CustomerList() {
	var customers []customer.CustomerInt

	table := tablewriter.CreateTable()
	table.AddHeaders("Customer Name", "Rate", "Contact Name", "Contact Email")

	customers = customer.CustomerList()

	fmt.Println("List existing customers: ")

	for _, customer := range customers {
		table.AddRow(customer.GetName(), customer.GetRate(), customer.GetContactName(), customer.GetContactEmail())
	}
	fmt.Println(table.Render())

}
