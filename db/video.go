package db

import (
	"errors"
	"shadiao/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// CreateVideo 创建视频
func CreateVideo(video *models.Video) error {
	if err := DB.Create(video).Error; err != nil {
		log.Error().Err(err).Int("aid", video.AID).Msg("创建视频失败")
		return err
	}
	log.Info().Uint("id", video.ID).Int("aid", video.AID).Str("title", video.Title).Msg("创建视频成功")
	return nil
}

// GetVideoByID 根据ID获取视频（过滤未审核视频，包含所有关联）
func GetVideoByID(id uint) (*models.Video, error) {
	var video models.Video
	if err := DB.
		Preload("Author").
		Preload("Tags").
		Preload("Categories").
		Where("id = ? AND reviewed = ?", id, true).
		First(&video).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Uint("id", id).Msg("视频未找到或未审核")
		} else {
			log.Error().Err(err).Uint("id", id).Msg("获取视频失败")
		}
		return nil, err
	}
	return &video, nil
}

// 根据ID获取视频（包含未审核视频，包含所有关联）
func GetVideoByIDAny(id uint) (*models.Video, error) {
	var video models.Video
	if err := DB.
		Preload("Author").
		Preload("Tags").
		Preload("Categories").
		Where("id = ? ", id).
		First(&video).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Uint("id", id).Msg("视频未找到或未审核")
		} else {
			log.Error().Err(err).Uint("id", id).Msg("获取视频失败")
		}
		return nil, err
	}
	return &video, nil
}

// GetVideoByAID 根据AID获取视频（过滤未审核视频）
func GetVideoByAID(aid int) (*models.Video, error) {
	var video models.Video
	if err := DB.
		Preload("Author").
		Preload("Tags").
		Preload("Categories").
		Where("aid = ? AND reviewed = ?", aid, true).
		First(&video).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Int("aid", aid).Msg("视频未找到或未审核")
		} else {
			log.Error().Err(err).Int("aid", aid).Msg("获取视频失败")
		}
		return nil, err
	}
	return &video, nil
}

// GetAllVideos 获取所有已审核视频
func GetAllVideos() ([]models.Video, error) {
	var videos []models.Video
	if err := DB.
		Preload("Author").
		Preload("Tags").
		Preload("Categories").
		Where("reviewed = ?", true).
		Find(&videos).Error; err != nil {

		log.Error().Err(err).Msg("获取所有视频失败")
		return nil, err
	}
	return videos, nil
}

// UpdateVideo 更新视频
func UpdateVideo(video *models.Video) error {
	if err := DB.Save(video).Error; err != nil {
		log.Error().Err(err).Uint("id", video.ID).Msg("更新视频失败")
		return err
	}
	log.Info().Uint("id", video.ID).Int("aid", video.AID).Msg("更新视频成功")
	return nil
}

// DeleteVideo 删除视频
func DeleteVideo(id uint) error {
	if err := DB.Delete(&models.Video{}, id).Error; err != nil {
		log.Error().Err(err).Uint("id", id).Msg("删除视频失败")
		return err
	}
	log.Info().Uint("id", id).Msg("删除视频成功")
	return nil
}

// AddTagsToVideo 为视频添加标签
func AddTagsToVideo(videoID uint, tagIDs []uint) error {
	video, err := GetVideoByID(videoID)
	if err != nil {
		return err
	}

	var tags []models.Tag
	if err := DB.Find(&tags, tagIDs).Error; err != nil {
		log.Error().Err(err).Msg("查找标签失败")
		return err
	}

	if err := DB.Model(video).Association("Tags").Append(tags); err != nil {
		log.Error().Err(err).Uint("videoID", videoID).Msg("添加标签失败")
		return err
	}

	log.Info().Uint("videoID", videoID).Interface("tagIDs", tagIDs).Msg("已为视频添加标签")
	return nil
}

// RemoveTagsFromVideo 从视频移除标签
func RemoveTagsFromVideo(videoID uint, tagIDs []uint) error {
	video, err := GetVideoByID(videoID)
	if err != nil {
		return err
	}

	var tags []models.Tag
	if err := DB.Find(&tags, tagIDs).Error; err != nil {
		log.Error().Err(err).Msg("查找标签失败")
		return err
	}

	if err := DB.Model(video).Association("Tags").Delete(tags); err != nil {
		log.Error().Err(err).Uint("videoID", videoID).Msg("移除标签失败")
		return err
	}

	log.Info().Uint("videoID", videoID).Interface("tagIDs", tagIDs).Msg("已从视频移除标签")
	return nil
}

// AddCategoriesToVideo 为视频添加分类
func AddCategoriesToVideo(videoID uint, categoryIDs []uint) error {
	video, err := GetVideoByID(videoID)
	if err != nil {
		return err
	}

	var categories []models.Category
	if err := DB.Find(&categories, categoryIDs).Error; err != nil {
		log.Error().Err(err).Msg("查找分类失败")
		return err
	}

	if err := DB.Model(video).Association("Categories").Append(categories); err != nil {
		log.Error().Err(err).Uint("videoID", videoID).Msg("添加分类失败")
		return err
	}

	log.Info().Uint("videoID", videoID).Interface("categoryIDs", categoryIDs).Msg("已为视频添加分类")
	return nil
}

