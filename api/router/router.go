package router

import (
	"seeyou-go/api/controllers"
	"seeyou-go/api/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/loginByEmail", controllers.LoginByEmail)
		auth.POST("/register", controllers.Register)
		auth.POST("/registerByEmail", controllers.RegisterByEmail)
		auth.GET("/logout", middlewares.AuthMiddleware(), controllers.Logout)
		auth.POST("/updateUserInfo", middlewares.AuthMiddleware(), controllers.UpdateUserInfo)
		auth.GET("/userInfo", middlewares.AuthMiddleware(), middlewares.GetUserInfoMiddleware(), controllers.GetUserInfo)
	}
	common := r.Group("/api/common")
	{
		common.GET("/sendEmailCode", controllers.SendEmailCode)
	}
	likes := r.Group("/api/likes")
	{
		likes.GET("/like", middlewares.AuthMiddleware(), controllers.LikePost)
		likes.GET("/cancelLike", middlewares.AuthMiddleware(), controllers.CancelLikePost)
	}

	return r
}
