package derpigo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
Post is an individual forum post.
*/
type Post struct {
	ID      int64  `json:"id"`
	TopicID int64  `json:"topic_id"`
	Author  string `json:"author"`
	Subject string `json:"subject"`

	// A warning to people. The forums are a very scary place on Derpibooru.
	// Some areas are containement. This is the unformatted Textile version
	// of the body.
	Body string `json:"body"`
}

/*
GetThreadByName returns a Thread based on the given thread name.
*/
func (c *Connection) GetThreadByName(ctx context.Context, name string) ([]Post, error) {
	if strings.Count(name, "/") != 1 {
		return nil, ErrNeedsOneSlash
	}

	data, _, err := c.apiCall(ctx, http.MethodGet, name+".json", url.Values{}, nil, 200)
	if err != nil {
		return nil, err
	}

	t := []Post{}

	err = json.Unmarshal(data, &t)

	return t, err
}

/*
Fourm is one of the discussion forums on Derpibooru.
*/
type Forum struct {
	Topics []struct {
		Slug            string    `json:"slug"`
		Title           string    `json:"title"`
		Sticky          bool      `json:"sticky"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		ForumID         string    `json:"forum_id"`
		LastRepliedToAt string    `json:"last_replied_to_at"`
		ID              string    `json:"id"`
	} `json:"topics"`
}

/*
GetForum returns a forum structure, all ready to go!

Please note that the creators of this library are not responsible for any mental
scarring that may result thanks to usage of this site's API.
*/
func (c *Connection) GetForum(ctx context.Context, name string) (*Forum, error) {
	data, _, err := c.apiCall(ctx, http.MethodGet, name+".json", url.Values{}, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}

	f := &Forum{}

	err = json.Unmarshal(data, f)

	return f, err
}
