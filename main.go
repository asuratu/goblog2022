package main

import (
	"goblog/app/global"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	"goblog/pkg/logger"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()
	bootstrap.InitRouter()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(global.Router))
	logger.LogError(err)
}
