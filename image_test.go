package derpigo

import "testing"

func TestGetImage(t *testing.T) {
	ctx, myC := setup()

	const imageID = 912673 // https://derpibooru.org/912673

	i, _, err := myC.GetImage(ctx, imageID)
	if err != nil {
		t.Fatal(err)
	}

	if i.ID != imageID {
		t.Fatalf("ID is wrong: expected %v, got %d", imageID, i.ID)
	}
}
