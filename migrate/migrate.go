package main

import (
	"example/API/initializers"
	"example/API/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Foaie{})
}
