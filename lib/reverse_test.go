package lib

import "testing"

func TestReverse_WillReverseAcceptableString(t *testing.T) {
	original := "hello"
	expected := "olleh"
	actual, _ := Reverse(original)
	if actual != expected {
		t.Fatalf("actual: %s was not equal to expected %s", actual, expected)
	}
}
