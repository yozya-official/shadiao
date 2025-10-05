package service

import (
	"net/http"
	"shadiao/db"
	"shadiao/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

// ============== Category Service ==============

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	var input models.Category

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON for create category")
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 验证必填字段
	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "分类名称不能为空"})
		return
	}
	if input.Type == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "分类类型不能为空"})
		return
	}

	// 创建分类
	if err := db.CreateCategory(&input); err != nil {
		log.Error().Err(err).Str("name", input.Name).Msg("Failed to create category")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建分类失败"})
		return
	}

	log.Info().Uint("id", input.ID).Str("name", input.Name).Msg("Category created successfully")
	c.JSON(http.StatusCreated, gin.H{
		"message":  "分类创建成功",
		"category": input,
	})
}

// GetAllCategories 获取所有分类
func GetAllCategories(c *gin.Context) {
	categories, err := db.GetAllCategories()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all categories")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取分类列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// GetCategoriesByType 根据类型获取分类
func GetCategoriesByType(c *gin.Context) {
	categoryType := c.Param("type")

	if categoryType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "分类类型不能为空"})
		return
	}

	categories, err := db.GetCategoriesByType(categoryType)
	if err != nil {
		log.Error().Err(err).Str("type", categoryType).Msg("Failed to get categories by type")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取分类列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// GetCategoryByID 根据 ID 获取分类
func GetCategoryByID(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的分类 ID"})
		return
	}

	category, err := db.GetCategoryByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Category not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "分类不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

// UpdateCategory 更新分类
func UpdateCategory(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的分类 ID"})
		return
	}

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON for update category")
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 获取现有分类
	category, err := db.GetCategoryByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Category not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "分类不存在"})
		return
	}

	// 更新字段
	if input.Name != "" {
		category.Name = input.Name
	}
	if input.Description != "" {
		category.Description = input.Description
	}
	if input.Type != "" {
		category.Type = input.Type
	}

	// 更新分类
	if err := db.UpdateCategory(category); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to update category")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新分类失败"})
		return
	}

	log.Info().Uint("id", id).Msg("Category updated successfully")
	c.JSON(http.StatusOK, gin.H{
		"message":  "更新成功",
		"category": category,
	})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的分类 ID"})
		return
	}

	// 检查分类是否存在
	_, err := db.GetCategoryByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Category not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "分类不存在"})
		return
	}

	// 检查是否有关联视频
	videoCount, _ := db.CountVideosByCategory(id)
	if videoCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "该分类下还有视频，无法删除",
			"videoCount": videoCount,
		})
		return
	}

	// 删除分类
	if err := db.DeleteCategory(id); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to delete category")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除分类失败"})
		return
	}

	log.Info().Uint("id", id).Msg("Category deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// GetCategoryVideos 获取分类关联的所有视频
func GetCategoryVideos(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的分类 ID"})
		return
	}

	videos, err := db.GetVideosByCategory(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to get category videos")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取分类视频失败"})
		return
	}

	// 获取总数
	total, _ := db.CountVideosByCategory(id)

	c.JSON(http.StatusOK, gin.H{
		"videos": videos,
		"total":  total,
	})
}
