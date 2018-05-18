package services

import (
	"beego-postgres-blog/common"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"beego-postgres-blog/models"
	"time"
	"fmt"
)

func IndexSystem() ( *models.Systems) {
	cache := common.Cache()
	key := "index:system:data"
	res := cache.Get(key)
	fmt.Println("打个断言")
	if res == nil {
		system,_ := models.GetSystemsById(1)
		if system == nil {
			var systemCreate = models.Systems{
				Title	:	"叶落山城秋",
			}
			id,_ := models.AddSystems(&systemCreate)
			system,_ = models.GetSystemsById(id)
		}
		timeoutDuration := 24 * 30 * time.Hour
		fmt.Println("打个断言2")
		data ,_ := json.Marshal(system)
		cache.Put(key,data,timeoutDuration)
		return system
	}

	var err error
	var system  *models.Systems
	string1,_ := redis.String(res,err)
	json.Unmarshal([]byte(string1),&system)
	fmt.Println("打个断言3")
	return system
}