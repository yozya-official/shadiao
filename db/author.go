package db

import (
	"errors"
	"shadiao/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// =================== 作者管理 ===================

// CreateAuthor 创建作者
func CreateAuthor(author *models.Author) error {
	if err := DB.Create(author).Error; err != nil {
		log.Error().Err(err).Int("uid", author.UID).Msg("创建作者失败")
		return err
	}
	log.Info().Uint("id", author.ID).Int("uid", author.UID).Str("name", author.Name).Msg("创建作者成功")
	return nil
}

// GetAuthorByID 根据ID获取作者（包含视频、标签、分类）
func GetAuthorByID(id uint) (*models.Author, error) {
	var author models.Author
	if err := DB.
		Preload("Videos", "reviewed = ?", true). // 仅加载已审核视频
		Preload("Videos.Author").
		Preload("Videos.Tags").
		Preload("Videos.Categories").
		First(&author, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Uint("id", id).Msg("未找到对应的作者")
		} else {
			log.Error().Err(err).Uint("id", id).Msg("获取作者信息失败")
		}
		return nil, err
	}
	log.Info().Uint("id", author.ID).Str("name", author.Name).Msg("获取作者信息成功")
	return &author, nil
}

// GetAuthorByUID 根据B站UID获取作者（包含视频、标签、分类）
func GetAuthorByUID(uid int) (*models.Author, error) {
	var author models.Author
	if err := DB.
		Preload("Videos", "reviewed = ?", true).
		Preload("Videos.Author").
		Preload("Videos.Tags").
		Preload("Videos.Categories").
		Where("uid = ?", uid).
		First(&author).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Int("uid", uid).Msg("未找到对应的作者")
		} else {
			log.Error().Err(err).Int("uid", uid).Msg("获取作者信息失败")
		}
		return nil, err
	}
	log.Info().Uint("id", author.ID).Int("uid", uid).Str("name", author.Name).Msg("获取作者信息成功")
	return &author, nil
}

// GetAllAuthors 获取所有作者（预加载视频数量）
func GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	if err := DB.
		Preload("Videos", "reviewed = ?", true).
		Find(&authors).Error; err != nil {

		log.Error().Err(err).Msg("获取所有作者失败")
		return nil, err
	}
	log.Info().Int("count", len(authors)).Msg("获取所有作者成功")
	return authors, nil
}

// UpdateAuthor 更新作者信息
func UpdateAuthor(author *models.Author) error {
	if err := DB.Save(author).Error; err != nil {
		log.Error().Err(err).Uint("id", author.ID).Msg("更新作者失败")
		return err
	}
	log.Info().Uint("id", author.ID).Int("uid", author.UID).Str("name", author.Name).Msg("更新作者成功")
	return nil
}

// DeleteAuthor 删除作者（自动级联删除关联视频）
func DeleteAuthor(id uint) error {
	if err := DB.Delete(&models.Author{}, id).Error; err != nil {
		log.Error().Err(err).Uint("id", id).Msg("删除作者失败")
		return err
	}
	log.Info().Uint("id", id).Msg("删除作者成功（级联删除其视频）")
	return nil
}
