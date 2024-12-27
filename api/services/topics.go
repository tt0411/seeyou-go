package services

import (
	"math"
	"seeyou-go/api/models"
	"seeyou-go/global"
	"seeyou-go/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTopicList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	var topics []struct {
		models.Topics
		ImgPath string `json:"img_path"`
	}
	var total int64
	err := global.DB.Table("t_topics").Select("t_topics.*, t_file.path as img_path").
		Joins("LEFT JOIN t_file ON t_topics.img_id = t_file.id").Where("t_topics.status = 1").
		Count(&total).
		Offset((page - 1) * pageSize).Limit(pageSize).Scan(&topics).Error
	if err != nil {
		utils.ResponseError(ctx, "获取话题列表失败", nil)
		return
	}
	var data interface{}
	if topics == nil {
		data = []struct{}{}
	} else {
		data = topics
	}
	utils.ResponseOk(ctx, "获取话题列表成功", gin.H{
		"data":      data,
		"total":     total,
		"page":      page,
		"pageSize":  pageSize,
		"totalPage": int64(math.Ceil(float64(total) / float64(pageSize))),
	})
}

func GetTopicInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		utils.ResponseError(ctx, "ID不能为空", nil)
		return
	}
	var topic struct {
		models.Topics
		ImgPath string `json:"img_path"`
	}
	err := global.DB.Table("t_topics").Select("t_topics.*, t_file.path as img_path").
		Joins("LEFT JOIN t_file ON t_topics.img_id = t_file.id").
		Where("t_topics.id = ? AND t_topics.status = 1", id).Scan(&topic).Error
	if err != nil {
		utils.ResponseError(ctx, "获取话题信息失败"+err.Error(), nil)
		return
	}
	if topic.Topics == (models.Topics{}) {
		utils.ResponseError(ctx, "话题不存在", nil)
		return
	}
	utils.ResponseOk(ctx, "获取话题信息成功", topic)
}

func AddTopic(ctx *gin.Context) {
	topic := models.Topics{}
	if err := ctx.ShouldBindJSON(&topic); err != nil {
		utils.ResponseError(ctx, "参数错误"+err.Error(), nil)
		return
	}
	userId, ok := ctx.Get("userId")
	if !ok {
		utils.ResponseError(ctx, "用户不存在", nil)
		return
	}
	topic.CreatedBy = userId.(int)
	// 判断话题名称是否存在
	var count int64
	global.DB.Model(&models.Topics{}).Where("topic_name = ?", topic.TopicName).Count(&count)
	if count > 0 {
		utils.ResponseError(ctx, "话题名称已存在", nil)
		return
	}
	err := global.DB.Create(&topic).Error
	if err != nil {
		utils.ResponseError(ctx, "数据库错误", nil)
		return
	}
	utils.ResponseOk(ctx, "添加话题成功", nil)
}

func UpdateTopic(ctx *gin.Context) {
	var input struct {
		ID        int    `json:"id"`
		Category  string `json:"category,omitempty"`
		TopicName string `json:"topic_name,omitempty"`
		TopicDesc string `json:"topic_desc,omitempty"`
		ImgID     string `json:"img_id,omitempty"`
		IsTop     int    `json:"is_top,omitempty"`
		IsHot     int    `json:"is_hot,omitempty"`
		Sort      int    `json:"sort,omitempty"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, "请求参数错误: "+err.Error(), nil)
		return
	}

	if input.ID == 0 {
		utils.ResponseError(ctx, "ID不能为空", nil)
		return
	}

	id := input.ID
	// 判断话题是否存在
	var count int64
	global.DB.Model(&models.Topics{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		utils.ResponseError(ctx, "话题不存在", nil)
		return
	}
	if input.TopicName != "" {
		// 判断话题名称是否存在
		var count1 int64
		global.DB.Model(&models.Topics{}).Where("topic_name = ? AND id != ?", input.TopicName, id).Count(&count1)
		if count1 > 0 {
			utils.ResponseError(ctx, "话题名称已存在", nil)
			return
		}
	}
	topic := models.Topics{}
	if input.TopicName != "" {
		topic.TopicName = input.TopicName
	}
	if input.TopicDesc != "" {
		topic.TopicDesc = input.TopicDesc
	}
	if input.ImgID != "" {
		topic.ImgID = input.ImgID
	}
	if input.IsTop != 0 {
		topic.IsTop = input.IsTop
	}
	if input.IsHot != 0 {
		topic.IsHot = input.IsHot
	}
	if input.Sort != 0 {
		topic.Sort = input.Sort
	}
	if input.Category != "" {
		topic.Category = input.Category
	}
	topic.UpdatedAt = time.Now()
	err := global.DB.Model(&models.Topics{}).Where("id = ?", id).Updates(&topic).Error
	if err != nil {
		utils.ResponseError(ctx, "更新话题失败", nil)
		return
	}
	utils.ResponseOk(ctx, "更新话题成功", nil)
}

func DeleteTopic(ctx *gin.Context) {
	if ctx.Query("id") == "" {
		utils.ResponseError(ctx, "参数错误", nil)
		return
	}
	id, _ := strconv.Atoi(ctx.Query("id"))
	err := global.DB.Model(&models.Topics{}).Where("id = ?", id).Update("status", 0).Error
	if err != nil {
		utils.ResponseError(ctx, "删除话题失败", nil)
		return
	}
	utils.ResponseOk(ctx, "删除话题成功", nil)
}
