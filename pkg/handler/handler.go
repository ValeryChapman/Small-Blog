package handler

import (
	"github.com/ValeryChapman/blog-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// routes
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		lists := api.Group("/articles")
		{
			lists.POST("/", h.userIdentity, h.createArticle)
			lists.GET("/", h.getAllArticles)
			lists.GET("/:id", h.getArticleById)
			lists.PUT("/:id", h.userIdentity, h.updateArticle)
			lists.DELETE("/:id", h.userIdentity, h.deleteArticle)
		}
	}

	return router
}

//created_at timestamp    with time zone not null,
