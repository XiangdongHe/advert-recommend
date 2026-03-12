package rpc

import (
	"log"

	"gitee.com/HeXiangdong/AdvertRecommend/recommend-service/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var UserClient userservice.Client

func InitUserRPC() {

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	cli, err := userservice.NewClient("user-service", client.WithResolver(r))

	if err != nil {
		panic(err)
	}

	UserClient = cli
}
