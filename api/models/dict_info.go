package models

import "gorm.io/gorm"

// DictInfo 表示字典表
type DictInfo struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"` // 自增id
	TypeID   int    `gorm:"not null" json:"type_id"`            // 字典分类id
	ParentID int    `json:"parent_id"`                          // 级联字典id
	Label    string `gorm:"size:50;not null" json:"label"`      // 字典名称
	Value    string `gorm:"size:50;not null" json:"value"`      // 字典值
	Order    int    `json:"order"`                              // 排序
	Remark   string `gorm:"size:255" json:"remark"`             // 备注
}

func (DictInfo) TableName() string {
	return "t_dict_info" // 指定表名为 t_dict_info
}
