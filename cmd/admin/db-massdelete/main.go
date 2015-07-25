/*
Command db-massdelete handles mass deletion of derpibooru images.

Please use this sparingly.

Usage:

	Usage of ./db-massdelete:
	  -keyFile="/home/xena/.local/share/within/db.cadance.key": file with the derpibooru key to use
	  -reason="": reason to use when deleting images

Then give it the image ID's you want to delete.

	./db-massdelete -reason "OP is a duck" 123 325 1561 136324
*/
package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/Xe/derpigo"
)

var (
	keyFile = flag.String("keyFile", "/home/xena/.local/share/within/db.cadance.key", "file with the derpibooru key to use")
	reason  = flag.String("reason", "", "reason to use when deleting images")
)

func main() {
	flag.Parse()

	if *reason == "" {
		log.Fatal("Need a reason")
	}

	key, err := ioutil.ReadFile(*keyFile)
	if err != nil {
		log.Fatal(err)
	}

	c := derpigo.New(string(key))

	for _, i := range flag.Args() {
		err := c.DeleteImage(i, *reason)
		if err != nil {
			panic(err)
		}

		log.Printf("Deleted %s because %s", i, *reason)
	}
}
