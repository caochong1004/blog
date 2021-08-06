package service

import (
	"github.com/longjoy/blog/dao/db"
	"github.com/longjoy/blog/model"
)

func GetProductList(params *model.RequestByProductList) (list []*model.ShopProductList, err error) {
	list, err = db.GetProductList(params)
	if err != nil {
		return
	}
	return
}

func CopyProduct(productId int64)(err error)  {
	//获取商品详情
	productDetail, err := db.GetProductById(productId)
	if err != nil {
		return
	}

	//获取商品sku详情
	skuDetail, err := db.GetSkuByProductId(productId)
	if err != nil {
		return
	}

	//获取商品素材
	materials, err  := db.GetMaterialByProductId(productId)
	if err != nil {
		return
	}

	err = db.CopyProduct(productDetail, skuDetail, materials)
	if err != nil {
		return
	}
	return
}