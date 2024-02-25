package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOppeningHandler(ctx *gin.Context) {
	request := struct {
		role string
	}{}

	ctx.BindJSON(&request)

	logger.Infof("request received: %v", request)
}
""			