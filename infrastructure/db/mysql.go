package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ConfigDb is the configuration for the database.


type Mysql struct {
}

func (s Mysql) TransactionDb() (*sql.Tx, error) {
	db, err := s.ConnectionDb()
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s Mysql) ConnectionDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("STRING_CONNECTION"))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 1)

	return db, nil
}

func (s Mysql) CheckDb() bool {
	db, err := s.ConnectionDb()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	return true
}