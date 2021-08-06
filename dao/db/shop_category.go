package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/longjoy/blog/model"
	"log"
	"time"
)

/**
添加商品分类
*/
func AddCategory(gcy *model.GoodsCategoryDetail)(category_id int64, err error)  {
	current_time := time.Now()
	sql := "insert into shop_category(name, level, parent_id, material_id, seo_keyword, seo_title, seo_describe, display_wap, status, robot_status, sort, type, `describe`, created_at, updated_at) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := DBWeb.Exec(sql, gcy.Name, gcy.Level, gcy.ParentId,
		gcy.MaterialId, gcy.SeoKeyword, gcy.SeoTitle,
		gcy.SeoDescribe, gcy.DisplayWap, gcy.Status, gcy.RobotStatus, gcy.Sort, gcy.Type, gcy.Describe, current_time, current_time)
	if err != nil {
		log.Printf("service_exec_err=", err)
		return
	}
	category_id , err = result.LastInsertId()
	if err != nil {
		log.Printf("service_err=", err)
		return
	}
	return
}

func GetShopCategoryById(catId int64) (gcy *model.GoodsCategoryDetail, err error)  {
	gcy = &model.GoodsCategoryDetail{}
	sql := "select id, name, level,parent_id,material_id, type, seo_keyword, `describe`, seo_title, seo_describe, display_wap, status, robot_status, sort from shop_category where id =?"
	err = DBWeb.Get(gcy, sql, catId)
	if err != nil {
		return
	}
	return
}

/**
商品分类复制
*/
func CopyCategory(catId int64) (category_id int64, err error) {
	gcy, err := GetShopCategoryById(catId)
	if err != nil {
		log.Printf("copy_err=", err)
		return
	}
	//复制
	current_time := time.Now()
	sql := "insert into shop_category(`name`, `level`, parent_id, material_id, seo_keyword, seo_title, seo_describe, display_wap, status, robot_status, sort, `type`, `describe`, created_at, updated_at) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := DBWeb.Exec(sql, gcy.Name, gcy.Level, gcy.ParentId,
		gcy.MaterialId, gcy.SeoKeyword, gcy.SeoTitle,
		gcy.SeoDescribe, gcy.DisplayWap, gcy.Status, gcy.RobotStatus, gcy.Sort, gcy.Type, gcy.Describe, current_time, current_time)
	if err != nil {
		log.Printf("service_exec_err=", err)
		return
	}
	category_id , err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

func GetShopCategoryLists(page, page_size int) (gcl []*model.GoodsCategory, err error)  {

	offset := (page - 1) * page_size
	sql := "select id, name, parent_id, level, status, robot_status from shop_category where status = 1 order by sort desc , id desc limit ?,? "
	err = DBWeb.Select(&gcl, sql, offset, page_size)
	if err != nil {
		log.Printf("cat_list_err=", err)
		return
	}
	return
}

func GetChildCat(cat_ids []string) (catId []int64, err error)  {
	sqlStr, args, err := sqlx.In("select id from shop_category where status = 1 and  (id in(?) or parent_id in (?))", cat_ids,cat_ids)
	if err != nil {
		log.Printf("cat_id_in_err=", err)
		return
	}
	err = DBWeb.Select(&catId, sqlStr, args...)
	if err != nil {
		log.Printf("cat_id_err=", err)
		return
	}
	return
}

func UpdateCategoryStatus(catIds []int64, status int64)(rows int64, err error)  {
	sqlStr, args, err := sqlx.In("UPDATE shop_category set status = ?  WHERE id in (?)", status, catIds)
	if err != nil {
		log.Printf("update_category_status_in_err=", err)
		return
	}
	result, err := DBWeb.Exec(sqlStr, args...)
	if err != nil {
		log.Printf("update_category_status_exec_err=", err)
		return
	}
	rows, err  = result.RowsAffected()
	if err != nil {
		return
	}
	return
}

func DbParentCategory(cat_id int64)(cbp []*model.CategoryByParent, err error)  {
	sql := "select id, name from shop_category where status = 1 and parent_id = ?"
	err = DBWeb.Select(&cbp, sql, cat_id)
	if err != nil {
		return
	}
	return
}

func EditCategory(sc *model.GoodsCategoryDetail) (rows int64, err error)  {
	sql := "update  shop_category set name = ?, level=?,parent_id=?,material_id=?,`type`=?, `describe`=?, seo_keyword=?,seo_title=?," +
		"seo_describe=?, display_wap=?, status=?, robot_status=?, sort=?, updated_at=? where id=?"
	result, err := DBWeb.Exec(sql, sc.Name, sc.Level, sc.ParentId, sc.MaterialId, sc.Type, sc.Describe, sc.SeoKeyword, sc.SeoTitle,
		sc.SeoDescribe, sc.DisplayWap, sc.Status, sc.RobotStatus, sc.Sort, time.Now(), sc.ID)
	rows,err  = result.RowsAffected()
	if err != nil {
		return
	}
	return
}

func CategoryDetail(cat_id int64)(cyd *model.CategoryDetail, err error)  {
	cyd = &model.CategoryDetail{}
	sql := "select sc.name, sc.level, sc.parent_id, sc.material_id, sc.type, sc.describe, sc.describe, sc.seo_title, sc.seo_keyword" +
		", sc.display_wap, sc.seo_describe, sc.status, sc.robot_status, sc.sort, m.url  from shop_category as sc left join shop_material as m " +
		"on sc.material_id = m.id where sc.id = ?"

	err = DBWeb.Get(cyd, sql, cat_id)
	if err != nil {
		log.Printf("cyd db err = ", err)
		return
	}
	return
}
