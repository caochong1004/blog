package model

type GoodsCategory struct {
	ID int64 `db:"id" form:"id" json:"id"`
	ParentId int `db:"parent_id" form:"parent_id" json:"parent_id"`
	Name string `db:"name" form:"name" json:"name"`
	Level int `db:"level" form:"level" json:"level"`
	Status int `db:"status" form:"status" json:"status"`
	RobotStatus int `db:"robot_status" form:"robot_status" json:"robot_status"`

}

type CategoryByParent struct {
	ID int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}


type GoodsCategoryDetail struct {
	GoodsCategory
	MaterialId int `db:"material_id" form:"material_id" json:"material_id"`
	Type int `db:"type" form:"type" json:"type"`
	Describe string `db:"describe" form:"describe" json:"describe"`
	SeoKeyword string `db:"seo_keyword" form:"seo_keyword" json:"seo_keyword"`
	SeoTitle  string `db:"seo_title" form:"seo_title" json:"seo_title"`
	SeoDescribe string `db:"seo_describe" form:"seo_describe" json:"seo_describe"`
	DisplayWap string `db:"display_wap" form:"display_wap" json:"display_wap"`
	Sort  int  `db:"sort" form:"sort" json:"sort"`
}

type CategoryDetail struct {
	GoodsCategoryDetail
	ShopMaterialUrl

}

type CategoryId struct {
	ID int `db:"id" json:"id"`
}
