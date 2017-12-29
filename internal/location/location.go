package location

import (
	"context"
	"fmt"

	"github.com/jnewmano/public-notices/internal/location/locationbing"
	"github.com/jnewmano/public-notices/internal/location/locationgoogle"
	"github.com/jnewmano/public-notices/internal/location/locationtypes"
)

type LocationProvider interface {
	AddressLocation(context.Context, string) (*locationtypes.Location, error)
}

type LocationClient struct {
	google LocationProvider
	bing   LocationProvider
}

func New(ctx context.Context, googleAPIKey string, bingAPIKey string) (*LocationClient, error) {

	g, err := locationgoogle.New(ctx, googleAPIKey)
	if err != nil {
		return nil, err
	}

	b, err := locationbing.New(ctx, bingAPIKey)
	if err != nil {
		return nil, err
	}

	l := LocationClient{
		google: g,
		bing:   b,
	}

	return &l, nil
}

func (l *LocationClient) AddressLocation(ctx context.Context, a string) (*locationtypes.Location, error) {

	// try to lookup location first with google
	location, err := l.google.AddressLocation(ctx, a)
	if err == nil {
		return location, nil
	}

	fmt.Printf("Unable to lookup location with google, trying bing: [%s]\n", err)

	// if that fails, then try bing
	location, err = l.bing.AddressLocation(ctx, a)
	if err == nil {
		return location, nil
	}

	fmt.Printf("Unable to lookup location with bing: [%s]\n", err)

	return nil, err

}
