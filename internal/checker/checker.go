package checker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/jnewmano/public-notices/internal/download"
)

type Processor func(context.Context, string, io.Reader) error

type Checker struct {
	Processors []Processor
}

func New(processors ...Processor) (*Checker, error) {
	c := Checker{
		Processors: processors,
	}

	return &c, nil
}

func (c *Checker) Do(ctx context.Context, url string, lastTag string) (string, error) {

	fmt.Println("Checking document", url)
	info, err := download.Head(ctx, url)
	if err != nil {
		return "", err
	}

	tag := info.ETag

	// use etag to check to see if the remote file hasn't changed
	if tag == lastTag {
		return tag, nil
	}

	fmt.Println("Downloading document", url)
	r, err := download.Download(ctx, url)
	if err != nil {
		return "", err
	}
	defer r.Close()

	name := url + "#" + tag

	fmt.Println("Running processors")
	for i, v := range c.Processors {
		fmt.Println("Processor", i)
		buff := bytes.NewBuffer(nil)

		r2 := io.TeeReader(r, buff)

		err := v(ctx, name, r2)
		if err != nil {
			// TODO: it's not great to abandon the entire pipeline if one stage fails
			return "", fmt.Errorf("processor error [%d] [%s]", i, err)
		}

		r = ioutil.NopCloser(buff)
	}

	return tag, nil
}
