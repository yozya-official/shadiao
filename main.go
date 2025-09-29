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

	// API 路由组
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

	r.Run("0.0.0.0:8000")

}
