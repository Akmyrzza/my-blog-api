package sqlite

import (
	"context"
	"fmt"

	"github.com/Akmyrzza/blog-api/internal/entity"
	"github.com/Akmyrzza/blog-api/pkg/util"
)

func (s *Sqlite) CreateUser(ctx context.Context, u *entity.User) error {
	statement, err := s.Sqldb.Prepare("INSERT INTO users (username, firstname, lastname, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("sqlite insert into users table err: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(u.Username, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return fmt.Errorf("sqlite insert into users table err: %w", err)
	}

	return nil
}

func (s *Sqlite) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	statement, err := s.Sqldb.Prepare("SELECT id, username, firstname, lastname, password FROM users WHERE username = ?")
	if err != nil {
		return nil, fmt.Errorf("sqlite select from users table err: %w", err)
	}
	defer statement.Close()

	row, err := statement.Query(username)
	if err != nil {
		return nil, fmt.Errorf("sqlite select from users table err: %w", err)
	}

	var user entity.User
	for row.Next() {
		err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Password)
		if err != nil {
			return nil, err
		}

		if err := util.ComparePassword([]byte(user.Password), password); err != nil {
			return nil, err
		}

		user.Password = password
	}

	return &user, nil
}

func (s *Sqlite) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	statement, err := s.Sqldb.Prepare("SELECT id, username, firstname, lastname, password FROM users WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("sqlite select from users table err: %w", err)
	}
	defer statement.Close()

	row, err := statement.Query(id)
	if err != nil {
		return nil, fmt.Errorf("sqlite select from users table err: %w", err)
	}

	var user entity.User
	for row.Next() {
		err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Password)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (s *Sqlite) UpdateUser(ctx context.Context, id int64, u *entity.User) error {
	statement, err := s.Sqldb.Prepare("UPDATE users SET username = ?, firstname = ?, lastname = ?, password =? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("sqlite update users table err: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(u.Username, u.FirstName, u.LastName, u.Password, id)
	if err != nil {
		return fmt.Errorf("sqlite update users table err: %w", err)
	}

	return nil
}

func (s *Sqlite) DeleteUser(ctx context.Context, id int64) error {
	statement, err := s.Sqldb.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return fmt.Errorf("sqlite delete from users table err: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return fmt.Errorf("sqlite delete from users table err: %w", err)
	}

	return nil
}
