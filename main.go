package main

import (
	"github.com/astaxie/beego"
	"wenxuan/log"
	_ "wenxuan/routers"
)

func main() {
	_ = log.Init("conf/log.json")
	beego.Run()
}

