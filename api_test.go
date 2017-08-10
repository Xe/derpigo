package derpigo

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func setup() (context.Context, *Connection) {
	return context.Background(), New(WithAPIKey(os.Getenv("DB_API_KEY")))
}

func TestAPICall(t *testing.T) {
	ctx, myC := setup()

	const imageID = "912673" // https://derpibooru.org/912673

	_, _, err := myC.apiCall(ctx, http.MethodGet, imageID+".json", url.Values{}, nil, 200)
	if err != nil {
		t.Fatal(err)
	}
}
