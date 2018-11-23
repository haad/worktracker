package customer

import (
	"fmt"

	"github.com/haad/worktracker/model/project"
	"github.com/haad/worktracker/sql"
)

type CustomerInt interface {
	GetID() uint
	GetName() string
	GetContactEmail() string
	GetContactName() string
	GetRate() uint
}

func CustomerCreate(name string, rate uint, contact string, email string) {
	fmt.Println("Creating customer:", name, "with default rate:", rate)
	sql.DBc.Create(&sql.Customer{Name: name, Rate: rate, ContactEmail: email, ContactName: contact})
}

func CustomerDelete(name string) {
	var customer sql.Customer

	sql.DBc.Set("gorm:auto_preload", true).Where("name = ?", name).First(&customer)

	fmt.Println("Deleting projects:")
	for _, p := range customer.Projects {
		project.ProjectDelete(p.GetID())
		//		sql.DBc.Unscoped().Delete(&project)
	}

	fmt.Println("Deleting customer:", name)
	sql.DBc.Unscoped().Delete(&customer)
}

func CustomerList() []CustomerInt {
	var customers []sql.Customer
	var cint []CustomerInt

	sql.DBc.Set("gorm:auto_preload", true).Find(&customers)

	for _, c := range customers {
		cint = append(cint, c)
	}

	return cint
}
