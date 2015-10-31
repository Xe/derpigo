package derpigo

import "testing"

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

	if thread.Topics[0].TopicID != 8813 {
		t.Fatal("The time wasting thread is different. Panic!")
	}
}

func TestGetInvalidForum(t *testing.T) {
	_, err := myC.GetForum("allah")
	if err == nil {
		t.Fatal("Derpibooru has converted to islam. WTF?")
	}
}

func TestGetForum(t *testing.T) {
	_, err := myC.GetForum("dis")
	if err != nil {
		t.Fatal(err)
	}
}
