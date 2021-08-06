package db

import (
	"github.com/longjoy/blog/model"
	"time"
)

func InsertArticle(article *model.ArticleDetail)(articleId int64, err error)  {
	now := time.Now()
	sqlStr := `insert into article(
					category_id, content, title, 
					view_count, comment_count, 
					username, status, summary,
					create_time, update_time) values (?,?,?,?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlStr, article.Article.CategoryId, article.Content, article.Article.Title,
		article.Article.ViewCount, article.Article.CommentCount, article.Article.UserName, article.Article.Status, article.Article.Summary, article.Article.Create_time, now)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

func GetArticleById(articleId int64)(article *model.ArticleDetail, err error)  {
	article = &model.ArticleDetail{}
	sqlStr := `select id, category_id, content, title, view_count, comment_count, username, status, summary, create_time from article where status = 1 and id = ?`
	err = DB.Get(article, sqlStr, articleId)
	if err != nil {
		return
	}
	return
}

func GetAllArticle()(article []*model.Article, err error)  {
	sqlStr := `select id, category_id, title, view_count, comment_count, username, status, summary, create_time from article where status = 1`

	err = DB.Select(&article, sqlStr)
	if err != nil{
		return
	}

	return
}

func GetArticlesByCat(catId int64)(articles []*model.Article, err error)  {
	sqlStr := `select id, category_id, title, view_count, comment_count, username, status, summary, create_time from article where status =1 and category_id = ? `

	err = DB.Select(&articles, sqlStr, catId)
	if err != nil {
		return
	}
	return
}