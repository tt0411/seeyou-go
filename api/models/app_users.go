package models

import (
	"time"

	"gorm.io/gorm"
)

// AppUser 表示 app 用户表
type AppUser struct {
	gorm.Model
	ID            int        `gorm:"primaryKey;autoIncrement" json:"id"` // 自增id
	UserNo        string     `gorm:"size:12;not null" json:"user_no"`    // 账号
	Phone         string     `gorm:"size:11" json:"phone"`               // 手机号
	NickName      string     `gorm:"size:50" json:"nick_name"`           // 昵称
	RealName      string     `gorm:"size:50" json:"real_name"`           // 真实姓名
	Signature     string     `gorm:"size:255" json:"signature"`          // 个性签名
	Password      string     `gorm:"size:255" json:"password"`           // 密码
	Intro         string     `gorm:"size:255" json:"intro"`              // 简介
	Age           int        `json:"age"`                                // 年龄
	Avatar        string     `gorm:"size:255" json:"avatar"`             // 头像
	Constellation string     `gorm:"size:6" json:"constellation"`        // 星座
	Gender        int        `gorm:"default:3" json:"gender"`            // 性别 1、男，2、女，3、保密
	School        string     `gorm:"size:255" json:"school"`             // 毕业学校
	Birthday      *time.Time `json:"birthday"`                           // 生日
	Province      string     `gorm:"size:50" json:"province"`            // 省份
	City          string     `gorm:"size:50" json:"city"`                // 城市
	QQ            string     `gorm:"size:20" json:"qq"`                  // QQ
	Email         string     `gorm:"size:50;not null" json:"email"`      // 邮箱
	State         int        `gorm:"default:1;not null" json:"state"`    // 状态 0、停用，1、启用
	Remark        string     `gorm:"size:255" json:"remark"`             // 备注
	LastLogin     time.Time  `json:"last_login"`                         // 上次登录时间
	CreatedAt     time.Time  `json:"created_at"`                         // 创建时间
	UpdatedAt     time.Time  `json:"updated_at"`                         // 更新时间
	DeletedAt     *time.Time `json:"deleted_at"`                         // 删除时间
}

func (AppUser) TableName() string {
	return "t_app_users" // 指定表名为 t_app_users
}
