package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	UserName string `json:"user_name" form:"userName"`
	UserAge  string `json:"user_age" form:"userAge"`
}

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/getRequest", func(context *gin.Context) {
			query := context.Query("uid")
			// 默认值
			defaultQuery := context.DefaultQuery("page", "1")
			context.String(http.StatusOK, "收到参数：%v,page:%v", query, defaultQuery)
		})

		apiRouters.POST("/formPostRequest", func(context *gin.Context) {
			userName := context.PostForm("userName")
			userAge := context.DefaultPostForm("userAge", "20")

			user := &user{}
			err := context.ShouldBind(user)
			if err != nil {
				return
			} else {
				fmt.Println(user)
			}

			context.String(http.StatusOK, "后端收到的参数：userName:{%d},userAge:{%d}", userName, userAge)
		})

	}

}
