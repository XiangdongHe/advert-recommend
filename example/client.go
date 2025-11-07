package main

import (
	"AdvertRecommend/kitex_gen/advert"
	"AdvertRecommend/kitex_gen/advert/advertservice"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cloudwego/kitex/client"
)

func createAdPlanExample() {
	// 创建客户端
	c, err := advertservice.NewClient(
		"advertservice",
		client.WithHostPorts("127.0.0.1:8888"),
	)
	if err != nil {
		log.Fatalf("❌ 创建客户端失败: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n请输入用户ID（输入 q 退出）：")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // 去掉换行符

		if input == "q" {
			fmt.Println("退出程序。")
			break
		}

		userId, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("⚠️ 输入无效，请输入数字类型的用户ID。")
			continue
		}
		// 查询用户兴趣
		getReq := &advert.GetUserInterestsRequest{
			UserId: userId,
		}
		getResp, err := c.GetUserInterests(context.Background(), getReq)
		if err != nil {
			log.Printf("❌ 查询用户兴趣失败: %v", err)
			continue
		}
		// 打印用户兴趣信息
		if len(getResp.Interests) == 0 {
			fmt.Println("⚠️ 暂无兴趣信息。")
			continue
		}

		fmt.Println("\n📢 用户兴趣信息：")
		for _, interst := range getResp.Interests {
			fmt.Printf("%s、", interst.Tag)
		}

		// 查询广告推荐
		getRcReq := &advert.GetAdvertRecommendRequest{
			UserId: userId,
		}

		getRcResp, err := c.GetAdvertRecommend(context.Background(), getRcReq)
		if err != nil {
			log.Printf("❌ 查询失败: %v", err)
			continue
		}

		// 打印广告创意信息
		if len(getRcResp.Adverts) == 0 {
			fmt.Println("⚠️ 暂无匹配的广告创意。")
			continue
		}

		fmt.Println("\n📢 推荐广告创意列表：")
		for _, ad := range getRcResp.Adverts {
			fmt.Printf("- 创意ID: %d | 类型: %d | 标题: %s | 描述: %s | 媒体URL: %s\n",
				ad.CreativeId, ad.CreativeType, ad.Title, ad.Description, ad.MediaUrl)
		}
	}
}

func main() {
	createAdPlanExample()
}
