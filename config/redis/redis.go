package redis

import (
	// "github.com/garyburd/redigo/redis"
	"context"
	redis "github.com/go-redis/redis/v8"
)
// REDIGO
// type RedisCli struct {
// 	conn redis.Conn
// }

// GO REDIS
type RedisCli struct {
	client *redis.Client
}

var instanceRedisCli *RedisCli = nil
var ctx = context.Background()

func Connect() (conn *RedisCli) {
	if instanceRedisCli == nil {
		instanceRedisCli = new(RedisCli)
		// REDIGO
		// var err error

		// instanceRedisCli.conn, err = redis.Dial("tcp", ":6379")

		// if err != nil {
		// 	panic(err)
		// }

		// if _, err := instanceRedisCli.conn.Do("AUTH", "kuncoro"); err != nil {
		// 	instanceRedisCli.conn.Close()
		// 	panic(err)
		// }

		// GO REDIS
		instanceRedisCli.client = redis.NewClient(&redis.Options{
			Addr: "redis:6379",
			Password: "",
			DB: 0,
		})
	}

	return instanceRedisCli
}
func (redisCli *RedisCli) SetValue(key string, value string, expiration ...interface{}) error {
	// REDIGO
	// _, err := redisCli.conn.Do("SET", key, value)

	// if err == nil && expiration != nil {
	// 	redisCli.conn.Do("EXPIRE", key, expiration[0])
	// }

	// GO REDIS
	err := redisCli.client.Set(ctx, key, value, 0).Err()
    if err != nil {
        panic(err)
    }

	return err
}

func (redisCli *RedisCli) GetValue(key string) (interface{}, error) {
	// REDIGO
	// return redisCli.conn.Do("GET", key)

	//GO REDIS
	return redisCli.client.Get(ctx, key).Result()
}
