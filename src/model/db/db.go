// package main

package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type url struct {
	OriUrl     string `db:"oriUrl"`
	ShortenUrl string `db:"shortenUrl"`
	ExpireAt   string `db:"expireAt"`
}

var (
	UserName string = "dcard"
	Password string = "DcardPass"
	Addr     string = "127.0.0.1"
	Port     int    = 3306
	Database string = "dcard"
)

func Init() {
	if os.Getenv("DB_USERNAME") != "" {
		UserName = os.Getenv("DB_USERNAME")
	}
	if os.Getenv("DcardPass") != "" {
		Password = os.Getenv("DcardPass")
	}
	if os.Getenv("DB_HOST") != "" {
		Addr = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_DATABASE") != "" {
		Database = os.Getenv("DB_DATABASE")
	}
	if os.Getenv("DB_USERNAME") != "" {
		UserName = os.Getenv("DB_USERNAME")
	}
	if _, e := strconv.Atoi(os.Getenv("DB_PORT")); e == nil {
		Port, e = strconv.Atoi(os.Getenv("DB_PORT"))
	}
	client := Connect()
	client.Close()
	return
}

func Connect() *sql.DB {
	fmt.Println(Password)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	return db
}

func SetShortenUrl(oriUrl, shortenUrl, expireAt string) {
	db := Connect()
	defer db.Close()
	_, err := db.Exec("insert into url(oriUrl,shortenUrl,expirtedAt) values(?,?,?)", oriUrl, shortenUrl, expireAt)
	if err != nil {
		panic(err)
	}
}

func GetOriUrl(shortenUrl, expireAt string) (oriUrl string) {
	db := Connect()
	defer db.Close()
	err := db.QueryRow("select oriUrl from dcard.url where oriUrl = ? and expirtedAt > ?", shortenUrl, expireAt).Scan(&oriUrl)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return
}

func GetShortenUrlbyExp(oriUrl string, expireAt string) (shortenUrl string) {
	db := Connect()
	defer db.Close()
	err := db.QueryRow("select shortenUrl from dcard.url where oriUrl = ? and expirtedAt = ?", oriUrl, expireAt).Scan(&shortenUrl)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return
}

func OriExited(oriUrl string) bool {
	db := Connect()
	defer db.Close()
	var res url
	err := db.QueryRow("SELECT * FROM url WHERE oriUrl = ?", oriUrl).Scan(&res.OriUrl, &res.ShortenUrl, &res.ExpireAt)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return res.OriUrl != ""
}
