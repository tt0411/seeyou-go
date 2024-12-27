package services

import (
	"context"
	"fmt"
	"seeyou-go/api/models"
	"seeyou-go/config"
	"seeyou-go/global"
	"seeyou-go/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var input struct {
		UserNo   string `json:"user_no" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, "请求参数错误", nil)
		return
	}
	var user models.AppUser
	if err := global.DB.Where("user_no = ?", input.UserNo).First(&user).Error; err != nil {
		utils.ResponseError(ctx, "用户不存在", nil)
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		utils.ResponseError(ctx, "密码错误", nil)
		return
	}

	token, err := utils.GenerateToken(strconv.Itoa(user.ID))
	if err != nil {
		utils.ResponseError(ctx, "生成token失败", nil)
		return
	}
	// 更新用户最后登录时间
	user.LastLogin = time.Now()
	global.DB.Save(&user)
	global.RedisDB.Set(context.Background(), fmt.Sprintf("online:%s", strconv.Itoa(user.ID)), 1, time.Minute*time.Duration(config.AppConfig.App.TokenTimeout))
	utils.ResponseOk(ctx, "登录成功", gin.H{"token": token})
}

func LoginByEmail(ctx *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, "请求参数错误", nil)
		return
	}
	// 从redis中获取验证码
	code, err := global.RedisDB.Get(context.Background(), fmt.Sprintf("email_code:%s", input.Email)).Result()
	if err != nil {
		utils.ResponseError(ctx, "验证码错误", nil)
		return
	}
	if code != input.Code {
		utils.ResponseError(ctx, "验证码错误", nil)
		return
	}
	// 从数据库中获取用户
	var user models.AppUser
	if err := global.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.ResponseError(ctx, "用户不存在", nil)
		return
	}
	token, err := utils.GenerateToken(strconv.Itoa(user.ID))
	if err != nil {
		utils.ResponseError(ctx, "生成token失败", nil)
		return
	}
	// 更新用户最后登录时间
	user.LastLogin = time.Now()
	global.DB.Save(&user)
	global.RedisDB.Set(context.Background(), fmt.Sprintf("online:%s", strconv.Itoa(user.ID)), 1, time.Minute*time.Duration(config.AppConfig.App.TokenTimeout))
	global.RedisDB.Del(context.Background(), fmt.Sprintf("email_code:%s", input.Email))
	utils.ResponseOk(ctx, "登录成功", gin.H{"token": token})
}

func Register(ctx *gin.Context) {
	var user models.AppUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseError(ctx, "请求参数错误: "+err.Error(), nil)
		return
	}
	randomNo := utils.RandomNumber()
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.ResponseError(ctx, "密码加密错误", nil)
		return
	}
	if user.NickName == "" {
		user.NickName = "用户" + randomNo
	}
	user.Password = hashedPwd
	user.State = 1
	user.Gender = 3
	user.UserNo = randomNo
	user.LastLogin = time.Now()
	// 判断手机号是否已注册
	if err := global.DB.Where("phone = ?", user.Phone).First(&user).Error; err == nil {
		utils.ResponseError(ctx, "手机号已注册", nil)
		return
	}
	if err := global.DB.AutoMigrate(&user); err != nil {
		utils.ResponseError(ctx, "数据库错误", nil)
		return
	}
	if err := global.DB.Create(&user).Error; err != nil {
		utils.ResponseError(ctx, "创建用户失败", nil)
		return
	}
	token, err := utils.GenerateToken(strconv.Itoa(user.ID))
	if err != nil {
		utils.ResponseError(ctx, "token创建失败", nil)
		return
	}
	global.RedisDB.Set(context.Background(), fmt.Sprintf("online:%s", strconv.Itoa(user.ID)), 1, time.Minute*time.Duration(config.AppConfig.App.TokenTimeout))
	utils.ResponseOk(ctx, "注册成功", gin.H{"token": token})
}

func RegisterByEmail(ctx *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, "请求参数错误", nil)
		return
	}
	// 从redis中获取验证码
	code, err := global.RedisDB.Get(context.Background(), fmt.Sprintf("email_code:%s", input.Email)).Result()
	if err != nil {
		utils.ResponseError(ctx, "验证码错误", nil)
		return
	}
	if code != input.Code {
		utils.ResponseError(ctx, "验证码错误", nil)
		return
	}
	// 从数据库中获取用户
	var user models.AppUser
	if err := global.DB.Where("email = ?", input.Email).First(&user).Error; err == nil {
		utils.ResponseOk(ctx, "邮箱已绑定，请直接登录", nil)
		return
	}
	randomNo := utils.RandomNumber()
	user.Email = input.Email
	user.UserNo = randomNo
	user.LastLogin = time.Now()
	if err := global.DB.AutoMigrate(&user); err != nil {
		utils.ResponseError(ctx, "数据库错误", nil)
		return
	}
	if err := global.DB.Create(&user).Error; err != nil {
		utils.ResponseError(ctx, "创建用户失败", nil)
		return
	}
	token, err := utils.GenerateToken(user.UserNo)
	if err != nil {
		utils.ResponseError(ctx, "token创建失败", nil)
		return
	}
	global.RedisDB.Set(context.Background(), fmt.Sprintf("online:%s", user.UserNo), 1, time.Minute*time.Duration(config.AppConfig.App.TokenTimeout))
	global.RedisDB.Del(context.Background(), fmt.Sprintf("email_code:%s", input.Email))
	utils.ResponseOk(ctx, "注册成功", gin.H{"token": token})
}

func UpdateUserInfo(ctx *gin.Context) {
	var input struct {
		NickName  string `json:"nick_name,omitempty"`
		Signature string `json:"signature,omitempty"`
		Avatar    string `json:"avatar,omitempty"`
		RealName  string `json:"real_name,omitempty"`
		Birthday  string `json:"birthday,omitempty"`
		Province  string `json:"province,omitempty"`
		City      string `json:"city,omitempty"`
		QQ        string `json:"qq,omitempty"`
		School    string `json:"school,omitempty"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, "请求参数错误: "+err.Error(), nil)
		return
	}

	userId, exists := ctx.Get("userId")
	if !exists {
		utils.ResponseError(ctx, "用户不存在", nil)
		return
	}

	var user models.AppUser
	if err := global.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		utils.ResponseError(ctx, "用户不存在", nil)
		return
	}

	// 更新用户信息
	if input.NickName != "" {
		user.NickName = input.NickName
	}
	if input.Signature != "" {
		user.Signature = input.Signature
	}
	if input.Avatar != "" {
		user.Avatar = input.Avatar
	}
	if input.RealName != "" {
		user.RealName = input.RealName
	}
	if input.School != "" {
		user.School = input.School
	}
	if input.QQ != "" {
		// 判断QQ是否合法
		if !utils.IsValidQQNumber(input.QQ) {
			utils.ResponseError(ctx, "QQ号码不合法", nil)
			return
		}
		// 判断QQ是否已存在
		var existingUser models.AppUser
		if err := global.DB.Where("qq = ?", input.QQ).First(&existingUser).Error; err == nil {
			utils.ResponseError(ctx, "QQ号码已被注册", nil)
			return
		}
		user.QQ = input.QQ
	}
	if input.Province != "" {
		user.Province = input.Province
	}
	if input.City != "" {
		user.City = input.City
	}
	if input.Birthday != "" {
		birthday, err := time.Parse("2006-01-02", input.Birthday)
		if err != nil {
			utils.ResponseError(ctx, "生日格式错误", nil)
			return
		}
		user.Birthday = &birthday

		// 更新年龄
		user.Age = int(time.Now().Year() - birthday.Year())
		if time.Now().YearDay() < birthday.YearDay() {
			user.Age-- // 如果还没到生日，年龄减一
		}

		// 更新星座
		month := birthday.Month()
		day := birthday.Day()
		switch {
		case (month == 1 && day >= 20) || (month == 2 && day <= 18):
			user.Constellation = "水瓶座"
		case (month == 2 && day >= 19) || (month == 3 && day <= 20):
			user.Constellation = "双鱼座"
		case (month == 3 && day >= 21) || (month == 4 && day <= 19):
			user.Constellation = "白羊座"
		case (month == 4 && day >= 20) || (month == 5 && day <= 20):
			user.Constellation = "金牛座"
		case (month == 5 && day >= 21) || (month == 6 && day <= 21):
			user.Constellation = "双子座"
		case (month == 6 && day >= 22) || (month == 7 && day <= 22):
			user.Constellation = "巨蟹座"
		case (month == 7 && day >= 23) || (month == 8 && day <= 22):
			user.Constellation = "狮子座"
		case (month == 8 && day >= 23) || (month == 9 && day <= 22):
			user.Constellation = "处女座"
		case (month == 9 && day >= 23) || (month == 10 && day <= 23):
			user.Constellation = "天秤座"
		case (month == 10 && day >= 24) || (month == 11 && day <= 22):
			user.Constellation = "天蝎座"
		case (month == 11 && day >= 23) || (month == 12 && day <= 21):
			user.Constellation = "射手座"
		case (month == 12 && day >= 22) || (month == 1 && day <= 19):
			user.Constellation = "摩羯座"
		}
	}

	user.UpdatedAt = time.Now()
	err := global.DB.Model(&models.AppUser{}).Where("id = ?", userId).Updates(&user).Error
	if err != nil {
		utils.ResponseError(ctx, "更新用户信息失败", nil)
		return
	}
	utils.ResponseOk(ctx, "用户信息更新成功", nil)
}

func Logout(ctx *gin.Context) {
	req, _ := ctx.Get("userId")
	userId := req.(string)
	global.RedisDB.Del(context.Background(), fmt.Sprintf("online:%s", userId))
	global.RedisDB.Del(context.Background(), fmt.Sprintf("token:%s", userId))
	utils.ResponseOk(ctx, "退出登录成功", nil)
}

func GetUserInfo(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		utils.ResponseError(ctx, "用户不存在", nil)
		return
	}
	utils.ResponseOk(ctx, "获取用户信息成功", user)
}
