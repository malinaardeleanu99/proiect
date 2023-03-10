package main

import (
	"example/API/controllers"
	"example/API/initializers"
	"example/API/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/foi", middleware.RequireAuth, controllers.FoiCreate)
	r.GET("/foi", middleware.RequireAuth, controllers.FoiIndex)
	r.GET("/foi/:ID", middleware.RequireAuth, controllers.FoiShow)
	r.PUT("/foi/:ID", middleware.RequireAuth, controllers.FoaieUpdate)
	r.DELETE("/foi/:ID", middleware.RequireAuth, controllers.FoiDelete)

	r.Run()
}
