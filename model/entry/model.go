package entry

import (
	"log"
	"time"

	"github.com/haad/worktracker/sql"
	"github.com/haad/worktracker/wtime"
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

func EntCreate(name string, desc string, dura string, startDate string, projectID uint,
	billable bool, tags string) {
	var project sql.Project

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

	if err := sql.GetProjectByID(projectID, &project); err != nil {
		log.Println("Project with id: ", projectID, "not found. Error:", err.Error())
		return
	}

	log.Println("Creating entry for a project: ", project.GetName(), "on customer: ", project.GetCustomerName())
	sql.DBc.Create(&sql.Entry{Name: name, Desc: desc, StartDate: startD.Unix(), Duration: int64(d.Seconds()), Billable: billable, ProjectID: project.ID})
}

func EntDelete(id uint) {
	var entry sql.Entry

	sql.DBc.Where("ID = ?", id).First(&entry)

	log.Println("Deleting entry:", entry.GetName())
	sql.DBc.Unscoped().Delete(&entry)
}

func EntList(projectID uint, startDate string) []EntryInt {
	var entries []sql.Entry
	var eint []EntryInt

	// Try to select all entries for a given projectID if it doesn't exist print all
	if sql.DBc.Set("gorm:auto_preload", true).Where("project_id = ?", projectID).Find(&entries).RecordNotFound() {
		log.Println("Entries with project id: ", projectID, "not found.")

		sql.DBc.Set("gorm:auto_preload", true).Find(&entries)
	}

	for _, e := range entries {
		// if startDate filter string was provided use it, to filter entries visible to user
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
