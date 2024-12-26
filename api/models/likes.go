package models

import (
	"time"

	"gorm.io/gorm"
)

// Like 表示用户点赞记录
type Likes struct {
	gorm.Model
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`                       // 点赞记录的唯一标识
	UserID     string    `gorm:"not null" json:"user_id"`                                  // 点赞用户的ID
	TargetType string    `gorm:"type:enum('post', 'comment');not null" json:"target_type"` // 点赞对象的类型
	TargetID   string    `gorm:"not null" json:"target_id"`                                // 点赞对象的ID
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`              // 点赞时间
}

// TableName 设置表名为 t_likes
func (Likes) TableName() string {
	return "t_likes"
}
