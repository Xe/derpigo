package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/Xe/derpigo"
)

var (
	keyFile = flag.String(
		"keyFile",
		"/home/xena/.local/share/within/db.key",
		"file with the derpibooru key to use",
	)
)

func main() {
	flag.Parse()

	key, err := ioutil.ReadFile(*keyFile)
	if err != nil {
		log.Fatal(err)
	}

	c := derpigo.New(string(key))
	_ = c
}
