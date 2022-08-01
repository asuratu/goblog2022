package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
)

// RegisterUserRoutes 注册网页相关路由
func RegisterUserRoutes(r *mux.Router) {
	// 用户认证
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", auc.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", auc.DoRegister).Methods("POST").Name("auth.doregister")

	// 中间件：强制内容类型为 HTML
	//r.Use(middlewares.ForceHTML)
}
