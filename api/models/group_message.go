package models

import (
	"time"

	"gorm.io/gorm"
)

// GroupMessage 表示群消息表
type GroupMessage struct {
	gorm.Model
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"` // 主键
	GroupID   int       `gorm:"not null" json:"group_id"`           // 群id
	SendID    int       `gorm:"not null" json:"send_id"`            // 发送者id
	SendMsg   string    `gorm:"type:text;not null" json:"send_msg"` // 发送消息
	MsgType   int       `gorm:"not null" json:"msg_type"`           // 消息类型  1：文字，2：图片，3：音频，4：视频，5：位置
	CreatedAt time.Time `gorm:"not null" json:"created_at"`         // 发送时间
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`         // 修改时间
}

func (GroupMessage) TableName() string {
	return "t_group_message" // 指定表名为 t_group_message
}
