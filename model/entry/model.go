package entry

import (
	"fmt"
	"github.com/haad/worktracker/sql"
)

type EntryInt interface {
	GetID() uint
	GetName() string
	GetDesc() string
	GetDuration() int64
	GetCustomerName() string
	GetProjectName() string
}

func EntCreate(name string, desc string, dura int64, projectName string, customerName string, billable bool, tags string) {
	var project sql.Project

	fmt.Println("P:", projectName, "C:", customerName)
	if err := sql.GetProjectByName(customerName, projectName, &project); err != nil {
		fmt.Println("Project: ", projectName, "with customer: ", customerName, "not found. Error:", err.Error())
		return
	}

	fmt.Println("Creating entry for a project: ", project, "on customer: ", project.GetCustomerName())
	sql.DBc.Create(&sql.Entry{Name: name, Desc: desc, Duration: dura, Billable: billable, ProjectID: project.ID})
}

func EntDelete(id uint) {
	var entry sql.Entry

	sql.DBc.Where("ID = ?", id).First(&entry)

	fmt.Println("Deleting entry:", entry.GetName())
	sql.DBc.Unscoped().Delete(&entry)
}

func EntList() []EntryInt {
	var entries []sql.Entry
	var eint []EntryInt

	sql.DBc.Set("gorm:auto_preload", true).Find(&entries)

	for _, e := range entries {
		fmt.Println("Project ID: ", e.ProjectID)
		eint = append(eint, e)
	}

	return eint
}
