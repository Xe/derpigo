package derpigo

import (
	"os"
	"testing"
)

func TestNewAPI(t *testing.T) {
	myC := New(
		os.Getenv("DB_API_KEY"),
	)

	_, err := myC.getJson("912673.json", 200)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetImage(t *testing.T) {
	myC := New(
		os.Getenv("DB_API_KEY"),
	)

	i, err := myC.GetImage(912673)
	if err != nil {
		t.Fatal(err)
	}

	if i.IDNumber != 912673 {
		t.Fatalf("ID is wrong: expected 912673, got %d", i.IDNumber)
	}
}
