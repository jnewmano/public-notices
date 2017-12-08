package address

import (
	"context"
	"fmt"
	"regexp"
)

var addressRegex = regexp.MustCompile(`located at (approximately )?(\b.*\b) (in|changing)`)

type Address struct {
	Approximate bool
	Location    string

	IndexStart int
	IndexEnd   int
}

func ExtractAddress(ctx context.Context, s string) (*Address, error) {

	// assume that the contents are consistently formatted
	// located at
	// located at approximately

	// terminated with
	// changing, in

	// if this fails, try to use natural language processing???

	matches := addressRegex.FindAllStringSubmatchIndex(s, -1)

	if len(matches) == 0 {
		return nil, fmt.Errorf("no matches")
	}

	isApproximate := matches[0][2] >= 0
	start := matches[0][4]
	end := matches[0][5]

	a := Address{
		Approximate: isApproximate,
		IndexStart:  start,
		IndexEnd:    end,
		Location:    s[start:end],
	}

	return &a, nil

}
