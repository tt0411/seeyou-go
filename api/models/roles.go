package models

import (
	"time"

	"gorm.io/gorm"
)

// Role 表示角色表
type Role struct {
	gorm.Model
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`   // 自增id
	RoleName  string    `gorm:"size:20;not null" json:"role_name"`    // 角色名称
	IsDefault int       `gorm:"default:0;not null" json:"is_default"` // 是否为默认角色 0、否，1、是
	Remark    string    `gorm:"size:255" json:"remark"`               // 备注
	CreatedAt time.Time `gorm:"not null" json:"created_at"`           // 创建时间
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`           // 更新时间
}

func (Role) TableName() string {
	return "t_roles" // 指定表名为 t_roles
}
