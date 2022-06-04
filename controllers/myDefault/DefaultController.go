package myDefault

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewsController struct{}

func (controller NewsController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/news.html", gin.H{
		"news": "新闻信息",
	})

}
