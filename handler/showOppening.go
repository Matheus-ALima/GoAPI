package handler

import (
	"net/http"

	"github.com/Matheus-ALima/GoAPI.git/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOppeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Oppening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	sendSuccess(ctx, "show-opening", &opening)
}
