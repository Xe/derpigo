# derpigo
--
    import "github.com/Xe/derpigo"

Pacakge derpigo is a set of dead simple Derpibooru [1] API bindings for Go
programs.

[1]: https://derpibooru.org

## Usage

```go
var (
	ErrNeedsOneSlash = errors.New("derpigo: this needs one slash in its invocation")
)
```
derpigo-specific errors.

#### type Connection

```go
type Connection struct {
}
```

Connection models the connection to the Derpibooru API.

#### func  New

```go
func New(options ...Option) (c *Connection)
```
New creates a new connection to the Derpibooru API.

#### func (*Connection) DeleteImage

```go
func (c *Connection) DeleteImage(id, why string) (err error)
```
DeleteImage deletes an image from the booru. Needs assistant/mod/admin API key.

This is useful for handling morons scriptedly. There seems to be a lot of them
lately. This is a shame.

#### func (*Connection) GetFilter

```go
func (c *Connection) GetFilter(ctx context.Context, id int64) (f *Filter, err error)
```
GetFilter returns a filter or an error.

#### func (*Connection) GetForum

```go
func (c *Connection) GetForum(ctx context.Context, name string) (*Forum, error)
```
GetForum returns a forum structure, all ready to go!

Please note that the creators of this library are not responsible for any mental
scarring that may result thanks to usage of this site's API.

#### func (*Connection) GetImage

```go
func (c *Connection) GetImage(ctx context.Context, id int) (*Image, []Interaction, error)
```
GetImage grabs image information with the api key of the recieving Connection.
If something fails it returns an error.

#### func (*Connection) GetThreadByName

```go
func (c *Connection) GetThreadByName(ctx context.Context, name string) ([]Post, error)
```
GetThreadByName returns a Thread based on the given thread name.

#### func (*Connection) GetUser

```go
func (c *Connection) GetUser(ctx context.Context, id string) (*User, error)
```
GetUser returns information on a user based on their ID.

#### type DupeReportModifier

```go
type DupeReportModifier struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	CommentCount int    `json:"comment_count"`
	UploadCount  int    `json:"upload_count"`
	PostCount    int    `json:"post_count"`
	TopicCount   int    `json:"topic_count"`
}
```

DupeReportModifier is the weighting of the an image duplicate report.

#### type DuplicateReport

```go
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
```

DuplicateReport is a duplicate image report.

#### type Error

```go
type Error struct {
	Underlying error
	RequestID  string
}
```

Error is a combination of a Go error and a Derpibooru request ID to help with
debugging failed API calls with the Derpibooru staff.

#### func  NewError

```go
func NewError(underlying error, resp *http.Response) *Error
```
NewError wraps an error with the X-Request-Id.

#### func (*Error) Error

```go
func (e *Error) Error() string
```
Error satisfies the error interface.

#### type Filter

```go
type Filter struct {
	ID               int64   `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	HiddenTagIds     []int64 `json:"hidden_tag_ids"`
	SpoileredTagIds  []int64 `json:"spoilered_tag_ids"`
	SpoileredTags    string  `json:"spoilered_tags"`
	HiddenTags       string  `json:"hidden_tags"`
	HiddenComplex    string  `json:"hidden_complex"`
	SpoileredComplex string  `json:"spoilered_complex"`
	Public           bool    `json:"public"`
	System           bool    `json:"system"`
	UserCount        int     `json:"user_count"`
	UserID           int64   `json:"user_id"`
}
```

Filter is an image or tag filter.

These are really fucking important. If you filter things badly, shit's not gonna
be fun.

#### type Forum

```go
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
```

Fourm is one of the discussion forums on Derpibooru.

#### type Image

```go
type Image struct {
	ID               string             `json:"id"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	DuplicateReports []*DuplicateReport `json:"duplicate_reports"`
	FirstSeenAt      time.Time          `json:"first_seen_at"`
	UploaderID       interface{}        `json:"uploader_id"`
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
	IsRendered   bool          `json:"is_rendered"`
	IsOptimized  bool          `json:"is_optimized"`
	Interactions []interface{} `json:"interactions"`
}
```

Image is an image on the Booru.

#### type Interaction

```go
type Interaction struct {
	ID              int    `json:"id"`
	InteractionType string `json:"interaction_type"`
	Value           string `json:"value"`
	UserID          int    `json:"user_id"`
	ImageID         int    `json:"image_id"`
}
```

Interaction is the "hard copy" of user interactions on images. Possible kinds
include (but are not limited to):

    - down
    - up
    - faved

#### type Option

```go
type Option func(*Connection)
```

Option is a function that modifies the given Connection.

#### func  WithAPIKey

```go
func WithAPIKey(apiKey string) Option
```
WithAPIKey specifies a given API key for all API calls.

#### func  WithDomain

```go
func WithDomain(domain string) Option
```
WithDomain specifies a different base domain to do API calls against

#### type Post

```go
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
```

Post is an individual forum post.

#### type User

```go
type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	CommentCount int    `json:"comment_count"`
	UploadCount  int    `json:"upload_count"`
	PostCount    int    `json:"post_count"`
	TopicCount   int    `json:"topic_count"`
}
```

User represents one of the crazy, crazy people that populate this site.
