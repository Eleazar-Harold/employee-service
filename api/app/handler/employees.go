package handler

import (
	"encoding/json"
	"net/http"

	"github.com/eleazar-harold/employee-service/api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAllEmployees func
func GetAllEmployees(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employees := []model.Employee{}
	db.Find(&employees)
	respondJSON(w, http.StatusOK, employees)
}

// CreateEmployee func
func CreateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employee := model.Employee{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, employee)
}

// GetEmployee func
func GetEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

// UpdateEmployee func
func UpdateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

// DeleteEmployee func
func DeleteEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	if err := db.Delete(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// DisableEmployee func
func DisableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	employee.Disable()
	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

// EnableEmployee func
func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	employee.Enable()
	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

// getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
func getEmployeeOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Employee {
	employee := model.Employee{}
	if err := db.First(&employee, model.Employee{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &employee
}
