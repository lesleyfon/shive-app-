package routes

import (
	controller "shive/controllers"
	"shive/middleware"

	"github.com/gin-gonic/gin"
)

func GenreRouter(router *gin.Engine) {
	// Authenticates user
	router.Use(middleware.Authenticate())

	// Post route to create  a genre
	router.POST(
		"/genres/creategenre",
		controller.CreateGenre(),
	)
}