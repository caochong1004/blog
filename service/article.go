package service

import (
	"github.com/longjoy/blog/dao/db"
	"github.com/longjoy/blog/model"
)

func ArticleList()(articleReLists []*model.ArticleList)  {
	//先查询出文章的列表
	articleLists, err  := db.GetAllArticle()

	if err != nil {
		return
	}
	//获取所有categoryId
	categoryIds := GetCategoryIds(articleLists)
	//log.Printf("categoryids:%v\n", categoryIds)
	categoryLists, err  := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	for _,article := range articleLists {
		articleReList := &model.ArticleList{
			Article:*article,
		}
		for _,category := range categoryLists {
			if article.CategoryId == category.CategoryId {
				articleReList.Category = *category
				break
			}
		}
		articleReLists = append(articleReLists, articleReList)
	}
	return articleReLists
	
}

func GetCategoryIds(articles []*model.Article)(categoryIds []int64)  {
	//categoryIds = append(categoryIds,articleLIst[0].CategoryId)
	for _, article := range articles{
		catId := article.CategoryId
		flag := true
		for _, v := range categoryIds{
			if v == catId {
				flag = false
				break
			}
		}

		if flag {
			categoryIds = append(categoryIds, catId)
		}
	}
	return categoryIds
}

func GetArticlesByCat(catId int64)(articles []*model.Article)  {
	articles, err := db.GetArticlesByCat(catId)
	if err != nil {
		return
	}
	return
}