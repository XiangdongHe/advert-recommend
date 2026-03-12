package main

import (
	"fmt"
	"log"
	"net"

	"gitee.com/HeXiangdong/AdvertRecommend/user-service/config"
	"gitee.com/HeXiangdong/AdvertRecommend/user-service/database"
	"gitee.com/HeXiangdong/AdvertRecommend/user-service/handler"
	user "gitee.com/HeXiangdong/AdvertRecommend/user-service/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
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

	// 创建 Kitex 服务器地址
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		log.Fatalf("Failed to resolve address: %v", err)
	}
	log.Printf("Server listening on %s", addr.String())
	// 注册etcd
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	// 启动 Kitex 服务
	impl := handler.NewUserServiceImpl()
	svr := user.NewServer(
		impl,
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user-service"}),
		server.WithRegistry(r),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
