package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"sync"
)

type Connection struct {
	Conn *sql.DB
}

var instance *Connection
var once sync.Once

func ConnectDB() (*Connection, error) {
	var (
		host     = os.Getenv("DB_HOST")
		port, _  = strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname = %s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return nil, err
	}
	once.Do(func() {
		instance = &Connection{Conn: db}
	})
	return instance, nil
}
