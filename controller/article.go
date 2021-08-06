package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog/service"
	"strconv"
)

func GetArticles(c *gin.Context)  {
	list := service.ArticleList()
	c.JSON(200,list)
}

func GetArticleByCat(c *gin.Context)  {
	cat := c.Query("cat")
	catId, _ := strconv.Atoi(cat)
	articles := service.GetArticlesByCat(int64(catId))

	c.JSON(200, articles)
}
