package article

import (
	"goblog/app/global"
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/logger"
	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`

	UserID uint64 `gorm:"not null;index"`
	User   user.User
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

// CreatedAtDate 创建日期
func (a *Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}
