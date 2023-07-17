package service

import (
	"context"

	"github.com/Akmyrzza/blog-api/internal/entity"
)

func (m *Manager) CreateArticle(ctx context.Context, a *entity.Article) error {
	err := m.Repository.CreateArticle(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) UpdateArticle(ctx context.Context, id int64, a *entity.Article) error {
	err := m.Repository.UpdateArticle(ctx, id, a)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) DeleteArticle(ctx context.Context, id int64) error {
	err := m.Repository.DeleteArticle(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetArticle(ctx context.Context, id int64) (*entity.Article, error) {
	a, err := m.Repository.GetArticle(ctx, id)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (m *Manager) GetAllArticle(ctx context.Context) ([]entity.Article, error) {
	a, err := m.Repository.GetAllArticle(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}