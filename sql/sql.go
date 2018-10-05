package sql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const WTYPE_CHORE = 0
const WTYPE_FUN = 1
const WTYPE_OSS = 2

const WTYPE_WORK = 3
const WTYPE_WORK_MEET = 4
const WTYPE_WORK_SUPP = 5

type Customer struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Rate uint

	Projects []Project `gorm:"foreignkey:ProjRef"`

	ContactName  string `gorm:"unique"`
	ContactEmail string `gorm:"unique"`
}

type Project struct {
	gorm.Model
	Name string

	ProjRef uint

	Customer Customer `gorm:"foreignkey:CustRef"`
	CustRef  uint

	Entries []Entry `gorm:"foreignkey:EntryRef"`
}

type Entry struct {
	gorm.Model
	Name     string
	Duration int64
	Started  int64
	Ended    int64
	Type     uint
	Desc     string
	Billable bool

	EntryRef uint

	Project Project `gorm:"foreignkey:ProjRef"`
	ProjRef uint

	Tags []*Tag `gorm:"many2many:entry_tags;"`
}

type Tag struct {
	gorm.Model
	Name string `gorm:"unique"`

	Entries []*Entry `gorm:"many2many:entry_tags;"`
}

func DbInit(DbType string, DbPath string) (*gorm.DB, error) {
	fmt.Println("Initializing Database...")
	db, err := gorm.Open(DbType, DbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Migrate the schema
	fmt.Println("Running database automigration...")
	db.AutoMigrate(&Customer{}, &Project{}, &Entry{}, &Tag{})

	return db, err
}
