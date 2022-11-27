package handler

import (
	"github.com/4from5/TimeHack-webapi/pkg/service"
	"github.com/gin-contrib/cors"
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

	router.Use(cors.Default())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		user := api.Group("/user")
		{
			user.GET("/", h.getUsername)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", h.getCategories)
			categories.GET("/:id", h.getCategoryById)

			categories.POST("/", h.createCategory)

			categories.DELETE("/:id", h.deleteCategory)

			categories.PUT("/:id", h.updateCategory)
		}

		events := api.Group("/events")
		{
			events.GET("/", h.getEvents)
			events.GET("/:id", h.getEventById)

			events.POST("/:id", h.createEvent)
			events.POST("/schedule", h.getSchedule)

			events.DELETE("/:id", h.deleteEvent)

		}

		notions := api.Group("/notions")
		{
			notions.GET("/", h.getNotions)
			notions.GET("/:id", h.getNotionById)

			notions.POST("/", h.createNotion)

			notions.DELETE("/:id", h.deleteNotion)

		}
		tasks := api.Group("/tasks")
		{
			tasks.GET("/", h.getTasks)
			tasks.GET("/:id", h.getTaskById)

			tasks.POST("/", h.createTask)

			tasks.DELETE("/:id", h.deleteNotion)

		}
	}
	return router
}
