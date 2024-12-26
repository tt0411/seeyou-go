package models

import (
	"time"

	"gorm.io/gorm"
)

// UserRelation 表示用户关系表(关注，粉丝)
type UserRelation struct {
	gorm.Model
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"` // 自增主键
	UserID       int       `gorm:"not null" json:"user_id"`            // 用户id
	FollowerID   int       `gorm:"not null" json:"follower_id"`        // 被关注者用户id
	RelationType int       `gorm:"not null" json:"relation_type"`      // 关系类型 1：关注，2：粉丝
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`         // 创建时间
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`         // 修改时间
}

func (UserRelation) TableName() string {
	return "t_user_relation" // 指定表名为 t_user_relation
}
