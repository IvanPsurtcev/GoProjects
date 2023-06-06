package storage

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"metamask-auth/config"
	appCloser "metamask-auth/pkg/app_closer"
)

func InitDBConnection(ctx context.Context, c *config.Config, closer appCloser.Closer) (*sqlx.DB, error) {
	connectionUrl := c.GetDsn()
	db, err := sqlx.ConnectContext(ctx, "pgx", connectionUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	closer.AddCloser(func() {
		db.Close()
	}, "pgConnection")
	return db, nil
}
