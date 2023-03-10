package models

import (
	"time"

	"gorm.io/gorm"
)

type Foaie struct {
	gorm.Model
	Data          time.Time
	Sofer         string
	Proiect       string
	FirmaPres     string `gorm:"default:Alpenside"`
	Auto          string
	Marca         string
	IndexAlim     int64
	Schimb        string
	CantitateAlim int64
	Autoturism    string
	Observatii    string
	Status        string
	Curse         []Cursa `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; ForeignKey:FoaieID;"`
	Autor         uint
}
