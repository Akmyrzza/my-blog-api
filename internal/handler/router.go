package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	apiV1.POST("/user", h.createUser)
	apiV1.GET("/user/login", h.login)
	apiV1.GET("/user/:id", h.getUser)
	apiV1.PUT("/user/:id", h.updateUser)
	apiV1.DELETE("/user/:id", h.deleteUser)

	apiV1.POST("/category", h.createCategory)
	apiV1.GET("/category", h.getAllCategory)
	apiV1.GET("/category/:id", h.getCategory)

	apiV1.POST("/article", h.createArticle)
	apiV1.PUT("/article/:id", h.updateArticle)
	apiV1.DELETE("/article/:id", h.deleteArticle)
	apiV1.GET("/article/:id", h.getArticle)
	apiV1.GET("/article", h.getAllArticle)
	apiV1.GET("/article/user/:id", h.getAllArticleByUserID)
	
	return router
}
