package db

import (
	"github.com/longjoy/blog/model"
	"log"
	"testing"
)

func init()  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blogger?parseTime=true&loc=Local"
	err := Init(dsn)
	if err != nil {
		panic(err)
	}

}

func TestInsertCategory(t *testing.T) {
	category := &model.Category{
		CategoryNo: 4,
		CategoryName: "php",
	}
	categoryId, err := InsertCategory(category)
	if err != nil {
		panic(err)
	}
	log.Printf("插入数据成功 id: %d\n", categoryId)
}

func TestGetCategoryById(t *testing.T) {
	id, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	log.Println(id)
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1,4)
	list, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}
	for _,v := range list{
		log.Println(v)
	}
}

func TestGetAllCategoryList(t *testing.T) {
	list, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}
	for _,v := range list{
		log.Println(v)
	}
}
