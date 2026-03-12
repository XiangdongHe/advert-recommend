package main

import (
	"fmt"
	"log"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/advert"
)

// 这是一个简单的客户端测试示例
// 实际使用时需要使用 Kitex 生成的客户端代码

func main() {
	log.Println("AdvertRecommend Client Test Example")

	// 注意：这只是示例代码，展示如何构造请求
	// 实际使用需要通过 Kitex 客户端调用

	// 示例1: 创建广告计划请求
	createPlanReq := &advert.CreateAdPlanRequest{
		Name:          "测试广告计划",
		Objective:     "click",
		Budget:        10000.00,
		BidPrice:      "CPC:0.5",
		TargetingRule: `{"region":["北京","上海"],"age":[18,35]}`,
		StartTime:     "2024-01-01 00:00:00",
		EndTime:       "2024-03-31 23:59:59",
	}
	fmt.Printf("创建广告计划请求: %+v\n", createPlanReq)

	// 示例2: 更新广告计划请求
	status := int32(0) // 暂停
	updatePlanReq := &advert.UpdateAdPlanRequest{
		PlanId: 1,
		Status: &status,
	}
	fmt.Printf("更新广告计划请求: %+v\n", updatePlanReq)

	// 示例3: 查询广告计划请求
	getPlanReq := &advert.GetAdPlanRequest{
		PlanId: 1,
	}
	fmt.Printf("查询广告计划请求: %+v\n", getPlanReq)

	// 示例4: 分页查询广告计划列表
	listPlansReq := &advert.ListAdPlansRequest{
		Page:     1,
		PageSize: 10,
	}
	fmt.Printf("分页查询请求: %+v\n", listPlansReq)

	// 示例5: 创建广告创意请求
	createCreativeReq := &advert.CreateAdCreativeRequest{
		PlanId:       1,
		CreativeType: 1, // image
		MediaUrl:     "https://example.com/ad1.jpg",
		Title:        "限时优惠",
		Description:  "全场5折",
	}
	fmt.Printf("创建广告创意请求: %+v\n", createCreativeReq)

	// 示例6: 创建用户画像请求
	createProfileReq := &advert.CreateUserProfileRequest{
		UserId:     1001,
		Gender:     1, // male
		Age:        28,
		Region:     "北京",
		DeviceType: "iPhone 14",
	}
	fmt.Printf("创建用户画像请求: %+v\n", createProfileReq)

	// 示例7: 添加用户兴趣请求
	addInterestReq := &advert.AddUserInterestRequest{
		UserId: 1001,
		Tag:    "科技",
		Weight: 0.85,
	}
	fmt.Printf("添加用户兴趣请求: %+v\n", addInterestReq)

	// 示例8: 创建广告事件请求
	createEventReq := &advert.CreateAdEventRequest{
		UserId:     1001,
		CreativeId: 1,
		EventType:  2, // click
		Ts:         "2024-01-15 10:30:05",
		Extra:      `{"source":"feed","position":1}`,
	}
	fmt.Printf("创建广告事件请求: %+v\n", createEventReq)

	fmt.Println("\n===========================================")
	fmt.Println("提示: 这只是示例代码，展示请求结构")
	fmt.Println("实际使用时，需要:")
	fmt.Println("1. 使用 Kitex 工具生成客户端代码")
	fmt.Println("2. 创建客户端连接到服务器")
	fmt.Println("3. 调用相应的 RPC 方法")
	fmt.Println("===========================================")
}

// 以下是实际使用 Kitex 客户端的示例代码框架
// 需要先使用 kitex 工具生成客户端代码
