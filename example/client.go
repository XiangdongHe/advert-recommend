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
	"sync"
	"time"

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
		//// 查询用户兴趣
		//getReq := &advert.GetUserInterestsRequest{
		//	UserId: userId,
		//}
		//getResp, err := c.GetUserInterests(context.Background(), getReq)
		//if err != nil {
		//	log.Printf("❌ 查询用户兴趣失败: %v", err)
		//	continue
		//}
		//// 打印用户兴趣信息
		//if len(getResp.Interests) == 0 {
		//	fmt.Println("⚠️ 暂无兴趣信息。")
		//	continue
		//}
		//
		//fmt.Println("\n📢 用户兴趣信息：")
		//for _, interst := range getResp.Interests {
		//	fmt.Printf("%s、", interst.Tag)
		//}

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
			fmt.Printf("- 创意ID: %d | 类型: %d | 标题: %s | 描述: %s | 媒体URL: %s | 分值：%.2f\n ",
				ad.CreativeId, ad.CreativeType, ad.Title, ad.Description, ad.MediaUrl, ad.Weight)
		}
	}
}

func pressure() {

	newClient, err := advertservice.NewClient(
		"advertservice",
		client.WithHostPorts("127.0.0.1:8888"),
	)
	if err != nil {
		log.Fatalf("❌ 创建客户端失败: %v", err)
	}

	// 统计变量
	wg := sync.WaitGroup{}
	start := time.Now()
	success := int64(0)
	failed := int64(0)
	totalCount := 10000
	c := 10

	fmt.Printf("开始压测 RPC 接口: PredictCTR, 地址: %s, 并发: %d, 请求总数: %d\n", "127.0.0.1:8888", c, totalCount)

	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < totalCount/c; j++ {
				getRcReq := &advert.GetAdvertRecommendRequest{
					UserId: 1001,
				}

				_, err := newClient.GetAdvertRecommend(context.Background(), getRcReq)
				if err != nil {
					failed++
					continue
				}
				success++
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("\n✅ 压测完成\n")
	fmt.Printf("总请求数: %d\n", totalCount)
	fmt.Printf("成功数: %d, 失败数: %d\n", success, failed)
	fmt.Printf("总耗时: %.2fs\n", duration.Seconds())
	fmt.Printf("平均QPS: %.2f\n", float64(totalCount)/duration.Seconds())
	fmt.Printf("平均单请求耗时: %.2f ms\n", (duration.Seconds()*1000)/float64(totalCount))
}

func main() {
	pressure()
	//createAdPlanExample()
}
