package checker

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChecker(t *testing.T) {

	ctx := context.Background()

	count := 0
	f := func(ctx context.Context, name string, r io.Reader) error {
		count++
		out, _ := ioutil.ReadAll(r)
		if len(out) != 14 {
			t.Fatalf("unexpected body length [%d]\n", len(out))
		}

		return nil
	}

	c, err := New(f, f, f)
	if err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}

	testTag := "some tag"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Etag", testTag)
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	url := ts.URL
	tag := ""

	tag, err = c.Do(ctx, url, tag)
	if err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}

	if count != 3 {
		t.Fatalf("unexpected processor call count [%d]\n", count)
	}

	// we should not download and process again
	tag, err = c.Do(ctx, url, tag)
	if err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}

	if count != 3 {
		t.Fatalf("unexpected processor call count [%d]\n", count)
	}

}
