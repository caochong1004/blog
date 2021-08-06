package main

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog/controller"
	"github.com/longjoy/blog/dao/db"
)

func main()  {
	router := gin.Default()
	dsn := "root:123456@tcp(127.0.0.1:3306)/blogger?parseTime=true&loc=PRC"
	err := db.Init(dsn)
	if err != nil {
		panic(err)
	}
	dsn_web := "root:123456@tcp(127.0.0.1:3306)/shop_web?parseTime=true&loc=PRC"
	err = db.InitWeb(dsn_web)
	if err != nil {
		panic(err)
	}

	dsn_user := "root:123456@tcp(127.0.0.1:3306)/shop_user?parseTime=true&loc=PRC"
	err = db.InitUser(dsn_user)
	if err != nil {
		panic(err)
	}


	err = db.RedisInit()
	if err != nil {
		panic(err)
	}
	defer db.RedisConn.Close()

	/**************************商品分类********************************/
	//添加商品分类
	router.POST("/category-add", controller.AddCategory)
	//复制商品分类
	router.POST("/shop-category-copy", controller.GetShopCategory)
	//批量禁用商品分类
	router.POST("/update-category-status", controller.UpdateCategoryStatus)
	//商品分类列表
	router.GET("/category-list", controller.CategoryList)
	//获得所有父级分类
	router.GET("/category-parent", controller.GetParentCategory)
	//商品分类编辑
	router.POST("/category-edit", controller.EditCategory)
	//商品分类详情
	router.GET("/category-detail",controller.GetCategoryDetail)

	//登录
	router.POST("/login", controller.Login)


	/******************商品相关*************************************/
	//商品列表
	router.GET("/product-list", controller.ProductList)
	router.POST("/product-copy",controller.ProductCopy)

	/*****************************sku相关******************************/
	//某一商品下的sku信息
	router.GET("/sku-list-by-product", controller.SkuListByProduct)




	router.GET("/article", controller.GetArticles)
	router.GET("/articleByCat", controller.GetArticleByCat)
	router.GET("/category", controller.GetCategoryList)
	router.GET("/get",controller.Test)
	router.Run()
}
