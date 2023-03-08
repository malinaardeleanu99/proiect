package main

import (
	"example/API/controllers"
	"example/API/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/foi", controllers.FoiCreate)
	r.GET("/foi", controllers.FoiIndex)
	r.GET("/foi/:ID", controllers.FoiShow)
	r.PUT("/foi/:ID", controllers.FoaieUpdate)
	r.DELETE("/foi/:ID", controllers.FoiDelete)
	r.Run()
}
