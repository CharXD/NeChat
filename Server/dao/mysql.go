package dao

import (
	"NeChat/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitSQL() {
	var err error
	DB, err = sql.Open("mysql", config.DataSourceName)
	if err != nil {
		fmt.Println("[ERROR]Try to connect failed,", err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("[ERROR]Connect failed", err)
	}
	DB.SetConnMaxLifetime(-1)
}
