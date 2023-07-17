package repository

import (
	"context"

	"github.com/Akmyrzza/blog-api/internal/entity"
)

type Repository interface {
	//user
	CreateUser(ctx context.Context, u *entity.User) error
	Login(ctx context.Context, username, passwrod string) (*entity.User, error)
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	UpdateUser(ctx context.Context, id int64, u *entity.User) error
	DeleteUser(ctx context.Context, id int64) error

	//article
	CreateArticle(ctx context.Context, a *entity.Article) error
	UpdateArticle(ctx context.Context, id int64, a *entity.Article) error
	DeleteArticle(ctx context.Context, id int64) error
	GetArticle(ctx context.Context, id int64) (*entity.Article, error)
	GetAllArticle(ctx context.Context) ([]entity.Article, error)
	//GetArticlesByUserID(ctx context.Context, userID int64) ([]entity.Article, error)

	//category
	CreateCategory(ctx context.Context, c *entity.Category) error
	GetCategory(ctx context.Context) (entity.Category, error)
	GetAllCategory(ctx context.Context) ([]entity.Category, error)
}
