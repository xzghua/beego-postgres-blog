package main

import (
	_ "bee-go-myBlog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/lib/pq"

	"github.com/astaxie/beego/logs"
	"time"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:123456@127.0.0.1:5432/postgres?sslmode=disable")

}



func main() {
	beego.SetLogFuncCall(true)

	x := time.Date(2018, 01, 01, 17, 30, 20, 20, time.Local)
	str := x.Format("2000-01-01")
	logs.SetLogger(logs.AdapterFile,`{"filename":"`+beego.AppConfig.String("log_path")+str+`-project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	orm.Debug = true

	beego.Run()
}

