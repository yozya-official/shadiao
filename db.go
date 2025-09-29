package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Author struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UID       int       `gorm:"unique;not null" json:"uid"`
	Name      string    `gorm:"not null" json:"name"`
	Avatar    string    `json:"avatar"`
	Videos    []Video   `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE" json:"videos"` // 关联视频
	MasterID  *uint     `json:"masterId"`                                                      // 代表作视频ID，可为空
	Master    *Video    `gorm:"foreignKey:MasterID" json:"master"`                             // 代表作视频关联
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Video struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	AID         int    `gorm:"unique;not null" json:"aid"`
	Title       string `gorm:"not null" json:"title"`
	AuthorID    uint   `gorm:"not null" json:"authorId"` // 外键
	Author      Author `json:"author"`
	URL         string `gorm:"not null" json:"url"`
	Cover       string `json:"cover"`
	Description string `json:"description"`

	Duration    int  `json:"duration"` // 30 60 120 999
	Views       int  `json:"views"`    // 1 10 100 1000 9999
	IsOriginal  bool `gorm:"default:false" json:"isOriginal"`
	IsCompleted bool `gorm:"default:false" json:"isCompleted"`

	Background string         `json:"background"`
	World      string         `json:"world"`
	Style      datatypes.JSON `json:"-"`              // 存数据库用 JSON
	StyleList  []string       `gorm:"-" json:"style"` // 给前端用，自动序列化/反序列化
	HasSystem  bool           `gorm:"default:false" json:"hasSystem"`

	Reviewed bool `gorm:"default:false" json:"reviewed"`

	Ctime     time.Time `gorm:"autoCreateTime" json:"ctime"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (v *Video) BeforeSave(tx *gorm.DB) (err error) {
	if v.StyleList != nil {
		b, err := json.Marshal(v.StyleList)
		if err != nil {
			return err
		}
		v.Style = b
	}
	return nil
}

// 在 GORM hook 里反序列化
func (v *Video) AfterFind(tx *gorm.DB) (err error) {
	if len(v.Style) > 0 {
		err := json.Unmarshal(v.Style, &v.StyleList)
		if err != nil {
			return err
		}
	}
	return nil
}

var db *gorm.DB

func initDB() {
	var err error

	if envPath := os.Getenv("DATABASE_URL"); envPath != "" {
		dbPath = envPath
	}

	db, err = gorm.Open(sqlite.Open(dbPath+"?_foreign_keys=1"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// 自动迁移
	err = db.AutoMigrate(&Author{}, &Video{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
}
