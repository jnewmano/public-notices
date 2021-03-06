package tokenize

import (
	"bytes"
	"context"
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	ctx := context.Background()

	r := bytes.NewBuffer([]byte(sampleText))

	tokens, err := Tokenize(ctx, r)
	if err != nil {
		t.Fatalf("unable to split sample text: %s\n", err)
	}

	if len(tokens) != 3 {
		fmt.Printf("%#v\n", tokens)
		t.Fatalf("unexpected number of tokens [%d]\n", len(tokens))
	}

}

const sampleText = `
This is my sentence
broken onto two lines.



This the the second 
sentence broken onto two lines.

Final sentence.


`
