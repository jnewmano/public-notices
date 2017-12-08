package notice

import (
	"context"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseNotice(t *testing.T) {
	ctx := context.Background()

	f, err := ioutil.ReadFile("notices.txt")
	if err != nil {
		t.Fatalf("unable to read file: %s\n", err)
	}

	s := strings.Split(string(f), "\n")

	n, err := parseNotices(ctx, s)
	if err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}

	if len(n) != 8 {
		t.Fatalf("expected to find 8 notices, found %d\n", len(n))
	}
}
