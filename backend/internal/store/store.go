package store

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/inmemory"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/sql/postgres"
)

type Server interface {
	foodtinder.Server
	io.Closer
}

// Opts contains optional options for the store backend.
type Opts struct {
	// FileServer is the file server to use. If nil, file.InMemory will be used.
	FileServer foodtinder.FileServer
}

// Open opens a new database connection.
func Open(uri string, opts Opts) (Server, error) {
	url, err := url.Parse(uri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse URI")
	}

	switch url.Scheme {
	case "mock":
		var state inmemory.State
		url.Path = strings.TrimPrefix(url.String(), "mock://")
		url.Host = ""

		if url.Path != "" {
			// Hack because url.Parse confuses some of the path as the domain.
			b, err := os.ReadFile(url.Path)
			if err != nil {
				return nil, fmt.Errorf("cannot read mock database at %s: %w", url.Path, err)
			}

			if err := json.Unmarshal(b, &state); err != nil {
				return nil, errors.Wrap(err, "cannot parse mock database JSON")
			}
		}

		s := inmemory.NewServer(state, opts.FileServer)
		return NopCloser(s), nil

	case "postgres":
		return postgres.New(url, opts.FileServer)

	default:
		return nil, fmt.Errorf("unknown SQL scheme %q", url.Scheme)
	}
}

type nopCloser struct{ foodtinder.Server }

// NopCloser wraps a foodtinder.Server with a Close method that always returns
// nil.
func NopCloser(s foodtinder.Server) Server {
	return nopCloser{s}
}

func (n nopCloser) Close() error { return nil }
