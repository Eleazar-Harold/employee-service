package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Employee model struct
type Employee struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
}

// Disable employee func
func (e *Employee) Disable() {
	e.Status = false
}

// Enable employee func
func (e *Employee) Enable() {
	e.Status = true
}
