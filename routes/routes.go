package routes

import (
	"github.com/albuquerquealdry/chatops-babidi/controllers"
	"github.com/gin-gonic/gin"
)

func Request() {
	r := gin.Default()
	r.POST("/commitAction", controllers.CommitAction)
	r.Run(":8080")

}
