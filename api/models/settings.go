package models

import (
	"time"

	"gorm.io/gorm"
)

// Setting 表示系统设置表
type Setting struct {
	gorm.Model
	ID                 int       `gorm:"primaryKey;autoIncrement" json:"id"`   // id
	Title              string    `gorm:"size:50" json:"title"`                 // 系统名称
	TitleEn            string    `gorm:"size:100" json:"title_en"`             // 系统英文名称
	Slogan             string    `gorm:"size:255" json:"slogan"`               // 标语
	Logo               string    `gorm:"size:255" json:"logo"`                 // 系统logo
	Version            string    `gorm:"size:50" json:"version"`               // 版本号
	IconfontProjectUrl string    `gorm:"size:100" json:"iconfont_project_url"` // iconfont项目地址
	IconfontUrl        string    `gorm:"size:100" json:"iconfont_url"`         // iconfont js在线地址
	CreatedAt          time.Time `json:"created_at"`                           // 创建时间
	UpdatedAt          time.Time `json:"updated_at"`                           // 更新时间
}

func (Setting) TableName() string {
	return "t_settings" // 指定表名为 t_settings
}
