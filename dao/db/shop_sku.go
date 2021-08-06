package db

import (
	"github.com/longjoy/blog/model"
	"log"
)

func SkuListByProduct(pid int64)(sbp []*model.SkuByProduct, err error)  {
	sql := "select sku_id, specs_ids, price,original_price,cost_price,stock,remain,sold, created_at from shop_sku_detail where product_id=?"
	err = DBWeb.Select(&sbp, sql, pid)
	if err != nil{
		log.Printf("sku_list_err =", err)
		return
	}
	return
}

func GetSkuByProductId(productId int64)(skd []*model.SkuDetail, err error)  {
	sql := "select sku_id, specs_ids, price,original_price,cost_price,remain,sold, stock, created_at, sales,send_type,send_url from shop_sku_detail where product_id=?"

	err = DBWeb.Select(&skd, sql, productId)
	if err != nil {
		return
	}
	return
}
