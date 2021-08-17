package service

import (
	"github.com/ValeryChapman/blog-app"
	"github.com/ValeryChapman/blog-app/pkg/repository"
)

type ArticleService struct {
	repo repository.Article
}

func NewArticleService(repo repository.Article) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) Create(article blog.Article) (int, error) {
	return s.repo.Create(article)
}

func (s *ArticleService) GetAll() []blog.Article {
	return s.repo.GetAll()
}

func (s *ArticleService) GetById(articleId int) (blog.Article, error) {
	return s.repo.GetById(articleId)
}

func (s *ArticleService) Delete(articleId int) error {
	return s.repo.Delete(articleId)
}

func (s *ArticleService) Update(articleId int, input blog.UpdateArticleInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(articleId, input)
}
