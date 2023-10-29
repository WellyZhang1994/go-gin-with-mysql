package db

import (
	"github.com/WellyZhang1994/go-gin-with-mysql/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() {
	c := config.GetConfig()
	user := c.GetString("db.user")
	password := c.GetString("db.password")
	host := c.GetString("db.host")
	port := c.GetString("db.port")
	database := c.GetString("db.database")
	connectionDNS := user + ":" + password + "@tcp(" + host + ":" + port +")/" + database 
	dbConnect, err := sql.Open("mysql", connectionDNS)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	dbConnect.SetMaxOpenConns(10)
	dbConnect.SetMaxIdleConns(10)
	db = dbConnect
}

func GetDB() *sql.DB {
	return db
}
