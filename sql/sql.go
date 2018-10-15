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

var DBc *gorm.DB

type Customer struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Rate uint

	Projects []Project `gorm:"foreignkey:ProjRef"`

	ContactName  string
	ContactEmail string
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

func DbInit(DbType string, DbPath string) {
	var err error

	fmt.Println("Initializing Database...")
	DBc, err = gorm.Open(DbType, DbPath)
	if err != nil {
		DBc.Close()
		panic("DB open Failed")
	}
	//	defer DBc.Close()

	// Migrate the schema
	fmt.Println("Running database automigration...")
	DBc.AutoMigrate(&Customer{}, &Project{}, &Entry{}, &Tag{})

}

func DBPreload() {
	var customer Customer

	if err := DBc.Where("name = ?", "Cra").First(&customer).Error; err != nil {
		fmt.Println("Preloading database data")
		DBc.Create(&Customer{Name: "Cra", Rate: 40, ContactEmail: "pkouril@cra.cz", ContactName: "Premysl Kouril"})
		DBc.Create(&Customer{Name: "Pixel", Rate: 30, ContactEmail: "mderer@pixelfederation.com", ContactName: "Marek Derer"})
		DBc.Create(&Customer{Name: "IB", Rate: 40, ContactEmail: "wth@ibuildings.it", ContactName: ""})
		DBc.Create(&Customer{Name: "Freal", Rate: 40, ContactEmail: "jbombiak@vederie.sk", ContactName: "Jozef Bombiak"})
	} else {
		fmt.Println("Database data loaded already")
	}
}
