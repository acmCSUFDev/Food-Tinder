package postgres

import (
	"database/sql"
	"embed"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

//go:embed query/*.sql
var schemas embed.FS

func init() {
	goose.SetBaseFS(schemas)
}

func up(db *sql.DB) error {
	return goose.Up(db, "query")
}

// errorIsUniqueViolation returns true if the error returned from the PostgreSQL
// query is one that is a Unique Violation error.
func errorIsUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// https://www.postgresql.org/docs/9.4/errcodes-appendix.html
		return pgErr.Code == pgerrcode.UniqueViolation
	}
	return false
}

func nullstr(str string) sql.NullString {
	if str == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: str,
		Valid:  true,
	}
}
