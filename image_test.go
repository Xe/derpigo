package derpigo

import "testing"

func TestGetImage(t *testing.T) {
	ctx, myC := setup()

	i, _, err := myC.GetImage(ctx, 912673)
	if err != nil {
		t.Fatal(err)
	}

	if i.ID != "912673" {
		t.Fatalf("ID is wrong: expected 912673, got %s", i.ID)
	}
}
