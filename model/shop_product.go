package model

import (
	"database/sql"
	"time"
)

type ShopProduct struct {
	Id int64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Visitors int `json:"visitors" db:"visitors"`
	Pageviews int `json:"pageviews" db:"pageviews"`
	SaleType int `json:"sale_type" db:"sale_type"`
}

type ShopProductDetail struct {
	ShopProduct
	CatId int64 `json:"cat_id" db:"cat_id"`
	Type int64 `json:"type" db:"type"`
	Lite_name string `json:"lite_name" db:"lite_name"`
	SeoTitle string `json:"seo_title" db:"seo_title"`
	SeoKeyword string `json:"seo_keyword" db:"seo_keyword"`
	SeoDesc string `json:"seo_desc" db:"seo_desc"`
	MaterialId int64 `json:"material_id" db:"material_id"`
	VideoMaterialId int64 `json:"video_material_id" db:"video_material_id"`
	SalePoint string `json:"sale_point" db:"sale_point"`
	ShareDesc string `json:"share_desc" db:"share_desc"`
	Tags string `json:"tags" db:"tags"`
	ReduceType int `json:"reduce_type" db:"reduce_type"`
	IsHasSku int `json:"is_has_sku" db:"is_has_sku"`
	SendType int `json:"send_type" db:"send_type"`
	SendUrl string `json:"send_url" db:"send_url"`
	SaleTime sql.NullTime `json:"sale_time" db:"sale_time"`
	SaleService int `json:"sale_service" db:"sale_service"`
	Limit int `json:"limit" db:"limit"`
	LimitType int `json:"limit_type" db:"limit_type"`
	Sales int `json:"sales" db:"sales"`
	LimitNum int `json:"limit_num" db:"limit_num"`
	PeriodLimitType int `json:"period_limit_type" db:"period_limit_type"`
	Message string `json:"message" db:"message"`
	MessageIsRequire int `json:"message_is_require" db:"message_is_require"`
	Detail string `json:"detail" db:"detail"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

}

type ShopProductList struct {
	ShopProduct
	ShopMaterialUrl
}

type RequestByProductList struct {
	Name string `json:"name" form:"name"`
	CatId int `json:"cat_id" form:"cat_id"`
	MinSales int `json:"min_sales" form:"min_sales"`
	MaxSales int `json:"max_sales" form:"max_sales"`
	Type int `json:"type" form:"type"`
	Page int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

