package service

import (
	"github.com/Akmyrzza/blog-api/internal/config"
	"github.com/Akmyrzza/blog-api/internal/repository"
)

type Manager struct {
	Repository repository.Repository
	Config *config.Config
}

func New(repository repository.Repository, cfg *config.Config) *Manager {
	return &Manager{
		Repository: repository,
		Config: cfg,
	}
}