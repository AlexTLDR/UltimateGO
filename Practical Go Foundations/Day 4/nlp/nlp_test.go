package nlp

import (
	"reflect"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

/*
var tokanizeCases = []struct { // anonymus struct
	text   string
	tokens []string
}{
	{"Who's on first?", []string{"who", "s", "on", "first"}},
	{"", nil},
}
*/

type tokanizeCase struct {
	Text string
	Tokens[string]
}

func loadTokenizeCases(t *testing.T) []loadTokenizeCases {
	dara, err := ioutil.ReadFile("tokenize_cases.toml")
	var testCase struct {
		Cases []tokenizeCase
	}
}
func TestTokenizeTable(t *testing.T) {
	for _, tc := range loadTokenizeCases(t) {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
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
