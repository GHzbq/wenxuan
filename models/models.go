package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	FieldUserName = "UserName"
	FieldPassword = "Password"

	FieldArticleIndex = "ArticleIndex"
	FieldArticleTitle = "ArticleTitle"
	FieldArticleContent = "ArticleContent"
	FieldArticleCreationTime = "CreationTime"
	FieldArticleViewCount = "ViewCount"

	TableArticle = "Article"
)

func init() {
	// 设置数据库基本信息
	e := orm.RegisterDataBase("default", "mysql", "root:Root@123@tcp(47.102.208.185:3306)/wenxuan?charset=utf8")
	if e != nil {
		fmt.Printf("orm.RegisterDataBase failed, error = %v\n", e.Error())
		return
	}
	// 映射model数据
	orm.RegisterModel(new(UserData), new(Article))
	// 生成表 第二个参数表示强制更新表 会将表删除，重新生成
	e = orm.RunSyncdb("default", false, true)
	if e != nil {
		fmt.Printf("orm.RunSyncdb failed, error = %v\n", e.Error())
		return
	}
	logs.Info("init database successfully")
}