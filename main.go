package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(context *gin.Context) {
		params := context.Params
		for i := range params {
			fmt.Println(i)
		}
		//context.String(http.StatusOK, "值：%v", "你好哇")
		context.HTML(http.StatusOK, "index/index.html", gin.H{
			"title": "首页",
		})
	})

	r.GET("/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "admin/news.html", gin.H{
			"news": "新闻信息",
		})
	})

	r.GET("/getRequest", func(context *gin.Context) {
		query := context.Query("uid")
		// 默认值
		defaultQuery := context.DefaultQuery("page", "1")
		context.String(http.StatusOK, "收到参数：%v,page:%v", query, defaultQuery)
	})

	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
