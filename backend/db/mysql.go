package db

import (
	"ecomm/config"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func ConnectMysql() error {
	var err error
	dsn := config.Config.GetString("mysql.dsn")
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return errors.New("connect DB failed err: " + err.Error())
	}
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	log.Println("connect mysql success ...")
	return nil
}
