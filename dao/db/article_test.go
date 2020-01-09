package db

import (
	"fmt"
	"github.com/lihuicms-code-rep/ginblog/model"
	"testing"
	"time"
)


func TestInsertArticle(t *testing.T) {
	article := &model.Article{}
    article.CategoryID = 1
    article.CommentCnt = 0
    content := "i am new article content ....."
    article.Content = content
    article.Username = "孙笑川"
    article.CreateTime = time.Now()
    article.Title = "我是你哥哥"
    article.Summary = "浮生日记"
    article.ViewCnt = 1
    articleID, err := InsertArticle(article)
    if err != nil {
    	t.Logf("insert article err")
		return
	}

    t.Logf("insert article success, articleID=%d", articleID)
}


func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		fmt.Println("get article list error ", err)
		return
	}


	for i, v := range articleList {
		t.Logf("%d article %v", i+1, v)
	}
}

func TestGetArticleDetail(t *testing.T) {
	article, err := GetArticleDetail(12)
	if err != nil {
		fmt.Println("get article detail err", err)
		return
	}

	t.Logf("article detail:%#v", article)
}

func TestGetArticleListByCategoryID(t *testing.T) {
	articleList, err := GetArticleListByCategoryID(1, 1, 15)
	if err != nil {
		return
	}

	for _, v := range articleList {
		t.Logf("articleID:%d, article:%#v", v.Id, v)
	}
}