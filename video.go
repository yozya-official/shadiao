package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建视频（包含作者）
func createVideo(c *gin.Context) {
	var input struct {
		Video  Video  `json:"video"`
		Author Author `json:"author"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 确认作者存在，不存在则创建
	var author Author
	if err := db.Where("uid = ?", input.Author.UID).First(&author).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			author = input.Author
			if err := db.Create(&author).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	video := input.Video
	video.AuthorID = author.ID

	if err := db.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"video": video, "author": author})
}

// 获取所有视频
func getAllVideos(c *gin.Context) {
	var videos []Video
	if err := db.Preload("Author").Order("created_at DESC").Where("reviewed = ?", true).Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)
}

// 获取所有未审核视频
func getUnreviewedVideos(c *gin.Context) {
	var videos []Video
	if err := db.Preload("Author").
		Where("reviewed = ?", false). // 未审核
		Order("created_at DESC").
		Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)
}

// 修改视频审核状态（即修改reviewed字段）
func updateVideoReviewStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid video id"})
		return
	}

	var req struct {
		Reviewed bool `json:"reviewed"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&Video{}).Where("id = ?", id).Update("reviewed", req.Reviewed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update success"})
}

// 获取单个视频
func getVideoByID(c *gin.Context) {
	id := c.Param("id")
	var video Video

	if err := db.Preload("Author").Where("reviewed = ?", true).First(&video, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "视频不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, video)
}

// 更新视频
func updateVideo(c *gin.Context) {
	id := c.Param("id")
	var input Video
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var video Video
	if err := db.Where("reviewed = ?", true).First(&video, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "视频不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	video.Title = input.Title
	video.AID = input.AID
	video.URL = input.URL
	video.Cover = input.Cover
	video.Description = input.Description
	video.Duration = input.Duration
	video.Views = input.Views
	video.IsOriginal = input.IsOriginal
	video.IsCompleted = input.IsCompleted
	video.Background = input.Background
	video.World = input.World
	video.HasSystem = input.HasSystem
	video.Ctime = input.Ctime
	video.Style = input.Style

	if err := db.Save(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	video.Style = input.Style
	c.JSON(http.StatusOK, video)
}

// 删除视频
func deleteVideo(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Video{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
