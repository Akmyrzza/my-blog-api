package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Akmyrzza/blog-api/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createArticle(ctx *gin.Context) {
	var req entity.Article
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.CreateArticle(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "Article successfully created")
}

func (h *Handler) updateArticle(ctx *gin.Context) {
	var req entity.Article
	id := ctx.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.UpdateArticle(ctx, id64, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "Article successfully created")
}

func (h *Handler) deleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}

	err = h.srvs.DeleteArticle(ctx, id64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Article successfully deleted")
}

func (h *Handler) getArticle(ctx *gin.Context) {

	id := ctx.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	a, err := h.srvs.GetArticle(ctx, id64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, a)
}

func (h *Handler) getAllArticle(ctx *gin.Context) {

	a, err := h.srvs.GetAllArticle(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, a)
}