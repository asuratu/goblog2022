package main

import (
	"goblog/app/global"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/logger"
	"net/http"
)

func main() {
	bootstrap.SetupDB()
	bootstrap.InitRouter()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(global.Router))
	logger.LogError(err)
}
