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

	const threadName = "dis/the-time-wasting-thread" // https://derpibooru.org/forums/dis/topics/the-time-wasting-thread
	const threadTopicID = 8813

	thread, err := myC.GetThreadByName(ctx, threadName)
	if err != nil {
		t.Fatal(err)
	}

	if thread[0].TopicID != threadTopicID {
		t.Fatalf("expected thread id does not match: want %d: %d", threadTopicID, thread[0].TopicID)
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

	const forumName = "dis" // https://derpibooru.org/forums/dis

	_, err := myC.GetForum(ctx, forumName)
	if err != nil {
		t.Fatal(err)
	}
}
