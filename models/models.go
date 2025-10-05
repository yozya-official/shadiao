package models

import "time"

// 标签表
type Tag struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"not null;uniqueIndex:tag_id_name_type" json:"name"` // 联合唯一
	Type            string    `gorm:"not null;uniqueIndex:tag_id_name_type" json:"type"` // 联合唯一
	DisplayName     string    `json:"displayName"`
	TypeDisplayName string    `json:"typeDisplayName"`
	Icon            string    `json:"icon"`
	Description     string    `json:"description"`                         // 标签描述
	Videos          []Video   `gorm:"many2many:video_tags;" json:"videos"` // 多对多关联视频
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// 分类表
type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null;uniqueIndex:category_id_name_type" json:"name"` // 联合唯一
	Type        string    `gorm:"not null;uniqueIndex:category_id_name_type" json:"type"` // 联合唯一
	DisplayName string    `json:"displayName"`
	Description string    `json:"description"`                               // 分类描述
	Videos      []Video   `gorm:"many2many:video_categories;" json:"videos"` // 多对多关联视频
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// 作者表
type Author struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UID       int       `gorm:"unique;not null" json:"uid"`                                    // B站 UID
	Name      string    `gorm:"not null" json:"name"`                                          // 作者名称
	Avatar    string    `json:"avatar"`                                                        // 头像
	Videos    []Video   `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE" json:"videos"` // 关联视频
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// 视频表
type Video struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	AID         int    `gorm:"unique;not null" json:"aid"` // B站 AV 号
	Title       string `gorm:"not null" json:"title"`      // 视频标题
	AuthorID    uint   `gorm:"not null" json:"authorId"`   // 外键
	Author      Author `json:"author"`                     // 作者
	URL         string `gorm:"not null" json:"url"`        // 视频地址
	Cover       string `json:"cover"`                      // 封面
	Description string `json:"description"`                // 描述

	Duration int `json:"duration"` // 视频长度 30 60 120 999
	Views    int `json:"views"`    // 播放量 1 10 100 1000 9999

	Tags       []Tag      `gorm:"many2many:video_tags;" json:"tags"`             // 标签
	Categories []Category `gorm:"many2many:video_categories;" json:"categories"` // 分类

	Reviewed bool `gorm:"default:false" json:"reviewed"` // 是否审核

	Ctime     time.Time `gorm:"autoCreateTime" json:"ctime"` // 创建时间
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
