package data

import (
	"context"

	// Mysql driver

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

// InitDatabase initializes a database pool from models.Config
func (d *Data) initDatabase() error {
	var err error
	d.database, err = pgxpool.New(context.Background(), d.Config.DBconnStr())
	if err != nil {
		return err
	}

	err = d.database.Ping(context.Background())
	if err != nil {
		return err
	}
	d.Queries = db.New(d.database)
	return nil
}
