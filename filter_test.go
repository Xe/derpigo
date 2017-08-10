package derpigo

import "testing"

func TestGetFilter(t *testing.T) {
	ctx, myC := setup()

	const filterID = 50106 // https://derpibooru.org/filters/50106

	_, err := myC.GetFilter(ctx, filterID)
	if err != nil {
		t.Fatal(err)
	}
}
