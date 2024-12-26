package models

import (
	"time"

	"gorm.io/gorm"
)

// RoleUser 表示角色用户关系表
type RoleUser struct {
	gorm.Model
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"` // 自增id
	RoleID    int       `gorm:"not null" json:"role_id"`            // 角色id
	UserID    int       `gorm:"not null" json:"user_id"`            // 用户id
	CreatedAt time.Time `gorm:"not null" json:"created_at"`         // 创建时间
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`         // 更新时间
}

func (RoleUser) TableName() string {
	return "t_roles_users" // 指定表名为 t_roles_users
}
