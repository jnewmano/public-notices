package storage

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"cloud.google.com/go/storage"
)

type Storage struct {
	c *storage.Client
	b *storage.BucketHandle
}

func New(ctx context.Context, bucket string) (*Storage, error) {

	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("opening bucket:", bucket)
	b := c.Bucket(bucket)

	s := Storage{
		c: c,
		b: b,
	}

	return &s, nil
}

func (s *Storage) Write(ctx context.Context, source string, version string, f io.Reader) error {

	name := source + "_" + version

	name = url.PathEscape(name)

	fmt.Println("Storage - Getting object")
	o := s.b.Object(name)

	fmt.Println("Storage - Creating writer")
	w := o.NewWriter(ctx)

	fmt.Println("Storage - Copying")
	_, err := io.Copy(w, f)
	if err != nil {
		_ = w.Close()
		return err
	}

	fmt.Println("Storage - Closing")
	err = w.Close()
	if err != nil {
		return err
	}

	fmt.Println("Storage - Done")

	return nil
}
