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
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Xe/derpigo"
)

var (
	keyFile = flag.String("keyFile", "/home/xena/.local/share/within/db.cadance.key", "file with the derpibooru key to use")
	reason  = flag.String("reason", "", "reason to use when deleting images")
	needTag = flag.String("needtag", "", "optional tag an image must have to be deleted")
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
		id, err := strconv.Atoi(i)
		if err != nil {
			log.Printf("Bad number %s", i)
			continue
		}

		img, err := c.GetImage(id)
		if err != nil {
			log.Printf("couldn't fetch info on image %d: %v", id, err)
		}

		if *needTag != "" {
			ok := false

			for _, tag := range strings.Split(img.Tags, ", ") {
				if tag == *needTag {
					ok = true
				}
			}

			if !ok && *needTag != "" {
				log.Printf("Can't delete %d, doesn't have the tag %s", id, *needTag)
				continue
			}
		}

		fout, err := os.Create("/home/xena/pictures/derpi/" + img.ID + " " + img.Tags + "." + img.OriginalFormat)
		if err != nil {
			panic(err)
		}
		defer fout.Close()

		resp, err := http.Get("https:" + img.Image)
		if err != nil {
			log.Printf("could not download image: %v", err)
			continue
		}
		defer resp.Body.Close()

		io.Copy(fout, resp.Body)

		err = c.DeleteImage(i, *reason)
		if err != nil {
			panic(err)
		}

		log.Printf("Deleted %s because %s", i, *reason)
	}
}
