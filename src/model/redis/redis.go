package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func Init() error {
	client := Connect()
	defer client.Close()
	tv, _ := client.Get("total").Result()
	var val int64
	_, err := fmt.Sscan(tv, &val)
	if err != nil {
		return err
	}
	e := client.Set("total", 1, 0).Err()
	if e != nil {
		return err
	}
	return nil
}

func OriExited(ori string) bool {
	client := Connect()
	defer client.Close()
	_, err := client.SIsMember("ori", ori).Result()
	if err != nil {
		panic(err)
	}
	return err != redis.Nil
}

func GetOriUrl(shortenUrl string) string {
	client := Connect()
	defer client.Close()
	val, err := client.Get(shortenUrl).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	if err == redis.Nil {
		return ""
	}
	return val
}

func SetShortenUrl(ori, shorten string, t time.Duration) {
	client := Connect()
	defer client.Close()
	err := client.Set(shorten, ori, t).Err()
	if err != nil {
		panic(err)
	}
}

func GetCounter() (val int64) {
	client := Connect()
	defer client.Close()
	tv, _ := client.Get("total").Result()
	fmt.Sscan(tv, &val)
	return
}

func UpdateCounter() {
	client := Connect()
	client.Incr("total").Val()
}
