package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lihuicms-code-rep/ginblog/controller"
	"github.com/lihuicms-code-rep/ginblog/dao/db"
)

func init() {
	//连接mysql
	dsn := "root:alfredLH2019@12306@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.Init(dsn)
	if err != nil {
		panic(err)
	}

}
func main() {
	router := gin.Default()

	//加载静态文件
	router.Static("/static/", "./static")
	//加载模板
	router.LoadHTMLGlob("views/*")

	//路由
	router.GET("/", controller.IndexHandle)
	router.GET("/category", controller.CategoryList)

	//启动并服务
	router.Run(":8888")
}