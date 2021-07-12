package domain

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID         uint64 `gorm:"primary_key"`
	Name       string
	Email      string
	Nick       string
	Document   string
	Active     bool
	LeiCompany string
}
