package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	initDB()

	// 获取底层 *sql.DB 并关闭
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB:", err)
	}
	defer sqlDB.Close() // 程序退出时自动关闭连接池

	// 生成随机密钥
	key := InitAPIKey()
	fmt.Println("生成的密钥为: ", key)

	r := gin.Default()

	// CORS 中间件 - 必须放在最前面
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API 路由组 - 必须在静态文件之前定义
	api := r.Group("/api/v1")
	{
		// 视频路由
		api.GET("/videos/unreviewed", getUnreviewedVideos)
		api.POST("/videos", createVideo)
		api.GET("/videos", getAllVideos)
		api.GET("/video/parse", parseVideoURL)
		api.GET("/videos/:id", getVideoByID)
		api.PUT("/videos/:id", updateVideo)

		// 作者路由
		api.POST("/authors", createAuthor)
		api.GET("/authors", getAllAuthors)
		api.GET("/authors/:id", getAuthorByID)
		api.PUT("/authors/:id", updateAuthor)
		api.PUT("/authors/:id/master", setMaster)
	}

	// 需要授权的路由
	authorized := api.Group("/")
	authorized.Use(APIKeyAuthMiddleware())
	{
		authorized.PATCH("/videos/:id/review", updateVideoReviewStatus)
		authorized.DELETE("/videos/:id", deleteVideo)
		authorized.DELETE("/authors/:id", deleteAuthor)

		// 验证
		authorized.GET("/authorization", func(c *gin.Context) {
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

	r.Run("0.0.0.0:9000")
}
