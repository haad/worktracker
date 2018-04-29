package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/ogier/pflag"
)

type Customer struct {
	gorm.Model
	Name        string
	DefaultRate uint
	Projects    []Project
}

type Project struct {
	gorm.Model
	Name        string
	Rate        uint
	CustomerID  uint
	Customer    Customer
	WorkEntries []WorkEntry
}

type WorkEntry struct {
	gorm.Model
	Name        string
	Description string
	Hours       uint
	ProjectID   uint
	Project     Project
}

func main() {

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Program %s version:\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		pflag.PrintDefaults()
	}

	// CLI flags
	var _create = pflag.BoolP("create", "C", true, "Create flag.")
	var _delete = pflag.BoolP("delete", "D", false, "Delete flag.")
	var _update = pflag.BoolP("update", "U", false, "Update flag.")
	var _list = pflag.BoolP("list", "L", false, "List entries for project/customer/workentries flag.")
	var _dump = pflag.BoolP("dump", "M", false, "Dump whole dbflag.")

	var _customer_name = pflag.StringP("customer", "c", "", "Create customer if doesn't exist.")
	var _project_name = pflag.StringP("project", "p", "", "Create project if doesn't exist.")
	var _workentry_name = pflag.StringP("workentry", "w", "", "Create workentry if doesn't exist.")

	//var _rate = pflag.Uint16P("rate", "r", 0, "Set work hour rate to given value it's related to customer/project")

	// Parse the CLI flags
	pflag.Parse()

	if *_create {
		if *_customer_name != "" {
			log.Printf("customer name passed %s", *_customer_name)
		}
		if *_project_name != "" {
			log.Printf("project name passed %s", *_project_name)
		}
		if *_workentry_name != "" {
			log.Printf("workentry name passed %s", *_workentry_name)
		}
		log.Printf("create flag passed")
	}

	if *_delete {
		log.Printf("delete flag passed")
	}

	if *_update {
		log.Printf("update flag passed")
	}

	if *_list {
		log.Printf("list flag passed")
	}

	if *_dump {
		log.Printf("dump flag passed")
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Customer{}, &Project{}, &WorkEntry{})
	db.Close()
}
