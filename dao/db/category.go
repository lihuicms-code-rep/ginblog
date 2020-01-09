package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lihuicms-code-rep/ginblog/model"
)

//文章分类相关操作


//添加分类
func InsertCategory(category *model.Category) (int64, error) {
	sqlStr := "insert into category(category_name,category_no) value(?,?)"
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return -1, errors.New("insert category error")
	}

	categoryID, err := result.LastInsertId()
	if err != nil {
		return -1, errors.New("lastInsertId error")
	}

	return categoryID, nil
}


//获取单个分类
func GetCategoryByID(id int64) (*model.Category, error) {
	category := new(model.Category)
	sqlStr := "select id,category_name,category_no from category where id = ?"

	if err := DB.Get(category, sqlStr, id); err != nil {
		fmt.Println("Get error", err)
		return nil, err
	}

	return category, nil
}


//获取多个分类
func GetCategoryList(ids []int64) ([]*model.Category, error) {
	categories := make([]*model.Category, len(ids))
	//构建查询语句
	sqlStr, args, err := sqlx.In("select id,category_name,category_no from category where id in(?)", ids)
	if err != nil {
		return nil, errors.New("sqlStr error")
	}

	//查询
	err = DB.Select(&categories, sqlStr, args...)
	if err != nil {
		fmt.Println("select category list error, ", err)
		return nil, errors.New("select category list error")
	}

	return categories, nil
}


//获取所有分类
func GetAllCategoryList() ([]*model.Category, error) {
	categories := make([]*model.Category, 0)
	//构建查询语句
	sqlStr := "select id,category_name,category_no from category order by category_no asc"
	//查询
	err := DB.Select(&categories, sqlStr)
	if err != nil {
		return nil, errors.New("select all category error")
	}

	return categories, nil
}
