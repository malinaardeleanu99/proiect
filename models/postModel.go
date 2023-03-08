package models

import (
	"gorm.io/gorm"
)

type Foaie struct {
	gorm.Model
	Data  string
	Sofer string
}
