package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eleazar-harold/employee-service/api/app/handlers"
	"github.com/eleazar-harold/employee-service/api/app/services"
	"github.com/eleazar-harold/employee-service/api/config"
	"github.com/gorilla/mux"
)

// App has router and db instances
type App struct {
	sm *mux.Router
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// Initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
	)

	log.Printf("DB URI: %s\n", psqlInfo)

	es, err := services.NewEmployeeService(psqlInfo)
	must(err)

	defer es.Close()
	es.AutoMigrate()

	employeeH := handlers.NewEmployees(es)

	a.sm = mux.NewRouter()

	// Routing for handling the projects
	a.sm.Methods(http.MethodGet).Subrouter().HandleFunc("/employees", employeeH.GetAllEmployees)
	a.sm.Methods(http.MethodPost).Subrouter().HandleFunc("/employees", employeeH.CreateEmployee)
	a.sm.Methods(http.MethodGet).Subrouter().HandleFunc("/employees/{title}", employeeH.GetEmployee)
	a.sm.Methods(http.MethodPut).Subrouter().HandleFunc("/employees/{title}", employeeH.UpdateEmployee)
	a.sm.Methods(http.MethodDelete).Subrouter().HandleFunc("/employees/{title}", employeeH.DeleteEmployee)
	a.sm.Methods(http.MethodPut).Subrouter().HandleFunc("/employees/{title}/disable", employeeH.DisableEmployee)
	a.sm.Methods(http.MethodPut).Subrouter().HandleFunc("/employees/{title}/enable", employeeH.EnableEmployee)

	http.ListenAndServe(":"+os.Getenv("APPORT"), a.sm)
}
