package model

//文章分类 实体类
type Category struct {
	CategoryID   int64  `db:"id"`           //分类ID
	CategoryName string `db:"category_name"` //分类名
	CategoryNo   int    `db:"category_no"`  //排序
}
