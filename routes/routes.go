package routes

import (
	"os"

	"github.com/albuquerquealdry/chatops-babidi/controllers"
	"github.com/gin-gonic/gin"
)

func Request() {
	r := gin.Default()

	r.POST("/commitAction", controllers.CommitAction)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

}
