package service

import (
	"fmt"
	"net/http"
	"shadiao/db"
	"shadiao/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

// 视频数据(仅用于场景新视频)
type VideoData struct {
	ID          uint      `json:"id,omitempty"`
	Title       string    `json:"title"`       // 视频标题
	AID         int       `json:"aid"`         // B站 AV 号
	URL         string    `json:"url"`         // 视频地址
	Cover       string    `json:"cover"`       // 封面
	Duration    int       `json:"duration"`    // 视频长度
	Description string    `json:"description"` // 描述
	Views       int       `json:"views"`       // 播放量
	IsOriginal  bool      `json:"isOriginal"`  // 是否原创
	IsCompleted bool      `json:"isCompleted"` // 是否完结
	World       string    `json:"world"`       // 世界观
	Background  string    `json:"background"`  // 背景
	Style       []string  `json:"style"`       // 风格列表
	HasSystem   bool      `json:"hasSystem"`   // 是否系统内置
	Ctime       time.Time `json:"ctime,omitempty"`
}

// 新视频数据结构，包含作者
type VideoNewData struct {
	Video  VideoData     `json:"video"`
	Author models.Author `json:"author"`
}

// CreateVideo 创建视频（包含作者和标签）
// CreateVideo 创建视频（包含作者和标签）
func CreateVideo(c *gin.Context) {
	var input VideoNewData

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg("JSON 绑定失败")
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 验证必填字段
	if input.Video.AID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "视频 AID 不能为空"})
		return
	}
	if input.Author.UID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "作者 UID 不能为空"})
		return
	}

	// 检查视频是否已存在
	existingVideo, _ := db.GetVideoByAID(input.Video.AID)
	if existingVideo != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该视频已被推荐过了哦~"})
		return
	}

	// 查找或创建作者
	author, err := db.GetAuthorByUID(input.Author.UID)
	if err != nil {
		if err := db.CreateAuthor(&input.Author); err != nil {
			log.Error().Err(err).Int("uid", input.Author.UID).Msg("创建作者失败")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "创建作者失败"})
			return
		}
		author = &input.Author
	}

	var v models.Video
	v.AID = input.Video.AID
	v.AuthorID = author.ID
	v.Title = input.Video.Title
	v.URL = input.Video.URL
	v.Cover = input.Video.Cover
	v.Description = input.Video.Description
	v.Duration = input.Video.Duration
	v.Views = input.Video.Views
	v.Reviewed = false
	v.Ctime = input.Video.Ctime

	// ----------------------
	// 直接查询已有标签并赋值
	// ----------------------
	tagNames := map[string]string{
		"isOriginal":  fmt.Sprintf("%v", input.Video.IsOriginal),
		"isCompleted": fmt.Sprintf("%v", input.Video.IsCompleted),
		"hasSystem":   fmt.Sprintf("%v", input.Video.HasSystem),
		"background":  input.Video.Background,
		"world":       input.Video.World,
	}

	var tags []models.Tag
	for tType, tName := range tagNames {
		if tName == "" {
			continue
		}
		tag, err := db.GetTagByNameAndType(tName, tType)
		if err == nil && tag != nil {
			tags = append(tags, *tag)
		}
	}

	// style 数组
	for _, styleName := range input.Video.Style {
		tag, err := db.GetTagByNameAndType(styleName, "style")
		if err == nil && tag != nil {
			tags = append(tags, *tag)
		}
	}

	v.Tags = tags

	// 创建视频
	if err := db.CreateVideo(&v); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "unique constraint") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "该视频已被推荐过了哦~"})
			return
		}
		log.Error().Err(err).Int("aid", input.Video.AID).Msg("创建视频失败")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建视频失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "视频创建成功",
		"video":   v,
		"author":  author,
	})
}

// GetAllVideos 获取所有视频
func GetAllVideos(c *gin.Context) {
	// 支持分页参数

	videos, err := db.GetAllVideos()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all videos")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取视频列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"videos": videos,
		"total":  len(videos),
	})
}

// GetUnreviewedVideos 获取所有未审核视频
func GetUnreviewedVideos(c *gin.Context) {

	videos, err := db.GetUnreviewedVideos()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get unreviewed videos")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取未审核视频失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"videos": videos,
		"total":  len(videos),
	})
}

