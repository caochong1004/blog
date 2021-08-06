package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DBWeb *sqlx.DB
func InitWeb(dsn string) (err error) {
	DBWeb, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Printf("数据库连接失败, err:%s\n", err)
		return
	}

	err = DBWeb.Ping()
	if err != nil {
		return err
	}
	DBWeb.SetMaxIdleConns(100)
	DBWeb.SetConnMaxIdleTime(16)
	return nil
}
