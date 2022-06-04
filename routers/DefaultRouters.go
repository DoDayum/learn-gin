package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/controllers/myDefault"
	"net/http"
)

func DefaultRouters(r *gin.Engine) {

	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(context *gin.Context) {
			params := context.Params
			for i := range params {
				fmt.Println(i)
			}
			//context.String(http.StatusOK, "值：%v", "你好哇")
			context.HTML(http.StatusOK, "index/index.html", gin.H{
				"title": "首页",
			})
		})

		defaultRouters.GET("/news", myDefault.NewsController{}.Index)
	}

}
