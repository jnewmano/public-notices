package storage

import (
	"context"
	"io"

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

	b := c.Bucket(bucket)

	s := Storage{
		c: c,
		b: b,
	}

	return &s, nil
}

func (s *Storage) Write(ctx context.Context, name string, f io.Reader) error {

	o := s.b.Object(name)
	w := o.NewWriter(ctx)

	_, err := io.Copy(w, f)
	if err != nil {
		_ = w.Close()
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}
