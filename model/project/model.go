package project

import (
	"fmt"
	"github.com/haad/worktracker/sql"
	//"github.com/haad/worktracker/model/entry"
)

type ProjectInt interface {
	GetName() string
	GetCustomerName() string
}

func ProjectCreate(name string, customerName string) {
	var customer sql.Customer
	sql.DBc.Where("name = ?", customerName).First(&customer)

	fmt.Println("Creating project:", name, "with default rate:", customer)
	sql.DBc.Create(&sql.Project{Name: name, CustomerID: customer.ID})
}

func ProjectDelete(name string) {
	var project sql.Project
	sql.DBc.Where("name = ?", name).First(&project)

	fmt.Println("Deleting project:", name)
	sql.DBc.Unscoped().Delete(&project)
}

func ProjectList() []ProjectInt {
	var projects []sql.Project
	var pint []ProjectInt

	sql.DBc.Set("gorm:auto_preload", true).Find(&projects)

	for _, p := range projects {
		pint = append(pint, p)
	}

	return pint
}
