package main

import (
	"github.com/astaxie/beego"
	_ "wenxuan/routers"
	"wenxuan/utils"
)

func main() {
	_ = utils.Init("conf/log.json")
	beego.Run()
}

