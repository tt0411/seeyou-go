package models

import (
	"time"

	"gorm.io/gorm"
)

// Topic 表示话题表
type Topic struct {
	gorm.Model
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"` // 自增主键
	Category  int        `json:"category"`                           // 分类
	TopicName string     `gorm:"size:50;not null" json:"topic_name"` // 话题名称
	IsTop     int        `gorm:"default:0;not null" json:"is_top"`   // 是否置顶 0：不置顶，1：置顶
	Sort      int        `json:"sort"`                               // 排序
	Status    int        `gorm:"not null" json:"status"`             // 状态 1：待审核，2：审核通过，3：审核不通过
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`         // 创建时间
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`         // 更改时间
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`            // 删除时间
}

func (Topic) TableName() string {
	return "t_topics" // 指定表名为 t_topics
}
