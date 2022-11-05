package routes

import (
	"os"

	"github.com/albuquerquealdry/chatops-babidi/controllers"
	"github.com/gin-gonic/gin"
)

func Request() {
	r := gin.Default()

	r.POST("/commitAction", controllers.CommitAction)
	r.Run(":" + os.Getenv("PORT"))

}
