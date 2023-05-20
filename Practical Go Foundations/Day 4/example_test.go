package nlp_test

import (
	"fmt"

	"github.com/AlexTLDR/UltimateGO"
)

func ExampleTokenize() {
	text := "Who's on first"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	//Output:
	// [who s on first]
}
