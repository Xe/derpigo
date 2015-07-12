package derpigo

import "time"

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
	ID      string `json:"id"`
	TopicID string `json:"topic_id"`
	Author  string `json:"author"`
	Subject string `json:"subject"`

	// A warning to people. The forums are a very scary place on Derpibooru.
	// Some areas are containement. This is the unformatted Textile version
	// of the body.
	Body string `json:"body"`
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