// RemoveCategoriesFromVideo 从视频移除分类
func RemoveCategoriesFromVideo(videoID uint, categoryIDs []uint) error {
	video, err := GetVideoByID(videoID)
	if err != nil {
		return err
	}

	var categories []models.Category
	if err := DB.Find(&categories, categoryIDs).Error; err != nil {
		log.Error().Err(err).Msg("查找分类失败")
		return err
	}

	if err := DB.Model(video).Association("Categories").Delete(categories); err != nil {
		log.Error().Err(err).Uint("videoID", videoID).Msg("移除分类失败")
		return err
	}

	log.Info().Uint("videoID", videoID).Interface("categoryIDs", categoryIDs).Msg("已从视频移除分类")
	return nil
}

// CountVideosByAuthor 统计作者的视频数量
func CountVideosByAuthor(authorID uint) (uint, error) {
	var count int64
	if err := DB.Model(&models.Video{}).
		Where("author_id = ? AND reviewed = ?", authorID, true).
		Count(&count).Error; err != nil {

		log.Error().Err(err).Uint("authorID", authorID).Msg("统计作者视频数量失败")
		return 0, err
	}
	return uint(count), nil
}

// CountVideosByCategory 统计分类的视频数量
func CountVideosByCategory(categoryID uint) (uint, error) {
	var count int64
	if err := DB.Table("video_categories").
		Where("category_id = ?", categoryID).
		Count(&count).Error; err != nil {

		log.Error().Err(err).Uint("categoryID", categoryID).Msg("统计分类视频数量失败")
		return 0, err
	}
	return uint(count), nil
}

// GetVideosByCategory 根据分类ID获取视频
func GetVideosByCategory(categoryID uint) ([]models.Video, error) {
	var videos []models.Video
	if err := DB.
		Joins("JOIN video_categories vc ON vc.video_id = videos.id").
		Where("vc.category_id = ? AND reviewed = ?", categoryID, true).
		Preload("Author").
		Preload("Tags").
		Preload("Categories").
		Find(&videos).Error; err != nil {

		log.Error().Err(err).Uint("categoryID", categoryID).Msg("获取分类视频失败")
		return nil, err
	}
	return videos, nil
}

// CountVideosByTag 统计标签的视频数量
func CountVideosByTag(tagID uint) (uint, error) {
	var count int64
	if err := DB.Table("video_tags").
		Where("tag_id = ?", tagID).
		Count(&count).Error; err != nil {

		log.Error().Err(err).Uint("tagID", tagID).Msg("统计标签视频数量失败")
		return 0, err
	}
	return uint(count), nil
}

// GetVideosByTag 根据标签ID获取视频
func GetVideosByTag(tagID uint) ([]models.Video, error) {
	var videos []models.Video
	if err := DB.
		Joins("JOIN video_tags vt ON vt.video_id = videos.id").
		Where("vt.tag_id = ? AND reviewed = ?", tagID, true).
		Preload("Author").
		Preload("Tags").
		Preload("Categories").
		Find(&videos).Error; err != nil {

		log.Error().Err(err).Uint("tagID", tagID).Msg("获取标签视频失败")
		return nil, err
	}
	return videos, nil
}

// GetUnreviewedVideos 获取未审核视频
func GetUnreviewedVideos() ([]models.Video, error) {
	var videos []models.Video
	if err := DB.
		Preload("Author").
		Preload("Tags").
		Preload("Categories").
		Where("reviewed = ?", false).
		Find(&videos).Error; err != nil {

		log.Error().Err(err).Msg("获取未审核视频失败")
		return nil, err
	}
	return videos, nil
}

// GetVideosByAuthor 根据作者ID获取视频（已审核）
func GetVideosByAuthor(authorID uint) ([]models.Video, error) {
	var videos []models.Video
	if err := DB.
		Preload("Tags").
		Preload("Categories").
		Where("author_id = ? AND reviewed = ?", authorID, true).
		Find(&videos).Error; err != nil {

		log.Error().Err(err).Uint("authorID", authorID).Msg("获取作者视频失败")
		return nil, err
	}
	return videos, nil
}

// UpdateVideoTags 更新视频的标签（全量替换）
func UpdateVideoTags(videoID uint, tagIDs []uint) error {
	video := models.Video{ID: videoID}

	var tags []models.Tag
	if err := DB.Find(&tags, tagIDs).Error; err != nil {
		log.Error().Err(err).Msg("查找标签失败")
		return err
	}

	if err := DB.Model(&video).Association("Tags").Replace(tags); err != nil {
		log.Error().Err(err).Uint("videoID", videoID).Msg("更新视频标签失败")
		return err
	}

	log.Info().Uint("videoID", videoID).Msg("更新视频标签成功")
	return nil
}

// UpdateVideoCategories 更新视频的分类（全量替换）
func UpdateVideoCategories(videoID uint, categoryIDs []uint) error {
	video := models.Video{ID: videoID}

	var categories []models.Category
	if err := DB.Find(&categories, categoryIDs).Error; err != nil {
		log.Error().Err(err).Msg("查找分类失败")
		return err
	}

	if err := DB.Model(&video).Association("Categories").Replace(categories); err != nil {
		log.Error().Err(err).Uint("videoID", videoID).Msg("更新视频分类失败")
		return err
	}

	log.Info().Uint("videoID", videoID).Msg("更新视频分类成功")
	return nil
}
