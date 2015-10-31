package derpigo

import "testing"

func TestGetFilter(t *testing.T) {
	_, err := myC.GetFilter(50106)
	if err != nil {
		t.Fatal(err)
	}
}
