package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog/service"
	"strconv"
)

func SkuListByProduct(c *gin.Context)  {
	params := c.Query("product_id")
	pId, _ := strconv.ParseInt(params, 10, 64)
	product, err := service.GetSkuListByProduct(pId)
	if err != nil{
		c.JSON(500, err)
	}
	c.JSON(200, product)
}
