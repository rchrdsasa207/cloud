package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var createUserTable string = `CREATE TABLE IF NOT EXISTS "User" (
	"Id"	INTEGER,
	"Username"	TEXT,
	"Password"	TEXT,
	"Salt"	TEXT,
	PRIMARY KEY("Id" AUTOINCREMENT)
);`

type DB struct {
	db *sql.DB
}

func (db *DB) Close() {
	db.db.Close()
}

func OpenDB() *DB {
	db := &DB{}
	lightSql, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Fatal("Cannot open database: ", err)
	}
	db.db = lightSql

	_, err = db.db.Exec(createUserTable)
	if err != nil {
		log.Fatal("Could not create a user table: ", err)
	}

	return db
}
