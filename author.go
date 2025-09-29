package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建作者
func createAuthor(c *gin.Context) {
	var author Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 预加载代表作
	if author.MasterID != nil {
		db.Preload("Master").First(&author, author.ID)
	}

	c.JSON(http.StatusOK, author)
}

// 获取所有作者
func getAllAuthors(c *gin.Context) {
	var authors []Author
	if err := db.Preload("Master").Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authors)
}

// 根据ID获取作者
func getAuthorByID(c *gin.Context) {
	id := c.Param("id")
	var author Author

	if err := db.Preload("Master").Preload("Videos", "reviewed = ?", true).First(&author, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, author)
}

// 更新作者
func updateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author Author

	if err := db.First(&author, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var input struct {
		UID      int    `json:"uid"`
		Name     string `json:"name"`
		Avatar   string `json:"avatar"`
		MasterID *uint  `json:"masterId"` // 可以更新代表作
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author.UID = input.UID
	author.Name = input.Name
	author.Avatar = input.Avatar
	author.MasterID = input.MasterID

	if err := db.Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.Preload("Master").First(&author, author.ID)
	c.JSON(http.StatusOK, author)
}

// 设置作者代表作
func setMaster(c *gin.Context) {
	authorIDStr := c.Param("id")
	authorID, err := strconv.Atoi(authorIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var body struct {
		VideoID uint `json:"videoId"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var author Author
	if err := db.First(&author, authorID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// 检查视频是否存在且属于该作者
	var video Video
	if err := db.First(&video, body.VideoID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	if video.AuthorID != author.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Video does not belong to the author"})
		return
	}

	// 更新代表作
	author.MasterID = &video.ID
	if err := db.Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回带代表作的作者信息
	db.Preload("Master").First(&author, author.ID)
	c.JSON(http.StatusOK, author)
}

// 删除作者
func deleteAuthor(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := db.Delete(&Author{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
