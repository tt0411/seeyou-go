package models

import (
	"time"

	"gorm.io/gorm"
)

// Friend 表示好友表
type Friend struct {
	gorm.Model
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"` // 自增主键
	UserID    int       `gorm:"not null" json:"user_id"`            // 用户id
	FriendID  int       `gorm:"not null" json:"friend_id"`          // 好友id
	CreatedAt time.Time `gorm:"not null" json:"created_at"`         // 创建时间
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`         // 修改时间
}

func (Friend) TableName() string {
	return "t_friends" // 指定表名为 t_friends
}
