package article

import (
	"goblog/app/global"
	"goblog/pkg/logger"
	"strconv"
)

// Article 文章模型
type Article struct {
	ID    uint64
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a *Article) Link() string {
	showURL, err := global.Router.Get("articles.show").URL("id", strconv.FormatInt(int64(a.ID), 10))
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return showURL.String()
}
