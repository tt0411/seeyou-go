package models

import (
	"time"

	"gorm.io/gorm"
)

// Comment 表示评论表
type Comment struct {
	gorm.Model
	ID             int       `gorm:"primaryKey;autoIncrement" json:"id"`        // 自增主键
	ContentID      int       `gorm:"not null" json:"content_id"`                // 内容id
	CommentContent string    `gorm:"type:text;not null" json:"comment_content"` // 评论内容
	FromUID        int       `gorm:"not null" json:"from_uid"`                  // 评论者
	Likes          int       `gorm:"default:0;not null" json:"likes"`           // 点赞数量
	Status         int       `gorm:"not null" json:"status"`                    // 状态 1：正常，0：违规被封
	CreatedAt      time.Time `gorm:"not null" json:"created_at"`                // 创建时间
	UpdatedAt      time.Time `gorm:"not null" json:"updated_at"`                // 修改时间
}

func (Comment) TableName() string {
	return "t_comments" // 指定表名为 t_comments
}
