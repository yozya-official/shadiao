package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"shadiao/conf"
	"shadiao/db"
	"shadiao/middleware"
	"shadiao/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	// 配置日志
	logger := initLog()
	log.Logger = logger

	if envPath := os.Getenv("DATABASE_URL"); envPath != "" {
		conf.DbPath = envPath
	}

	db.Init(db.Config{
		DSN:          conf.DbPath,
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		MaxLifetime:  10,
	})

	// 生成随机密钥
	key := conf.InitAPIKey()
	fmt.Println("生成的密钥为: ", key)

	// 初始化Gin
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard
	r := gin.New()
	r.Use(middleware.GinLogger(logger), middleware.GinRecovery(logger))

	mode := strings.ToLower(os.Getenv("APP_MODE"))

	// 基于部署方式动态修改参数
	addr := ":9000"
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		// 全局使用
		addr = "0.0.0.0" + addr
		// CORS 中间件
		r.Use(middleware.Cors())
	}

	// API 路由组 - 必须在静态文件之前定义
	api := r.Group("/api/v1")

	// 公共查询接口
	public := api.Group("/")
	{
		// 视频查询
		public.GET("/videos", service.GetAllVideos)
		public.GET("/videos/unreviewed", service.GetUnreviewedVideos)
		public.GET("/video/parse", service.ParseVideoURL)
		public.GET("/videos/:id", service.GetVideoByID)

		// 作者查询
		public.GET("/authors", service.GetAllAuthors)
		public.GET("/authors/:id", service.GetAuthorByID)

		// 标签查询
		public.GET("/tags", service.GetAllTags)
		public.GET("/tags/:id", service.GetTagByID)
		public.GET("/tags/:id/videos", service.GetTagVideos)
	}

	// 需要授权的接口
	protected := api.Group("/")
	protected.Use(middleware.APIKeyAuthMiddleware())
	{
		// 视频操作
		protected.POST("/videos", service.CreateVideo)
		protected.PUT("/videos/:id", service.UpdateVideo)
		protected.DELETE("/videos/:id", service.DeleteVideo)
		protected.PATCH("/videos/:id/review", service.UpdateVideoReviewStatus)

		// 作者操作
		protected.POST("/authors", service.CreateAuthor)
		protected.PUT("/authors/:id", service.UpdateAuthor)
		protected.DELETE("/authors/:id", service.DeleteAuthor)

		// 标签操作
		protected.POST("/tags", service.CreateTag)
		protected.PUT("/tags/:id", service.UpdateTag)
		protected.DELETE("/tags/:id", service.DeleteTag)

		// 验证接口
		protected.GET("/authorization", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "验证通过"})
		})
	}

	// 静态资源路由
	r.Static("/assets", "./frontend/dist/assets")
	r.StaticFile("/favicon.svg", "./frontend/dist/favicon.svg")
	r.StaticFile("/logo.png", "./frontend/dist/logo.png")

	// 处理 Vue SPA 的所有其他路由 - 返回 index.html
	r.NoRoute(func(c *gin.Context) {
		// 排除 API 路由，避免返回 HTML 给 API 请求
		path := c.Request.URL.Path
		if len(path) >= 4 && path[:4] == "/api" {
			c.JSON(http.StatusNotFound, gin.H{"error": "API route not found"})
			return
		}

		// 其他所有路由返回 index.html，让 Vue Router 处理
		c.File("./frontend/dist/index.html")
	})

	log.Info().Msg("运行在: " + addr)

	r.Run(addr)
}
