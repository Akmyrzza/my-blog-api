package service

import (
	"context"

	"github.com/Akmyrzza/blog-api/internal/entity"
	"github.com/Akmyrzza/blog-api/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = m.Repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) Login(ctx context.Context, u *entity.Login) (*entity.User, error) {
	user, err := m.Repository.Login(ctx, u.Username, u.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	user, err := m.Repository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) UpdateUser(ctx context.Context, id int64, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	
	err = m.Repository.UpdateUser(ctx, id, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) DeleteUser(ctx context.Context, id int64) error {
	err := m.Repository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
