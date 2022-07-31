package routes

import (
	"goblog/app/http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	// 静态页面
	pc := controllers.PagesController{}
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// 自定义 404 页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
}
