package models

import (
	"time"

	"gorm.io/gorm"
)

// Check 表示审核表
type Check struct {
	gorm.Model
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`    // 自增主键
	Type        int       `gorm:"not null" json:"type"`                  // 类型 1：内容审核，2：主题审核
	ObjID       int       `gorm:"not null" json:"obj_id"`                // 审核对象id
	CheckerID   int       `gorm:"not null" json:"checker_id"`            // 审核人Id
	CheckerName string    `gorm:"size:255;not null" json:"checker_name"` // 审核人姓名
	Remark      string    `gorm:"size:255" json:"remark"`                // 审核意见
	Status      int       `gorm:"not null" json:"status"`                // 审核结果 0：待审核 1：审核通过，2：审核不通过
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`            // 审核时间
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`            // 修改时间
}

func (Check) TableName() string {
	return "t_check" // 指定表名为 t_check
}
