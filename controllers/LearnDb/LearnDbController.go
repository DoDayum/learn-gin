package LearnDb

import (
	"github.com/gin-gonic/gin"
	"learn-gin/models"
	"net/http"
)

type LearnDatabaseController struct{}

func (a LearnDatabaseController) Query(context *gin.Context) {

	var userList []models.User
	//models.DB.Debug().Find(&userList)
	models.DB.Debug().Where("name = ?", "demo02").Find(&userList)
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": userList,
	})
}
