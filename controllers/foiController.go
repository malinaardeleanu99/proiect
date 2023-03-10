package controllers

import (
	"example/API/initializers"
	"example/API/middleware"
	"example/API/models"
	"time"

	"github.com/gin-gonic/gin"
)

const formatDate = "2006-01-02"
const formatTime = "15:04"

func FoiCreate(c *gin.Context) {
	//get data off req body
	var body struct {
		Data          string
		Sofer         string
		Proiect       string
		FirmaPres     string
		Auto          string
		Marca         string
		IndexAlim     int64
		Schimb        string
		CantitateAlim int64
		Autoturism    string
		Observatii    string
		Status        string
		Curse         []models.Cursa
	}

	c.Bind(&body)

	//parse date and time
	data_parsed, err := time.Parse(formatDate, body.Data)
	if err != nil {
		c.Status(500)
	}
	var CurseParsed []models.Cursa
	var cursa models.Cursa

	for i := 0; i < len(body.Curse); i++ {
		oraStart, err1 := time.Parse(formatTime, body.Curse[i].OraStart)
		oraStop, err2 := time.Parse(formatTime, body.Curse[i].OraStop)

		if err1 != nil || err2 != nil {
			c.Status(500)
		}
		cursa = models.Cursa{
			Start:     body.Curse[i].Start,
			Stop:      body.Curse[i].Stop,
			Scop:      body.Curse[i].Scop,
			KmPlecare: body.Curse[i].KmPlecare,
			KmSosire:  body.Curse[i].KmSosire,
			OraStart:  oraStart.String(),
			OraStop:   oraStop.String(),
			Marfa:     body.Curse[i].Marfa,
			FoaieID:   body.Curse[i].FoaieID,
		}
		CurseParsed = append(CurseParsed, cursa)
	}

	userID := middleware.CheckUser(c)

	//create foaie
	foaie := models.Foaie{Data: data_parsed,
		Sofer:         body.Sofer,
		Proiect:       body.Proiect,
		FirmaPres:     body.FirmaPres,
		Auto:          body.Auto,
		Marca:         body.Marca,
		IndexAlim:     body.IndexAlim,
		Schimb:        body.Schimb,
		CantitateAlim: body.CantitateAlim,
		Autoturism:    body.Autoturism,
		Observatii:    body.Observatii,
		Status:        body.Status,
		Curse:         CurseParsed,
		Autor:         userID,
	}

	result := initializers.DB.Create(&foaie)

	//respond
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"foaie": foaie,
	})
}

func FoiIndex(c *gin.Context) {
	//get the foi
	var foi []models.Foaie
	var curse []models.Cursa
	initializers.DB.Find(&foi)

	for i := 0; i < len(foi); i++ {
		id := foi[i].ID
		initializers.DB.Find(&curse, "foaie_id = ?", id)
		foi[i].Curse = curse
		curse = nil
	}

	//respond
	c.JSON(200, gin.H{
		"foi": foi,
	})
}

func FoiUserIndex(c *gin.Context) {
	//get the foi
	var foi []models.Foaie
	var curse []models.Cursa

	userID := middleware.CheckUser(c)

	initializers.DB.Find(&foi, "autor = ?", userID)

	for i := 0; i < len(foi); i++ {
		id := foi[i].ID
		initializers.DB.Find(&curse, "foaie_id = ?", id)
		foi[i].Curse = curse
		curse = nil
	}

	//respond
	c.JSON(200, gin.H{
		"foi": foi,
	})
}

func FoiShow(c *gin.Context) {
	//get id off urls
	id := c.Param("ID")

	var foaie models.Foaie
	var curse []models.Cursa

	initializers.DB.First(&foaie, id)
	initializers.DB.Find(&curse, "foaie_id = ?", id)
	foaie.Curse = curse

	c.JSON(200, gin.H{
		"foaie": foaie,
	})

}

func FoaieUpdate(c *gin.Context) {
	id := c.Param("ID")

	var body struct {
		Data          string
		Sofer         string
		Proiect       string
		FirmaPres     string
		Auto          string
		Marca         string
		IndexAlim     int64
		Schimb        string
		CantitateAlim int64
		Autoturism    string
		Observatii    string
		Status        string
		Curse         []models.Cursa
	}

	c.Bind(&body)
	data_parsed, err := time.Parse(formatDate, body.Data)
	if err != nil {
		c.Status(500)
	}

	var foaie models.Foaie
	var curse []models.Cursa

	initializers.DB.First(&foaie, id)

	initializers.DB.Find(&curse, "foaie_id = ?", id)

	for i := 0; i < len(body.Curse); i++ {
		initializers.DB.Model(&curse[i]).Updates(models.Cursa{
			Start:     body.Curse[i].Start,
			Stop:      body.Curse[i].Stop,
			Scop:      body.Curse[i].Scop,
			KmPlecare: body.Curse[i].KmPlecare,
			KmSosire:  body.Curse[i].KmSosire,
			OraStart:  body.Curse[i].OraStart,
			OraStop:   body.Curse[i].OraStop,
			Marfa:     body.Curse[i].Marfa,
			FoaieID:   body.Curse[i].FoaieID,
		})
	}

	initializers.DB.Model(&foaie).Updates(models.Foaie{
		Data:          data_parsed,
		Sofer:         body.Sofer,
		Proiect:       body.Proiect,
		FirmaPres:     body.FirmaPres,
		Auto:          body.Auto,
		Marca:         body.Marca,
		IndexAlim:     body.IndexAlim,
		Schimb:        body.Schimb,
		CantitateAlim: body.CantitateAlim,
		Autoturism:    body.Autoturism,
		Observatii:    body.Observatii,
		Status:        body.Status,
		Curse:         body.Curse})
	c.JSON(200, gin.H{
		"foaie": foaie,
	})
}

func FoiDelete(c *gin.Context) {
	//get id off urls
	id := c.Param("ID")

	initializers.DB.Delete(&models.Foaie{}, id)
	initializers.DB.Delete(&models.Cursa{}, "foaie_id = ?", id)

	c.Status(200)
}
