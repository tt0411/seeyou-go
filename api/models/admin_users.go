package models

import (
	"time"

	"gorm.io/gorm"
)

// AdminUser 表示后台管理用户表
type AdminUser struct {
	gorm.Model
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"` // 自增id
	UserNo    string     `gorm:"size:15;not null" json:"user_no"`    // 账号
	Phone     string     `gorm:"size:11;not null" json:"phone"`      // 手机号
	Avatar    string     `gorm:"size:255" json:"avatar"`             // 头像
	Username  string     `gorm:"size:15" json:"username"`            // 用户名
	Password  string     `gorm:"size:255;not null" json:"password"`  // 密码
	Email     string     `gorm:"size:20" json:"email"`               // 邮箱
	Gender    int        `gorm:"default:3" json:"gender"`            // 性别 1、男，2、女，3、保密
	State     int        `gorm:"default:1;not null" json:"state"`    // 是否停用 1、启用，0、停用
	Type      int        `gorm:"default:2;not null" json:"type"`     // 类型 1、超级管理员 2、普通管理员
	Remark    string     `gorm:"size:255" json:"remark"`             // 备注
	CreatedAt time.Time  `json:"created_at"`                         // 创建时间
	UpdatedAt time.Time  `json:"updated_at"`                         // 更新时间
	DeletedAt *time.Time `json:"deleted_at"`                         // 删除时间
}

func (AdminUser) TableName() string {
	return "t_admin_users" // 指定表名为 t_admin_users
}
