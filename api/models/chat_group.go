package models

import (
	"time"

	"gorm.io/gorm"
)

// ChatGroup 表示群表
type ChatGroup struct {
	gorm.Model
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`   // 自增主键
	GroupUID   int       `gorm:"not null" json:"group_uid"`            // 群主id
	GroupName  string    `gorm:"size:50;not null" json:"group_name"`   // 群名称
	GroupCover string    `gorm:"size:255;not null" json:"group_cover"` // 群封面
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`           // 创建时间
	UpdatedAt  time.Time `gorm:"not null" json:"updated_at"`           // 修改时间
}

func (ChatGroup) TableName() string {
	return "t_chat_group" // 指定表名为 t_chat_group
}
