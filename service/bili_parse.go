package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// API 响应结构
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

// 响应结构
type ViewResp struct {
	Code int `json:"code"`
	Data struct {
		Bvid    string `json:"bvid"`
		Aid     int64  `json:"aid"`
		Title   string `json:"title"`
		Pic     string `json:"pic"`
		Pubdate int64  `json:"pubdate"`
		Cid     int64  `json:"cid"`
		Tname   string `json:"tname"`
		Desc    string `json:"desc"`
		Owner   struct {
			Mid  int64  `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			View    int64 `json:"view"`
			Danmaku int64 `json:"danmaku"`
			Reply   int64 `json:"reply"`
			Fav     int64 `json:"fav"`
			Coin    int64 `json:"coin"`
			Share   int64 `json:"share"`
			Like    int64 `json:"like"`
		} `json:"stat"`
		Pages []struct {
			Cid       int64  `json:"cid"`
			Page      int    `json:"page"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"dimension"`
		} `json:"pages"`
	} `json:"data"`
}

type PlayurlResp struct {
	Code int `json:"code"`
	Data struct {
		Durl []struct {
			Order  int    `json:"order"`
			Length int    `json:"length"`
			Size   int64  `json:"size"`
			Url    string `json:"url"`
		} `json:"durl"`
		AcceptQuality []int `json:"accept_quality"`
	} `json:"data"`
}

// 请求封装
func httpGet(url string) ([]byte, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Referer", "https://www.bilibili.com")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func getAVID(rawURL string) (string, error) {
	realURL := rawURL

	// 处理短链 b23.tv
	b23Re := regexp.MustCompile(`https?://b23\.tv/[^\s]+`)
	if match := b23Re.FindString(rawURL); match != "" {
		resolved, err := resolveShortURL(match)
		if err != nil {
			return "", fmt.Errorf("解析短链失败: %w", err)
		}
		realURL = resolved
	}

	// 尝试匹配 AV 号
	avRe := regexp.MustCompile(`av(\d+)`)
	if matches := avRe.FindStringSubmatch(realURL); len(matches) > 1 {
		return matches[1], nil
	}

	// 尝试匹配 BV 号并转换为 AV 号
	bvRe := regexp.MustCompile(`BV[0-9A-Za-z]+`)
	if matches := bvRe.FindStringSubmatch(realURL); len(matches) > 0 {
		avid := bvid2Avid(matches[0])
		return fmt.Sprintf("%d", avid), nil
	}

	return "", fmt.Errorf("无法从 URL 中提取视频 AV 号")
}

// 解析短链
func resolveShortURL(url string) (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	location := resp.Header.Get("Location")
	if location == "" {
		return url, nil
	}

	return location, nil
}

// 格式化文件大小
func formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

// 核心解析逻辑
func scrapeVideo(aid string) (*ParseVideoResponse, error) {
	resp := &ParseVideoResponse{
		Code:       0,
		Msg:        "success",
		ApiVersion: "1.0",
		Timestamp:  float64(time.Now().Unix()),
		Created:    time.Now().Unix(),
	}

	// 设置解析参数
	resp.Data.ParseParams.Type = "video"
	resp.Data.ParseParams.Platform = "BiliBili"
	resp.Data.ParseParams.ID = aid
	resp.Data.ParseParams.VideoType = "av"
	resp.Data.ParseParams.N = 1

	// 1. 获取视频详情
	viewURL := fmt.Sprintf("https://api.bilibili.com/x/web-interface/view?aid=%s", aid)
	viewBody, err := httpGet(viewURL)
	if err != nil {
		return nil, fmt.Errorf("获取视频信息失败: %v", err)
	}

	var viewData ViewResp
	if err := json.Unmarshal(viewBody, &viewData); err != nil {
		return nil, fmt.Errorf("解析视频信息失败: %v", err)
	}

	if viewData.Code != 0 {
		return nil, fmt.Errorf("bilibili API 返回错误: code=%d", viewData.Code)
	}

	// 填充作者信息
	resp.Data.Author.Name = viewData.Data.Owner.Name
	resp.Data.Author.Mid = viewData.Data.Owner.Mid
	resp.Data.Author.Avatar = viewData.Data.Owner.Face

	// 填充统计信息
	resp.Data.Stat.Like = int(viewData.Data.Stat.Like)
	resp.Data.Stat.Comment = int(viewData.Data.Stat.Reply)
	resp.Data.Stat.Collect = int(viewData.Data.Stat.Fav)
	resp.Data.Stat.Share = int(viewData.Data.Stat.Share)
	resp.Data.Stat.AwemeID = fmt.Sprintf("%s_%s", viewData.Data.Bvid, aid)
	resp.Data.Stat.Time = viewData.Data.Pubdate

	// 填充视频基本信息
	resp.Data.Item.Title = viewData.Data.Title
	resp.Data.Item.Cover = viewData.Data.Pic
	resp.Data.Item.Desc = viewData.Data.Desc
	resp.Data.Item.Tname = viewData.Data.Tname
	resp.Data.Item.Cid = viewData.Data.Cid

	// 从 pages 获取视频尺寸和时长
	if len(viewData.Data.Pages) > 0 {
		firstPage := viewData.Data.Pages[0]
		resp.Data.Item.Duration = float64(firstPage.Duration)
		resp.Data.Item.Width = firstPage.Dimension.Width
		resp.Data.Item.Height = firstPage.Dimension.Height
	}

	// 2. 获取播放地址
	cidStr := strconv.FormatInt(viewData.Data.Cid, 10)
	playURL := fmt.Sprintf("https://api.bilibili.com/x/player/playurl?avid=%s&cid=%s&qn=80&fnval=0", aid, cidStr)
	playBody, err := httpGet(playURL)

	if err == nil {
		var playData PlayurlResp
		if err := json.Unmarshal(playBody, &playData); err == nil {
			if playData.Code == 0 && len(playData.Data.Durl) > 0 {
				firstURL := playData.Data.Durl[0]
				resp.Data.Item.URL = firstURL.Url
				resp.Data.Item.Size = firstURL.Size
				resp.Data.Item.SizeStr = formatSize(firstURL.Size)
				resp.Data.Item.Duration = float64(firstURL.Length) / 1000.0 // 转为秒
				resp.Data.Item.Quality = "80"                               // 高清质量
			}
		}
	}

	// 如果无法获取播放地址，标记为未找到
	if resp.Data.Item.URL == "" {
		resp.Data.Item.URL = "NOT_FOUND"
		resp.Code = 1
		resp.Msg = "无法获取播放地址，可能需要登录或有地区限制"
	}

	return resp, nil
}

// Gin 路由处理函数
func ParseVideoURL(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 url 参数"})
		return
	}

	aid, err := getAVID(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法从 URL 中提取视频 ID"})
		return
	}

	// 执行视频解析
	result, err := scrapeVideo(aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "视频解析失败", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
