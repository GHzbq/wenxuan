package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

// Article 文章信息
type Article struct {
	ArticleIndex 	int64 		`orm:"pk;auto"`
	ArticleTitle 	string 		`orm:"size(20)"`
	ArticleMark     string      `orm:"size(2048)"`
	ArticleContent 	string 		`orm:"size(1024)"`
	CreationTime 	time.Time 	`orm:"auto_now_add;type(datetime)"` // YYYY-MM-DD HH:MM:SS
	ModifyTime   	time.Time 	`orm:"auto_now;type(date)"`         // YYYY-MM-DD
	ViewCount    	int32     	`orm:"default(0)"`
}

// InsertArticle 插入文章
func InsertArticle(article *Article) (e error) {
	if article == nil {
		return errors.New("article is nil")
	}
	// 1. 有orm对象
	o := orm.NewOrm()
	// 2. 插入数据
	id, e := o.Insert(article)
	if e != nil {
		logs.Error("o.Insert failed, error = %v", e.Error())
		return e
	}
	article.ArticleIndex = id
	logs.Debug("insert article = %v successfully", *article)
	return nil
}

// LoadArticleByID 通过文章index加载文章
func LoadArticleByID(index int64, cols ...string) (e error, article *Article) {
	// 1. 有orm对象
	o := orm.NewOrm()
	article = &Article{}
	article.ArticleIndex = index
	logs.Debug("cols = %v", cols)
	e = o.Read(article, cols...)
	if e != nil {
		logs.Error("o.Read failed, error = %v", e.Error())
		return e, nil
	} else {
		logs.Debug("load article successfully, articleIndex = %v", article.ArticleIndex)
		return nil, article
	}
}

// LoadAllArticles 加载所有的文章
func LoadAllArticles() (e error, articles []Article) {
	// 1. 有orm对象
	o := orm.NewOrm()
	// articles = make([]*Article, 0)
	//e = o.Read(articles, FieldArticleTitle, FieldArticleCreationTime, FieldArticleViewCount)
	count, e := o.QueryTable(TableArticle).All(&articles, FieldArticleIndex, FieldArticleTitle, FieldArticleCreationTime, FieldArticleViewCount)
	if e != nil {
		logs.Error("o.Read failed, error = %v", e.Error())
		return e, nil
	} else {
		logs.Debug("load %d articles.", count)
		return nil, articles
	}
}

// UpdateArticle 更新文章数据
func UpdateArticle(index int64, article *Article) (e error) {
	if index < 0 || article == nil {
		logs.Error("index = %v, article = %v", index, article)
		return errors.New("data invalid")
	}

	o := orm.NewOrm()
	art := &Article{ArticleIndex:index}
	e = o.Read(art)
	if e != nil {
		logs.Error("o.Read failed, error = %v", e.Error())
		return e
	}
	needUpdate := false
	cols := make([]string, 0)
	// 文章标题不同
	if art.ArticleTitle != article.ArticleTitle {
		art.ArticleTitle = article.ArticleTitle
		needUpdate = true
		cols = append(cols, FieldArticleTitle)
	}
	// 文章内容不同
	if art.ArticleContent != article.ArticleContent {
		art.ArticleContent = article.ArticleContent
		needUpdate = true
		cols = append(cols, FieldArticleContent)
	}

	// 如果需要更新
	if needUpdate == true {
		_, e = o.Update(art, cols...)
		if e != nil {
			logs.Error("o.Update failed, error = %v", e.Error())
			return e
		} else {
			logs.Info("o.Update successfully")
			return nil
		}
	}
	return nil
}

//func AddArticleViewCount(index int64) (e error) {
//	o := orm.NewOrm()
//	article := &Article{ArticleIndex:index}
//	o.ReadForUpdate()
//}

// DeleteArticle 删除文章
func DeleteArticle(index int64) (e error) {
	if index < 0 {
		logs.Error("index = %v", index)
		return errors.New("index < 0")
	}

	o := orm.NewOrm()
	article := &Article{ArticleIndex:index}
	_, e = o.Delete(article)
	if e != nil {
		logs.Error("o.Delete failed, error = %v", e.Error())
		return e
	}
	return nil
}