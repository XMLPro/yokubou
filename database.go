package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/crypto/bcrypt"
	"log"
)

type DB struct {
	db *sql.DB
}

func NewDB(dbName string) *DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}

func (db *DB) CreateUser(name, password string) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	string(hash)
}