// UpdateVideoReviewStatus 修改视频审核状态
func UpdateVideoReviewStatus(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的视频 ID"})
		return
	}

	var req struct {
		Reviewed bool `json:"reviewed"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON for review status")
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	// 获取视频
	video, err := db.GetVideoByIDAny(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Video not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "视频不存在"})
		return
	}

	// 更新审核状态
	video.Reviewed = req.Reviewed
	if err := db.UpdateVideo(video); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to update review status")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新审核状态失败"})
		return
	}

	log.Info().Uint("id", id).Bool("reviewed", req.Reviewed).Msg("Video review status updated")
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"video":   video,
	})
}

// GetVideoByID 获取单个视频
func GetVideoByID(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的视频 ID"})
		return
	}

	video, err := db.GetVideoByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Video not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "视频不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"video": video})
}

// GetVideoByAID 根据 AID 获取视频
func GetVideoByAID(c *gin.Context) {
	aid := cast.ToInt(c.Param("aid"))

	if aid == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的视频 AID"})
		return
	}

	video, err := db.GetVideoByAID(aid)
	if err != nil {
		log.Error().Err(err).Int("aid", aid).Msg("Video not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "视频不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"video": video})
}

// UpdateVideo 更新视频（包含标签转换）
func UpdateVideo(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的视频 ID"})
		return
	}

	var input VideoData
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg("JSON 绑定失败")
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 获取现有视频
	video, err := db.GetVideoByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("视频不存在")
		c.JSON(http.StatusNotFound, gin.H{"error": "视频不存在"})
		return
	}

	// 更新普通字段（非零值或非空字符串）
	if input.Title != "" {
		video.Title = input.Title
	}
	if input.AID != 0 {
		video.AID = input.AID
	}
	if input.URL != "" {
		video.URL = input.URL
	}
	if input.Cover != "" {
		video.Cover = input.Cover
	}
	if input.Description != "" {
		video.Description = input.Description
	}
	if input.Duration != 0 {
		video.Duration = input.Duration
	}
	if input.Views != 0 {
		video.Views = input.Views
	}

	// ====================== 标签处理 ======================
	tags := video.Tags // 保留已有标签

	// 布尔类型标签
	boolFields := map[string]bool{
		"isOriginal":  input.IsOriginal,
		"isCompleted": input.IsCompleted,
		"hasSystem":   input.HasSystem,
	}
	for typ, val := range boolFields {
		tag, _ := db.GetTagByNameAndType(fmt.Sprintf("%v", val), typ)
		if tag != nil && !containsTag(tags, tag.ID) {
			tags = append(tags, *tag)
		}
	}

	// 单值字段标签
	singleFields := map[string]string{
		"background": input.Background,
		"world":      input.World,
	}
	for typ, val := range singleFields {
		if val == "" {
			continue
		}
		tag, _ := db.GetTagByNameAndType(val, typ)
		if tag != nil && !containsTag(tags, tag.ID) {
			tags = append(tags, *tag)
		}
	}

	// 多值字段标签
	for _, style := range input.Style {
		tag, _ := db.GetTagByNameAndType(style, "style")
		if tag != nil && !containsTag(tags, tag.ID) {
			tags = append(tags, *tag)
		}
	}

	video.Tags = tags

	// 保存更新
	if err := db.UpdateVideo(video); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("更新视频失败")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新视频失败"})
		return
	}

	// 重新获取完整视频信息
	video, _ = db.GetVideoByID(id)

	log.Info().Uint("id", id).Msg("视频更新成功")
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"video":   video,
	})
}

// containsTag 辅助函数：检查 tag 是否已经在切片中
func containsTag(tags []models.Tag, id uint) bool {
	for _, t := range tags {
		if t.ID == id {
			return true
		}
	}
	return false
}

// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的视频 ID"})
		return
	}

	// 检查视频是否存在
	_, err := db.GetVideoByIDAny(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Video not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "视频不存在"})
		return
	}

	// 删除视频
	if err := db.DeleteVideo(id); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("Failed to delete video")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除视频失败"})
		return
	}

	log.Info().Uint("id", id).Msg("Video deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
