package project

import (
	"fmt"
	"github.com/haad/worktracker/sql"
	//"github.com/haad/worktracker/model/entry"
)

type ProjectInt interface {
	GetID() uint
	GetName() string
	GetCustomerName() string
}

func ProjectCreate(name string, customerName string) {
	var customer sql.Customer
	if err := sql.GetCustomerByName(customerName, &customer); err != nil {
		fmt.Println("Customer: ", customerName, "doesn't exist. Error: ", err.Error())
		return
	}

	fmt.Println("Creating project:", name, "with default rate:", customer)
	sql.DBc.Create(&sql.Project{Name: name, CustomerID: customer.GetID()})
}

func ProjectDelete(id uint) {
	var project sql.Project

	sql.DBc.Where("ID = ?", id).First(&project)

	fmt.Println("Deleting project:", project.GetName())
	sql.DBc.Unscoped().Delete(&project)
}

func ProjectList(customerName string) []ProjectInt {
	var projects []sql.Project
	var pint []ProjectInt
	var customer sql.Customer

	if customerName != "" {
		if err := sql.GetCustomerByName(customerName, &customer); err != nil {
			fmt.Println("Customer: ", customerName, "doesn't exist. Error: ", err.Error())
			panic("")
		}

		sql.DBc.Set("gorm:auto_preload", true).Where("customer_id = ?", customer.GetID()).Find(&projects)
	} else {
		sql.DBc.Set("gorm:auto_preload", true).Find(&projects)
	}

	for _, p := range projects {
		pint = append(pint, p)
	}

	return pint
}
