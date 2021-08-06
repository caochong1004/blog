package db

import (
	"github.com/longjoy/blog/model"
	"log"
	"testing"
	"time"
)

func init()  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blogger?parseTime=true&loc=PRC"
	err := Init(dsn)
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {
	articleDetail :=&model.ArticleDetail{}
	articleDetail.Article.CategoryId = 2
	articleDetail.Content = "CCCCCC"
	articleDetail.Article.Summary = "jkjjDFSDFSFSDFkkjkjkcccc"
	articleDetail.Article.Status = 1
	articleDetail.Article.ViewCount = 0
	articleDetail.Article.CommentCount = 0
	articleDetail.Article.Title = "title1"
	articleDetail.Article.UserName = "lisi"
	articleDetail.Article.Create_time = time.Now()
	article, err := InsertArticle(articleDetail)
	if err != nil {
		panic(err)
	}
	log.Println(article)

}

func TestGetArticleById(t *testing.T) {
	article, err := GetArticleById(4)
	if err != nil{
		panic(err)
	}
	log.Println(article)
}

func TestGetAllArticle(t *testing.T) {
	article, err := GetAllArticle()
	if err != nil{
		panic(err)
	}
	for _, a :=  range article{
		log.Println(a)
	}
}

func TestGetArticlesByCat(t *testing.T) {
	article, err := GetArticlesByCat(1)
	if err != nil{
		panic(err)
	}
	for _, a :=  range article{
		log.Println(a)
	}
}