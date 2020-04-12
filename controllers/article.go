package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"wenxuan/models"
)

// HandleNcGetAddArticle 获取添加文章页面
func (c *MainController) HandleNcGetAddArticle () {
//	c.Data["welcome"] = `
//## 只求极致
//[ **M** ] arkdown + E [ **ditor** ] = **Mditor**
//> Mditor 是一个简洁、易于集成、方便扩展、期望舒服的编写 markdown 的编辑器，仅此而已...`
	// c.TplName = "copy_from_chrome.html"
	c.TplName = "todo_add_article.html"
	//c.TplName = "image_cross_domain_upload.html"
	//c.TplName = "add_article.html"
	//c.TplName = "ali.html"
}

const (
	dealSuccess = "{\"errno\": 0}"
)

// HandleNcPostAddArticle 处理添加文章请求
func (c *MainController) HandleNcPostAddArticle () {
	// 1. 拿到数据
	var tmp struct{
		ArticleTitle string `json:"article_title"`
		ArticleMark string `json:"article_mark"`
		ArticleContent string `json:"article_content"`
	}
	data := c.Ctx.Input.RequestBody
	e := json.Unmarshal(data, &tmp)
	if e != nil {
		logs.Error("json.Unmarshal failed, error = %v", e.Error())
		//c.Redirect("/add_article", 302)
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Ctx.WriteString("{\"errno\": 1, \"errorMsg\": \"json unmarshal failed\"}")
		return
	}
	logs.Debug("data = %v, tmp = %v", data, tmp)
	
	// 2. 判断数据是否合法
	if tmp.ArticleTitle == "" {
		//c.Redirect("/add_article", 302)
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Ctx.WriteString("{\"errno\": 2, \"errorMsg\": \"article title is nil\"}")
		return
	}
	if tmp.ArticleMark == "" || tmp.ArticleContent == "" {
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Ctx.WriteString("{\"errno\": 3, \"errorMsg\": \"article content is nil\"}")
	}
	// 3. 插入数据
	article := &models.Article{
		ArticleTitle:   tmp.ArticleTitle,
		ArticleMark:    tmp.ArticleMark,
		ArticleContent: tmp.ArticleContent,
	}
	e = models.InsertArticle(article)
	if e != nil {
		logs.Error("models.InsertArticle failed, error = %v", e.Error())
		//c.Redirect("/add_article", 302)
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.Ctx.WriteString("{\"errno\": 4, \"errorMsg\": \"operate db failed\"}")
		return
	}

	// 4. 返回文章页面
	//c.Redirect("/index", 302)
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.WriteString(dealSuccess)
}

// HandleNcGetArticle 查看文章详情
func (c *MainController) HandleNcGetArticle() {
	// 1. 获取文章ID
	id, e := c.GetInt64("id")
	if e != nil {
		logs.Error("get id failed")
		c.Redirect("/index", 302)
		return
	}
	logs.Debug("id = %v", id)
	// 2. 查询数据库获取数据
	e, article := models.LoadArticleByID(id)
	if e != nil {
		logs.Error("models.LoadArticleByID failed, id = %v, error = %v", id, e.Error())
		c.Redirect("index", 302)
		return
	}
	// 3. 传递数据给试图
	c.Data["article"] = article
	c.TplName = "article_info.html"
}