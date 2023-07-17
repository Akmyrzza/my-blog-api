package sqlite

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Akmyrzza/blog-api/internal/entity"
)

func (s *Sqlite) CreateArticle(ctx context.Context, a *entity.Article) error {
	statement, err := s.Sqldb.Prepare("INSERT INTO articles (title, description, user_id, categories) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("sqlite insert into articles table err1: %w", err)
	}
	defer statement.Close()

	categories := `["`
	for i, v := range a.Categories {
		if i != 0 {
			categories = categories + `","`
		}
		categories = categories + strconv.Itoa(v)
	}
	categories = categories + `"]`
	fmt.Println(categories)
	_, err = statement.Exec(a.Title, a.Description, a.UserID, categories)
	if err != nil {
		return fmt.Errorf("sqlite insert into articles table err2: %w", err)
	}

	return nil
}

func (s *Sqlite) UpdateArticle(ctx context.Context, id int64, a *entity.Article) error {
	statement, err := s.Sqldb.Prepare("UPDATE articles SET title = ?, description = ?, user_id = ?, categories =? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("sqlite update articles table err: %w", err)
	}
	defer statement.Close()

	categories := `["`
	for i, v := range a.Categories {
		if i != 0 {
			categories = categories + `","`
		}
		categories = categories + strconv.Itoa(v)
	}
	categories = categories + `"]`
	fmt.Println(categories)

	_, err = statement.Exec(a.Title, a.Description, a.UserID, categories, id)
	if err != nil {
		return fmt.Errorf("sqlite update articles table err: %w", err)
	}

	return nil
}

func (s *Sqlite) DeleteArticle(ctx context.Context, id int64) error {
	statement, err := s.Sqldb.Prepare("DELETE FROM articles WHERE id = ?")
	if err != nil {
		return fmt.Errorf("sqlite delete from articles table err: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return fmt.Errorf("sqlite delete from articles table err: %w", err)
	}

	return nil
}

func (s *Sqlite) GetArticle(ctx context.Context, id int64) (*entity.Article, error) {
	statement, err := s.Sqldb.Prepare("SELECT id, title, description, user_id, categories FROM articles WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("sqlite select from articles table err: %w", err)
	}
	defer statement.Close()

	row, err := statement.Query(id)
	if err != nil {
		return nil, fmt.Errorf("sqlite select from articles table err: %w", err)
	}

	var article entity.Article
	for row.Next() {
		var categoriesJSON string
		err := row.Scan(&article.ID, &article.Title, &article.Description, &article.UserID, &categoriesJSON)
		if err != nil {
			return nil, err
		}

		var categories []string
		err = json.Unmarshal([]byte(categoriesJSON), &categories)
		if err != nil {
			return nil, err
		}

		for _, v := range categories {
			val, _ := strconv.Atoi(v)
			article.Categories = append(article.Categories, val)
		}
	}

	return &article, nil
}

func (s *Sqlite) GetAllArticle(ctx context.Context) ([]entity.Article, error) {
	statement, err := s.Sqldb.Prepare("SELECT id, title, description, user_id, categories FROM articles")
	if err != nil {
		return nil, fmt.Errorf("sqlite select from articles table err: %w", err)
	}
	defer statement.Close()

	row, err := statement.Query()
	if err != nil {
		return nil, fmt.Errorf("sqlite select from articles table err: %w", err)
	}

	var articles []entity.Article
	for row.Next() {
		var article entity.Article
		var categoriesJSON string

		err := row.Scan(&article.ID, &article.Title, &article.Description, &article.UserID, &categoriesJSON)
		if err != nil {
			return nil, err
		}

		var categories []string
		err = json.Unmarshal([]byte(categoriesJSON), &categories)
		if err != nil {
			return nil, err
		}

		for _, v := range categories {
			val, _ := strconv.Atoi(v)
			article.Categories = append(article.Categories, val)
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (s *Sqlite) GetAllArticleByUserID(ctx context.Context, userID int64) ([]entity.Article, error) {
	statement, err := s.Sqldb.Prepare("SELECT id, title, description, user_id, categories FROM articles WHERE user_id = ?")
	if err != nil {
		return nil, fmt.Errorf("sqlite select from articles table err: %w", err)
	}
	defer statement.Close()

	row, err := statement.Query(userID)
	if err != nil {
		return nil, fmt.Errorf("sqlite select from articles table err: %w", err)
	}

	var articles []entity.Article
	for row.Next() {
		var article entity.Article
		var categoriesJSON string

		err := row.Scan(&article.ID, &article.Title, &article.Description, &article.UserID, &categoriesJSON)
		if err != nil {
			return nil, err
		}

		var categories []string
		err = json.Unmarshal([]byte(categoriesJSON), &categories)
		if err != nil {
			return nil, err
		}

		for _, v := range categories {
			val, _ := strconv.Atoi(v)
			article.Categories = append(article.Categories, val)
		}
		
		articles = append(articles, article)
	}

	return articles, nil
}