package handler

import (
	"github.com/4from5/TimeHack-webapi/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(LiberalCORS)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		categories := api.Group("/categories")
		{
			categories.GET("/", h.getCategories)
			categories.GET("/:id", h.getCategoryById)

			categories.POST("/", h.createCategory)
		}

		events := api.Group("/events")
		{
			events.GET("/", h.getEvents)
			events.GET("/:id", h.getEventById)

			events.POST("/:id", h.createEvent)
		}

		notions := api.Group("/notions")
		{
			notions.GET("/", h.getNotions)
			notions.GET("/:id", h.getNotionById)
			notions.POST("/", h.createNotion)
		}
	}
	return router
}
