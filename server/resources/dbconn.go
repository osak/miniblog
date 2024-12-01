package resources

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "miniblog_dev:dev@/miniblog?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}
	return db, nil
}
