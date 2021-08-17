package service

import (
	"github.com/ValeryChapman/blog-app"
	"github.com/ValeryChapman/blog-app/pkg/repository"
)

// article
type Article interface {
	Create(article blog.Article) (int, error)
	GetAll() []blog.Article
	GetById(articleId int) (blog.Article, error)
	Delete(articleId int) error
	Update(articleId int, input blog.UpdateArticleInput) error
}

type Service struct {
	Article
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Article: NewArticleService(repos.Article),
	}
}
