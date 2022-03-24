package postgres

import (
	"database/sql"
	"net/url"
	"strconv"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/fileserver"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/sql/postgres/query"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
)

// Server implements a PostgreSQL driver.
type Server struct {
	sqlDB   *sql.DB
	queries *query.Queries
	idNode  *snowflake.Node
	fserver foodtinder.FileServer
}

type Opts struct {
	// SnowflakeNode is the node number to use for snowflake generation. Default
	// 0.
	SnowflakeNode int64
}

// New creates a new PostgreSQL server backend.
func New(url *url.URL, fserver foodtinder.FileServer) (*Server, error) {
	if fserver == nil {
		fserver = fileserver.InMemory(nil)
	}

	db, err := sql.Open("postgres", url.String())
	if err != nil {
		return nil, err
	}

	var idNode int64
	if node := url.Query().Get("id_node"); node != "" {
		n, err := strconv.ParseInt(node, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "invalid ?id_node")
		}
		idNode = n
	}

	n, err := snowflake.NewNode(idNode)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing snowflake node")
	}

	if err := up(db); err != nil {
		return nil, errors.Wrap(err, "error doing migrations up")
	}

	return &Server{
		fserver: fserver,
		sqlDB:   db,
		queries: query.New(db),
		idNode:  n,
	}, nil
}

func (s *Server) FileServer() foodtinder.FileServer {
	return s.fserver
}

func (s *Server) LoginServer() foodtinder.LoginServer {
	return (*LoginServer)(s)
}

func (s *Server) AuthorizerServer() foodtinder.AuthorizerServer {
	return (*AuthorizerServer)(s)
}

func (s *Server) AuthorizedServer(session *foodtinder.Session) foodtinder.AuthorizedServer {
	return &AuthorizedServer{
		Server:  s,
		session: session,
	}
}

func (s *Server) Close() error { return s.sqlDB.Close() }
