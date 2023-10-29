package models

import (
	"github.com/WellyZhang1994/go-gin-with-mysql/db"
)

type Health struct {
}

func (h Health) GetVersion() (string, error) {
	db := db.GetDB()
	var version string
    err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err2 != nil {
		panic(err2.Error())
	}
	return version, nil
}
