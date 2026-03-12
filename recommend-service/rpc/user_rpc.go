package rpc

import (
	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var UserClient userservice.Client

func InitUserRPC() {
	cli, err := userservice.NewClient(
		"user-service",
		client.WithHostPorts("127.0.0.1:8889"),
	)
	if err != nil {
		panic(err)
	}
	UserClient = cli
}
