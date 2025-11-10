package main

import (
	"AdvertRecommend/config"
	"AdvertRecommend/database"
	"AdvertRecommend/handler"
	"AdvertRecommend/kitex_gen/advert/advertservice"
	"AdvertRecommend/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func SyncAdDataToRedis() error {
	var adPlans []*models.AdPlan
	query := database.DB.Model(&models.AdPlan{})
	if err := query.Find(&adPlans).Error; err != nil {
		log.Fatalf("Failed to query plans: %v", err)
		return err
	}
	// 查出所有广告创意
	var planIDs []int64
	for _, p := range adPlans {
		planIDs = append(planIDs, p.PlanID)
	}
	var adCreatives []*models.AdCreative
	if err := database.DB.Find(&adCreatives).Error; err != nil {
		log.Fatalf("Failed to query adCreatives: %v", err)
		return err
	}

	// 将广告创意分配给对应的计划
	creativeMap := make(map[int64][]*models.AdCreative)
	for _, c := range adCreatives {
		creativeMap[c.PlanID] = append(creativeMap[c.PlanID], c)
	}
	for _, plan := range adPlans {
		plan.Creatives = creativeMap[plan.PlanID]
	}
	// 插入redis
	ctx := context.Background()
	for _, p := range adPlans {
		data, _ := json.Marshal(p)
		database.RDB.Set(ctx, fmt.Sprintf("ad:plan:%d", p.PlanID), data, 0)
		// 建立 tag → plan 反向索引
		var targeting map[string]string
		if err := json.Unmarshal([]byte(p.TargetingRule), &targeting); err != nil {
			continue
		}
		interest, _ := targeting["interest"]
		database.RDB.SAdd(ctx, fmt.Sprintf("interest:%s", interest), p.PlanID)
		for _, adCreative := range p.Creatives {
			adCreativeStr, _ := json.Marshal(p)
			database.RDB.Set(ctx, fmt.Sprintf("ad:creative:%d", adCreative.CreativeID), adCreativeStr, 0)
			database.RDB.SAdd(ctx, fmt.Sprintf("ad:plan:%d:creatives", p.PlanID), adCreative.CreativeID)
		}
	}

	return nil
}

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
	//if err := SyncAdDataToRedis(); err != nil {
	//	log.Fatalf("Failed to sync ad data to redis: %v", err)
	//}
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
