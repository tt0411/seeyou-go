package models

import (
	"time"

	"gorm.io/gorm"
)

// Reply 表示回复表
type Reply struct {
	gorm.Model
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`      // 自增主键
	CommentID    int       `gorm:"not null" json:"comment_id"`              // 评论id
	FromUID      int       `gorm:"not null" json:"from_uid"`                // 回复者id
	ReplyContent string    `gorm:"type:text;not null" json:"reply_content"` // 回复内容
	Likes        int       `gorm:"default:0;not null" json:"likes"`         // 点赞数量
	Status       int       `gorm:"not null" json:"status"`                  // 状态 1：正常，0：违规被封禁
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`              // 创建时间
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`              // 修改时间
}

func (Reply) TableName() string {
	return "t_replies" // 指定表名为 t_replies
}
