package main

import (
	_ "bee-go-myBlog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_"github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:123456@127.0.0.1:5432/postgres?sslmode=disable")

}

func main() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"/Users/zhu/go/src/bee-go-myBlog/test.log"}`)
	orm.Debug = true

	beego.Run()
}

