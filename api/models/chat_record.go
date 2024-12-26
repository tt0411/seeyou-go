package models

import (
	"time"

	"gorm.io/gorm"
)

// ChatRecord 表示一对一消息表
type ChatRecord struct {
	gorm.Model
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`     // 自增主键
	PostMessage string    `gorm:"type:text;not null" json:"post_message"` // 消息内容
	Status      int       `gorm:"not null" json:"status"`                 // 消息状态 0：未读，1：已读
	MsgType     int       `gorm:"not null" json:"msg_type"`               // 消息类型 1：文字，2：图片，3：音频，4：视频，5：位置
	FromID      int       `gorm:"not null" json:"from_id"`                // 发送者id
	ToID        int       `gorm:"not null" json:"to_id"`                  // 接收者id
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`             // 创建时间
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`             // 修改时间
}

func (ChatRecord) TableName() string {
	return "t_chat_record" // 指定表名为 t_chat_record
}
