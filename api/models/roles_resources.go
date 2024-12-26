package models

import (
	"time"

	"gorm.io/gorm"
)

// RoleResource 表示角色资源关系表
type RoleResource struct {
	gorm.Model
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`  // 自增主键
	RoleID     int       `gorm:"not null" json:"role_id"`             // 角色id
	ResourceID string    `gorm:"size:60;not null" json:"resource_id"` // 资源id
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`          // 创建时间
	UpdatedAt  time.Time `gorm:"not null" json:"updated_at"`          // 修改时间
}

func (RoleResource) TableName() string {
	return "t_roles_resources" // 指定表名为 t_roles_resources
}
