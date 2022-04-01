package services

import (
	model "dcardHw/src/model"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net/url"
	"time"
)

func GenerateShortenUrl(ori string, expireAt time.Time) (status int, shortenUrl string) {
	_, err := url.ParseRequestURI(ori)
	status = 0
	if err != nil {
		status = 1
		return
	}
	t := expireAt.Format("2006-01-02 15:04:05")
	shortenUrl = model.GetShortenUrlbyExp(ori, t)
	fmt.Println(expireAt)
	if shortenUrl == "" {
		id := model.GetCounter()
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(id))
		shortenUrl = base64.StdEncoding.EncodeToString(b)
		dt := time.Duration(expireAt.Sub(time.Now()))
		model.SetShortenUrl(ori, shortenUrl, t, dt)
		model.UpdateCounter()
	}
	return

}

func RedirectUrl(shortenUrl string) (s int, ori string) {
	ori = model.GetOriUrl(shortenUrl, time.Now().Format("2006-01-02 15:04:05"))
	s = 0
	if ori == "" {
		s = 1
	}
	return
}
