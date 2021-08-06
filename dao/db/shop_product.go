package db

import (
	"bytes"
	"fmt"
	"github.com/longjoy/blog/model"
	"log"
	"strconv"
	"strings"
	"time"
)

func GetProductList(params *model.RequestByProductList)(spl []*model.ShopProductList, err error)  {
	sql := "select p.id, p.name, p.visitors, p.pageviews,p.sale_type, m.url from shop_product as p left join shop_material m on" +
		" p.material_id = m.id where p.sale_type < 4"
	var sqlStr bytes.Buffer
	sqlStr.WriteString(sql)
	if params.Name != "" {
		where1 := " and  p.name like '%"
		sqlStr.WriteString(where1)
		sqlStr.WriteString(params.Name)
		sqlStr.WriteString("%'")
	}
	if params.CatId != 0 {
		where2 := " and p.cat_id = "
		sqlStr.WriteString(where2)
		sqlStr.WriteString(strconv.Itoa(params.CatId))
	}
	if params.MinSales != 0 {
		where3 := "  and p.sales >= "
		sqlStr.WriteString(where3)
		sqlStr.WriteString(strconv.Itoa(params.MinSales))
	}
	if params.MaxSales != 0 {
		where4 := " and p.sales <= "
		sqlStr.WriteString(where4)
		sqlStr.WriteString(strconv.Itoa(params.MaxSales))
	}

	if params.Type != 0 {
		if params.Type == 1 {
			where5 := " and  p.sale_type <= 2"
			sqlStr.WriteString(where5)
		}
		if params.Type == 2 {
			where5 := " and p.sale_type = 3"
			sqlStr.WriteString(where5)
		}
	}

	pageStr := " order by p.id desc  limit ?,? "
	sqlStr.WriteString(pageStr)
	sql = sqlStr.String()
	offset := (params.Page - 1) * params.PageSize
	err = DBWeb.Select(&spl, sql, offset, params.PageSize)
	if err != nil {
		log.Printf("pl_error=", err)
		return
	}
	return

}

func GetSaleTypeById(productId int64) (st int)  {
	sql := "select sale_type from shop_product where id = ?"
	err := DBWeb.Get(&st, sql, productId)
	if err != nil {
		log.Printf("sale_type_sql=",err)
		return 0
	}
	return
}

func GetProductById(productId int64)(productInfo *model.ShopProductDetail, err error)  {
	productInfo = &model.ShopProductDetail{}
	sql := "select * from shop_product where id = ?"
	err = DBWeb.Get(productInfo, sql, productId)
	if err != nil {
		log.Printf("product_info_err=", err)
		return
	}
	return
}

