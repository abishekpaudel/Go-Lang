package main

import (
	"go-crud-api/database"
	"go-crud-api/models"
	"go-crud-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.User{})

	routes.SetupRoutes(r)

	r.Run(":8080")
}
