package derpigo

import "testing"

func TestGetThreadByNameButNotBecauseMissingSlash(t *testing.T) {
	ctx, myC := setup()

	_, err := myC.GetThreadByName(ctx, "ponies")
	if err == nil {
		t.Fatal("This allowed invalid input")
	}
}

func TestGetThreadByName(t *testing.T) {
	ctx, myC := setup()

	thread, err := myC.GetThreadByName(ctx, "dis/the-time-wasting-thread")
	if err != nil {
		t.Fatal(err)
	}

	if thread[0].TopicID != 8813 {
		t.Fatalf("expected thread id does not match: want %d: %d", 8813, thread[0].TopicID)
	}
}

func TestGetInvalidForum(t *testing.T) {
	ctx, myC := setup()

	_, err := myC.GetForum(ctx, "allah")
	if err == nil {
		t.Fatal("invalid forum name resolves")
	}
}

func TestGetForum(t *testing.T) {
	ctx, myC := setup()

	_, err := myC.GetForum(ctx, "dis")
	if err != nil {
		t.Fatal(err)
	}
}
