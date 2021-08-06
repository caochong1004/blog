package model

import (
	"time"
)

type SkuByProduct struct {
	Skuid int64 `json:"sku_id" db:"sku_id"`
	SpecsIds string `json:"specs_ids" db:"specs_ids"`
	Price float64 `json:"price" db:"price"`
	OriginPrice float64 `json:"original_price" db:"original_price"`
	CostPrice float64 `json:"cost_price" db:"cost_price"`
	Sold int `json:"sold" db:"sold"`
	Stock int `json:"stock" db:"stock"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	//SaleType int `json:"sale_type"`
	//Name string `json:"name"`
	Remain int `json:"remain" db:"remain"`
}

type SkuByProductSpec struct {
	SkuByProduct
	Name string `json:"name"`
	SaleType int `json:"sale_type"`
	CreatedTime string `json:"created_time"`
}

type SkuDetail struct {
	SkuByProduct
	SendUrl string `json:"send_url" db:"send_url"`
	SendType int `json:"send_type" db:"send_type"`
	Sales int `json:"sales" db:"sales"`
	Code string `json:"code" db:"code"`

}


