package main

import (
	"github.com/dakong-yi/im-go-server/config"
	"github.com/dakong-yi/im-go-server/internal/app"
	"github.com/dakong-yi/im-go-server/internal/db"
)

func main() {
	// 加载配置
	config.LoadConfig()
	// 加载数据库
	db.InitDB()
	// 启动服务
	app.Run()
	// 阻塞主goroutine
	select {}
}
