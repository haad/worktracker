package entry

import (
	"fmt"
	"github.com/haad/worktracker/sql"
	"github.com/haad/worktracker/wtime"
	"time"
)

type EntryInt interface {
	GetID() uint
	GetName() string
	GetDesc() string
	GetDuration() int64
	GetSDate() int64
	GetSDateString() string
	GetDurationString() string
	GetCustomerName() string
	GetProjectName() string
}

func EntCreate(name string, desc string, dura string, startDate string, projectName string, customerName string,
	billable bool, tags string) {
	var project sql.Project

	fmt.Println(startDate, sql.ShortForm)
	if startDate == "" {
		startDate = time.Now().Format(sql.ShortForm)
	}

	startD, err := time.Parse(sql.ShortForm, startDate)
	if err != nil {
		panic(err)
	}

	d, err := time.ParseDuration(dura)
	if err != nil {
		panic(err)
	}

	if err := sql.GetProjectByName(customerName, projectName, &project); err != nil {
		fmt.Println("Project: ", projectName, "with customer: ", customerName, "not found. Error:", err.Error())
		return
	}

	fmt.Println("Creating entry for a project: ", project.GetName(), "on customer: ", project.GetCustomerName())
	sql.DBc.Create(&sql.Entry{Name: name, Desc: desc, StartDate: startD.Unix(), Duration: int64(d.Seconds()), Billable: billable, ProjectID: project.ID})
}

func EntDelete(id uint) {
	var entry sql.Entry

	sql.DBc.Where("ID = ?", id).First(&entry)

	fmt.Println("Deleting entry:", entry.GetName())
	sql.DBc.Unscoped().Delete(&entry)
}

func EntList(projectName string, customerName string, startDate string) []EntryInt {
	var entries []sql.Entry
	var eint []EntryInt

	var project sql.Project

	if projectName != "" && customerName != "" {
		if err := sql.GetProjectByName(customerName, projectName, &project); err != nil {
			fmt.Println("Project: ", projectName, "with customer: ", customerName, "not found. Error:", err.Error())
			// XXX: fix error handling
			panic("")
		}

		sql.DBc.Set("gorm:auto_preload", true).Where("project_id = ?", project.GetID()).Find(&entries)
	} else {
		sql.DBc.Set("gorm:auto_preload", true).Find(&entries)
	}

	for _, e := range entries {
		// if startDate fileter string was provided use it, to fileter entries visible to user
		if startDate != "" {
			if wtime.CompareStartDate(startDate, e.GetSDate()) {
				eint = append(eint, e)
			} else {
				continue
			}
		} else {
			eint = append(eint, e)
		}
	}

	return eint
}