func CopyProduct(p *model.ShopProductDetail, sd []*model.SkuDetail, ms []*model.ProductMaterial)(err error)  {
	tx, err := DBWeb.Begin()
	if err != nil {
		log.Printf("copy_product_1=",err)
		return
	}
	current := time.Now()
	//先复制商品表
	sqlByProduct := "insert into shop_product set name =?,cat_id=?,type=?,lite_name=?,seo_title=?,seo_keyword=?,seo_desc=?,material_id=?," +
		"video_material_id=?,sale_point=?,share_desc=?,tags=?,reduce_type=?,is_has_sku=?,send_type=?,send_url=?," +
		"sale_type=?,sale_time=?,sale_service=?,`limit`=?,limit_type=?,limit_num=?,period_limit_type=?,message=?,message_is_require=?," +
		"detail=?,created_at=?,updated_at=?"
	//log.Printf("sale_time=", time.Parse("2006-01-02 15:03:04",p.SaleTime))
	exec, err := tx.Exec(sqlByProduct, p.Name, p.CatId, p.Type, p.Lite_name, p.SeoTitle, p.SeoKeyword, p.SeoDesc, p.MaterialId, p.VideoMaterialId,
		p.SalePoint, p.ShareDesc, p.Tags, p.ReduceType, p.IsHasSku, p.SendType, p.SendUrl, p.SaleType, p.SaleTime, p.SaleService, p.Limit,
		p.LimitType, p.LimitNum,p.PeriodLimitType, p.Message, p.MessageIsRequire, p.Detail, current, current)
	if err != nil {
		log.Printf("insert_pro_err=", err)
		tx.Rollback()
		return
	}
	lastProductId, err := exec.LastInsertId()
	if err != nil {
		log.Printf("product insert err =", err)
		tx.Rollback()
		return
	}
	//商品sku表

	type SkuData struct{
		ProductId int64 `json:"product_id" db:"product_id"`
		SpecsIds string `json:"specs_ids" db:"specs_ids"`
		Price float64 `json:"price" db:"price"`
		OriginPrice float64 `json:"original_price" db:"original_price"`
		CostPrice float64 `json:"cost_price" db:"cost_price"`
		Sold int `json:"sold" db:"sold"`
		Stock int `json:"stock" db:"stock"`
		Remain int `json:"remain" db:"remain"`
		SendUrl string `json:"send_url" db:"send_url"`
		Code string `json:"code" db:"code"`
		SendType int `json:"send_type" db:"send_type"`
		Sales int `json:"sales" db:"sales"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	}
	skuDataSlice := make([]*SkuData, 0)
	for _, sku := range sd{
		s:= &SkuData{ProductId: lastProductId, SpecsIds: sku.SpecsIds, Price: sku.Price, OriginPrice: sku.OriginPrice,
				CostPrice: sku.CostPrice, Sold: sku.Sold,Stock: sku.Stock,Remain: sku.Remain,SendType: sku.SendType,
				SendUrl: sku.SendUrl,Sales: sku.Sales,CreatedAt: current,UpdatedAt: current,Code: sku.Code,
			}
			skuDataSlice = append(skuDataSlice,s)
	}

	valueStrings := make([]string, 0, len(skuDataSlice))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(skuDataSlice) * 14)
	// 遍历users准备相关数据
	for _, u := range skuDataSlice {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		valueArgs = append(valueArgs, u.ProductId)
		valueArgs = append(valueArgs, u.SpecsIds)
		valueArgs = append(valueArgs, u.Price)
		valueArgs = append(valueArgs, u.OriginPrice)
		valueArgs = append(valueArgs, u.CostPrice)
		valueArgs = append(valueArgs, u.Code)
		valueArgs = append(valueArgs, u.Stock)
		valueArgs = append(valueArgs, u.Remain)
		valueArgs = append(valueArgs, u.Sold)
		valueArgs = append(valueArgs, u.Sales)
		valueArgs = append(valueArgs, u.SendType)
		valueArgs = append(valueArgs, u.SendUrl)
		valueArgs = append(valueArgs, u.CreatedAt)
		valueArgs = append(valueArgs, u.UpdatedAt)
	}
	stmt := fmt.Sprintf("insert into shop_sku_detail (product_id, specs_ids, price,original_price,cost_price," +
		"code,stock,remain,sold,sales,send_type,send_url,created_at,updated_at) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err = tx.Exec(stmt, valueArgs...)
	if err != nil {
		log.Printf("sku insert err =", err)
		tx.Rollback()
		return
	}

	//处理商品关联的素材
	type ProductMaterial struct {
		ProductId int64 `json:"product_id" db:"product_id"`
		MaterialId int64 `json:"material_id" db:"material_id"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	}

	materialSlice := make([]*ProductMaterial, 0)
	for _, m := range ms {
		s1 := &ProductMaterial{ProductId: lastProductId, MaterialId: m.MaterialId, CreatedAt: current, UpdatedAt: current}
		materialSlice = append(materialSlice, s1)
	}

	valueStringsMaterial := make([]string, 0, len(materialSlice))
	// 存放values的slice
	valueArgsMaterial := make([]interface{}, 0, len(materialSlice) * 4)
	for _, mls := range materialSlice {
		// 此处占位符要与插入值的个数对应
		valueStringsMaterial = append(valueStringsMaterial, "(?,?,?,?)")
		valueArgsMaterial = append(valueArgsMaterial, mls.ProductId)
		valueArgsMaterial = append(valueArgsMaterial, mls.MaterialId)
		valueArgsMaterial = append(valueArgsMaterial, mls.CreatedAt)
		valueArgsMaterial = append(valueArgsMaterial, mls.UpdatedAt)
	}
	stmt1 := fmt.Sprintf("insert into shop_product_material (product_id, material_id, created_at,updated_at) values %s",
		strings.Join(valueStringsMaterial, ","))
	_, err = tx.Exec(stmt1, valueArgsMaterial...)
	if err != nil {
		log.Printf("material_insert_err=", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

