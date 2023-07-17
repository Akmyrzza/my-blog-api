package pgrepo

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
)

const usersTable = "users"

type Postgres struct {
	host     string
	port     string
	dbName   string
	username string
	password string
	Pool     *pgxpool.Pool
}

func New(opts ...Option) (*Postgres, error) {
	p := new(Postgres)

	for _, opt := range opts {
		opt(p)
	}

	q := url.Values{}
	q.Add("sslmode", "disable")

	u := url.URL{
		Scheme:   "postgresql",
		User:     url.UserPassword(p.username, p.password),
		Host:     fmt.Sprintf("%s:%s", p.host, p.port),
		Path:     p.dbName,
		RawQuery: q.Encode(),
	}

	poolConfig, err := pgxpool.ParseConfig(u.String())
	if err != nil {
		return nil, fmt.Errorf("pgxpool parse config err: %w", err)
	}

	p.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect err: %w", err)
	}

	return p, nil
}
