package db

import (
	"github.com/gomodule/redigo/redis"
	"log"
)



var RedisConn redis.Conn

func RedisInit() (err error) {
	option := redis.DialPassword("123456")
	RedisConn, err = redis.Dial("tcp", "127.0.0.1:6379", option)
	if err != nil {
		log.Println("conn redis failed,", err)
		return
	}

	log.Printf("redis conn success")

	return nil
}
