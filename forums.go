package derpigo

import (
	"encoding/json"
	"strings"
	"time"
)

/*
Thread is the structure that Derpibooru replies with after asking for
a forum thread.
*/
type Thread struct {
	Topics []Topic `json:"topics"`
}

/*
Topic is an individual forum topic with its replies.
*/
type Topic struct {
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
func (c *Connection) GetThreadByName(name string) (*Thread, error) {
	if strings.Count(name, "/") != 1 {
		return nil, ErrNeedsOneSlash
	}

	data, err := c.getJson(name+".json", 200)
	if err != nil {
		return nil, err
	}

	t := &Thread{}

	err = json.Unmarshal(data, t)

	return t, err
}

/*
Fourm is one of the discussion forums on Derpibooru.

This is where all the chaos and hell mix together into a big ball of horror.
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
func (c *Connection) GetForum(name string) (*Forum, error) {
	if len(name) > 4 {
		return nil, ErrTooLongForBoardName
	}

	data, err := c.getJson(name+".json", 200)
	if err != nil {
		return nil, err
	}

	f := &Forum{}

	err = json.Unmarshal(data, f)

	return f, err
}
