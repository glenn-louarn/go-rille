package config

import (
	redis "github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
	"log"
)

var db redis.Conn

func DatabaseInit() {
	var err error

	db, err = redis.Dial("tcp", "93.23.6.57:6379")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Do("AUTH", "lesgogos")

	if err != nil {
		log.Fatal(err)
	}
}

// Getter for db var
func Db() redis.Conn {
	return db
}
