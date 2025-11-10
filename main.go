package main

import (
	"AdvertRecommend/config"
	"AdvertRecommend/database"
	"AdvertRecommend/handler"
	"AdvertRecommend/kitex_gen/advert/advertservice"
	"fmt"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	// 加载配置
	cfg := config.GetDefaultConfig()
	config.Global = cfg
	log.Printf("Starting AdvertRecommend Service...")

	// 初始化数据库
	dbConfig := database.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		Charset:  cfg.Database.Charset,
	}

	if err := database.InitDB(dbConfig); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := database.InitRedis(); err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}

	// 创建服务处理器
	impl := handler.NewAdvertServiceImpl()

	// 创建 Kitex 服务器地址
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		log.Fatalf("Failed to resolve address: %v", err)
	}

	log.Printf("Server listening on %s", addr.String())

	// 启动 Kitex 服务
	svr := advertservice.NewServer(
		impl,
		server.WithServiceAddr(addr),
	)

	err = svr.Run()
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
