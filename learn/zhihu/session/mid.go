package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(PrintFullPath())
	app.GET("/", func(context *gin.Context) {
		fmt.Println("Hello World")
		context.Writer.WriteString("Hello World")
	})
	app.Run()
}

func PrintFullPath() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		// 路由处理前执行
		fmt.Printf("接收到请求：%s\n", path)
		context.Next()
		// 路由处理后执行
		fmt.Printf("请求处理完成: %s\n", path)
	}
}
