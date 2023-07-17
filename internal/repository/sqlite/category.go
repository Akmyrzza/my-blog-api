package sqlite

import (
	"context"
	"fmt"

	"github.com/Akmyrzza/blog-api/internal/entity"
)

func (s *Sqlite) CreateCategory(ctx context.Context, c *entity.Category) error {
	statement, err := s.Sqldb.Prepare("INSERT INTO categories (name) VALUES (?)")
	if err != nil {
		return fmt.Errorf("sqlite insert into categories table err: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(c.Name)
	if err != nil {
		return fmt.Errorf("sqlite insert into categories table err: %w", err)
	}

	return nil
}

func (s *Sqlite) GetCategory(ctx context.Context, id int64) (*entity.Category, error) {
	statement, err := s.Sqldb.Prepare("SELECT id, name FROM categories WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("sqlite select from categories table err: %w", err)
	}
	defer statement.Close()

	row, err := statement.Query(id)
	if err != nil {
		return nil, fmt.Errorf("sqlite select from categories table err: %w", err)
	}

	var category entity.Category
	for row.Next() {
		err := row.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
	}

	return &category, nil
}

func (s *Sqlite) GetAllCategory(ctx context.Context) ([]entity.Category, error) {
	statement, err := s.Sqldb.Prepare("SELECT id, name FROM categories")
	if err != nil {
		return nil, fmt.Errorf("sqlite select from categories table err: %w", err)
	}
	defer statement.Close()

	row, err := statement.Query()
	if err != nil {
		return nil, fmt.Errorf("sqlite select from categories table err: %w", err)
	}

	var categories []entity.Category
	for row.Next() {
		var category entity.Category
		err := row.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
