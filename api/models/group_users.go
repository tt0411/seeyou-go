package models

import (
	"time"

	"gorm.io/gorm"
)

// GroupUser 表示群成员表
type GroupUser struct {
	gorm.Model
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`     // 自增主键
	GroupID     int       `gorm:"not null" json:"group_id"`               // 群id
	UserID      int       `gorm:"not null" json:"user_id"`                // 用户id
	GroupUname  string    `gorm:"size:50" json:"group_uname"`             // 群内名
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`             // 加入时间
	NoreadCount int       `gorm:"default:0;not null" json:"noread_count"` // 未读消息数
	IDShield    int       `gorm:"default:0;not null" json:"id_shield"`    // 是否屏蔽群消息 0：不屏蔽，1：屏蔽
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`             // 修改时间
}

func (GroupUser) TableName() string {
	return "t_group_users" // 指定表名为 t_group_users
}
