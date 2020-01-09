package service

import (
	"github.com/lihuicms-code-rep/ginblog/dao/db"
	"github.com/lihuicms-code-rep/ginblog/model"
)

//获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		return
	}

	return
}
