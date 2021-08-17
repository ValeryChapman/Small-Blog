package blog

import (
	"errors"
	"time"
)

type Article struct {
	Id         int       `json:"id" db:"id"`
	Title      string    `json:"title" db:"title" binding:"required"`
	Text       string    `json:"text" db:"text"`
	Views      int       `json:"views" db:"views"`
	Created_at time.Time `json:"created_at" db:"created_at"`
}

type UpdateArticleInput struct {
	Title *string `json:"title"`
	Text  *string `json:"text"`
}

func (i UpdateArticleInput) Validate() error {
	if i.Title == nil && i.Text == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
