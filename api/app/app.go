package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eleazar-harold/employee-service/api/app/handler"
	"github.com/eleazar-harold/employee-service/api/app/model"
	"github.com/eleazar-harold/employee-service/api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Name,
		config.DB.Password,
	)

	log.Printf("DB URI: %s\n", dsn)

	a.DB, err = gorm.Open(config.DB.Dialect, dsn)
	if err != nil {
		log.Printf("Could not connect to %s database\n", config.DB.Dialect)
	} else {
		log.Printf("We are connected to the %s database\n", config.DB.Dialect)
	}

	a.DB.AutoMigrate(&model.Employee{}) //database migration
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{title}", a.GetEmployee)
	a.Put("/employees/{title}", a.UpdateEmployee)
	a.Delete("/employees/{title}", a.DeleteEmployee)
	a.Put("/employees/{title}/disable", a.DisableEmployee)
	a.Put("/employees/{title}/enable", a.EnableEmployee)
}

// Get wrapper method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.Methods(http.MethodGet).Subrouter().HandleFunc(path, f)
}

// Post wrapper method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.Methods(http.MethodPost).Subrouter().HandleFunc(path, f)
}

// Put wrapper method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.Methods(http.MethodPut).Subrouter().HandleFunc(path, f)
}

// Delete wrapper method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.Methods(http.MethodDelete).Subrouter().HandleFunc(path, f)
}

// GetAllEmployees handler
func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEmployees(a.DB, w, r)
}

// CreateEmployee handler
func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.CreateEmployee(a.DB, w, r)
}

// GetEmployee handler
func (a *App) GetEmployee(w http.ResponseWriter, r *http.Request) {
	handler.GetEmployee(a.DB, w, r)
}

// UpdateEmployee handler
func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.UpdateEmployee(a.DB, w, r)
}

// DeleteEmployee handler
func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DeleteEmployee(a.DB, w, r)
}

// DisableEmployee handler
func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DisableEmployee(a.DB, w, r)
}

// EnableEmployee handler
func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.EnableEmployee(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
