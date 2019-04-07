package project

import (
	"log"
	"time"

	"github.com/haad/worktracker/sql"
)

type ProjectInt interface {
	GetID() uint
	GetName() string
	GetEstimate() int64
	GetEstimateString() string
	GetWorkLoggedString() string
	GetCustomerName() string
	GetFinished() bool
}

func ProjectCreate(name string, estimate string, customerName string) {
	var customer sql.Customer
	var est int64 = 0

	if err := sql.GetCustomerByName(customerName, &customer); err != nil {
		log.Println("Customer: ", customerName, "doesn't exist. Error: ", err.Error())
		return
	}

	if estimate != "" {
		d, err := time.ParseDuration(estimate)
		if err != nil {
			panic(err)
		}

		est = int64(d.Seconds())
	}

	log.Println("Creating project:", name, "under customer:", customer)
	sql.DBc.Create(&sql.Project{Name: name, Estimate: est, CustomerID: customer.GetID()})
}

func ProjectEdit(id uint, name string, estimate string, finished bool) {
	var project sql.Project

	sql.DBc.Set("gorm:auto_preload", true).Where("ID = ?", id).First(&project)

	if estimate != "" {
		d, err := time.ParseDuration(estimate)
		if err != nil {
			panic(err)
		}

		project.Estimate = int64(d.Seconds())
	}

	if name != "" {
		project.Name = name
	}

	project.Finished = finished

	log.Println("Editing project:", name, "with estimate:", estimate)
	sql.DBc.Save(&project)
}

func ProjectDelete(id uint) {
	var project sql.Project

	sql.DBc.Set("gorm:auto_preload", true).Where("ID = ?", id).First(&project)

	for _, entry := range project.Entries {
		log.Println("Deleting entry:", entry.GetName())
		sql.DBc.Unscoped().Delete(&entry)
	}

	log.Println("Deleting project:", project.GetName())
	sql.DBc.Unscoped().Delete(&project)
}

func ProjectList(customerName string, includeFinished bool) []ProjectInt {
	var projects []sql.Project
	var pint []ProjectInt
	var customer sql.Customer

	if customerName != "" {
		if err := sql.GetCustomerByName(customerName, &customer); err != nil {
			log.Fatalln("Customer: ", customerName, "doesn't exist. Error: ", err.Error())
		}

		sql.DBc.Set("gorm:auto_preload", true).Where("customer_id = ?", customer.GetID()).Find(&projects)
	} else {
		sql.DBc.Set("gorm:auto_preload", true).Find(&projects)
	}

	for _, p := range projects {
		if !p.GetFinished() || includeFinished {
			pint = append(pint, p)
		}
	}

	return pint
}
