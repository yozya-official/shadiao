package main

import (
	"os"
	"shadiao/conf"
	"shadiao/db"
	"shadiao/service"

	"github.com/rs/zerolog/log"

	"github.com/Yuelioi/gkit/logx/zero"
	"github.com/Yuelioi/gkit/web/gin/middleware/apikey"
	"github.com/Yuelioi/gkit/web/gin/middleware/log/gzero"
	"github.com/Yuelioi/gkit/web/gin/middleware/ratelimit"
	"github.com/Yuelioi/gkit/web/gin/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// 配置日志
	logger := zero.Default()
	log.Logger = logger

	// 初始化数据库
	if envPath := os.Getenv("DATABASE_URL"); envPath != "" {
		conf.DbPath = envPath
	}

	db.Init(db.Config{
		DSN:          conf.DbPath,
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		MaxLifetime:  10,
	})

	// 初始化 API Key
	_ = conf.InitAPIKey()

	// gin.DefaultWriter = io.Discard
	// gin.DefaultErrorWriter = io.Discard

	// 服务器配置
	cfg := server.ServerConfig{
		Addr:      ":9000",
		Logger:    logger,
		Mode:      os.Getenv("APP_MODE"),
		APIPrefix: "/api/v1",
		Middlewares: []gin.HandlerFunc{
			gzero.Default(logger),
			gzero.GinRecovery(logger),
			ratelimit.Default(),
		},
		EnableCORS: true,
		SPAPath:    "./frontend/dist",
	}

	// 启动服务器
	err := server.Start(cfg, func(api *gin.RouterGroup) {
		// 公共接口
		public := api.Group("/")
		{
			public.GET("/videos", service.GetAllVideos)
			public.POST("/videos", service.CreateVideo)
			public.GET("/videos/unreviewed", service.GetUnreviewedVideos)
			public.GET("/video/parse", service.ParseVideoURL)
			public.GET("/videos/:id", service.GetVideoByID)

			public.GET("/authors", service.GetAllAuthors)
			public.GET("/authors/:id", service.GetAuthorByID)

			public.GET("/tags", service.GetAllTags)
			public.GET("/tags/:id", service.GetTagByID)
			public.GET("/tags/:id/videos", service.GetTagVideos)
		}

		// 需要授权的接口
		protected := api.Group("/")
		protected.Use(apikey.Default(conf.IsValidAPIKey))
		{
			// 视频操作
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
		}
	})

	if err != nil {
		logger.Fatal().Err(err).Msg("服务启动失败")
	}
}
