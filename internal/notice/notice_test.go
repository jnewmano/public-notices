package notice

import (
	"context"
	"io/ioutil"
	"strings"
	"testing"
	"time"
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

func TestParseNoticeDate(t *testing.T) {

	ctx := context.Background()

	line := `On December 14th, 2017 at 7:00 p.m. at the  City Council Chambers located at 15 North 300 East, the
 City Planning Commission will hold a public hearing to receive public comment on the following items:`

	d, err := parseDate(ctx, line)
	if err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}

	if d.Day() != 14 || d.Month() != time.December || d.Year() != 2017 {
		t.Fatalf("unexpected date: %s\n", d)
	}
}
