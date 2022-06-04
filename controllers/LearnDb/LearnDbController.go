package LearnDb

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
go get -u gorm.io/gorm && go get -u gorm.io/driver/sqlite

*/
type LearnDatabaseController struct{}

func (a LearnDatabaseController) Query(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "查询功能",
	})
}
