package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Akmyrzza/blog-api/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCategory(ctx *gin.Context) {
	var req entity.Category

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.CreateCategory(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Category successfully created"})
}

func (h *Handler) getAllCategory(ctx *gin.Context) {

	categories, err := h.srvs.GetAllCategory(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (h *Handler) getCategory(ctx *gin.Context) {

	id := ctx.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 
	}
	
	category, err := h.srvs.GetCategory(ctx, id64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, category)
}
