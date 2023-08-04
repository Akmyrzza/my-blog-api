package service

import (
	"github.com/Akmyrzza/blog-api/internal/config"
	"github.com/Akmyrzza/blog-api/internal/repository"
	"github.com/Akmyrzza/blog-api/pkg/jwttoken"
)

type Manager struct {
	Repository repository.Repository
	Token      *jwttoken.JWTToken
	Config     *config.Config
}

func New(repository repository.Repository, token *jwttoken.JWTToken, cfg *config.Config) *Manager {
	return &Manager{
		Repository: repository,
		Token:      token,
		Config:     cfg,
	}
}
