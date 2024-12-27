package controllers

import (
	"seeyou-go/api/services"

	"github.com/gin-gonic/gin"
)

func GetTopicList(ctx *gin.Context) {
	services.GetTopicList(ctx)
}

func GetTopicInfo(ctx *gin.Context) {
	services.GetTopicInfo(ctx)
}

func AddTopic(ctx *gin.Context) {
	services.AddTopic(ctx)
}

func UpdateTopic(ctx *gin.Context) {
	services.UpdateTopic(ctx)
}

func DeleteTopic(ctx *gin.Context) {
	services.DeleteTopic(ctx)
}
