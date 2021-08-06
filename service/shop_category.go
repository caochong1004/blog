package service

import (
	"github.com/longjoy/blog/dao/db"
	"github.com/longjoy/blog/model"
)

/**
添加商品分类
 */
func AddCategory(gcy *model.GoodsCategoryDetail)(category_id int64)  {
	category_id, err  := db.AddCategory(gcy)
	if err != nil {
		return
	}
	return
}

/**
商品分类复制
 */
func CopyCategory(catId int64) (category_id int64) {
	category_id, err := db.CopyCategory(catId)
	if err != nil {
		return
	}
	return
}

func GetCategoryLists(page, page_size int) (gcl []*model.GoodsCategory)  {
	gcl, err := db.GetShopCategoryLists(page, page_size)
	if err != nil {
		return
	}
	return

}

func UpdateCategoryStatus(cat_ids []string, status int64) (rows int64, err error)  {
	//获取这些分类下所有子分类
	categoryIds, err := db.GetChildCat(cat_ids)
	if err != nil {
		return
	}
	//修改状态
	rows, err = db.UpdateCategoryStatus(categoryIds, status)
	if err != nil {
		return
	}
	return
}

func ParentCategory(cat_id int64)(cbp []*model.CategoryByParent, err error)  {
	cbp, err = db.DbParentCategory(cat_id)
	if err != nil {
		return
	}
	return
}

func EditCategory(sc *model.GoodsCategoryDetail) (rows int64, err error) {
	rows, err = db.EditCategory(sc)
	if err != nil {
		 return
	}
	return
}

func CategoryDetail(cat_id int64) (cyd *model.CategoryDetail, err error) {
	cyd, err = db.CategoryDetail(cat_id)
	if err != nil {
		return
	}
	return
}
