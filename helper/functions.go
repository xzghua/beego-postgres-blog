package helper

import (
	"time"
	"math/rand"
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
