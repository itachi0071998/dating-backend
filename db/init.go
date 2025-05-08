package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	log.Println("Db connection started with db url as ", os.Getenv("DB_URL"))  // Log every health check request
	dsn := os.Getenv("DB_URL")
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return err
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
		return err
	}
	fmt.Println("Connected to the SQL Server!")
	return nil
}