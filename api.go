package derpigo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

/*
GetImage grabs image information with the api key of the recieving Connection.
If something fails it returns an error.
*/
func (c *Connection) GetImage(id int) (*Image, error) {
	data, err := c.getJson(fmt.Sprintf("%d.json", id), 200)
	if err != nil {
		return nil, err
	}

	i := &Image{}

	err = json.Unmarshal(data, i)
	if err != nil {
		return nil, err
	}

	return i, nil
}
