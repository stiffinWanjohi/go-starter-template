package db

import (
	"database/sql"
	"fmt"
	"os"

	// postgresql driver
	_ "github.com/lib/pq"
)

type SQLOperations interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type DBManager struct {
	*sql.DB
}

func (db *DBManager) InTransaction(
	operations func(SQLOperations) error,
) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err = operations(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}

		return err
	}

	return tx.Commit()
}

func NewPostgresDBManager() *DBManager {
	return NewPostgresDBManagerWithParameters(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
}

func NewPostgresDBManagerWithParameters(
	username,
	password,
	host,
	databaseName string,
) *DBManager {
	return newDBManagerWithParameters("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=disable",
		username, password, databaseName, host,
	))
}

func NewPostgresDBManagerWithURL(databaseURL string) *DBManager {
	return newDBManagerWithParameters("postgres", databaseURL)
}

func NewPostgresDBManagerWithUrl(databaseUrl string) *DBManager {
	return newDBManagerWithParameters("postgres", databaseUrl)
}

func newDBManagerWithParameters(driverName string, databaseUrl string) *DBManager {

	if databaseUrl == "" {
		panic("database url is required")
	}

	db, err := sql.Open(driverName, databaseUrl)
	if err != nil {
		panic(fmt.Sprintf("sql.Open failed because err=[%v]", err))
	}

	return &DBManager{db}
}
