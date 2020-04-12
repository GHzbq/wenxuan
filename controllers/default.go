package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"wenxuan/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// HandleNcGetIndex 获取主页面
func (c *MainController) HandleNcGetIndex() {
	e, articles := models.LoadAllArticles()
	if e != nil {
		logs.Error("models.LoadAllArticles failed, error = %v", e.Error())
		c.Ctx.WriteString("load data failed")
		//c.Redirect("/not_found", 302)
		return
	}
	c.Data["articles"] = articles
	c.TplName = "index.html"
}