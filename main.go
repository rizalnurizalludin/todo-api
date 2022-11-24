package main

import (
	"github.com/rizalnurizalludin/todo-api/models"
	"github.com/rizalnurizalludin/todo-api/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.ActivityGroup{}, &models.TODO{})

	r := routes.SetupRoutes(db)
	r.Run("localhost:3030")
}
