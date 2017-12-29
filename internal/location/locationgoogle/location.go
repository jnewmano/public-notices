package locationgoogle

import (
	"context"
	"fmt"

	"github.com/jnewmano/public-notices/internal/location/locationtypes"
	"googlemaps.github.io/maps"
)

type GoogleLocation struct {
	c *maps.Client
}

func New(ctx context.Context, key string) (*GoogleLocation, error) {

	c, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		return nil, err
	}

	l := GoogleLocation{
		c: c,
	}

	return &l, nil
}

func (l *GoogleLocation) AddressLocation(ctx context.Context, a string) (*locationtypes.Location, error) {

	r := maps.GeocodingRequest{
		Address: a,
	}

	resp, err := l.c.Geocode(ctx, &r)
	if err != nil {
		return nil, fmt.Errorf("unable to get geocode information: %s\n", err)
	}

	if len(resp) != 1 {
		return nil, fmt.Errorf("expected one result, got [%d] [%s]", len(resp), a)
	}

	ll := resp[0].Geometry.Location

	loc := locationtypes.Location{
		Latitude:  ll.Lat,
		Longitude: ll.Lng,
	}

	return &loc, nil

}
