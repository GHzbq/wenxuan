package routers

import (
	"wenxuan/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:HandleNcGetIndex")
	beego.Router("/index", &controllers.MainController{}, "get:HandleNcGetIndex")
	beego.Router("/index.html", &controllers.MainController{})

    beego.Router("/login", &controllers.MainController{}, "get:HandleNcGetLogin;post:HandleNcPostLogin")
    beego.Router("/register", &controllers.MainController{}, "get:HandleNcGetRegister;post:HandleNcPostRegister")

    beego.Router("/add_article", &controllers.MainController{}, "get:HandleNcGetAddArticle;post:HandleNcPostAddArticle")
	beego.Router("/article", &controllers.MainController{}, "get:HandleNcGetArticle")

    beego.Router("/upload_picture", &controllers.MainController{}, "get:HandleNcGetUploadPicture")
}
