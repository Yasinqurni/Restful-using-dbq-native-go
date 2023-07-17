package config

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DbInit(env *Config) *sql.DB {

	dbHost := env.DbHost
	dbPort := env.DbPort
	dbUser := env.DbUsername
	dbPass := env.DbPassword
	dbName := env.DbDatabase

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	db, err := sql.Open(`mysql`, dsn)

	if err != nil {
		return nil
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
