package derpigo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// derpigo-specific errors.
var (
	ErrNeedsOneSlash = errors.New("derpigo: this needs one slash in its invocation")
)

/*
Connection models the connection to the Derpibooru API.
*/
type Connection struct {
	apiKey string // API key for all DB communication
	domain string // domain name to communicate with
}

// Option is a function that modifies the given Connection.
type Option func(*Connection)

// WithAPIKey specifies a given API key for all API calls.
func WithAPIKey(apiKey string) Option {
	// automatically trim newline from end of API key.
	if strings.HasSuffix(apiKey, "\n") {
		apiKey = strings.Split(apiKey, "\n")[0]
	}

	return func(c *Connection) {
		c.apiKey = apiKey
	}
}

// WithDomain specifies a different base domain to do API calls against
func WithDomain(domain string) Option {
	return func(c *Connection) {
		c.domain = domain
	}
}

// New creates a new connection to the Derpibooru API.
func New(options ...Option) (c *Connection) {

	c = &Connection{domain: "derpibooru.org"}

	for _, opt := range options {
		opt(c)
	}

	return
}

func (c *Connection) apiCall(ctx context.Context, method, route string, args url.Values, body interface{}, wantResponseCode int) ([]byte, []Interaction, error) {
	var (
		buf *bytes.Buffer = bytes.NewBuffer(nil)
		req *http.Request
		err error
	)

	purl, err := url.Parse(fmt.Sprintf("https://%s/%s", c.domain, route))
	if err != nil {
		return nil, nil, err
	}

	if ak := c.apiKey; ak != "" {
		args.Add("key", ak)
	}

	purl.RawQuery = args.Encode()
	urlStr := purl.String()

	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, nil, err
		}

		req, err = http.NewRequest(method, urlStr, buf)
	} else {
		req, err = http.NewRequest(method, urlStr, nil)
	}

	if err != nil {
		return nil, nil, err
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	type interactionsWrapper struct {
		Interactions []Interaction `json:"interactions"`
	}

	iw := interactionsWrapper{}

	// if this fails it's okay
	_ = json.Unmarshal(data, &iw)

	if resp.StatusCode != wantResponseCode {
		purl.RawQuery = ""

		return nil, nil, NewError(
			fmt.Errorf(
				"derpigo: expected code %d for %s, got %d",
				wantResponseCode,
				purl.String(),
				resp.StatusCode,
			),
			resp,
		)
	}

	return data, iw.Interactions, nil
}

// Interaction is the "hard copy" of user interactions on images. Possible kinds include (but are not limited to):
//
//     - down
//     - up
//     - faved
type Interaction struct {
	ID              int    `json:"id"`
	InteractionType string `json:"interaction_type"`
	Value           string `json:"value"`
	UserID          int    `json:"user_id"`
	ImageID         int    `json:"image_id"`
}
