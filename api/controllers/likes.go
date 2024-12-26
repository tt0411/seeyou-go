package controllers

import (
	"seeyou-go/api/services"

	"github.com/gin-gonic/gin"
)

func LikePost(ctx *gin.Context) {
	services.LikePost(ctx)
}

func CancelLikePost(ctx *gin.Context) {
	services.CancelLikePost(ctx)
}
