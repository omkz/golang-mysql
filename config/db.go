package config

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_blog")
	if err != nil {
		return nil, err
	}
	return db, nil
}

