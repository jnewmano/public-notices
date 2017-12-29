package locationbing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jnewmano/public-notices/internal/location/locationtypes"
)

type BingLocation struct {
	apiKey string
}

func New(ctx context.Context, apiKey string) (*BingLocation, error) {
	b := BingLocation{
		apiKey: apiKey,
	}

	return &b, nil
}

func (b *BingLocation) AddressLocation(ctx context.Context, a string) (*locationtypes.Location, error) {

	u := url.URL{
		Scheme: "https",
		Host:   "dev.virtualearth.net",
		Path:   "REST/v1/Locations",
	}

	args := url.Values{}
	args.Add("q", a)
	args.Add("key", b.apiKey)
	u.RawQuery = args.Encode()

	// http: //dev.virtualearth.net/REST/v1/Locations?q=address&key=keyValue

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := LocationResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if len(result.ResourceSets) != 1 {
		return nil, fmt.Errorf("expected one resource set result, got [%d] [%s]", len(result.ResourceSets), a)
	}

	if len(result.ResourceSets[0].Resources) != 1 {
		return nil, fmt.Errorf("expected one resources result, got [%d] [%s]", len(result.ResourceSets[0].Resources), a)
	}

	p := result.ResourceSets[0].Resources[0].Point

	loc := locationtypes.Location{
		Latitude:  p.Coordinates[0],
		Longitude: p.Coordinates[1],
	}

	return &loc, nil
}

type LocationResponse struct {
	AuthenticationResultCode string `json:"authenticationResultCode"`
	BrandLogoURI             string `json:"brandLogoUri"`
	Copyright                string `json:"copyright"`
	ResourceSets             []struct {
		EstimatedTotal int `json:"estimatedTotal"`
		Resources      []struct {
			Type  string    `json:"__type"`
			Bbox  []float64 `json:"bbox"`
			Name  string    `json:"name"`
			Point struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"point"`
			Address struct {
				AddressLine      string `json:"addressLine"`
				AdminDistrict    string `json:"adminDistrict"`
				AdminDistrict2   string `json:"adminDistrict2"`
				CountryRegion    string `json:"countryRegion"`
				FormattedAddress string `json:"formattedAddress"`
				Locality         string `json:"locality"`
				PostalCode       string `json:"postalCode"`
			} `json:"address"`
			Confidence    string `json:"confidence"`
			EntityType    string `json:"entityType"`
			GeocodePoints []struct {
				Type              string    `json:"type"`
				Coordinates       []float64 `json:"coordinates"`
				CalculationMethod string    `json:"calculationMethod"`
				UsageTypes        []string  `json:"usageTypes"`
			} `json:"geocodePoints"`
			MatchCodes []string `json:"matchCodes"`
		} `json:"resources"`
	} `json:"resourceSets"`
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	TraceID           string `json:"traceId"`
}
