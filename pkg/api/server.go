package http

import (
	"github.com/gin-gonic/gin"

	handler "NurtureNest/pkg/api/handler"
	middleware "NurtureNest/pkg/api/middleware"
)

type ServerHTTP struct {
	engine              *gin.Engine
	userHandler         *handler.UserHandler
	tutorHandler        *handler.TutorHandler
	goalTrackingHandler *handler.GoalTrackingHandler
}

func NewServerHTTP(userHandler *handler.UserHandler, tutorHandler *handler.TutorHandler,
	chatBotHandler *handler.ChatBotHandler, goalTrackingHandler *handler.GoalTrackingHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "200",
		})
	})

	//register user
	engine.POST("user/register", userHandler.UserRegister)
	engine.GET("/users", userHandler.UserFindAll)

	// Request JWT
	engine.POST("/login", userHandler.UserLogin)

	// Auth middleware
	api := engine.Group("/api", middleware.JwtMiddleware())

	api.GET("users", userHandler.UserFindAll)
	api.GET("users/:id", userHandler.FindByID)
	api.POST("users", userHandler.Save)
	api.DELETE("users/:id", userHandler.Delete)

	api.POST("/chatbot/generate", chatBotHandler.GenerateChatBot)
	api.GET("/chatbot/chats", chatBotHandler.GetAllHistoryChatBot)

	api.POST("/goal/set", goalTrackingHandler.SetGoal)
	api.PATCH("/goal/result", goalTrackingHandler.GoalResult)

	api.POST("tutor/", tutorHandler.Save)

	return &ServerHTTP{
		engine:              engine,
		userHandler:         userHandler,
		tutorHandler:        tutorHandler,
		goalTrackingHandler: goalTrackingHandler,
	}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
