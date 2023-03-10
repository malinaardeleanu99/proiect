package models

import (
	"gorm.io/gorm"
)

type Cursa struct {
	gorm.Model
	Start     string
	Stop      string
	Scop      string
	KmPlecare int64
	KmSosire  int64
	OraStart  string
	OraStop   string
	Marfa     string
	FoaieID   uint
}
