package db

import (
	"database/sql"
	"ecomm/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectMysql() *sql.DB {
	dsn := config.Config.GetString("mysql.dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil || db.Ping() != nil {
		fmt.Println(fmt.Errorf("connect to mysql have error %s", err.Error()))
		return nil
	}
	return db
}
