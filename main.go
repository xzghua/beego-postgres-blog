package main

import (
	_ "beego-postgres-blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/lib/pq"

	"time"
	"github.com/astaxie/beego/logs"
	"beego-postgres-blog/common"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	postgresPwd := beego.AppConfig.String("postgrespass")
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:"+postgresPwd+"@127.0.0.1:5432/postgres?sslmode=disable")

}



func main() {
	beego.SetLogFuncCall(true)
	o := orm.NewOrm()
	o.Using("default")
	str := time.Now().Format("2006-01-02")
	logs.SetLogger(logs.AdapterFile,`{"filename":"`+beego.AppConfig.String("log_path")+str+`-project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	orm.Debug = true
	beego.BConfig.Log.AccessLogs = false

	beego.AddFuncMap("iphpt_rem",common.Rem)
	beego.Run()
}

