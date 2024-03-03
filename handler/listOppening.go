package handler

import (
	"net/http"

	"github.com/Matheus-ALima/GoAPI.git/schemas"
	"github.com/gin-gonic/gin"
)

func ListOppeningHandler(ctx *gin.Context) {
	openings := []schemas.Oppening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openigs", openings)
}
