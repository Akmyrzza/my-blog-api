package api

import "github.com/Akmyrzza/blog-api/internal/entity"

type RegisterRequest struct {
	entity.User
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
