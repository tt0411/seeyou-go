package services

import (
	"context"
	"fmt"
	"seeyou-go/api/models"
	"seeyou-go/global"
	"seeyou-go/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func LikePost(ctx *gin.Context) {
	targetId := ctx.Query("target_id")
	targetType := ctx.Query("target_type")
	req, _ := ctx.Get("userId")
	userId := req.(string)
	key := fmt.Sprintf("likes:%s:%s:%s", userId, targetType, targetId)
	// 不能重复点赞
	exists, err := global.RedisDB.Exists(context.Background(), key).Result()
	if err != nil || exists == 1 {
		utils.ResponseError(ctx, "不能重复点赞", nil)
		return
	}
	err = global.RedisDB.Set(context.Background(), key, userId, 0).Err()
	if err != nil {
		utils.ResponseError(ctx, "点赞失败", nil)
		return
	}
	utils.ResponseOk(ctx, "点赞成功", nil)
}

func CancelLikePost(ctx *gin.Context) {
	targetId := ctx.Query("target_id")
	targetType := ctx.Query("target_type")
	req, _ := ctx.Get("userId")
	userId := req.(string)
	key := fmt.Sprintf("likes:%s:%s:%s", userId, targetType, targetId)
	err := global.RedisDB.Del(context.Background(), key).Err()
	if err != nil {
		utils.ResponseError(ctx, "取消点赞失败", nil)
		return
	}
	utils.ResponseOk(ctx, "取消点赞成功", nil)
}

// 将redis中的点赞数据同步到数据库
func SyncLikesToMySQL() {
	iter := global.RedisDB.Scan(context.Background(), 0, "likes:*", 0).Iterator()
	for iter.Next(context.Background()) {
		key := iter.Val()
		_, err := global.RedisDB.Get(context.Background(), key).Result()
		if err != nil {
			continue
		}
		userId := strings.Split(key, ":")[1]
		targetType := strings.Split(key, ":")[2]
		targetId := strings.Split(key, ":")[3]
		err = global.DB.Create(&models.Likes{UserID: userId, TargetType: targetType, TargetID: targetId}).Error
		if err != nil {
			fmt.Println("同步点赞数据到数据库失败", err)
			continue
		} else {
			global.RedisDB.Del(context.Background(), key).Err()
		}
	}
}

// 获取点赞数
func GetLikeCount(targetId string, targetType string) (int64, error) {
	var count int64
	err := global.DB.Model(&models.Likes{}).Where("target_id = ? AND target_type = ?", targetId, targetType).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
