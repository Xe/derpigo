package derpigo

import "testing"

func TestGetUserInformation(t *testing.T) {
	ctx, myC := setup()

	const username = "Xena" // https://derpibooru.org/profiles/Xena - Creator of this library

	u, err := myC.GetUser(ctx, username)
	if err != nil {
		t.Fatal(err)
	}

	if u.Name != username {
		t.Fatalf("Wrong name was looked up: got %s, expected %s", u.Name, username)
	}
}
