package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog/service"
)

func GetCategoryList(c *gin.Context)  {
	categoryLists := service.GetCategoryList()
	c.JSON(200, categoryLists)
}

func Test(c *gin.Context)  {
	err  := service.TestInsert()
	c.JSON(200, err)
}
