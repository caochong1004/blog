package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/longjoy/blog/model"
	"strconv"
	"strings"
	"sync"
	"time"
)

func InsertCategory(category *model.Category) (categoryId int64, err error)  {
	sqlStr := `insert into category(category_name, category_no, create_time, update_time) values (?, ?, ? ,?)`
	now := time.Now()
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo, now, now)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

func GetCategoryById(categoryId int64)(category *model.Category, err error)  {
	category = &model.Category{}
	sqlStr := `select id, category_name, category_no from category where id = ?`
	err = DB.Get(category, sqlStr, categoryId)
	if err != nil {
		return
	}
	return
}

func GetCategoryList(categoryIds []int64)(categoryList []*model.Category, err error)  {
	sqlStr, args, err := sqlx.In("select id, category_name, category_no from category where id in(?) order by id desc ", categoryIds)
	if err != nil {
		return
	}
	err = DB.Select(&categoryList, sqlStr, args...)
	if err != nil {
		return
	}
	return
}

func GetAllCategoryList()(categoryList []*model.Category, err error)  {
	sqlStr := `Select id, category_name, category_no from category `
	err = DB.Select(&categoryList, sqlStr)
	if err !=nil {
		 return
	}
	return
}

func InsertTest() (err error)  {
	type Test struct {
		Id int `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
	}
	var wg sync.WaitGroup
	wg.Add(6)
	for i := 1; i <=6; i++ {
		go func() {
			testSlice := make([]*Test, 0)
			for i := 1; i <= 60000; i++{
				s:= &Test{
					Name: "name_"+strconv.Itoa(i),
				}
				testSlice = append(testSlice,s)
			}
			valueStrings := make([]string, 0, len(testSlice))
			// 存放values的slice
			valueArgs := make([]interface{}, 0, len(testSlice) * 1)
			for _, u := range testSlice {
				// 此处占位符要与插入值的个数对应
				valueStrings = append(valueStrings, "(?)")
				valueArgs = append(valueArgs,u.Name)

			}
			stmt := fmt.Sprintf("insert into test (name) VALUES %s",
				strings.Join(valueStrings, ","))

			_, err = DB.Exec(stmt, valueArgs...)
			wg.Done()
		}()
	}

	wg.Wait()
	if err != nil {
		return
	}
	return
}
