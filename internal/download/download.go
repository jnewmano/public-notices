package download

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultTimeout   = time.Second * 30
	defaultUserAgent = "Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19"
)

type Info struct {
	URL          string
	LastModified string
	ContentType  string

	ETag    string
	Expires string
}

// Download returns an io.ReadCloser of the contents at the URL,
// if the status code is a 2xx.
// Caller is responsible for closing the reader
func Download(ctx context.Context, url string) (io.ReadCloser, error) {

	req, err := NewRequest(ctx, "GET", url)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		resp.Body.Close()
		return nil, fmt.Errorf("bad status code: %d\n", resp.StatusCode)
	}

	return resp.Body, nil
}

func Head(ctx context.Context, url string) (*Info, error) {

	ctx, done := context.WithTimeout(ctx, defaultTimeout)
	defer done()

	client := http.DefaultClient

	req, err := NewRequest(ctx, "HEAD", url)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	lastModified := resp.Header.Get("Last-Modified")
	contentType := resp.Header.Get("Content-Type")
	eTag := resp.Header.Get("Etag")
	expires := resp.Header.Get("Expires")

	i := Info{
		URL:          url,
		LastModified: lastModified,
		ContentType:  contentType,
		ETag:         eTag,
		Expires:      expires,
	}

	return &i, nil
}

func NewRequest(ctx context.Context, method string, url string) (*http.Request, error) {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", defaultUserAgent)

	return req, nil
}
