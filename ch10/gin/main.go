package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// go modules 添加依赖 删除未使用的依赖项
// go get, go mod tidy, go install
// 升级 go get -u 升级到最新的次要版本或者修订版本
// go get -u=patch 升级到最新的修订版本
// go get github.com/go-redis/redis/v8@v1.0.0 会修改go.mod文件
// A项目，仓库是project-A，但是go.mod设置的是github.com/bobby/A
// B项目由于依赖了A项目，import的github.com/bobby/project-A, go get命令的时候，package和代码仓库的名称不一样，replace
// go.mod edit -replace github.com/bobby/A=github.com/bobby/project-A@v1.0.0
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
