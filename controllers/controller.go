package controllers

import (
	"github.com/albuquerquealdry/chatops-babidi/service"
	"github.com/gin-gonic/gin"
)

type CommitStructure struct {
	Repository string `json:"repository"`
	CommitHash string `json:"commitHash"`
	Branch     string `json:"branch"`
	UserName   string `json:"userName"`
	Timestamp  string `json:"timestamp"`
}

type PullStructure struct {
	Repository   string `json:"repository"`
	OriginBranch string `json:"originBranch"`
	FutureBranch string `json:"futureBranch"`
	UserName     string `json:"userName"`
	Timestamp    string `json:"timestamp"`
}

func CommitAction(c *gin.Context) {

	var requestBody CommitStructure
	if err := c.BindJSON(&requestBody); err != nil {
	}
	service.TelegramCommit(requestBody.Repository, requestBody.CommitHash, requestBody.Branch, requestBody.UserName, requestBody.Timestamp)
}

func PullAction(c *gin.Context) {

	var requestBody PullStructure
	if err := c.BindJSON(&requestBody); err != nil {
	}
	service.TelegramPull(requestBody.Repository, requestBody.OriginBranch, requestBody.FutureBranch, requestBody.UserName, requestBody.Timestamp)
}
