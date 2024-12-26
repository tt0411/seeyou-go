package models

import (
	"time"

	"gorm.io/gorm"
)

// UserSetting 表示用户设置表
type UserSetting struct {
	gorm.Model
	ID              int       `gorm:"primaryKey;autoIncrement" json:"id"`          // 自增主键
	UserID          int       `gorm:"not null" json:"user_id"`                     // 用户id
	IsLikedNotice   int       `gorm:"default:1;not null" json:"is_liked_notice"`   // 是否开启点赞消息通知 1：开启，0：关闭
	IsCommentNotice int       `gorm:"default:1;not null" json:"is_comment_notice"` // 是否开启评论消息通知 1：开启，0：关闭
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`                  // 创建时间
	UpdatedAt       time.Time `gorm:"not null" json:"updated_at"`                  // 修改时间
}

func (UserSetting) TableName() string {
	return "t_user_settings" // 指定表名为 t_user_settings
}
