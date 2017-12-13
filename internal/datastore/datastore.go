package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
)

type DataStore struct {
	client *datastore.Client
}

func New(ctx context.Context, projectID string) (*DataStore, error) {

	c, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	d := DataStore{
		client: c,
	}

	return &d, nil
}

func (d *DataStore) Future(ctx context.Context, kind string, values interface{}) error {

	q := datastore.NewQuery(kind)

	filter := "Date >"
	v := time.Now()

	q = q.Filter(filter, v)

	_, err := d.client.GetAll(ctx, q, values)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataStore) Put(ctx context.Context, kind string, name string, value interface{}) error {

	// Creates a Key instance.
	key := datastore.NameKey(kind, name, nil)

	// Saves the new entity.
	_, err := d.client.Put(ctx, key, value)
	if err != nil {
		return err
	}

	return nil

}
