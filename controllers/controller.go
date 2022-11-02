package controllers

import "github.com/gin-gonic/gin"

func CommitAction(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensage" : "commit action",
	})
}