package models

import (
	"time"

	"gorm.io/gorm"
)

// UserStop 表示用户封号表
type UserStop struct {
	gorm.Model
	ID            int        `gorm:"primaryKey;autoIncrement" json:"id"` // 自增主键
	UserID        int        `gorm:"not null" json:"user_id"`            // 用户id
	IsStop        int        `gorm:"not null" json:"is_stop"`            // 是否封号 0：否，1：是
	StopStartTime *time.Time `json:"stop_start_time"`                    // 开始封号时间
	StopEndTime   *time.Time `json:"stop_end_time"`                      // 结束封号时间
	Remark        string     `gorm:"size:255" json:"remark"`             // 备注
	Operator      int        `json:"operator"`                           // 操作人
	CreatedAt     time.Time  `gorm:"not null" json:"created_at"`         // 创建时间
	UpdatedAt     time.Time  `gorm:"not null" json:"updated_at"`         // 修改时间
}

func (UserStop) TableName() string {
	return "t_user_stop" // 指定表名为 t_user_stop
}
