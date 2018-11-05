package entry

import (
	"fmt"
	"github.com/haad/worktracker/sql"
)

func EntCreate(name string, desc string, dura int64, projectName string, customerName string, billable bool, tags string) {
	var project sql.Project
	var customer sql.Customer

	sql.GetCustomerByName(customerName, &customer)

	sql.DBc.Where("name = ? AND CustomerID = ?", projectName, customer.GetID()).First(&project)
	fmt.Println("Creating entry:", name, "with desc:", desc, "duration was:", dura,
		"Entry belongs to :", projectName, "and is billable:", billable, "with tags: ", tags)

	sql.DBc.Create(&sql.Entry{Name: name, Desc: desc, Duration: dura, Billable: billable, ProjectID: project.ID})
}

func EntDelete(id uint) {
	fmt.Println("Deleting entry:", id)
}
