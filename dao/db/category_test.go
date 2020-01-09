package db

import (
	"fmt"
	"testing"
)

//先连接
func init() {
	//parseTime=true的作用, 将mysql中时间类型,自动解析为go中的时间类型
	dsn := "root:alfredLH2019@12306@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dsn)
	if err != nil {
		panic(err)
	}
}

func TestGetCategoryByID(t *testing.T) {
	category, err := GetCategoryByID(1)
	if err != nil {
		fmt.Println("GetCategoryByID error ", err)
		panic(err)
	}

	t.Logf("category:%#v", category)

}

func TestGetCategoryList(t *testing.T) {
	var ids = []int64{1, 2, 3}
	categoryList, err := GetCategoryList(ids)
	if err != nil {
		panic(err)
	}

	for _, v := range categoryList {
		if v == nil {
			continue
		}
		t.Logf("id:%d, category:%#v", v.CategoryID, v)
	}
}

func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}

	for _, v := range categoryList {
		t.Logf("id:%d, category:%#v", v.CategoryID, v)
	}
}
