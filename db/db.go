package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

const (
	HOST = "database"
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	createTablesIfNotExists(db)
	return db, nil
}

func createTablesIfNotExists(db Database) error {
	log.Printf("Apply create tables migration")
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "000001_create_tables.up.sql")

	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Printf("Error while reading migration file: %s", ioErr)
		return ioErr
	}
	sql := string(c)
	_, err := db.Conn.Exec(sql)
	if err != nil {
		log.Printf("Error while executing migration: %s", err)
		return err
	}
	return nil
}
