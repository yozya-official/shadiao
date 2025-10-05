package db

import (
	"errors"
	"shadiao/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// =================== 标签管理 ===================

// CreateTag 创建标签
func CreateTag(tag *models.Tag) error {
	if err := DB.Create(tag).Error; err != nil {
		log.Error().Err(err).Str("name", tag.Name).Msg("创建标签失败")
		return err
	}
	log.Info().Uint("id", tag.ID).Str("name", tag.Name).Msg("创建标签成功")
	return nil
}

// GetTagByID 根据ID获取标签（预加载视频、作者、分类）
func GetTagByID(id uint) (*models.Tag, error) {
	var tag models.Tag
	if err := DB.
		Preload("Videos", "reviewed = ?", true). // 仅加载已审核视频
		Preload("Videos.Author").
		Preload("Videos.Tags").
		Preload("Videos.Categories").
		First(&tag, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Uint("id", id).Msg("未找到对应的标签")
		} else {
			log.Error().Err(err).Uint("id", id).Msg("获取标签失败")
		}
		return nil, err
	}
	log.Info().Uint("id", tag.ID).Str("name", tag.Name).Msg("获取标签成功")
	return &tag, nil
}

// GetTagByName 根据名称获取标签（预加载视频）
func GetTagByName(name string) (*models.Tag, error) {
	var tag models.Tag
	if err := DB.
		Preload("Videos", "reviewed = ?", true).
		Preload("Videos.Author").
		Preload("Videos.Categories").
		Where("name = ?", name).
		First(&tag).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Str("name", name).Msg("未找到对应的标签")
		} else {
			log.Error().Err(err).Str("name", name).Msg("获取标签失败")
		}
		return nil, err
	}
	log.Info().Uint("id", tag.ID).Str("name", tag.Name).Msg("获取标签成功")
	return &tag, nil
}

// GetAllTags 获取所有标签（可选预加载视频数量）
func GetAllTags() ([]models.Tag, error) {
	var tags []models.Tag
	if err := DB.
		Preload("Videos", "reviewed = ?", true).
		Find(&tags).Error; err != nil {

		log.Error().Err(err).Msg("获取所有标签失败")
		return nil, err
	}
	log.Info().Int("count", len(tags)).Msg("获取所有标签成功")
	return tags, nil
}

// UpdateTag 更新标签信息
func UpdateTag(tag *models.Tag) error {
	if err := DB.Save(tag).Error; err != nil {
		log.Error().Err(err).Uint("id", tag.ID).Msg("更新标签失败")
		return err
	}
	log.Info().Uint("id", tag.ID).Str("name", tag.Name).Msg("更新标签成功")
	return nil
}

// DeleteTag 删除标签（会自动解除与视频的关联）
func DeleteTag(id uint) error {
	if err := DB.Delete(&models.Tag{}, id).Error; err != nil {
		log.Error().Err(err).Uint("id", id).Msg("删除标签失败")
		return err
	}
	log.Info().Uint("id", id).Msg("删除标签成功")
	return nil
}

func GetTagByNameAndType(name, typ string) (*models.Tag, error) {
	var tag models.Tag
	if err := DB.Where("name = ? AND type = ?", name, typ).First(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}
