package repository

import (
	"github.com/ValeryChapman/blog-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	GetUser(token string) error
}

type Article interface {
	Create(article blog.Article) (int, error)
	GetAll() []blog.Article
	GetById(articleId int) (blog.Article, error)
	Delete(articleId int) error
	Update(articleId int, input blog.UpdateArticleInput) error
}

type Repository struct {
	Authorization
	Article
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Article: NewArticlePostgres(db),
	}
}
