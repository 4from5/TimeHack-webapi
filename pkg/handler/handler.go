package handler

import (
	"github.com/4from5/TimeHack-webapi/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		api.Group("categories")
		{
			auth.GET("/", h.getCategories)
		}
		api.Group("events")
		api.Group("notions")
		api.Group("tasks")
	}
	return router
}
