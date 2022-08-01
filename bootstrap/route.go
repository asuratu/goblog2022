package bootstrap

import (
	"goblog/app/global"
	"goblog/routes"

	"github.com/gorilla/mux"
)

// InitRouter 路由初始化
func InitRouter() {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	global.Router = router
}
