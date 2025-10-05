package db

import (
	"fmt"
	"shadiao/models"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Config 数据库配置
type Config struct {
	DSN          string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
}

// Init 初始化数据库连接
func Init(cfg Config) error {
	var err error

	// 配置 GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 使用 zerolog，关闭 GORM 默认日志
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}

	// 连接数据库（这里以 SQLite 为例，可根据需要改为 MySQL/PostgreSQL）
	DB, err = gorm.Open(sqlite.Open(cfg.DSN), gormConfig)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		return fmt.Errorf("failed to connect database: %w", err)
	}

	// 获取底层 sql.DB
	sqlDB, err := DB.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get database instance")
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifetime)

	log.Info().Msg("Database connected successfully")

	// 自动迁移
	if err := AutoMigrate(); err != nil {
		return err
	}

	return nil
}

// AutoMigrate 自动迁移数据库表
func AutoMigrate() error {
	log.Info().Msg("Starting database migration")

	err := DB.AutoMigrate(
		&models.Tag{},
		&models.Category{},
		&models.Author{},
		&models.Video{},
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to migrate database")
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Info().Msg("Database migration completed")
	return nil
}

// Close 关闭数据库连接
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get database instance")
		return err
	}

	if err := sqlDB.Close(); err != nil {
		log.Error().Err(err).Msg("Failed to close database")
		return err
	}

	log.Info().Msg("Database connection closed")
	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
