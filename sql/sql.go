package sql

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// Gorm Documentation
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/haad/worktracker/wtime"
)

// DBc connection to our local DB
var DBc *gorm.DB

const ShortForm = "02/01/2006"

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
	Name     string
	Estimate int64

	CustomerID uint `json:"-"`

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

// GetEstimate getter for project Planned hours
func (p Project) GetEstimate() int64 {
	return p.Estimate
}

// GetDurationString gets duretion converted to string for entry
func (p Project) GetEstimateString() string {
	return wtime.GetDurantionString(p.Estimate)
}

func (p Project) GetWorkLoggedString() string {
	var workedTime int64 = 0

	for _, entry := range p.Entries {
		workedTime += entry.GetDuration()
	}

	return wtime.GetDurantionString(workedTime)
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

func (p Project) MarshalJSON() ([]byte, error) {
	basicProject := struct {
		ID         uint    `json:"id"`
		Name       string  `json:"name"`
		Estimate   string  `json:"estimate"`
		LoggedTime string  `json:"logged_time"`
		Entries    []Entry `json:"entries`
	}{
		ID:         p.GetID(),
		Name:       p.Name,
		Estimate:   p.GetEstimateString(),
		LoggedTime: p.GetWorkLoggedString(),
		Entries:    p.Entries,
	}

	return json.Marshal(basicProject)
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

	Tags []*Tag `gorm:"many2many:entry_tags;" json:"-"`
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

func (e Entry) GetSDateString() string {
	tm := time.Unix(e.StartDate, 0)

	return tm.Format(ShortForm)
}

// GetDuration getter for entry
func (e Entry) GetDuration() int64 {
	return e.Duration
}

// GetDurationString gets duretion converted to string for entry
func (e Entry) GetDurationString() string {
	return wtime.GetDurantionString(e.Duration)
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

func (e Entry) MarshalJSON() ([]byte, error) {
	basicEntry := struct {
		Name      string `json:"name"`
		Duration  string `json:"duration"`
		StartDate string `json:"start_date"`
		Desc      string `json:"desc"`
		Billable  bool   `json:"billable"`
	}{
		Name:      e.Name,
		Duration:  e.GetDurationString(),
		StartDate: e.GetSDateString(),
		Desc:      e.Desc,
		Billable:  e.Billable,
	}

	return json.Marshal(basicEntry)
}

type Tag struct {
	gorm.Model `json:"-"`
	Name       string `gorm:"unique"`

	Entries []*Entry `gorm:"many2many:entry_tags;"`
}

// DBInit initialize database connection and setups GORM
func DBInit(DbType string, DbPath string) {
	var err error

	fmt.Println("Initializing Database...")
	DBc, err = gorm.Open(DbType, DbPath)
	if err != nil {
		DBc.Close()
		log.Fatalln(err)
	}
	//	defer db.Close()

	// Migrate the schema
	log.Println("Running database automigration...")
	DBc.AutoMigrate(&Customer{}, &Project{}, &Entry{})

	DBc.LogMode(false)
}

// DBPreload will load db with preload data
// TODO: Make sure we load only data which is needed when it's needed.
func DBPreload() {

	log.Println("Preloading database data if needed...")

	DBc.Where(Customer{Name: "Cra"}).FirstOrCreate(&Customer{Name: "Cra", Rate: 40,
		ContactEmail: "pkouril@cra.cz", ContactName: "Premysl Kouril"})
	DBc.Where(Customer{Name: "Pixel"}).FirstOrCreate(&Customer{Name: "Pixel", Rate: 37,
		ContactEmail: "bbernat@pixelfederation.com", ContactName: "Branislav Bernat"})
	DBc.Where(Customer{Name: "Pygmalios"}).FirstOrCreate(&Customer{Name: "Pygmalios", Rate: 50,
		ContactEmail: "m.holler@pygmalios.com", ContactName: "Marian Holler"})
	DBc.Where(Customer{Name: "IB"}).FirstOrCreate(&Customer{Name: "IB", Rate: 40,
		ContactEmail: "alessandra@ibuildings.it", ContactName: "Alesandra Pretromilli"})
	DBc.Where(Customer{Name: "Freal"}).FirstOrCreate(&Customer{Name: "Freal", Rate: 40,
		ContactEmail: "jbombiak@vederie.sk", ContactName: "Jozef Bombiak"})
}
