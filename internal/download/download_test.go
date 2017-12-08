package download

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownload(t *testing.T) {
	ctx := context.Background()

	testTag := "some tag"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Etag", testTag)
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	fmt.Println(ts.URL + "/some/path")
	url := ts.URL

	info, err := Head(ctx, url)
	if err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}

	if info.ETag != testTag {
		t.Fatalf("unexpected tag header [%s]\n", info.ETag)
	}

	ctx, done := context.WithTimeout(ctx, defaultTimeout)
	defer done()

	r, err := Download(ctx, url)
	if err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}
	defer r.Close()

	out, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("unable to read body: %s\n", err)
	}

	if len(out) == 0 {
		t.Fatalf("did not read any bytes")
	}

}
