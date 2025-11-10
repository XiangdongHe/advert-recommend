package database

import (
	"AdvertRecommend/models"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var RDB *redis.Client

// Config 数据库配置
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Charset  string
}

// InitDB 初始化数据库连接
func InitDB(config Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.Charset,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connected successfully")
	return nil
}

func InitRedis() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "h545466093",
		DB:           0,
		PoolSize:     50,
		MinIdleConns: 10,
	})
	if err := RDB.Ping(context.Background()).Err(); err != nil {
		return fmt.Errorf("failed to connect redis: %v", err)
	}
	return nil
}

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate() error {
	err := DB.AutoMigrate(
		&models.AdPlan{},
		&models.AdCreative{},
		&models.UserProfileBase{},
		&models.UserProfileInterest{},
		&models.UserAdEventLog{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")
	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
