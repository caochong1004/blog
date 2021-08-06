package service

import (
	"github.com/longjoy/blog/dao/db"
	"github.com/longjoy/blog/model"
)

func GetCategoryList()(categoryLists []*model.Category)  {
	categoryLists, err := db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}

func TestInsert()(err error)  {
	err = db.InsertTest()
	if err != nil {
		return
	}
	return
}
