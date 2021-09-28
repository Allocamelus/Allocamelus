package data

import (
	"database/sql"
	"errors"
	"sync"
	"time"

	// Mysql driver

	"github.com/allocamelus/allocamelus/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	"k8s.io/klog/v2"
)

// Prepare defines a function to prepare a sql statement
type Prepare = func(sql string) *sql.Stmt

type PrepareCallback = func(Prepare)

// PrepareQueue adds easy tools to queue sql for prepare in
// the init() function or any time Before the database is ready
type PrepareQueue struct {
	mu       sync.Mutex
	Prepared bool
	items    map[int]PrepareCallback
}

// NewPrepareQueue
func NewPrepareQueue() *PrepareQueue {
	p := new(PrepareQueue)
	p.items = map[int]PrepareCallback{}
	return p
}

// Add stmt pointer and query to queue
//
// Must not be called after prepareAll
func (p *PrepareQueue) Add(stmt **sql.Stmt, q string) {
	p.mu.Lock()
	if p.Prepared {
		logger.Fatal(errors.New("error: PrepareQueue.Add called after prepareAll"))
	}
	p.items[len(p.items)] = func(p Prepare) {
		*stmt = p(q)
	}
	p.mu.Unlock()
}

// prepareAll prepares all queries in queue
func (p *PrepareQueue) prepareAll(d *Data) {
	p.mu.Lock()
	p.Prepared = true
	// New WaitGroup
	var wg sync.WaitGroup

	for _, callback := range p.items {
		// Add one to WaitGroup
		wg.Add(1)

		// New goroutine for each queue callback
		go func(cb PrepareCallback) {
			cb(d.Prepare)
			wg.Done()
		}(callback)
	}

	// Wait for all routines to be done
	wg.Wait()
	p.mu.Unlock()
}

// PrepareQueuer global to allow queuing
var PrepareQueuer = NewPrepareQueue()

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
	d.database.SetConnMaxLifetime(5 * time.Second)

	// Open doesn't open a connection. Validate DSN data:
	err = d.database.Ping()
	if err != nil {
		return err
	}
	// PrepareAll
	PrepareQueuer.prepareAll(d)
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
