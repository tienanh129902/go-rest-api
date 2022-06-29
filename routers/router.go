package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/tienanh129902/go-rest-api/controllers"
	middleware "github.com/tienanh129902/go-rest-api/middlewares"
	"github.com/tienanh129902/go-rest-api/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	route := router.Group("/api/v1")
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
			question.GET("/:id", utils.AuthenOnly, controllers.GET_QuestionById)
			question.DELETE("/:id", utils.AuthenOnly, controllers.DEL_QuestionById)
		}
		play := route.Group("/play")
		{
			play.GET("", utils.AuthenOnly, controllers.GET_AllQuestions)
			play.POST("/submit", utils.AuthenOnly, controllers.POST_UserSubmit)
		}
		score := route.Group("/score")
		{
			score.GET("/:userid", utils.AuthenOnly, controllers.GET_ScoreBoardByUserId)
		}
	}
	return
}
