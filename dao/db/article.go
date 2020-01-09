package db

import (
	"errors"
	"fmt"
	"github.com/lihuicms-code-rep/ginblog/model"
)

//文章相关操作

//插入文章
func InsertArticle(article *model.Article) (int64, error) {
	if article == nil {
		return -1, errors.New("article is nil")
	}

	sqlStr := "insert into article(content,summary,title,username,category_id,view_count,comment_count) values(?,?,?,?,?,?,?)"
	result, err := DB.Exec(sqlStr, article.Content, article.Summary, article.Title, article.Username,
		article.CategoryID, article.ViewCnt, article.CommentCnt)
	if err != nil {
		fmt.Println("insert article error", err)
		return -1, errors.New("insert article error")
	}

	articleID, err := result.LastInsertId()
	return articleID, err
}

//获取文章列表,分页处理
func GetArticleList(pageNum, pageSize int) ([]*model.Article, error) {
	if pageNum < 0 || pageSize < 0 {
		return nil, errors.New("param error")
	}

	var al = make([]*model.Article, 0)

	sqlStr := "select id,summary,title,view_count,create_time,comment_count,username,category_id from article where status = 1 order by create_time desc limit ?,?"

	err := DB.Select(&al, sqlStr, pageNum-1, pageSize)
	return al, err
}


//根据文章id,查询单个文章
func GetArticleDetail(articleID int64) (detail *model.Article, err error) {
	if articleID < 0 {
		return
	}

	sqlStr := "select id,content,summary,title,view_count,comment_count,username,category_id,create_time from article where id = ? and status = 1"


	err = DB.Get(&detail, sqlStr, articleID)
	return
}

//根据分类id,查询这一类文章
func GetArticleListByCategoryID(categoryID int64, pageNum, pageSize int) (articleList []*model.Article, err error) {
	if categoryID < 0 || pageNum < 0 || pageSize < 0 {
		return
	}

	sqlStr := `select id,summary,title,view_count,create_time,comment_count,username,category_id
               from article where status = 1 and category_id = ? order by create_time desc limit ?,?
              `
	err = DB.Select(&articleList, sqlStr, categoryID, pageNum, pageSize)
	return
}