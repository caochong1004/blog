package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB
func Init(dsn string) (err error) {
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Printf("数据库连接失败, err:%s\n", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		return err
	}
	DB.SetMaxIdleConns(100)
	DB.SetConnMaxIdleTime(16)
	return nil
}
