package controllers

import (
	"seeyou-go/api/services"

	"github.com/gin-gonic/gin"
)

func SendEmailCode(ctx *gin.Context) {
	services.SendEmailCode(ctx)
}