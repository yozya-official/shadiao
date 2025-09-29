package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

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

type ParseVideoResponse struct {
	Code       int     `json:"code"`
	Msg        string  `json:"msg"`
	ApiVersion string  `json:"api_version"`
	Timestamp  float64 `json:"timestamp"`
	Created    int64   `json:"created"`
	Data       struct {
		ParseParams struct {
			Type      string `json:"type"`
			Platform  string `json:"platform"`
			ID        string `json:"id"`
			VideoType string `json:"video_type"`
			N         int    `json:"n"`
		} `json:"parse_params"`
		Author struct {
			Name   string `json:"name"`
			Mid    int64  `json:"mid"`
			Avatar string `json:"avatar"`
		} `json:"author"`
		Stat struct {
			Like    int    `json:"like"`
			Comment int    `json:"comment"`
			Collect int    `json:"collect"`
			Share   int    `json:"share"`
			AwemeID string `json:"aweme_id"`
			Time    int64  `json:"time"`
		} `json:"stat"`
		Item struct {
			Title    string  `json:"title"`
			Cover    string  `json:"cover"`
			Desc     string  `json:"desc"`
			Tname    string  `json:"tname"`
			URL      string  `json:"url"`
			Quality  string  `json:"quality"`
			FPS      int     `json:"fps"`
			Bitrate  string  `json:"bitrate"`
			Duration float64 `json:"duration"`
			Size     int64   `json:"size"`
			SizeStr  string  `json:"size_str"`
			Height   int     `json:"height"`
			Width    int     `json:"width"`
			Cid      int64   `json:"cid"`
		} `json:"item"`
	} `json:"data"`
}

func parseVideoURL(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 url 参数"})
		return
	}

	client := &http.Client{
		// 禁止自动跳转
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var realUrl = url

	// 转换短链
	if strings.Contains(url, "b23.tv") {
		realUrl = CleanBilibiliURL(resp.Header.Get("Location"))
		if realUrl == "" {
			c.JSON(http.StatusOK, gin.H{
				"url":    url,
				"status": resp.Status,
				"note":   "没有发现跳转",
			})
			return
		}
	} else if strings.Contains(url, "www.bilibili.com/video/BV") {
		// 转换bv号
		realUrl = ConvertBvUrlToAv(url)
	}

	// 调用第三方 API
	apiResp, err := client.Get("https://api.hk.sisi.us.kg/API/Parse/?url=" + realUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to request parse API", "detail": err.Error()})
		return
	}
	defer apiResp.Body.Close()

	body, err := io.ReadAll(apiResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read response", "detail": err.Error()})
		return
	}

	var data ParseVideoResponse
	if err := json.Unmarshal(body, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse JSON", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
