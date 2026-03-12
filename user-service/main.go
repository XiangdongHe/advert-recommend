package main

import (
	"log"

	"gitee.com/HeXiangdong/AdvertRecommend/user-service/handler"
	user "gitee.com/HeXiangdong/AdvertRecommend/user-service/kitex_gen/user/userservice"
)

func main() {
	impl := handler.NewUserServiceImpl()
	svr := user.NewServer(impl)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
