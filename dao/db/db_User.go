package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DBUser *sqlx.DB
func InitUser(dsn string) (err error) {
	DBUser, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Printf("数据库连接失败, err:%s\n", err)
		return
	}

	err = DBUser.Ping()
	if err != nil {
		return err
	}
	DBUser.SetMaxIdleConns(100)
	DBUser.SetConnMaxIdleTime(16)
	return nil
}
