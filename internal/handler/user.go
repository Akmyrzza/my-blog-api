package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Akmyrzza/blog-api/internal/entity"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *Handler) createUser(ctx *gin.Context) {
	var req entity.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User successfully created"})
}

func (h *Handler) login(ctx *gin.Context) {
	var req entity.Login

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	user, err := h.srvs.Login(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) getUser(ctx *gin.Context) {

	id := ctx.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	user, err := h.srvs.GetUser(ctx, id64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	var req entity.User

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.UpdateUser(ctx, id64, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User successfully updated"})
}

func (h *Handler) deleteUser(ctx *gin.Context) {

	id := ctx.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.DeleteUser(ctx, id64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
