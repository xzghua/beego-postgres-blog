package common

import (
	"time"
	"math/rand"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

func GoMerge(arr1 []interface{},arr2 []interface{}) ([]interface{}) {
	for _,val := range arr2 {
		arr1 = append(arr1,val)
	}
	return arr1
}


func GoRepeat(str string, num int64) (string) {
	var i int64
	newStr := ""
	if num != 0 {
		for i = 0; i < num; i++ {
			newStr += str
		}
	}
	return newStr
}


//生成随机字符串
func GetRandomString(l int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


func MyPaginate(page int64,tableName string,env string) (totalPage int64,lastPage int64,currentPage int64,nextPage int64 ){
	o := orm.NewOrm()
	var cnt int64
	if env == "console" {
		cnt,_ = o.QueryTable(tableName).Count()
	} else {
		cache := Cache()
		key := "index:paginate:data"
		res := cache.Get(key)
		if res == nil {
			cnt,_ = o.QueryTable(tableName).Filter("deleted_at__isnull",true).Count()
			timeoutDuration := 24 * 30 * time.Hour
			data ,_ := json.Marshal(cnt)
			cache.Put(key,data,timeoutDuration)
		} else {
			var err error
			string1,err := redis.String(res,err)
			json.Unmarshal([]byte(string1),&cnt)
		}
	}
	limit, _ := beego.AppConfig.Int64("page_offset")
	res := round(cnt,limit)
	totalPage = res

	if page - 1 <= 0 {
		lastPage = 1
	} else {
		lastPage = page - 1
	}

	if page >= res {
		currentPage = res
	} else {
		currentPage = page
	}

	if page + 1 >= res {
		nextPage = res
	} else {
		nextPage = page + 1
	}

	return totalPage,lastPage,currentPage,nextPage
}

func round(a int64,b int64) int64 {
	rem := a % b
	dis := a / b
	if rem > 0 {
		return dis+1
	} else {
		return dis
	}
}