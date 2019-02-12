package customer

import (
	"log"

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

func CustomerCreate(name string, rate int, contact string, email string) {
	log.Println("Creating customer:", name, "with default rate:", rate)
	sql.DBc.Create(&sql.Customer{Name: name, Rate: uint(rate), ContactEmail: email, ContactName: contact})
}

func CustomerEdit(name string, rate int, contact string, email string) {
	var customer sql.Customer

	log.Println("Editing customer:", name, "with default rate:", rate)
	sql.DBc.Set("gorm:auto_preload", true).Where("name = ?", name).First(&customer)

	if rate >= 0 && uint(rate) != customer.GetRate() {
		customer.Rate = uint(rate)
	}

	if email != "" && email != customer.GetContactEmail() {
		customer.ContactEmail = email
	}

	if contact != "" && contact != customer.GetContactName() {
		customer.ContactName = contact
	}

	sql.DBc.Save(&customer)
}

func CustomerDelete(name string) {
	var customer sql.Customer

	sql.DBc.Set("gorm:auto_preload", true).Where("name = ?", name).First(&customer)

	log.Println("Deleting projects:")
	for _, p := range customer.Projects {
		project.ProjectDelete(p.GetID())
		//		sql.DBc.Unscoped().Delete(&project)
	}

	log.Println("Deleting customer:", name)
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
