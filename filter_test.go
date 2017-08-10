package derpigo

import "testing"

func TestGetFilter(t *testing.T) {
	ctx, myC := setup()

	_, err := myC.GetFilter(ctx, 50106)
	if err != nil {
		t.Fatal(err)
	}
}
