package tokenize

import (
	"bufio"
	"context"
	"io"
	"strings"
)

// Tokenize takes a reader and returns a string slice
// Where each element in the string slice is a block of text
// not separated by more than one new line.
func Tokenize(ctx context.Context, r io.Reader) ([]string, error) {

	b := bufio.NewReader(r)

	tokens := []string{}
	current := ""
	count := 0

	for {
		l, err := b.ReadString('\n')
		if err == io.EOF {
			if len(current) > 0 {
				tokens = append(tokens, current)
			}
			break

		} else if err != nil {
			return nil, err
		}

		l = strings.TrimSpace(l)
		if len(l) == 0 {
			if len(current) == 0 {
				continue
			}

			count = 0
			tokens = append(tokens, current)
			current = ""
			continue
		}

		sp := " "
		if count == 0 {
			sp = ""
		}
		current = current + sp + string(l)
		count++

	}

	return tokens, nil

}
