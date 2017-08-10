package derpigo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

/*
User represents one of the crazy, crazy people that populate this site.
*/
type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	CommentCount int    `json:"comment_count"`
	UploadCount  int    `json:"upload_count"`
	PostCount    int    `json:"post_count"`
	TopicCount   int    `json:"topic_count"`
}

/*
GetUser returns information on a user based on their ID.
*/
func (c *Connection) GetUser(ctx context.Context, id string) (*User, error) {
	data, _, err := c.apiCall(ctx, http.MethodGet, "profiles/"+id+".json", url.Values{}, nil, 200)
	if err != nil {
		return nil, err
	}

	u := &User{}

	err = json.Unmarshal(data, u)

	return u, err
}
