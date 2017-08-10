package derpigo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Filter is an image or tag filter.
//
// These are really fucking important. If you filter things badly, shit's
// not gonna be fun.
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

// GetFilter returns a filter or an error.
func (c *Connection) GetFilter(ctx context.Context, id int64) (f *Filter, err error) {
	data, _, err := c.apiCall(ctx, http.MethodGet, fmt.Sprintf("filters/%d.json", id), url.Values{}, nil, 200)
	if err != nil {
		return nil, err
	}

	f = &Filter{}

	err = json.Unmarshal(data, f)

	return f, err
}
