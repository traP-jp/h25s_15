package coredb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func New() (*DB, error) {
	database, ok := os.LookupEnv("NS_MARIADB_DATABASE")
	if !ok {
		return nil, errors.New("NS_MARIADB_DATABASE environment variable not set")
	}

	host, ok := os.LookupEnv("NS_MARIADB_HOSTNAME")
	if !ok {
		return nil, errors.New("NS_MARIADB_HOSTNAME environment variable not set")
	}

	port, ok := os.LookupEnv("NS_MARIADB_PORT")
	if !ok {
		return nil, errors.New("NS_MARIADB_PORT environment variable not set")
	}

	user, ok := os.LookupEnv("NS_MARIADB_USER")
	if !ok {
		return nil, errors.New("NS_MARIADB_USER environment variable not set")
	}
	password, ok := os.LookupEnv("NS_MARIADB_PASSWORD")
	if !ok {
		return nil, errors.New("NS_MARIADB_PASSWORD environment variable not set")
	}

	conf := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 host + ":" + port,
		DBName:               database,
		AllowNativePasswords: true,
		Loc:                  time.Local,
		ParseTime:            true,
	}
	dsn := conf.FormatDSN()

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	return &DB{
		db: db,
	}, nil
}

type IDB interface {
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type dbKey struct{}

var dbKeyInstance = &dbKey{}

func (d *DB) DB(ctx context.Context) IDB {
	if db, ok := ctx.Value(dbKeyInstance).(IDB); ok {
		return db
	}

	return d.db
}

func (d *DB) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}
	}()

	ctx = context.WithValue(ctx, dbKeyInstance, tx)

	if err := fn(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
