package bootstrap

import (
	"embed"
	"goblog/app/global"
	"goblog/routes"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute(staticFS embed.FS) {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	routes.RegisterUserRoutes(router)

	// 静态资源
	sub, _ := fs.Sub(staticFS, "public")
	router.PathPrefix("/").Handler(http.FileServer(http.FS(sub)))

	global.Router = router
}
