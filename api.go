package derpigo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// derpigo-specific errors.
var (
	ErrNeedsOneSlash       = errors.New("derpigo: this needs one slash in its invocation")
	ErrTooLongForBoardName = errors.New("derpigo: this is too long to be a board name")
	ErrNotSpecified        = errors.New("derpigo: some real bad shit happened")
)

/*
Connection models the connection to the Derpibooru API.
*/
type Connection struct {
	apiKey string // API key for all DB communication
}

// New creates a new connection to the Derpibooru API.
func New(apikey string) (c *Connection) {
	if strings.HasSuffix(apikey, "\n") {
		apikey = strings.Split(apikey, "\n")[0]
		log.Printf("Had to trim newline from api key?")
	}

	c = &Connection{
		apiKey: apikey,
	}

	return
}

/*
getJson gets the raw json from the API as a byteslice. It will return the byte slice
representing the json and an error if the underlying call failed. The error will be a
derpigo.Error to make debugging the API easier for the Derpibooru staff.
*/
func (c *Connection) getJson(fragment string, expected int) (data []byte, err error) {
	resp, err := http.Get("https://derpibooru.org/" + fragment + "?key=" + c.apiKey)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != expected {
		return nil, NewError(
			fmt.Errorf(
				"derpigo: expected code %d for https://derpibooru.org/%s, got %d",
				expected,
				fragment,
				resp.StatusCode,
			),
			resp,
		)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(err, resp)
	}

	return
}
