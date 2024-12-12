package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Ayush330/server/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitalizeSql() {
	var err error
	fmt.Println(config.GetSqlAddress())
	db, err = sql.Open("mysql", config.GetSqlAddress())
	if err != nil {
		fmt.Println(err)
		log.Panic("Cannot open sql pool")
	}
	db.SetMaxOpenConns(config.MAXM_CONNECTIONS)
	err = db.Ping()
	if err != nil {
		log.Panicf("Cannot connect to database: %v", err) // Log the error if connection fails
	}

	log.Println("Database connection established successfully!")
}
