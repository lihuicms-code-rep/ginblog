package model

import "time"

//文章 实体类
type Article struct {
	Id         int64     `db:"id"`            //文章ID
	CategoryID int64     `db:"category_id"`   //所属分类ID
	Content    string    `db:"content"`       //文章内容
	Title      string    `db:"title"`         //标题
	ViewCnt    uint32    `db:"view_count"`    //浏览次数
	CommentCnt uint32    `db:"comment_count"` //评论数
	Username   string    `db:"username"`      //作者
	Status     int64     `db:"status"`        //文章状态
	Summary    string    `db:"summary"`       //摘要
	CreateTime time.Time `db:"create_time"`   //创建时间
	UpdateTime time.Time `db:"update_time"`   //修改时间
}

//文章上下页
type ArticleRecord struct {
	Article
	Category
}
