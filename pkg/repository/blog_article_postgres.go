package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/ValeryChapman/blog-app"
	"github.com/jmoiron/sqlx"
)

type ArticlePostgres struct {
	db *sqlx.DB
}

func NewArticlePostgres(db *sqlx.DB) *ArticlePostgres {
	return &ArticlePostgres{db: db}
}

// create a new list
func (r *ArticlePostgres) Create(list blog.Article) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createArticleQuery := fmt.Sprintf(`INSERT INTO %s (title, text, created_at) VALUES ($1, $2, $3) RETURNING id`, articlesTable)
	row := tx.QueryRow(createArticleQuery, list.Title, list.Text, time.Now())
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

// get list of all articles
func (r *ArticlePostgres) GetAll() []blog.Article {
	var articles []blog.Article

	query := fmt.Sprintf(`SELECT * FROM %s`, articlesTable)
	r.db.Select(&articles, query)

	return articles
}

// get article by id
func (r *ArticlePostgres) GetById(articleId int) (blog.Article, error) {
	var article blog.Article

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, articlesTable)
	err := r.db.Get(&article, query, articleId)

	return article, err
}

// delete article by id
func (r *ArticlePostgres) Delete(articleId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, articlesTable)
	_, err := r.db.Exec(query, articleId)

	return err
}

// update article data
func (r *ArticlePostgres) Update(articleId int, input blog.UpdateArticleInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argId))
		args = append(args, *input.Text)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%d", articlesTable, setQuery, articleId)

	_, err := r.db.Exec(query, args...)
	return err
}
