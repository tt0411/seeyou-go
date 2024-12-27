package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"regexp"
	"seeyou-go/config"
	"seeyou-go/global"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func Response(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ResponseError(c *gin.Context, msg string, data interface{}) {
	c.JSON(-1, gin.H{
		"code": -1,
		"msg":  msg,
		"data": data,
	})
}

func ResponseOk(c *gin.Context, msg string, data interface{}) {
	c.JSON(0, gin.H{
		"code": 0,
		"msg":  msg,
		"data": data,
	})
}

func IsValidQQNumber(qq string) bool {
	pattern := regexp.MustCompile(`^[1-9][0-9]{4,10}$`)
	return pattern.MatchString(qq)
}

func RandomNumber(digits ...int) string {
	// 设置默认值为8
	d := 8
	if len(digits) > 0 {
		d = digits[0]
	}

	// 创建一个新的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 计算生成的随机数范围
	min := int(math.Pow(10, float64(d-1)))
	max := int(math.Pow(10, float64(d))) - 1

	// 生成随机数
	randomNumber := r.Intn(max-min+1) + min

	// 将数字转换为字符串
	return fmt.Sprintf("%0*d", d, randomNumber)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Minute * time.Duration(config.AppConfig.App.TokenTimeout)).Unix(),
	})
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	err = global.RedisDB.Set(context.Background(), fmt.Sprintf("token:%s", userId), signedToken, time.Minute*time.Duration(config.AppConfig.App.TokenTimeout)).Err()
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token解析错误")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["userId"].(string)
		if !ok {
			return "", errors.New("token中用户信息解析失败")
		}
		return userId, nil
	}
	return "", err
}

func SendEmail(target string, content string) error {
	// 从配置文件中获取SMTP服务器信息
	emailConfig := config.AppConfig.Mail

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(emailConfig.User, "去见APP官方"))
	m.SetHeader("To", target)
	m.SetHeader("Subject", "去见APP - 遇见不一样的人生")
	m.SetBody("text/html", content)

	d := gomail.NewDialer(emailConfig.Smtp, emailConfig.SmtpPort, emailConfig.User, emailConfig.Password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("邮件发送失败:", err)
		return err
	}
	return nil
}
