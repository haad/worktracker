package cmd

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/haad/worktracker/model/customer"
)

func init() {

	var customerName string
	var customerContactName string
	var customerEmail string
	var rate int

	var custCmd = &cobra.Command{
		Use:     "customer",
		Aliases: []string{"c"},
		Short:   "Manipulate worktracker customers",
		Long:    `Create/Update/List/Delete customers`,
	}

	var custCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create customer with given name",
		Long:  `Create new customers`,
		Run: func(cmd *cobra.Command, args []string) {
			customer.CustomerCreate(customerName, rate, customerContactName, customerEmail)
		},
	}
	custCreateCmd.Flags().IntVarP(&rate, "rate", "r", 0, "Default rate for a given customer")
	custCreateCmd.Flags().StringVarP(&customerName, "name", "n", "", "customer name to work with")
	custCreateCmd.Flags().StringVarP(&customerContactName, "contact-name", "C", "", "customer contact name")
	custCreateCmd.Flags().StringVarP(&customerEmail, "email", "e", "", "customer contact email")
	custCreateCmd.MarkFlagRequired("name")
	custCreateCmd.MarkFlagRequired("rate")

	var custEditCmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit customer with given name",
		Long:  `Edit customers information in DB`,
		Run: func(cmd *cobra.Command, args []string) {
			customer.CustomerEdit(customerName, rate, customerContactName, customerEmail)
		},
	}
	custEditCmd.Flags().IntVarP(&rate, "rate", "r", -1, "Default rate for a given customer")
	custEditCmd.Flags().StringVarP(&customerName, "name", "n", "", "customer name to work with")
	custEditCmd.Flags().StringVarP(&customerContactName, "contact-name", "C", "", "customer contact name")
	custEditCmd.Flags().StringVarP(&customerEmail, "email", "e", "", "customer contact email")
	custEditCmd.MarkFlagRequired("name")

	var custDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete customer with given name",
		Long:  `Delete existing customers`,
		Run: func(cmd *cobra.Command, args []string) {
			customer.CustomerDelete(customerName)
		},
	}
	custDeleteCmd.Flags().StringVarP(&customerName, "name", "n", "", "customer name to work with")
	custDeleteCmd.MarkFlagRequired("name")

	var custListCmd = &cobra.Command{
		Use:   "list",
		Short: "List customers",
		Long:  `List created customers`,
		Run: func(cmd *cobra.Command, args []string) {
			customerList()
		},
	}

	rootCmd.AddCommand(custCmd)
	custCmd.AddCommand(custCreateCmd)
	custCmd.AddCommand(custEditCmd)
	custCmd.AddCommand(custDeleteCmd)
	custCmd.AddCommand(custListCmd)
}

func customerList() {
	var customers []customer.CustomerInt

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Customer Name", "Rate", "Contact Name", "Contact Email"})
	// table.AddTitle("Customer List")

	customers = customer.CustomerList()

	for _, c := range customers {
		table.Append([]string{strconv.FormatUint(uint64(c.GetID()), 10), c.GetName(), string(c.GetRate()), c.GetContactName(), c.GetContactEmail()})
	}
	table.Render()
}
