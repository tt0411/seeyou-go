package models

import (
	"time"

	"gorm.io/gorm"
)

// Content 表示内容表
type Content struct {
	gorm.Model
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"` // 自增主键
	Title     string     `gorm:"type:text" json:"title"`             // 标题
	Content   string     `gorm:"type:text" json:"content"`           // 内容
	Type      int        `gorm:"not null" json:"type"`               // 类型 1：微内容，2：文章，3：视频
	VideoID   string     `gorm:"size:255" json:"video_id"`           // 视频id
	ImgIDs    string     `gorm:"size:255" json:"img_ids"`            // 图片ids，多个用 ',' 隔开
	TopicIDs  string     `gorm:"size:255" json:"topic_ids"`          // 话题ids,多个用 ',' 隔开
	OpenType  int        `gorm:"not null" json:"open_type"`          // 公开类型 1：公开(所有人可见)，2：私密(仅自己可见)，3：粉丝可见，4：好友可见
	IsComment int        `gorm:"not null" json:"is_comment"`         // 是否开启评论 1：开启，0：不开启
	IsDraft   int        `gorm:"not null" json:"is_draft"`           // 是否草稿 0：否，1：是
	IsTop     int        `gorm:"default:0;not null" json:"is_top"`   // 是否置顶 0：不置顶，1：置顶
	Likes     int        `gorm:"default:0;not null" json:"likes"`    // 点赞数量
	Sort      int        `json:"sort"`                               // 排序
	Status    int        `gorm:"not null" json:"status"`             // 状态 1：待审核，2：审核中，3：审核通过，4：审核不通过 5: 不需要审核
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`         // 创建时间
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`         // 更改时间
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`            // 删除时间
}

func (Content) TableName() string {
	return "t_contents" // 指定表名为 t_contents
}
