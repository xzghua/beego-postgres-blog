package common

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"fmt"
)

func Cache() cache.Cache {
	host := beego.AppConfig.String("redishost")
	port := beego.AppConfig.String("redisport")
	//pwd := beego.AppConfig.String("redispassword")

	cache2, err := cache.NewCache("redis", `{"conn": "`+host+`:`+port+`"}`)
	if err != nil {
		fmt.Println(err,"redis有问题")
	}
	return cache2
}