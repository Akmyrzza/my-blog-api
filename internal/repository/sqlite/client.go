package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Sqldb  *sql.DB
	dbname string
}

func New(dbname string) (*Sqlite, error) {
	s := new(Sqlite)

	sqldb, err := sql.Open("sqlite3", dbname)
	if err != nil {
		return nil, fmt.Errorf("sqlite open err: %w", err)
	}

	s.dbname = dbname
	s.Sqldb = sqldb

	_, err = s.Sqldb.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		firstname TEXT NOT NULL,
		lastname TEXT NOT NULL,
		password TEXT NOT NULL
	);`)
	if err != nil {
		return nil, fmt.Errorf("sqlite users table err: %w", err)
	}

	_, err = s.Sqldb.Exec(`CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);`)
	if err != nil {
		return nil, fmt.Errorf("sqlite categories table err: %w", err)
	}

	_, err = s.Sqldb.Exec(`CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		user_id INTEGER,
		categories INTEGER
	);`)
	if err != nil {
		return nil, fmt.Errorf("sqlite categories table err: %w", err)
	}

	return s, nil

}
