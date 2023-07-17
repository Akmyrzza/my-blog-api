package service

import (
	"context"

	"github.com/Akmyrzza/blog-api/internal/entity"
)

func (m *Manager) CreateCategory(ctx context.Context, c *entity.Category) error {

	err := m.Repository.CreateCategory(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetAllCategory(ctx context.Context) ([]entity.Category, error) {
	categories, err := m.Repository.GetAllCategory(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (m *Manager) GetCategory(ctx context.Context, id int64) (*entity.Category, error) {
	user, err := m.Repository.GetCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}