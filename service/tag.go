package service

import (
	"net/http"
	"shadiao/db"
	"shadiao/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

// ============== Tag Service ==============

// CreateTag 创建标签
func CreateTag(c *gin.Context) {
	var input models.Tag

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON for create tag")
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 验证必填字段
	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "标签名称不能为空"})
		return
	}

	// 检查标签名是否已存在
	existing, _ := db.GetTagByName(input.Name)
	if existing != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "标签名称已存在"})
		return
	}

	// 创建标签
	if err := db.CreateTag(&input); err != nil {
		log.Error().Err(err).Str("name", input.Name).Msg("Failed to create tag")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建标签失败"})
		return
	}

	log.Info().Uint("id", input.ID).Str("name", input.Name).Msg("Tag created successfully")
	c.JSON(http.StatusCreated, gin.H{
		"message": "标签创建成功",
		"tag":     input,
	})
}

// GetAllTags 获取所有标签
func GetAllTags(c *gin.Context) {
	tags, err := db.GetAllTags()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all tags")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取标签列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tags": tags})
}

// GetTagByID 根据 ID 获取标签
func GetTagByID(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的标签 ID"})
		return
	}

	tag, err := db.GetTagByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Tag not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "标签不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tag": tag})
}

// UpdateTag 更新标签
func UpdateTag(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的标签 ID"})
		return
	}

	var input models.Tag
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON for update tag")
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 获取现有标签
	tag, err := db.GetTagByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Tag not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "标签不存在"})
		return
	}

	// 更新字段
	if input.Name != "" {
		// 检查新名称是否已被其他标签使用
		existing, _ := db.GetTagByName(input.Name)
		if existing != nil && existing.ID != id {
			c.JSON(http.StatusBadRequest, gin.H{"error": "标签名称已存在"})
			return
		}
		tag.Name = input.Name
	}
	if input.Description != "" {
		tag.Description = input.Description
	}

	if input.TypeDisplayName != "" {
		tag.TypeDisplayName = input.TypeDisplayName
	}

	if input.DisplayName != "" {
		tag.DisplayName = input.DisplayName
	}

	if input.Icon != "" {
		tag.Icon = input.Icon
	}

	// 更新标签
	if err := db.UpdateTag(tag); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to update tag")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新标签失败"})
		return
	}

	log.Info().Uint("id", id).Msg("Tag updated successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"tag":     tag,
	})
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的标签 ID"})
		return
	}

	// 检查标签是否存在
	_, err := db.GetTagByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Tag not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "标签不存在"})
		return
	}

	// 检查是否有关联视频
	videoCount, _ := db.CountVideosByTag(id)
	if videoCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "该标签下还有视频，无法删除",
			"videoCount": videoCount,
		})
		return
	}

	// 删除标签
	if err := db.DeleteTag(id); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to delete tag")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除标签失败"})
		return
	}

	log.Info().Uint("id", id).Msg("Tag deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// GetTagVideos 获取标签关联的所有视频
func GetTagVideos(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的标签 ID"})
		return
	}

	videos, err := db.GetVideosByTag(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to get tag videos")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取标签视频失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"videos": videos,
		"total":  len(videos),
	})
}
