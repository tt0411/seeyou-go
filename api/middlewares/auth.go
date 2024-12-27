package middlewares

import (
	"context"
	"fmt"
	"seeyou-go/api/models"
	"seeyou-go/config"
	"seeyou-go/global"
	"seeyou-go/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			utils.Response(c, 401, "token不存在", "")
			c.Abort()
			return
		}
		userId, err := utils.ParseToken(token)
		if err != nil {
			utils.Response(c, 401, "token解析失败"+err.Error(), "")
			c.Abort()
			return
		}
		// 从redis里校验token是否存在
		exists, err := global.RedisDB.Exists(context.Background(), fmt.Sprintf("token:%s", userId)).Result()
		if err != nil || exists == 0 {
			utils.Response(c, 401, "token不存在或已过期", "")
			c.Abort()
			return
		}
		c.Set("userId", userId)
		// 更新redis中的在线状态的过期时间
		global.RedisDB.Expire(context.Background(), fmt.Sprintf("online:%s", userId), time.Minute*time.Duration(config.AppConfig.App.TokenTimeout))
		// 更新redis中的token过期时间
		global.RedisDB.Expire(context.Background(), fmt.Sprintf("token:%s", userId), time.Minute*time.Duration(config.AppConfig.App.TokenTimeout))
		c.Next()
	}
}

type UserResponse struct {
	ID        int    `json:"id"`
	UserNo    string `json:"user_no"`
	Phone     string `json:"phone"`
	NickName  string `json:"nick_name"`
	RealName  string `json:"real_name"`
	Signature string `json:"signature"`
	Intro     string `json:"intro"`
	Age       int    `json:"age"`
	Avatar    string `json:"avatar"`
	Gender    int    `json:"gender"`
	State     int    `json:"state"`
	Remark    string `json:"remark"`
	LastLogin string `json:"last_login"`
}

func GetUserInfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户名
		req, _ := c.Get("userId")
		userId := req.(string)
		// 从db中获取用户信息
		var user models.AppUser
		if err := global.DB.Where("id = ?", userId).First(&user).Error; err != nil {
			utils.Response(c, 401, "用户不存在", "")
			c.Abort()
			return
		}
		response := UserResponse{
			ID:        user.ID,
			UserNo:    user.UserNo,
			Phone:     user.Phone,
			NickName:  user.NickName,
			RealName:  user.RealName,
			Signature: user.Signature,
			Intro:     user.Intro,
			Age:       user.Age,
			Avatar:    user.Avatar,
			Gender:    user.Gender,
			State:     user.State,
			Remark:    user.Remark,
			LastLogin: user.LastLogin.Format(time.RFC3339), // 格式化时间
		}

		if len(response.Phone) >= 7 {
			response.Phone = response.Phone[:3] + "****" + response.Phone[7:] // 中间4位用*替代
		}
		c.Set("user", response)
		c.Next()
	}
}
