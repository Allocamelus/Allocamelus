package data

import (
	"database/sql"

	// Mysql driver
	_ "github.com/go-sql-driver/mysql"
	"k8s.io/klog/v2"
)

// Prepare defines a function to prepare a sql statement
type Prepare = func(sql string) *sql.Stmt

// InitDatabase initializes a database pool from models.Config
func (d *Data) initDatabase() error {
	var err error
	dataSource := d.Config.Db.UserName + ":" + d.Config.Db.Password + "@" + d.Config.Db.Net + "(" + d.Config.Db.Server + ")/" + d.Config.Db.Name
	d.database, err = sql.Open("mysql", dataSource)
	if err != nil {
		return err
	}

	d.database.SetMaxOpenConns(100)
	d.database.SetMaxIdleConns(10)
	d.database.SetConnMaxLifetime(10)

	// Open doesn't open a connection. Validate DSN data:
	err = d.database.Ping()
	if err != nil {
		return err
	}
	return nil
}

// Select is a query wrapper
func (d *Data) Select(sql string, args ...interface{}) (*sql.Rows, error) {
	if len(args) != 0 {
		return d.database.Query(sql, args...)
	}
	return d.database.Query(sql)
}

// SelectRow is a queryRow wrapper
func (d *Data) SelectRow(sql string, args ...interface{}) *sql.Row {
	return d.database.QueryRow(sql, args...)
}

// Exec is a Exec wrapper
func (d *Data) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return d.database.Exec(sql, args...)
}

// Prepare prepares a query and returns the stmt
// Logs own errors
func (d *Data) Prepare(sql string) *sql.Stmt {
	stmt, err := d.database.Prepare(sql)
	if err != nil {
		klog.Fatal(err)
	}
	return stmt
}
