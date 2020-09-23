package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eleazar-harold/employee-service/api/app/models"
	"github.com/eleazar-harold/employee-service/api/app/services"
	"github.com/gorilla/mux"
)

/*
NewEmployees func used to create a new users controller
This function will if templates are not parsed correctly
and should only be used during initial setup.
*/
func NewEmployees(es *services.EmployeeService) *Employees {
	return &Employees{
		es: es,
	}
}

/*
Employees struct type
*/
type Employees struct {
	es *services.EmployeeService
}

// GetAllEmployees func
func (e *Employees) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees := e.es.GetAll()
	respondJSON(w, http.StatusOK, employees)
}

// CreateEmployee func
func (e *Employees) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := e.es.Create(&employee); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, employee)
}

// GetEmployee func
func (e *Employees) GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee, err := e.es.ByName(name)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
	respondJSON(w, http.StatusOK, employee)
}

// UpdateEmployee func
func (e *Employees) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee, err := e.es.ByName(name)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := e.es.Update(employee); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

// DeleteEmployee func
func (e *Employees) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee, err := e.es.ByName(name)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := e.es.Delete(employee.ID); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// DisableEmployee func
func (e *Employees) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee, err := e.es.ByName(name)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	employee.Disable()
	if err := e.es.Update(employee); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

// EnableEmployee func
func (e *Employees) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee, err := e.es.ByName(name)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	employee.Enable()
	if err := e.es.Update(employee); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}
