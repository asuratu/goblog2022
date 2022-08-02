package routes

import (
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {
	// 静态页面
	pc := controllers.PagesController{}
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// 自定义 404 页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)

	// 首页, 重定向到文章列表页
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/articles", http.StatusFound)
	}).Methods("GET").Name("home")

	// 文章相关页面
	ac := new(controllers.ArticlesController)
	// 文章详情页面
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	// 文章列表页面
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	// 创建文章页面
	r.HandleFunc("/articles/create", middlewares.Auth(ac.Create)).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", middlewares.Auth(ac.Store)).Methods("POST").Name("articles.store")
	// 编辑文章页面
	r.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(ac.Edit)).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(ac.Update)).Methods("POST").Name("articles.update")
	// 删除文章
	r.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(ac.Delete)).Methods("POST").Name("articles.delete")

	// 文章分类
	cc := new(controllers.CategoriesController)
	r.HandleFunc("/categories/create", middlewares.Auth(cc.Create)).Methods("GET").Name("categories.create")
	r.HandleFunc("/categories", middlewares.Auth(cc.Store)).Methods("POST").Name("categories.store")
	r.HandleFunc("/categories/{id:[0-9]+}", cc.Show).Methods("GET").Name("categories.show")

	// --- 全局中间件 ---

	// 中间件：强制内容类型为 HTML
	//r.Use(middlewares.ForceHTML)

	// 开始会话
	r.Use(middlewares.StartSession)
}
