package db

import "database/sql"

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{db: db}
}

func (db *DB) Close() error {
	return db.db.Close()
}
