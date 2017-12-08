package address

import (
	"context"
	"testing"
)

var test1 = `Public hearing and consideration of Bruce Mackay’s request for approval of a conditional use for the Bel Aire Senior Living site plan, a 3.5-acre assisted living facility located at approximately Chapel Ridge Road & Shady Bend Lane in a Planned Community zone.`
var expected1 = `Chapel Ridge Road & Shady Bend Lane`

var test2 = `Public hearing and consideration of Bruce Mackay’s request for approval of a conditional use for the Bel Aire Senior Living site plan, a 3.5-acre assisted living facility located at Chapel Ridge Road & Shady Bend Lane in a Planned Community zone.`
var expected2 = `Chapel Ridge Road & Shady Bend Lane`

var test3 = `abc`
var expected3 = ``

func TestExtractAddress(t *testing.T) {
	ctx := context.Background()

	a, err := ExtractAddress(ctx, test1)
	if err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}

	if a.Approximate == false {
		t.Fatalf("expected address to be approximate")
	}

}
