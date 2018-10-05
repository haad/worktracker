package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CustomerName string
var CustomerContactName string
var CustomerEmail string
var Rate int

func init() {
	rootCmd.AddCommand(CustCmd)
	CustCmd.AddCommand(CustCreateCmd)

	CustCreateCmd.Flags().IntVarP(&Rate, "rate", "r", 0, "Default rate for a given customer")
	CustCreateCmd.Flags().StringVarP(&CustomerName, "name", "n", "", "Customer name to work with")
	CustCreateCmd.Flags().StringVarP(&CustomerContactName, "contact-name", "C", "", "Customer contact name")
	CustCreateCmd.Flags().StringVarP(&CustomerEmail, "email", "e", "", "Customer contact email")
	CustCreateCmd.MarkFlagRequired("name")
	CustCreateCmd.MarkFlagRequired("rate")
}

var CustCmd = &cobra.Command{
	Use:   "customer",
	Short: "Manipulate worktracker customers",
	Long:  `All software has versions. This is Worktrackers's`,
}

var CustCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create customer with given name",
	Long:  `All software has versions. This is Worktrackers's`,
	Run: func(cmd *cobra.Command, args []string) {
		CustomerCreate(CustomerName, Rate)
	},
}

func CustomerCreate(name string, rate int) {
	fmt.Println("Creating customer:", name, "with default rate:", rate)
}
