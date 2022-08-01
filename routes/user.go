package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
)

// RegisterUserRoutes 注册网页相关路由
func RegisterUserRoutes(r *mux.Router) {
	auc := new(controllers.AuthController)

	// 用户注册
	r.HandleFunc("/auth/register", auc.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", auc.DoRegister).Methods("POST").Name("auth.doregister")

	// 用户登录
	r.HandleFunc("/auth/login", auc.Login).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", auc.DoLogin).Methods("POST").Name("auth.dologin")

	// 用户登出
	r.HandleFunc("/auth/logout", auc.Logout).Methods("POST").Name("auth.logout")

	// --- 全局中间件 ---

	// 中间件：强制内容类型为 HTML
	//r.Use(middlewares.ForceHTML)

	// 开始会话
	r.Use(middlewares.StartSession)
}
