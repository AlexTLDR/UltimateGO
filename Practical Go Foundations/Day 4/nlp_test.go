package nlp

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	test := "What's on second?"
	expected := []string{"what", "s", "on", "second"}
	tokens := Tokenize(text)
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
}
