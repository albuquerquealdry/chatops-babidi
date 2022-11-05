package controllers

import (
	"github.com/albuquerquealdry/chatops-babidi/service"
	"github.com/gin-gonic/gin"
)

type ReqStructure struct {
	Repository string `json:"repository"`
	CommitHash string `json:"commitHash"`
	Branch     string `json:"branch"`
	UserName   string `json:"userName"`
	Timestamp  string `json:"timestamp"`
}

func CommitAction(c *gin.Context) {

	var requestBody ReqStructure
	if err := c.BindJSON(&requestBody); err != nil {
	}
	service.TelegramRander(requestBody.Repository, requestBody.CommitHash, requestBody.Branch, requestBody.UserName, requestBody.Timestamp)
}
