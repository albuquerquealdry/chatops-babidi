package routes

import (
	"fmt"
	"os"

	"github.com/albuquerquealdry/chatops-babidi/controllers"
	"github.com/gin-gonic/gin"
)

func Request() {
	r := gin.Default()

	r.POST("/commitAction", controllers.CommitAction)
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	r.Run(":" + port)

}
