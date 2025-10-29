package main

import (
	"AdvertRecommend/kitex_gen/advert"
	"AdvertRecommend/kitex_gen/advert/advertservice"
	"context"
	"log"

	"github.com/cloudwego/kitex/client"
)

func createAdPlanExample() {
	// 创建客户端
	c, err := advertservice.NewClient(
		"advertservice",
		client.WithHostPorts("127.0.0.1:8888"),
	)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	// 创建广告计划
	createReq := &advert.CreateAdPlanRequest{
		Name:          "测试广告计划",
		Objective:     "click",
		Budget:        10000.00,
		BidPrice:      "CPC:0.5",
		TargetingRule: `{"region":["北京","上海"],"age":[18,35]}`,
		StartTime:     "2024-03-01 00:00:00",
		EndTime:       "2024-05-31 23:59:59",
	}

	createResp, err := c.CreateAdPlan(context.Background(), createReq)
	if err != nil {
		log.Fatalf("RPC 调用失败: %v", err)
	}

	log.Printf("✅ 创建成功! Code=%d, PlanID=%d", createResp.BaseResp.Code, createResp.PlanId)

	// 查询广告计划
	getReq := &advert.GetAdPlanRequest{
		PlanId: createResp.PlanId,
	}

	getResp, err := c.GetAdPlan(context.Background(), getReq)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	log.Printf("✅ 查询成功! 计划名称=%s, 预算=%.2f", getResp.AdPlan.Name, getResp.AdPlan.Budget)
}

func main() {
	createAdPlanExample()
}
