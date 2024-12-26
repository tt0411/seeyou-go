package models

import (
	"time"

	"gorm.io/gorm"
)

// Message 表示消息表
type Message struct {
	gorm.Model
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`    // 自增主键
	ReceiverID int       `gorm:"not null" json:"receiver_id"`           // 接收者id
	FriendID   int       `json:"friend_id"`                             // 申请添加好友id(type = 3时用到)
	MsgContent string    `gorm:"type:text;not null" json:"msg_content"` // 消息内容
	Type       int       `gorm:"not null" json:"type"`                  // 消息类型 1：评论消息， 2：点赞消息，3：好友添加消息，4：系统消息
	IsRead     int       `gorm:"not null" json:"is_read"`               // 是否已读 1：已读，0：未读
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`            // 创建时间
	UpdatedAt  time.Time `gorm:"not null" json:"updated_at"`            // 修改时间
}

func (Message) TableName() string {
	return "t_messages" // 指定表名为 t_messages
}
