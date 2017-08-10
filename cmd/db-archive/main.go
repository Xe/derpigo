package main

import (
	"context"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/Xe/derpigo"
)

var (
	keyFile = flag.String("keyfile", os.Getenv("HOME")+"/.local/share/within/db.cadance.key", "file with the derpibooru key to use")
	outDir  = flag.String("outdir", ".", "directory to download images to")
)

func main() {
	flag.Parse()

	key, err := ioutil.ReadFile(*keyFile)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c := derpigo.New(derpigo.WithAPIKey(string(key)))

	for _, i := range flag.Args() {
		id, err := strconv.Atoi(i)
		if err != nil {
			if strings.HasPrefix(i, "http") {
				url, err := url.Parse(i)
				if err != nil {
					log.Printf("error reading %s: %#v", i, err)
					continue
				}

				realid, err := strconv.Atoi(url.Path[1:])
				if err != nil {
					log.Printf("I don't understand %s", url.Path)
					continue
				}

				id = realid
			}
		}

		img, _, err := c.GetImage(ctx, id)
		if err != nil {
			log.Printf("couldn't fetch info on image %d: %v", id, err)
		}

		tags := strings.Split(img.Tags, ", ")
		if len(tags) > 10 {
			tags = tags[0:11]
		}

		foutpath := *outDir + img.ID + " " + strings.Join(tags, ", ") + "." + img.OriginalFormat

		fout, err := os.Create(foutpath)
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

		log.Printf("downloaded %v to %v", i, foutpath)
	}
}
