package derpigo

import (
	"encoding/json"
	"fmt"
)

/*
Image is an image on the Booru.
*/
type Image struct {
	ID               string             `json:"id"`
	IDNumber         int                `json:"id_number"`
	CreatedAt        string             `json:"created_at"`
	UpdatedAt        string             `json:"updated_at"`
	DuplicateReports []*DuplicateReport `json:"duplicate_reports"`
	FileName         string             `json:"file_name"`
	Description      string             `json:"description"`
	Uploader         string             `json:"uploader"`
	Image            string             `json:"image"`
	Score            int                `json:"score"`
	Upvotes          int                `json:"upvotes"`
	Downvotes        int                `json:"downvotes"`
	Faves            int                `json:"faves"`
	CommentCount     int                `json:"comment_count"`
	Tags             string             `json:"tags"`
	TagIds           []string           `json:"tag_ids"`
	Width            int                `json:"width"`
	Height           int                `json:"height"`
	AspectRatio      float64            `json:"aspect_ratio"`
	OriginalFormat   string             `json:"original_format"`
	MimeType         string             `json:"mime_type"`
	Sha512Hash       string             `json:"sha512_hash"`
	OrigSha512Hash   string             `json:"orig_sha512_hash"`
	SourceURL        string             `json:"source_url"`
	License          string             `json:"license"`
	Representations  struct {
		ThumbTiny  string `json:"thumb_tiny"`
		ThumbSmall string `json:"thumb_small"`
		Thumb      string `json:"thumb"`
		Small      string `json:"small"`
		Medium     string `json:"medium"`
		Large      string `json:"large"`
		Tall       string `json:"tall"`
		Full       string `json:"full"`
	} `json:"representations"`
	IsRendered  bool `json:"is_rendered"`
	IsOptimized bool `json:"is_optimized"`
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
	if err != nil {
		return nil, err
	}

	return i, nil
}
