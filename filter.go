package derpigo

// Filter is an image or tag filter.
//
// These are really fucking important. If you filter things badly, shit's
// not gonna be fun.
type Filter struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	HiddenTagIds     []string `json:"hidden_tag_ids"`
	SpoileredTagIds  []string `json:"spoilered_tag_ids"`
	SpoileredTags    string   `json:"spoilered_tags"`
	HiddenTags       string   `json:"hidden_tags"`
	HiddenComplex    string   `json:"hidden_complex"`
	SpoileredComplex string   `json:"spoilered_complex"`
	Public           bool     `json:"public"`
	System           bool     `json:"system"`
	UserCount        int      `json:"user_count"`
	UserID           string   `json:"user_id"`
}
