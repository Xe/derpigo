package derpigo

import (
	"os"
	"testing"
)

var myC *Connection

func TestNewAPI(t *testing.T) {
	myC = New(
		os.Getenv("DB_API_KEY"),
	)

	_, err := myC.getJson("912673.json", 200)
	if err != nil {
		t.Fatal(err)
	}
}
