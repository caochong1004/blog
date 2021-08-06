package db

import "github.com/longjoy/blog/model"

func GetMaterialByProductId(productId int64)(ms []*model.ProductMaterial, err error)  {
	sql := "select material_id from shop_product_material where status = 1 and product_id =?"
	err = DBWeb.Select(&ms, sql, productId)
	if err != nil {
		return
	}
	return
}
