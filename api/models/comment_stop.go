package models

import (
	"time"

	"gorm.io/gorm"
)

// CommentStop 表示用户禁言表
type CommentStop struct {
	gorm.Model
	ID               int        `gorm:"primaryKey;autoIncrement" json:"id"` // 自增主键
	GroupID          int        `gorm:"not null" json:"group_id"`           // 群id
	UserID           int        `gorm:"not null" json:"user_id"`            // 用户id
	IsStop           int        `gorm:"not null" json:"is_stop"`            // 是否禁言 0：否，1：是
	CommentStartTime *time.Time `json:"comment_start_time"`                 // 开始禁言时间
	CommentEndTime   *time.Time `json:"comment_end_time"`                   // 结束禁言时间
	Remark           string     `gorm:"size:255" json:"remark"`             // 备注
	Operator         int        `json:"operator"`                           // 操作人
	CreatedAt        time.Time  `gorm:"not null" json:"created_at"`         // 创建时间
	UpdatedAt        time.Time  `gorm:"not null" json:"updated_at"`         // 修改时间
}

func (CommentStop) TableName() string {
	return "t_comment_stop" // 指定表名为 t_comment_stop
}
