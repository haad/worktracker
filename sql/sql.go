package sql

import (
	"fmt"
	"strconv"
	"time"

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

// GetID getter for customer ID
func (c Customer) GetID() uint {
	return c.ID
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

// GetCustomerByName return customer if exists
func GetCustomerByName(customerName string, customer *Customer) error {
	if DBc.Where("name = ?", customerName).First(customer).RecordNotFound() {
		//fmt.Println("Customer: ", customerName, "doesn't exist")
		return fmt.Errorf("customer %s doesn't exists", customerName)
	}
	return nil
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

	DBc.Where("id = ?", p.CustomerID).Find(&customer)

	return customer.Name
}

// GetCustomerRate for instance of a project
func (p Project) GetCustomerRate() uint {
	var customer Customer

	DBc.Where("id = ?", p.CustomerID).Find(&customer)

	return customer.Rate
}

// GetName getter for project Name
func (p Project) GetName() string {
	return p.Name
}

// GetID getter for project ID
func (p Project) GetID() uint {
	return p.ID
}

// GetProjectByName searchs for a project by it's name and customer's name to whitch it belongs
func GetProjectByName(customerName string, projectName string, project *Project) error {
	var customer Customer

	if err := GetCustomerByName(customerName, &customer); err != nil {
		return fmt.Errorf("customer %s doesn't exists", customerName)
	}

	if DBc.Where("name = ? AND customer_id = ?", projectName, customer.GetID()).First(project).RecordNotFound() {
		return fmt.Errorf("project name %s for customer %s is missing ", projectName, customerName)
	}

	return nil
}

// Entry definitions with it's Byters
type Entry struct {
	gorm.Model
	Name      string
	Duration  int64
	StartDate int64
	EndDate   int64
	Desc      string
	Billable  bool

	ProjectID uint

	Tags []*Tag `gorm:"many2many:entry_tags;"`
}

// GetName getter for entry Name
func (e Entry) GetName() string {
	return e.Name
}

// GetID getter for entry ID
func (e Entry) GetID() uint {
	return e.ID
}

// GetDesc getter for entry
func (e Entry) GetDesc() string {
	return e.Desc
}

// GetSDate getter for entry
func (e Entry) GetSDate() int64 {
	return e.StartDate
}

// GetDuration getter for entry
func (e Entry) GetDuration() int64 {
	return e.Duration
}

// GetDuration getter for entry
func (e Entry) GetDurationString() string {
	d, err := time.ParseDuration(strconv.FormatInt(e.Duration, 10) + "s")

	if err != nil {
		return ""
	}
	return d.String()
}

// GetProjectName for instance of a project
func (e Entry) GetProjectName() string {
	var project Project

	DBc.Where("id = ?", e.ProjectID).Find(&project)

	return project.Name
}

// GetCustomerName for instance of a project
func (e Entry) GetCustomerName() string {
	var project Project

	DBc.Where("id = ?", e.ProjectID).Find(&project)

	return project.GetCustomerName()
}

type Tag struct {
	gorm.Model
	Name string `gorm:"unique"`

	Entries []*Entry `gorm:"many2many:entry_tags;"`
}

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
