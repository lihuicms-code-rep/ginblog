package service

import (
	"github.com/lihuicms-code-rep/ginblog/dao/db"
	"github.com/lihuicms-code-rep/ginblog/model"
)

//获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1.获取文章列表
	articleList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil || len(articleList) <= 0 {
		return
	}


	//2.获取文章对应的分类
	categoryIDS := GetCategoryIDS(articleList)
	categoryList, err := db.GetCategoryList(categoryIDS)

	//3.聚合数据,文章和分类关联起来
	for _, article := range articleList {
		record := &model.ArticleRecord{
			Article:*article,
		}

		cid := article.CategoryID
		//根据分类id获取category
		for _, c := range categoryList {
			if cid == c.CategoryID {
				record.Category = *c
				break
			}
		}

		articleRecordList = append(articleRecordList, record)
	}

	return
}


//根据多个文章id获取对应的分类id的结合
func GetCategoryIDS(articleList []*model.Article) (ids []int64) {
	for _, article := range articleList {
		categoryID := article.CategoryID
		//去重
		if idExitsInSlice(categoryID, ids) == false {
			ids = append(ids, categoryID)
		}
	}

	return
}

func idExitsInSlice(id int64, slice []int64) bool {
	for _, item := range slice {
		if id == item {
			return true
		}
	}

	return false
}


//根据分类id,获取该类文章和他们的文章信息
func GetArticleRecordListById(categoryID int64, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1.获取文章列表
	articleList, err := db.GetArticleListByCategoryID(categoryID, pageNum, pageSize)
	if err != nil || len(articleList) <= 0 {
		return
	}


	//2.获取文章对应的分类
	categoryIDS := GetCategoryIDS(articleList)
	categoryList, err := db.GetCategoryList(categoryIDS)

	//3.聚合数据,文章和分类关联起来
	for _, article := range articleList {
		record := &model.ArticleRecord{
			Article:*article,
		}

		cid := article.CategoryID
		//根据分类id获取category
		for _, c := range categoryList {
			if cid == c.CategoryID {
				record.Category = *c
				break
			}
		}

		articleRecordList = append(articleRecordList, record)
	}

	return
}