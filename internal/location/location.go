package location

import (
	"context"
	"fmt"

	"googlemaps.github.io/maps"
)

type Location struct {
	Latitude  float64
	Longitude float64
}

type LocationClient struct {
	c *maps.Client
}

func New(ctx context.Context, key string) (*LocationClient, error) {

	c, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		return nil, err
	}

	l := LocationClient{
		c: c,
	}

	return &l, nil
}

func (l *LocationClient) AddressLocation(ctx context.Context, a string) (*Location, error) {

	r := maps.GeocodingRequest{
		Address: a,
	}

	resp, err := l.c.Geocode(ctx, &r)
	if err != nil {
		return nil, fmt.Errorf("unable to get geocode information: %s\n", err)
	}

	if len(resp) != 1 {
		return nil, fmt.Errorf("expected one result, got [%d]", len(resp))
	}

	ll := resp[0].Geometry.Location

	loc := Location{
		Latitude:  ll.Lat,
		Longitude: ll.Lng,
	}

	return &loc, nil

}
