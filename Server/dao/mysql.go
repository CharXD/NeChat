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
	dataSourceName := config.ServerConfig.SQL.User + ":" + config.ServerConfig.SQL.Pass + "@tcp(" + config.ServerConfig.SQL.Host + ":" + config.ServerConfig.SQL.Port + ")/" + config.ServerConfig.SQL.Database + "?charset=utf8mb4&parseTime=True"
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("[ERROR]Try to connect failed,", err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("[ERROR]Connect failed", err)
	}
	DB.SetConnMaxLifetime(-1)
}
