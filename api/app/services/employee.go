package services

import (
	"errors"

	"github.com/eleazar-harold/employee-service/api/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is returned when a resource cannot be found
	ErrNotFound = errors.New("models: resource not found")

	// ErrInvalidID is returned when an invalid ID is provided
	// to a method like Delete.
	ErrInvalidID = errors.New("models: ID provided was invalid")
)

// EmployeeService struct
type EmployeeService struct {
	db *gorm.DB
}

// NewEmployeeService func
func NewEmployeeService(connectionInfo string) (*EmployeeService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	return &EmployeeService{db: db}, nil
}

// ByID will look up a employee by id provided
// 1 - employee, nil
// 2 - nil, ErrNoFound
// 3 - nil, otherError
func (es *EmployeeService) ByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	db := es.db.Where("id = ?", id)
	err := first(db, &employee)
	return &employee, err
}

// ByName will look up a employee by name provided
// 1 - employee, nil
// 2 - nil, ErrNoFound
// 3 - nil, otherError
func (es *EmployeeService) ByName(name string) (*models.Employee, error) {
	var employee models.Employee
	db := es.db.Where("name = ?", name)
	err := first(db, &employee)
	return &employee, err
}

// ByCity will look up a employee by name provided
// 1 - employee, nil
// 2 - nil, ErrNoFound
// 3 - nil, otherError
func (es *EmployeeService) ByCity(city string) (*[]models.Employee) {
	employees := []models.Employee{}
	es.db.Find(&employees).Where("city = ?", city)
	return &employees
}

// GetAll will look up all employees 
func (es *EmployeeService) GetAll() (*[]models.Employee) {
	employees := []models.Employee{}
	es.db.Find(&employees)
	return &employees
}

// first will query using the provided gorm.DB
// and it will get the first item returned
// and place it in the dst. if nothing is found
// in the query it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}


// Create will add employee into the database
func (es *EmployeeService) Create(employee *models.Employee) error {
	return es.db.Create(&employee).Error
}

// Update will update employee provisioned with employee object
func (es *EmployeeService) Update(employee *models.Employee) error {
	return es.db.Save(&employee).Error
}

// Delete will delete employee with provided ID
func (es *EmployeeService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}

	employee := models.Employee{Model: gorm.Model{ID: id}}
	return es.db.Delete(&employee).Error
}

// Close the UserService database connection
func (es *EmployeeService) Close() error {
	return es.db.Close()
}

// DestructiveReset drops table and rebuilds it
func (es *EmployeeService) DestructiveReset() error {
	if err := es.db.DropTableIfExists(&models.Employee{}).Error; err != nil {
		return err
	}
	return es.AutoMigrate()
}

// AutoMigrate will attempt to Automigrate the users table
func (es *EmployeeService) AutoMigrate() error {
	if err := es.db.AutoMigrate(&models.Employee{}).Error; err != nil {
		return err
	}
	return nil
}
