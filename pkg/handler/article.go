package handler

import (
	"net/http"
	"strconv"

	"github.com/ValeryChapman/blog-app"
	"github.com/gin-gonic/gin"
)

// article id structure
type getArticleIdResponse struct {
	Id int `json:"id"`
}

// create a list
func (h *Handler) createArticle(c *gin.Context) {

	var input blog.Article
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, 2000, err.Error())
		return
	}

	id, err := h.services.Article.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, 2000, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data": getArticleIdResponse{
			Id: id,
		},
	})
}

// list of articles structure
type getAllArticlesResponse struct {
	Data []blog.Article `json:"data"`
}

// get all articles
func (h *Handler) getAllArticles(c *gin.Context) {

	lists := h.services.Article.GetAll()

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data": getAllArticlesResponse{
			Data: lists,
		},
	})

}

// get an article by id
func (h *Handler) getArticleById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, 2000, "invalid id param")
		return
	}

	article, err := h.services.Article.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, 2000, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   article,
	})
}

// delete article by id
func (h *Handler) deleteArticle(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, 2000, "invalid id param")
		return
	}

	err = h.services.Article.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, 2000, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// update list data by list id
func (h *Handler) updateArticle(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, 2000, "invalid id param")
		return
	}

	var input blog.UpdateArticleInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, 2000, err.Error())
		return
	}

	if err := h.services.Article.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, 2000, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
