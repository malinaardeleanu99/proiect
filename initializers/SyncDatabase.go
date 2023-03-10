package initializers

import "example/API/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Foaie{})
	DB.AutoMigrate(&models.Cursa{})
	DB.AutoMigrate(&models.User{})
}
