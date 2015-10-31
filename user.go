package derpigo

import "encoding/json"

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

This is kinda spartan, but it will do for now.
*/
func (c *Connection) GetUser(id string) (*User, error) {
	data, err := c.getJson("profiles/"+id+".json", 200)
	if err != nil {
		return nil, err
	}

	u := &User{}

	err = json.Unmarshal(data, u)

	return u, err
}
