package services

import (
	"context"
	"fmt"
	"os"
	"seeyou-go/api/models"
	"seeyou-go/global"
	"seeyou-go/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SendEmailCode(ctx *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}

	input.Email = ctx.Query("email")
	if input.Email == "" {
		utils.ResponseError(ctx, "请求参数错误: email 是必需的", nil)
		return
	}

	// 生成验证码
	code := utils.RandomNumber(6)

	// 确保 code 是字符串类型
	content := fmt.Sprintf(`<div style="display: flex;flex-direction: column;justify-content: center;align-items: center;
                    width: 300px;height: 300px;box-shadow: 0px 0px 10px #ccc;border-radius: 30px;margin: 66px auto;">
                  <img width="100" src="https://avatars.githubusercontent.com/u/35050738?v=4" alt="">
                  <span style="line-height: 36px;padding: 0 10px;">来自【去见APP - 遇见不一样的人生】邮箱验证码(有效时长5分钟)</span>
                  <div style="font-weight: 600;font-size: 22px;line-height: 46px;">%s</div>
                </div>`, code)

	// 发送验证码到邮箱
	err := utils.SendEmail(input.Email, content)
	if err != nil {
		utils.ResponseError(ctx, "发送验证码失败", nil)
		return
	}

	// 将验证码存储到 Redis 以便后续验证
	global.RedisDB.Set(context.Background(), fmt.Sprintf("email_code:%s", input.Email), code, 5*time.Minute) // 5分钟有效期
	utils.ResponseOk(ctx, "验证码发送成功", nil)
}

func UploadFile(ctx *gin.Context) {
	category := ctx.Param("category")
	if category == "" {
		utils.ResponseError(ctx, "请求参数错误: category 是必需的", nil)
		return
	}
	req, _ := ctx.Get("userId")
	userId := req.(string)
	// 获取文件
	file, err := ctx.FormFile("file")
	if err != nil {
		utils.ResponseError(ctx, "获取文件失败", nil)
		return
	}
	// 只能上传图片
	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
		utils.ResponseError(ctx, "只能上传图片", nil)
		return
	}
	filePath := fmt.Sprintf("uploads/%s", file.Filename)
	// 保存文件
	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		utils.ResponseError(ctx, "保存文件失败", nil)
		return
	}
	// md5校验数据唯一性
	md5, err := utils.GetFileMD5(filePath)
	if err != nil {
		utils.ResponseError(ctx, "获取文件md5失败"+err.Error(), nil)
		return
	}
	var files models.File
	// 查询数据库中是否存在相同md5的文件
	if err := global.DB.Where("md5 = ?", md5).First(&files).Error; err == nil {
		if files.Name != file.Filename {
			os.Remove(filePath)
		}
		utils.ResponseOk(ctx, "文件上传成功", files.Path)
		return
	}
	// 存储文件信息到数据库
	err = global.DB.Create(&models.File{
		Path:       filePath,
		Category:   category,
		Size:       int(file.Size),
		Name:       file.Filename,
		Ext:        file.Filename[strings.LastIndex(file.Filename, ".")+1:],
		MD5:        md5,
		UploaderID: userId,
	}).Error
	if err != nil {
		utils.ResponseError(ctx, "保存文件信息失败"+err.Error(), nil)
		return
	}
	utils.ResponseOk(ctx, "文件上传成功", filePath)
}
