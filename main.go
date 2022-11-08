package main

import (
	"os"
	"server/server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.GET("")
	router.GET("/environment/getEnvironment/:name", routes.GetEnvironment)
	router.POST("/character/create", routes.AddCharacter)
	router.DELETE("/character/deleteCharacter/:id", routes.DeleteCharacter)
	router.PUT("/character/updateCharacter/:id", routes.UpdateCharacter)
	router.GET("/character/getCharacter/:name", routes.Getcharacter)
	router.GET("/character/getCharacters", routes.GetCharacters)
	router.POST("/environment/create", routes.AddEnvironment)

	router.POST("")

	router.Run(":" + port)

}
