package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog/model"
	"github.com/longjoy/blog/service"
	"strconv"
)

func ProductList(c *gin.Context)  {
	params := model.RequestByProductList{}
	c.Bind(&params)
	list, err := service.GetProductList(&params)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, list)
}

func ProductCopy(c *gin.Context)  {
	params := c.PostForm("product_id")
	productId, _ := strconv.ParseInt(params, 10, 64)
	err := service.CopyProduct(productId)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, 1)
}
