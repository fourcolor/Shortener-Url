package model

import (
	db "dcardHw/src/model/db"
	redis "dcardHw/src/model/redis"
	"time"
)

func GetOriUrl(shortenUrl, exp string) (url string) {
	url = redis.GetOriUrl(shortenUrl)
	if url == "" {
		url = db.GetOriUrl(shortenUrl, exp)
	}
	return
}

func SetShortenUrl(ori, shorten, exp string, t time.Duration) {
	redis.SetShortenUrl(ori, shorten, t)
	db.SetShortenUrl(ori, shorten, exp)
}

func GetShortenUrlbyExp(ori, exp string) (url string) {
	url = db.GetShortenUrlbyExp(ori, exp)
	return
}

func GetCounter() (val int64) {
	return redis.GetCounter()
}

func UpdateCounter() {
	redis.UpdateCounter()
}

func Init() (e error) {

	e = redis.Init()
	if e != nil {
		return
	}
	db.Init()
	return
}
