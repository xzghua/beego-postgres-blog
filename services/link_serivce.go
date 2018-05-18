package services

import (
	"beego-postgres-blog/common"
	"beego-postgres-blog/models"
	"time"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

func IndexLinkList() ([]interface{}) {
	cache := common.Cache()
	key := "index:link:list"
	res := cache.Get(key)
	if res == nil {
		links,_ := models.IndexAllLink()
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(links)
		cache.Put(key,data,timeoutDuration)
		return links
	}

	var err error
	var links  []interface{}
	string1,_ := redis.String(res,err)
	json.Unmarshal([]byte(string1),&links)
	return links
}