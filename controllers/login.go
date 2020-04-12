package controllers

import (
	"github.com/astaxie/beego/logs"
	"wenxuan/models"
)

// HandleNcGetLogin 获取登录页面
func (c *MainController) HandleNcGetLogin() {
	c.TplName = "login.html"
}

// HandleNcPostLogin 处理登录请求
func (c *MainController) HandleNcPostLogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		logs.Error("login info invalid, username = %v, password = %v", username, password)
		c.Redirect("/login", 302)
		return
	}

	// 加载用户数据
	e, user := models.LoadUser(username)
	if e != nil {
		logs.Error("models.LoadUser failed, error = %v", e.Error())
		c.Redirect("/login", 302)
		return
	}

	// 验证密码
	if user == nil || password != user.Password {
		logs.Error("verify password failed")
		c.Redirect("/login", 302)
		return
	}
	c.Redirect("/index", 302)
}

// HandleNcGetRegister 获取注册页面
func (c *MainController) HandleNcGetRegister() {
	c.TplName = "register.html"
}

// HandleNcPostRegister 处理注册请求
func (c* MainController) HandleNcPostRegister() {
	username := c.GetString("username")
	password := c.GetString("password")
	email := c.GetString("email")
	if username == "" || password == "" || email == "" {
		logs.Error("register info invalid, username = %v, password = %v, email = %v", username, password, email)
		c.Redirect("/register", 302)
		return
	}
	user := &models.UserData{
		ID:       0,
		UserName: username,
		Password: password,
		Email:    email,
	}
	e := models.InsertUser(user)
	if e != nil {
		c.Redirect("/register", 302)
		return
	}
	c.Redirect("/login", 302)
}