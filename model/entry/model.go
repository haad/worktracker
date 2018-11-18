package entry

import (
	"fmt"
	"github.com/haad/worktracker/sql"
	"time"
	//"github.com/haad/worktracker/time"
)

type EntryInt interface {
	GetID() uint
	GetName() string
	GetDesc() string
	GetDuration() int64
	GetDurationString() string
	GetCustomerName() string
	GetProjectName() string
}

func EntCreate(name string, desc string, dura string, projectName string, customerName string,
	billable bool, tags string) {
	var project sql.Project
	var err error

	d, err := time.ParseDuration(dura)
	if err != nil {
		panic(err)
	}

	if err := sql.GetProjectByName(customerName, projectName, &project); err != nil {
		fmt.Println("Project: ", projectName, "with customer: ", customerName, "not found. Error:", err.Error())
		return
	}

	fmt.Println("Creating entry for a project: ", project, "on customer: ", project.GetCustomerName())
	sql.DBc.Create(&sql.Entry{Name: name, Desc: desc, Duration: int64(d.Seconds()), Billable: billable, ProjectID: project.ID})
}

func EntDelete(id uint) {
	var entry sql.Entry

	sql.DBc.Where("ID = ?", id).First(&entry)

	fmt.Println("Deleting entry:", entry.GetName())
	sql.DBc.Unscoped().Delete(&entry)
}

func EntList(projectName string, customerName string) []EntryInt {
	var entries []sql.Entry
	var eint []EntryInt

	var project sql.Project

	if projectName != "" && customerName != "" {
		if err := sql.GetProjectByName(customerName, projectName, &project); err != nil {
			fmt.Println("Project: ", projectName, "with customer: ", customerName, "not found. Error:", err.Error())
			panic("")
		}

		sql.DBc.Set("gorm:auto_preload", true).Where("project_id = ?", project.GetID()).Find(&entries)
	} else {
		sql.DBc.Set("gorm:auto_preload", true).Find(&entries)
	}

	for _, e := range entries {
		eint = append(eint, e)
	}

	return eint
}
