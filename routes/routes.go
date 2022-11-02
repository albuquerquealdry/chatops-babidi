package routes

import "github.com/gin-gonic/gin"
import "github.com/albuquerquealdry/chatops-babidi/controllers"

func Request() {
	r := gin.Default()
	r.GET("/commitAction", controllers.CommitAction)
	r.Run(":3001")
}