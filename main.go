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
	router.POST("/personagem/criar", routes.AddPersonagem)
	router.DELETE("/personagem/deletarPersonagem/:id", routes.DeletePersonagem)
	router.PUT("/personagem/updatePersonagem/:id", routes.UpdatePersonagem)
	router.GET("/personagem/getPersonagem/:nome", routes.GetPersonagem)
	router.GET("/personagem/getPersonagens", routes.GetPersonagens)

	router.POST("")

	router.Run(":" + port)

}
