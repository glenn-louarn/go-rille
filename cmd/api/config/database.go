package config

import (
	redis "github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var pool *redis.Pool

var IP_DATABASE = "93.23.6.57:6379"
var PW = "lesgogos"

func DatabaseInit() {
	var err error

	pool = &redis.Pool{
		MaxIdle: 100000,
		MaxActive: 1000000,
		Dial: func() (redis.Conn, error) {
			db, err := redis.Dial("tcp", IP_DATABASE)
			if err != nil {
				log.Printf("ERROR: fail init redis: %s", err.Error())
				os.Exit(1)
			}
			_, err = db.Do("AUTH", PW)
			return db, err
		},
	}

	if err != nil {
		log.Fatal(err)
	}
}

// Getter for db var
func Db()  redis.Conn{
	conn := pool.Get()
	return conn
}
