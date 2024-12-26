package models

import "gorm.io/gorm"

// DictType 表示字典分类表
type DictType struct {
	gorm.Model
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"` // 自增id
	Key    string `gorm:"size:50;not null" json:"key"`        // 字典分类key
	Value  string `gorm:"size:50" json:"value"`               // 字典分类名称
	Remark string `gorm:"size:255" json:"remark"`             // 备注
	Order  int    `json:"order"`                              // 排序
}

func (DictType) TableName() string {
	return "t_dict_type" // 指定表名为 t_dict_type
}
