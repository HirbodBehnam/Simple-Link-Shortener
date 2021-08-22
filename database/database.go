package database

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const databaseTimeout = time.Second

var MainDatabase Database

type Database struct {
	db *sql.DB
}

// New creates a new database connection
func New(dbUrl string) Database {
	var db Database
	var err error
	db.db, err = sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	db.db.SetConnMaxLifetime(time.Minute * 3)
	db.db.SetMaxOpenConns(50)
	db.db.SetMaxIdleConns(5)
	err = db.db.Ping() // test the connection
	if err != nil {
		log.Fatal("Cannot ping the database:", err)
	}
	return db
}

// Close closes database connection
func (d *Database) Close() {
	_ = d.db.Close()
}

// AddLink adds a link to database
func (d *Database) AddLink(ctx context.Context, key, link string) error {
	ctx, cancel := context.WithTimeout(ctx, databaseTimeout)
	_, err := d.db.ExecContext(ctx, "INSERT INTO `links` (`key`, `link`) VALUES (?,?)", key, link)
	cancel()
	return err
}

// GetLink gets a link from database
func (d *Database) GetLink(ctx context.Context, key string) (link string, err error) {
	ctx, cancel := context.WithTimeout(ctx, databaseTimeout)
	err = d.db.QueryRowContext(ctx, "SELECT `link` FROM `links` WHERE `key`=?", key).Scan(&link)
	cancel()
	return
}

// DeleteLink deletes a link in database
func (d *Database) DeleteLink(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, databaseTimeout)
	_, err := d.db.ExecContext(ctx, "DELETE FROM `links` WHERE `key`=?", key)
	cancel()
	return err
}
