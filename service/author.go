package service

import (
	"net/http"
	"shadiao/db"
	"shadiao/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

// CreateAuthor 创建作者
func CreateAuthor(c *gin.Context) {
	var input models.Author

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON for create author")
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 验证必填字段
	if input.UID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UID 不能为空"})
		return
	}
	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "作者名称不能为空"})
		return
	}

	// 检查 UID 是否已存在
	existing, _ := db.GetAuthorByUID(input.UID)
	if existing != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该 UID 已存在"})
		return
	}

	// 创建作者
	if err := db.CreateAuthor(&input); err != nil {
		log.Error().Err(err).Int("uid", input.UID).Msg("Failed to create author")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建作者失败"})
		return
	}

	log.Info().Uint("id", input.ID).Int("uid", input.UID).Msg("Author created successfully")
	c.JSON(http.StatusCreated, gin.H{
		"message": "作者创建成功",
		"author":  input,
	})
}

// GetAllAuthors 获取所有作者
func GetAllAuthors(c *gin.Context) {
	authors, err := db.GetAllAuthors()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all authors")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取作者列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"authors": authors})
}

// GetAuthorByID 根据 ID 获取作者
func GetAuthorByID(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的作者 ID"})
		return
	}

	author, err := db.GetAuthorByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Author not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "作者不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"author": author})
}

// GetAuthorByUID 根据 UID 获取作者
func GetAuthorByUID(c *gin.Context) {
	uid := cast.ToInt(c.Param("uid"))

	if uid == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 UID"})
		return
	}

	author, err := db.GetAuthorByUID(uid)
	if err != nil {
		log.Error().Err(err).Int("uid", uid).Msg("Author not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "作者不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"author": author})
}

// UpdateAuthor 更新作者信息
func UpdateAuthor(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的作者 ID"})
		return
	}

	var input models.Author
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON for update author")
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 获取现有作者
	author, err := db.GetAuthorByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Author not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "作者不存在"})
		return
	}

	// 更新字段
	if input.Name != "" {
		author.Name = input.Name
	}
	if input.Avatar != "" {
		author.Avatar = input.Avatar
	}
	// UID 不允许修改

	// 更新作者
	if err := db.UpdateAuthor(author); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to update author")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新作者失败"})
		return
	}

	log.Info().Uint("id", id).Msg("Author updated successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"author":  author,
	})
}

// DeleteAuthor 删除作者
func DeleteAuthor(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的作者 ID"})
		return
	}

	// 检查作者是否存在
	_, err := db.GetAuthorByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Author not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "作者不存在"})
		return
	}

	// 检查是否有关联视频
	videoCount, _ := db.CountVideosByAuthor(id)
	if videoCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "该作者下还有视频，无法删除",
			"videoCount": videoCount,
		})
		return
	}

	// 删除作者
	if err := db.DeleteAuthor(id); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to delete author")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除作者失败"})
		return
	}

	log.Info().Uint("id", id).Msg("Author deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// GetAuthorVideos 获取作者的所有视频
func GetAuthorVideos(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的作者 ID"})
		return
	}

	videos, err := db.GetVideosByAuthor(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to get author videos")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取作者视频失败"})
		return
	}

	// 获取总数
	total, _ := db.CountVideosByAuthor(id)

	c.JSON(http.StatusOK, gin.H{
		"videos": videos,
		"total":  total,
	})
}
