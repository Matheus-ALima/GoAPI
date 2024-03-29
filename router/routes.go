package router

import (
	"github.com/Matheus-ALima/GoAPI.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//Initalize Handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/oppening", handler.ShowOppeningHandler)
		v1.POST("/oppening", handler.CreateOppeningHandler)
		v1.DELETE("/oppening", handler.DeleteOppeningHandler)
		v1.PUT("/oppening", handler.UpdateOppeningHandler)
		v1.GET("/oppenings", handler.ListOppeningHandler)
	}
}
