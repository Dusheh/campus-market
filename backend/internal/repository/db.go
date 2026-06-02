package repository

import (
	"fmt"
	"log"

	"github.com/Dusheh/campus-market/internal/config"
	"github.com/Dusheh/campus-market/internal/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB 初始化数据库连接
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	// 自动迁移
	if err := db.AutoMigrate(
		&model.User{},
		&model.Service{},
		&model.Goods{},
		&model.Demand{},
		&model.Order{},
		&model.Message{},
	); err != nil {
		return nil, fmt.Errorf("数据库迁移失败: %w", err)
	}

	log.Println("数据库连接成功，表结构已同步")
	return db, nil
}

// InitRedis 初始化 Redis 连接
func InitRedis(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	log.Println("Redis 连接成功")
	return rdb, nil
}