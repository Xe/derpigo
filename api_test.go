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

func TestGetImage(t *testing.T) {
	i, err := myC.GetImage(912673)
	if err != nil {
		t.Fatal(err)
	}

	if i.IDNumber != 912673 {
		t.Fatalf("ID is wrong: expected 912673, got %d", i.IDNumber)
	}
}

func TestGetThreadByNameButNotBecauseMissingSlash(t *testing.T) {
	_, err := myC.GetThreadByName("ponies")
	if err == nil {
		t.Fatal("This allowed invalid input")
	}
}

func TestGetThreadByName(t *testing.T) {
	thread, err := myC.GetThreadByName("dis/the-time-wasting-thread")
	if err != nil {
		t.Fatal(err)
	}

	if thread.Topics[0].TopicID != "5161dd617f123bd25900013d" {
		t.Fatal("The time wasting thread is different. Panic!")
	}
}
