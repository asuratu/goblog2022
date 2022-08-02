package bootstrap

import (
	"goblog/app/global"
	"goblog/routes"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	routes.RegisterUserRoutes(router)
	global.Router = router
}
