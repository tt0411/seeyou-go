package models

import (
	"time"

	"gorm.io/gorm"
)

// File 表示文件表
type File struct {
	gorm.Model
	ID         int        `gorm:"primaryKey;autoIncrement" json:"id"`       // 自增主键
	Path       string     `gorm:"size:255;default:'';not null" json:"path"` // 文件路径
	Category   string     `gorm:"size:100;not null" json:"category"`        // 分类
	Name       string     `gorm:"size:255;default:'';not null" json:"name"` // 文件原始名称
	Ext        string     `gorm:"size:50;default:'';not null" json:"ext"`   // 文件后缀
	Size       int        `gorm:"default:0;not null" json:"size"`           // 文件大小
	MD5        string     `gorm:"size:32;default:'';not null" json:"md5"`   // 文件md5
	CreateTime *time.Time `json:"create_time"`                              // 创建时间
	UpdateTime *time.Time `gorm:"autoUpdateTime" json:"update_time"`        // 更新时间
	DeleteTime *time.Time `json:"delete_time"`                              // 删除时间
}

func (File) TableName() string {
	return "t_file" // 指定表名为 t_file
}