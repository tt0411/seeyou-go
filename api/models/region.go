package models

import "gorm.io/gorm"

// Region 表示区域表
type Region struct {
	gorm.Model
	RegionCode string  `gorm:"size:20" json:"region_code"` // 区域编码
	RegionName string  `gorm:"size:40" json:"region_name"` // 区域名称
	ParentCode int     `json:"parent_code"`                // 区域上级标识
	SimpleName string  `gorm:"size:40" json:"simple_name"` // 地名简称
	Level      int     `json:"level"`                      // 区域等级
	CityCode   string  `gorm:"size:20" json:"city_code"`   // 城市编码
	ZipCode    string  `gorm:"size:20" json:"zip_code"`    // 邮政编码
	MerName    string  `gorm:"size:100" json:"mer_name"`   // 组合名称
	Lng        float64 `json:"lng"`                        // 经度
	Lat        float64 `json:"lat"`                        // 纬度
	Pinyin     string  `gorm:"size:100" json:"pinyin"`     // 拼音
}

func (Region) TableName() string {
	return "t_region" // 指定表名为 t_region
}
