package models

import (
	"time"

	"gorm.io/gorm"
)

// Resource 表示资源表
type Resource struct {
	gorm.Model
	ID        string    `gorm:"primaryKey;size:60;not null" json:"id"`         // 主键id
	Name      string    `gorm:"size:60;not null" json:"name"`                  // 资源名称
	NameCN    string    `gorm:"size:60;not null" json:"name_cn"`               // 资源中文名称
	Path      string    `gorm:"size:100;default:''" json:"path"`               // 资源路径
	ParentID  string    `gorm:"size:50;default:'0';not null" json:"parent_id"` // 父菜单id, 一级菜单为 0
	Kind      int       `gorm:"not null" json:"kind"`                          // 资源类型 1-菜单, 2-操作, 3-外部菜单, 4-目录, 5-按钮
	Icon      string    `gorm:"size:255" json:"icon"`                          // 图标
	Hidden    int       `gorm:"default:0;not null" json:"hidden"`              // 是否隐藏此路由 1-是, 0-否
	Status    int       `gorm:"default:1;not null" json:"status"`              // 菜单状态 1-启用, 2-停用
	Component string    `gorm:"size:100;default:''" json:"component"`          // 路由对应的组件
	Location  string    `gorm:"size:255" json:"location"`                      // 外部菜单url
	Perms     string    `gorm:"size:100" json:"perms"`                         // 按钮,操作权限标识
	Access    string    `gorm:"size:60;default:'permission'" json:"access"`    // 路由权限标识
	Remark    string    `gorm:"size:255" json:"remark"`                        // 备注
	Sort      int       `gorm:"default:0;not null" json:"sort"`                // 排序
	CreatedAt time.Time `gorm:"not null" json:"created_at"`                    // 创建时间
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`                    // 更新时间
}

func (Resource) TableName() string {
	return "t_resources" // 指定表名为 t_resources
}
