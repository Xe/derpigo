package derpigo

import "testing"

func TestGetUserInformation(t *testing.T) {
	u, err := myC.GetUser("Xena") // Creator of this library
	if err != nil {
		t.Fatal(err)
	}

	if u.Name != "Xena" {
		t.Fatalf("Wrong name was looked up: got %s, expected Xena")
	}
}
