package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "myproject/routers"
	_ "github.com/mysql"
)

func main() {
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/beegodb?charset=utf8")
	beego.Run()
}

