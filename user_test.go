package derpigo

import "testing"

func TestGetUserInformation(t *testing.T) {
	u, err := myC.GetUser("5139964d7f123b8997000291") // Creator of this library
	if err != nil {
		t.Fatal(err)
	}

	if u.Name != "Xena" {
		t.Fatalf("Wrong name was looked up: got %s, expected Xena")
	}
}
