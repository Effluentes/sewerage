package postgresql

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var (
	pgInstance *sql.DB
	pgOnce     sync.Once
	pgErr      error
)

func GetPostgresDB(connStr string) (*sql.DB, error) {
	pgOnce.Do(func() {
		pgInstance, pgErr = sql.Open("postgres", connStr)
		if pgErr != nil {
			pgErr = fmt.Errorf("postgres init failed: %v", pgErr)
			return
		}
		if pgErr = pgInstance.Ping(); pgErr != nil {
			pgErr = fmt.Errorf("postgres ping failed: %v", pgErr)
		}
	})
	return pgInstance, pgErr
}