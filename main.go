package main

import (
	"flag"
	"fmt"

	"github.com/codescot/hue"
)

func main() {
	var (
		username string
		address  string
	)

	flag.StringVar(&username, "username", "", "-username=a1b2c3d4e")
	flag.StringVar(&address, "ip", "192.168.0.20", "-ip=127.0.0.1")
	flag.Parse()

	h := hue.Hue{
		AppName:  "huely",
		Address:  address,
		Username: username,
	}

	if h.Username == "" {
		h.Authenticate()
	}

	fmt.Println(h)
}
