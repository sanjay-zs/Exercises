package driver

import (
	"database/sql"
	"go-sql-driver/mysql"
)

// ConnectToDB function To establish database connection
func ConnectToDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 "sanjay",
		Passwd:               "12345678",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "test",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
