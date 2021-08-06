package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog/model"
	"github.com/longjoy/blog/service"
	"strconv"
)

func AddCategory(c *gin.Context)  {
	var shop_category model.GoodsCategoryDetail
	c.Bind(&shop_category)
	cat_id := service.AddCategory(&shop_category)
	c.JSON(200, cat_id)
}

/**
商品分类复制
 */
func GetShopCategory(c *gin.Context)  {
	catID := c.PostForm("cat_id")
	cat_id, err := strconv.ParseInt(catID, 10, 64)
	if err != nil {
		c.JSON(500, "error")
	}
	result := service.CopyCategory(cat_id)
	c.JSON(200, result)
}

func CategoryList(c *gin.Context)  {
	page := c.DefaultQuery("page", "1")
	page_size := c.DefaultQuery("page_size", "20")
	page_int, err  := strconv.Atoi(page)
	if err != nil {
		c.JSON(500, err)
	}
	page_size_int, err  := strconv.Atoi(page_size)
	if err != nil {
		c.JSON(500, err)
	}

	result := service.GetCategoryLists(page_int, page_size_int)
	c.JSON(200, result)
}

func UpdateCategoryStatus(c *gin.Context)  {
	cat_ids := c.PostFormArray("cat_id")
	status := c.DefaultPostForm("type", "1")
	cat_status, err  := strconv.ParseInt(status, 10, 64)
	if err != nil {
		c.JSON(500, err)
	}
    rows, err := service.UpdateCategoryStatus(cat_ids, cat_status)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, rows)
}

func GetParentCategory(c *gin.Context)  {
	param := c.Query("cat_id")
	cat_id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		c.JSON(500, err)
	}
	result, err := service.ParentCategory(cat_id)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, result)
}

func EditCategory(c *gin.Context)  {
	var shop_category model.GoodsCategoryDetail
	c.Bind(&shop_category)
	rows, err := service.EditCategory(&shop_category)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, rows)
}

func GetCategoryDetail(c *gin.Context)  {
	p := c.Query("cat_id")
	cat_id, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		c.JSON(500, err)
	}
	detail, err := service.CategoryDetail(cat_id)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, detail)
}

