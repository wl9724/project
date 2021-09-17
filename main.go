package main

import (
	_ "project/routers"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	driver_name, err := web.AppConfig.String("mysqlurls")
	if err != nil {
		return
	}
	maxIdle := 30
	maxConn := 30
	// 注册连接
	orm.RegisterDataBase("default", "mysql", driver_name, orm.MaxIdleConnections(maxIdle), orm.MaxOpenConnections(maxConn))
	// 自动建表
	orm.RunSyncdb("default", false, true)
	orm.DefaultTimeLoc = time.UTC
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	web.Run()
}
