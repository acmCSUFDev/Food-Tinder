package sql

//go:generate sqlc generate

import (
	"database/sql"
	"fmt"
	"io"
	"net/url"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/fs"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/sql/postgres"
)

type Server interface {
	foodtinder.Server
	io.Closer
}

// Open opens a new SQL connection.
func Open(url *url.URL, fs fs.Assets) (Server, error) {
	sqlDB, err := sql.Open(url.Scheme, url.String())
	if err != nil {
		return nil, err
	}

	var s Server

	switch url.Scheme {
	case "postgres":
		s, err = postgres.New(sqlDB, fs)
	default:
		sqlDB.Close()
		return nil, fmt.Errorf("unknown SQL scheme %q", url.Scheme)
	}

	return s, err
}
