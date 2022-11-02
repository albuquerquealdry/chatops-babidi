package controllers

import (
	"github.com/albuquerquealdry/chatops-babidi/service"
	"github.com/gin-gonic/gin"
)

func CommitAction(c *gin.Context) {
	service.TelegramRander()

	c.JSON(200, gin.H{
		"mensage": "Success Request",
	})
}
