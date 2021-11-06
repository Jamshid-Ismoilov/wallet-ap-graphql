package redis

import (
	"github.com/go-redis/redis"
)

var Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "myredispassword",
	DB:       0,
})