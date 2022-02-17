package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

type dbSQLite struct {
	Conn *sql.DB
	Err  error
}

func (_d *dbSQLite) Init() {
	_d.Conn, _d.Err = _d.Connect()
}

func (_d *dbSQLite) Connect() (*sql.DB, error) {
	cwd, _ := os.Getwd()
	dbName := os.Getenv("SQLITE_DBNAME")
	dbFilePath := path.Dir(cwd) + "/database/" + dbName

	// Check Path File
	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		fmt.Println("Create new database" + dbName)

		file, err := os.Create(dbFilePath) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}

		file.Close()

		fmt.Println("Database " + dbName + " created")
	}

	SQLiteDB, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// defer SQLiteDB.Close() // Defer Closing the database

	// Creating Tables
	_d.createTables(SQLiteDB)

	return SQLiteDB, nil
}

func (_d *dbSQLite) createTables(db *sql.DB) {
	// Users Table
	query := `
		CREATE TABLE IF NOT EXISTS "users" (
			"id"	INTEGER,
			"name"	TEXT NOT NULL UNIQUE,
			"phone"	TEXT NOT NULL,
			"password"	TEXT NOT NULL UNIQUE,
			"role"	TEXT NOT NULL,
			"created_at" Timestamp With Time Zone NOT NULL,
			PRIMARY KEY("id" AUTOINCREMENT)
		);
	`

	stmt, err := db.Prepare(query) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec() // Execute SQL Statements
}

var SQLite = &dbSQLite{}
