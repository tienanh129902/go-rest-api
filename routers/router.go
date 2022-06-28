package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/tienanh129902/go-rest-api/controllers"
	middleware "github.com/tienanh129902/go-rest-api/middlewares"
	"github.com/tienanh129902/go-rest-api/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	route := router.Group("/api")
	route.Use(
		middleware.CORSMiddleware,
		middleware.AuthenMiddleware,
	)
	{
		auth := route.Group("/auth")
		{
			auth.POST("/login", controllers.POST_Login)
			auth.POST("/signup", controllers.POST_Register)
			auth.POST("/logout", utils.AuthenOnly, controllers.POST_Logout)
		}
		user := route.Group("/user")
		{
			user.GET("/:id", utils.AuthenOnly, controllers.GET_User)
			user.PATCH("", utils.AuthenOnly, controllers.PATCH_User)
			user.GET("/me", utils.AuthenOnly, controllers.GET_Me)
		}
		question := route.Group("/question")
		{
			question.POST("", utils.AuthenOnly, controllers.POST_CreateQuestion)
			question.GET("", utils.AuthenOnly, controllers.GET_AllQuestions)
			question.GET("/:id", utils.AuthenOnly, controllers.GET_QuestionById)
			question.DELETE("/:id", utils.AuthenOnly, controllers.DEL_QuestionById)
		}
		answer := route.Group("/answer")
		{
			answer.POST("", utils.AuthenOnly)
			answer.GET("", utils.AuthenOnly)
		}
	}
	return
}
