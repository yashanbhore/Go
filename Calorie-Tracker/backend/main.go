package main

import (
	"Calorie-Tracker/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"

	router := gin.new()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id", routes.GetEntryById)
	router.GET("/ingredients/:ingredients", routes.GetEntriesByIngredients)

	router.PUT("/ingredients/update/:id",routes.UpdateIngredients)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)
	
	router.RUN(":"+port)
}
