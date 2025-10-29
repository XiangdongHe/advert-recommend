package main

import (
	"AdvertRecommend/config"
	"AdvertRecommend/database"
	"AdvertRecommend/handler"
	"AdvertRecommend/kitex_gen/advert"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	// 加载配置
	cfg := config.GetDefaultConfig()
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

	// 自动迁移数据库表结构（可选，生产环境建议手动管理）
	// if err := database.AutoMigrate(); err != nil {
	//     log.Fatalf("Failed to migrate database: %v", err)
	// }

	// 创建服务处理器
	impl := handler.NewAdvertServiceImpl()

	// 创建 Kitex 服务器
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	log.Printf("Server listening on %s", addr)

	// 启动服务
	svr := server.NewServer(
		server.WithServiceAddr(listener.Addr()),
	)

	err = svr.RegisterService(getServiceInfo(), impl)
	if err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	err = svr.Run()
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// getServiceInfo 获取服务信息
func getServiceInfo() *ServiceInfo {
	return &ServiceInfo{
		ServiceName: "AdvertService",
	}
}

// ServiceInfo 服务信息结构
type ServiceInfo struct {
	ServiceName string
}

func (s *ServiceInfo) GetServiceName() string {
	return s.ServiceName
}

func (s *ServiceInfo) GetMethods() map[string]interface{} {
	return map[string]interface{}{
		"CreateAdPlan":        createAdPlanHandler,
		"UpdateAdPlan":        updateAdPlanHandler,
		"GetAdPlan":           getAdPlanHandler,
		"ListAdPlans":         listAdPlansHandler,
		"DeleteAdPlan":        deleteAdPlanHandler,
		"CreateAdCreative":    createAdCreativeHandler,
		"UpdateAdCreative":    updateAdCreativeHandler,
		"GetAdCreative":       getAdCreativeHandler,
		"ListAdCreatives":     listAdCreativesHandler,
		"DeleteAdCreative":    deleteAdCreativeHandler,
		"CreateUserProfile":   createUserProfileHandler,
		"UpdateUserProfile":   updateUserProfileHandler,
		"GetUserProfile":      getUserProfileHandler,
		"DeleteUserProfile":   deleteUserProfileHandler,
		"AddUserInterest":     addUserInterestHandler,
		"UpdateUserInterest":  updateUserInterestHandler,
		"GetUserInterests":    getUserInterestsHandler,
		"DeleteUserInterest":  deleteUserInterestHandler,
		"CreateAdEvent":       createAdEventHandler,
		"GetUserAdEvents":     getUserAdEventsHandler,
		"GetCreativeAdEvents": getCreativeAdEventsHandler,
	}
}

// 简化的处理器函数
func createAdPlanHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).CreateAdPlan(ctx, req.(*advert.CreateAdPlanRequest))
}

func updateAdPlanHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).UpdateAdPlan(ctx, req.(*advert.UpdateAdPlanRequest))
}

func getAdPlanHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).GetAdPlan(ctx, req.(*advert.GetAdPlanRequest))
}

func listAdPlansHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).ListAdPlans(ctx, req.(*advert.ListAdPlansRequest))
}

func deleteAdPlanHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).DeleteAdPlan(ctx, req.(*advert.DeleteAdPlanRequest))
}

func createAdCreativeHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).CreateAdCreative(ctx, req.(*advert.CreateAdCreativeRequest))
}

func updateAdCreativeHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).UpdateAdCreative(ctx, req.(*advert.UpdateAdCreativeRequest))
}

func getAdCreativeHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).GetAdCreative(ctx, req.(*advert.GetAdCreativeRequest))
}

func listAdCreativesHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).ListAdCreatives(ctx, req.(*advert.ListAdCreativesRequest))
}

func deleteAdCreativeHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).DeleteAdCreative(ctx, req.(*advert.DeleteAdCreativeRequest))
}

func createUserProfileHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).CreateUserProfile(ctx, req.(*advert.CreateUserProfileRequest))
}

func updateUserProfileHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).UpdateUserProfile(ctx, req.(*advert.UpdateUserProfileRequest))
}

func getUserProfileHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).GetUserProfile(ctx, req.(*advert.GetUserProfileRequest))
}

func deleteUserProfileHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).DeleteUserProfile(ctx, req.(*advert.DeleteUserProfileRequest))
}

func addUserInterestHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).AddUserInterest(ctx, req.(*advert.AddUserInterestRequest))
}

func updateUserInterestHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).UpdateUserInterest(ctx, req.(*advert.UpdateUserInterestRequest))
}

func getUserInterestsHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).GetUserInterests(ctx, req.(*advert.GetUserInterestsRequest))
}

func deleteUserInterestHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).DeleteUserInterest(ctx, req.(*advert.DeleteUserInterestRequest))
}

func createAdEventHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).CreateAdEvent(ctx, req.(*advert.CreateAdEventRequest))
}

func getUserAdEventsHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).GetUserAdEvents(ctx, req.(*advert.GetUserAdEventsRequest))
}

func getCreativeAdEventsHandler(ctx context.Context, handler interface{}, req interface{}) (interface{}, error) {
	return handler.(advert.AdvertService).GetCreativeAdEvents(ctx, req.(*advert.GetCreativeAdEventsRequest))
}
