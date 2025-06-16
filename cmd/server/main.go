package main

// 使用gin框架写一个服务器
import (
	"aitools/internal/api"
)

func main() {
	// 设置路由
	r := api.SetupRouter()

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
