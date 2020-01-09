package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lihuicms-code-rep/ginblog/service"
	"net/http"
	"strconv"
)

//访问主页控制器
func IndexHandle(c *gin.Context) {

	//1.加载文章数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}

	//2.加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}

	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":articleRecordList,
		"category_list":categoryList,
	})
}

//分类信息
func CategoryList(c *gin.Context) {
	categoryID, err := strconv.ParseInt(c.Query("category_id"), 10, 64)
    if err != nil {
    	c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//根据分类id获取该分类下文章
	articleRecordList, err := service.GetArticleRecordListById(int64(categoryID), 0, 15)
    if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//再次加载所有分裂数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}

	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":articleRecordList,
		"category_list":categoryList,
	})


}
