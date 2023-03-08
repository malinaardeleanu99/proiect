package controllers

import (
	"example/API/initializers"
	"example/API/models"

	"github.com/gin-gonic/gin"
)

func FoiCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Data  string
		Sofer string
	}

	c.Bind(&body)

	//create foaie
	foaie := models.Foaie{Data: body.Data, Sofer: body.Sofer}

	result := initializers.DB.Create(&foaie)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return

	c.JSON(200, gin.H{
		"foaie": foaie,
	})
}

func FoiIndex(c *gin.Context) {
	//get the foi
	var foi []models.Foaie
	initializers.DB.Find(&foi)

	//respond
	c.JSON(200, gin.H{
		"foi": foi,
	})
}

func FoiShow(c *gin.Context) {
	//get id off urls
	id := c.Param("ID")

	var foaie models.Foaie
	initializers.DB.First(&foaie, id)

	c.JSON(200, gin.H{
		"foaie": foaie,
	})

}

func FoaieUpdate(c *gin.Context) {
	id := c.Param("ID")

	var body struct {
		Data  string
		Sofer string
	}

	c.Bind(&body)

	var foaie models.Foaie
	initializers.DB.First(&foaie, id)

	initializers.DB.Model(&foaie).Updates(models.Foaie{
		Data:  body.Data,
		Sofer: body.Sofer,
	})
	c.JSON(200, gin.H{
		"foaie": foaie,
	})
}

func FoiDelete(c *gin.Context) {
	//get id off urls
	id := c.Param("ID")

	initializers.DB.Delete(&models.Foaie{}, id)

	c.Status(200)
}
