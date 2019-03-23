# derpigo

[API Docs](https://godoc.org/within.website/derpigo) [License](https://github.com/Xe/derpigo/blob/master/LICENSE)

## Installation

```console
$ go get -u github.com/Xe/derpigo
```

## Usage Example

```go
package main

import (
    "context"
    "os"

    "within.website/derpigo"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    c := derpigo.New(derpigo.WithAPIKey(os.Getenv("DERPI_API_KEY")))

    const imgID = 1330414 // https://derpibooru.org/1330414

    img, interactions, err := c.GetImage(ctx, imgID)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%#v", img)
    log.Printf("#%v", interactions)
}
```
