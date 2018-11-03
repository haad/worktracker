package sql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// Gorm Documentation
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DBc connection to our local DB
var DBc *gorm.DB

// Customer definitions with it's getters
type Customer struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Rate uint

	Projects []Project `gorm:"foreignkey:CustomerID"`

	ContactName  string
	ContactEmail string
}

// GetName getter for customer name
func (c Customer) GetName() string {
	return c.Name
}

// GetRate getter for customer rate
func (c Customer) GetRate() uint {
	return c.Rate
}

// GetContactName getter for customer contact
func (c Customer) GetContactName() string {
	return c.ContactName
}

// GetContactEmail getter for customer email
func (c Customer) GetContactEmail() string {
	return c.ContactEmail
}

// Project definitions with it's getters
type Project struct {
	gorm.Model
	Name string

	//Customer   Customer
	CustomerID uint

	Entries []Entry `gorm:"foreignkey:ProjectID"`
}

// GetCustomerName for instance of a project
func (p Project) GetCustomerName() string {
	var customer Customer

	DBc.Where("ID = ?", p.CustomerID).Find(&customer)

	return customer.Name
}

// GetName getter for project Name
func (p Project) GetName() string {
	return p.Name
}

// Entry definitions with it's getters
type Entry struct {
	gorm.Model
	Name     string
	Duration int64
	Started  int64
	Ended    int64
	Desc     string
	Billable bool

	ProjectID uint

	//Tags []*Tag `gorm:"many2many:entry_tags;"`
}

//type Tag struct {
//	gorm.Model
//	Name string `gorm:"unique"`
//
//	Entries []*Entry `gorm:"many2many:entry_tags;"`
//}

// DBInit initialize database connection and setups GORM
func DBInit(DbType string, DbPath string) {
	var err error

	fmt.Println("Initializing Database...")
	DBc, err = gorm.Open(DbType, DbPath)
	if err != nil {
		DBc.Close()
		panic("DB open Failed")
	}
	//	defer db.Close()

	// Migrate the schema
	fmt.Println("Running database automigration...")
	DBc.AutoMigrate(&Customer{}, &Project{}, &Entry{})

	DBc.LogMode(false)
}

// DBPreload will load db with preload data
// TODO: Make sure we load only data which is needed when it's needed.
func DBPreload() {

	fmt.Println("Preloading database data if needed...")

	DBc.Where(Customer{Name: "Cra"}).FirstOrCreate(&Customer{Name: "Cra", Rate: 40,
		ContactEmail: "pkouril@cra.cz", ContactName: "Premysl Kouril"})
	DBc.Where(Customer{Name: "Pixel"}).FirstOrCreate(&Customer{Name: "Pixel", Rate: 37,
		ContactEmail: "bbernat@pixelfederation.com", ContactName: "Branislav Bernat"})
	DBc.Where(Customer{Name: "IB"}).FirstOrCreate(&Customer{Name: "IB", Rate: 40,
		ContactEmail: "alessandra@ibuildings.it", ContactName: "Alesandra Pretromilli"})
	DBc.Where(Customer{Name: "Freal"}).FirstOrCreate(&Customer{Name: "Freal", Rate: 40,
		ContactEmail: "jbombiak@vederie.sk", ContactName: "Jozef Bombiak"})
}
