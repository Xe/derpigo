package derpigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

/*
Image is an image on the Booru.
*/
type Image struct {
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DuplicateReports []struct {
		ID                 int         `json:"id"`
		State              string      `json:"state"`
		Reason             string      `json:"reason"`
		ImageID            int         `json:"image_id"`
		DuplicateOfImageID int         `json:"duplicate_of_image_id"`
		UserID             interface{} `json:"user_id"`
		Modifier           struct {
			ID           int           `json:"id"`
			Name         string        `json:"name"`
			Slug         string        `json:"slug"`
			Role         string        `json:"role"`
			Description  string        `json:"description"`
			AvatarURL    string        `json:"avatar_url"`
			CreatedAt    time.Time     `json:"created_at"`
			CommentCount int           `json:"comment_count"`
			UploadsCount int           `json:"uploads_count"`
			PostCount    int           `json:"post_count"`
			TopicCount   int           `json:"topic_count"`
			Links        []interface{} `json:"links"`
			Awards       []struct {
				ImageURL  string    `json:"image_url"`
				Title     string    `json:"title"`
				ID        int       `json:"id"`
				Label     string    `json:"label"`
				AwardedOn time.Time `json:"awarded_on"`
			} `json:"awards"`
		} `json:"modifier"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"duplicate_reports"`
	FirstSeenAt     time.Time   `json:"first_seen_at"`
	UploaderID      interface{} `json:"uploader_id"`
	FileName        string      `json:"file_name"`
	Description     string      `json:"description"`
	Uploader        string      `json:"uploader"`
	Image           string      `json:"image"`
	Score           int         `json:"score"`
	Upvotes         int         `json:"upvotes"`
	Downvotes       int         `json:"downvotes"`
	Faves           int         `json:"faves"`
	CommentCount    int         `json:"comment_count"`
	Tags            string      `json:"tags"`
	TagIds          []string    `json:"tag_ids"`
	Width           int         `json:"width"`
	Height          int         `json:"height"`
	AspectRatio     float64     `json:"aspect_ratio"`
	OriginalFormat  string      `json:"original_format"`
	MimeType        string      `json:"mime_type"`
	Sha512Hash      string      `json:"sha512_hash"`
	OrigSha512Hash  string      `json:"orig_sha512_hash"`
	SourceURL       string      `json:"source_url"`
	Representations struct {
		ThumbTiny  string `json:"thumb_tiny"`
		ThumbSmall string `json:"thumb_small"`
		Thumb      string `json:"thumb"`
		Small      string `json:"small"`
		Medium     string `json:"medium"`
		Large      string `json:"large"`
		Tall       string `json:"tall"`
		Full       string `json:"full"`
	} `json:"representations"`
	IsRendered   bool          `json:"is_rendered"`
	IsOptimized  bool          `json:"is_optimized"`
	Interactions []interface{} `json:"interactions"`
}

/*
DuplicateReport is a duplicate image report.
*/
type DuplicateReport struct {
	ID                  string              `json:"id"`
	State               string              `json:"state"`
	Reason              string              `json:"reason"`
	ImageIDNumber       int                 `json:"image_id_number"`
	TargetImageIDNumber int                 `json:"target_image_id_number"`
	User                interface{}         `json:"user"`
	CreatedAt           string              `json:"created_at"`
	Modifier            *DupeReportModifier `json:"modifier"`
}

/*
DupeReportModifier is the weighting of the an image duplicate report.
*/
type DupeReportModifier struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	CommentCount int    `json:"comment_count"`
	UploadCount  int    `json:"upload_count"`
	PostCount    int    `json:"post_count"`
	TopicCount   int    `json:"topic_count"`
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

	return i, err
}

/*
DeleteImage deletes an image from the booru. Needs assistant/mod/admin API key.

This is useful for handling morons scriptedly. There seems to be a lot of them
lately. This is a shame.
*/
func (c *Connection) DeleteImage(id, why string) (err error) {
	u, err := url.Parse(fmt.Sprintf("https://derpibooru.org/images/%s.json", id))
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("key", c.apiKey)
	q.Set("deletion_reason", why)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		panic(err)
	}

	cli := &http.Client{}

	resp, err := cli.Do(req)
	if err != nil {
		return NewError(err, resp)
	}

	return err
}
