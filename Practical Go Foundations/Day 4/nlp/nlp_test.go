package nlp

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

var tokanizeCases = []struct { // anonymus struct
	text   string
	tokens []string
}{
	{"Who's on first?", []string{"who", "s", "on", "first"}},
	{"", nil},
}

func TestTokenizeTable(t *testing.T) {
	for _, tc := range tokanizeCases {
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			require.Equal(t, tc.tokens, tokens)
		})
	}
}

func TestTokenize(t *testing.T) {
	test := "What's on second?"
	expected := []string{"what", "s", "on", "second"}
	tokens := Tokenize(text)
	require.DeepEqual(t, expected, tokens)
	/* Before testify
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
	*/
}
