package routers

import (
	"github.com/gin-gonic/gin"
	"learn-gin/controllers/LearnDb"
)

func LearnDbRouters(r *gin.Engine) {
	dbRouters := r.Group("/learnBd")
	{
		dbRouters.POST("/query", LearnDb.LearnDatabaseController{}.Query)
	}
}
