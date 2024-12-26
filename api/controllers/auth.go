package controllers

import (
	"seeyou-go/api/services"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	services.Login(ctx)
}

func LoginByEmail(ctx *gin.Context) {
	services.LoginByEmail(ctx)
}

func Register(ctx *gin.Context) {
	services.Register(ctx)
}

func RegisterByEmail(ctx *gin.Context) {
	services.RegisterByEmail(ctx)
}

func Logout(ctx *gin.Context) {
	services.Logout(ctx)
}

func UpdateUserInfo(ctx *gin.Context) {
	services.UpdateUserInfo(ctx)
}

func GetUserInfo(ctx *gin.Context) {
	services.GetUserInfo(ctx)
}
