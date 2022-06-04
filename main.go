package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/routers"
)

func main() {
	r := gin.Default()
	// 读取静态文件
	r.LoadHTMLGlob("templates/**/*")

	routers.DefaultRouters(r)
	routers.ApiRouters(r)
	routers.LearnDbRouters(r)

	// 服务开始运行
	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
