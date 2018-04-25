package services

import (
	"bee-go-myBlog/common"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"bee-go-myBlog/models"
	"time"
)

func IndexSystem() ( *models.Systems) {
	cache := common.Cache()
	key := "index:system:data"
	res := cache.Get(key)
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
		data ,_ := json.Marshal(system)
		cache.Put(key,data,timeoutDuration)
		return system
	}

	var err error
	var system  *models.Systems
	string1,_ := redis.String(res,err)
	json.Unmarshal([]byte(string1),&system)
	return system
}