package main

import (
	"api/categories"
	"api/models"
	"api/vendors"

	"github.com/gin-gonic/gin"
)

/**
how to query from the database (sqlite)
how to upsert and validate
how to test it
how to authenticate

*/

func main() {

	models.ConnectDatabase()
	router := gin.Default()

	router.GET("/categories", categories.Categories)
	router.POST("/categories", categories.Create)
	router.PUT("/categories/:id", categories.Update)
	router.DELETE("/categories/:id", categories.Delete)

	router.GET("/vendors", vendors.Vendors)
	router.POST("/vendors", vendors.Create)
	router.PUT("/vendors/:id", vendors.Update)
	router.DELETE("/vendors/:id", vendors.Delete)

	router.Run("localhost:8080")
}
