package db

import (
	"errors"
	"shadiao/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// =================== 分类管理 ===================

// CreateCategory 创建分类
func CreateCategory(category *models.Category) error {
	if err := DB.Create(category).Error; err != nil {
		log.Error().Err(err).Str("name", category.Name).Msg("创建分类失败")
		return err
	}
	log.Info().Uint("id", category.ID).Str("name", category.Name).Msg("创建分类成功")
	return nil
}

// GetCategoryByID 根据ID获取分类（预加载视频）
func GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := DB.
		Preload("Videos", "reviewed = ?", true). // 只加载审核通过的视频（可按需去掉）
		Preload("Videos.Author").
		Preload("Videos.Tags").
		First(&category, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Uint("id", id).Msg("未找到对应的分类")
		} else {
			log.Error().Err(err).Uint("id", id).Msg("获取分类失败")
		}
		return nil, err
	}

	log.Info().Uint("id", category.ID).Str("name", category.Name).Msg("获取分类成功")
	return &category, nil
}

// GetCategoriesByType 根据类型获取分类
func GetCategoriesByType(categoryType string) ([]models.Category, error) {
	var categories []models.Category
	if err := DB.
		Where("type = ?", categoryType).
		Find(&categories).Error; err != nil {

		log.Error().Err(err).Str("type", categoryType).Msg("按类型获取分类失败")
		return nil, err
	}

	log.Info().Str("type", categoryType).Int("count", len(categories)).Msg("按类型获取分类成功")
	return categories, nil
}

// GetAllCategories 获取所有分类
func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := DB.
		Preload("Videos", "reviewed = ?", true).
		Find(&categories).Error; err != nil {

		log.Error().Err(err).Msg("获取所有分类失败")
		return nil, err
	}

	log.Info().Int("count", len(categories)).Msg("获取所有分类成功")
	return categories, nil
}

// UpdateCategory 更新分类信息
func UpdateCategory(category *models.Category) error {
	if err := DB.Save(category).Error; err != nil {
		log.Error().Err(err).Uint("id", category.ID).Msg("更新分类失败")
		return err
	}
	log.Info().Uint("id", category.ID).Str("name", category.Name).Msg("更新分类成功")
	return nil
}

// DeleteCategory 删除分类（自动解除与视频的关联）
func DeleteCategory(id uint) error {
	if err := DB.Delete(&models.Category{}, id).Error; err != nil {
		log.Error().Err(err).Uint("id", id).Msg("删除分类失败")
		return err
	}
	log.Info().Uint("id", id).Msg("删除分类成功")
	return nil
}
