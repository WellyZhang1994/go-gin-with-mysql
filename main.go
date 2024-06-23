package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/WellyZhang1994/go-gin-with-mysql/config"
	"github.com/WellyZhang1994/go-gin-with-mysql/db"
	"github.com/WellyZhang1994/go-gin-with-mysql/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
	
}
