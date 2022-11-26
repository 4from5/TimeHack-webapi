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

//func LiberalCORS(c *gin.Context) {
//	c.Header("Access-Control-Allow-Origin", "*")
//	if c.Request.Method == "OPTIONS" {
//		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
//			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
//		}
//		c.AbortWithStatus(http.StatusOK)
//	}
//}

//func CORSMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Header("Access-Control-Allow-Credentials", "true")
//		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
//
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(204)
//			return
//		}
//
//		c.Next()
//	}
//}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())
	//router.Use(LiberalCORS)

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
			events.POST("/schedule", h.getSchedule)

		}

		notions := api.Group("/notions")
		{
			notions.GET("/", h.getNotions)
			notions.GET("/:id", h.getNotionById)

			notions.POST("/", h.createNotion)
		}
		tasks := api.Group("/tasks")
		{
			tasks.GET("/", h.getTasks)
			tasks.GET("/:id", h.getTaskById)

			tasks.POST("/", h.createTask)
		}
	}
	return router
}
