package derpigo

import "testing"

func TestGetImage(t *testing.T) {
	i, err := myC.GetImage(912673)
	if err != nil {
		t.Fatal(err)
	}

	if i.ID != "912673" {
		t.Fatalf("ID is wrong: expected 912673, got %s", i.ID)
	}
}
