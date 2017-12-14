package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jnewmano/public-notices/internal/checker"
	"github.com/jnewmano/public-notices/internal/datastore"
	"github.com/jnewmano/public-notices/internal/location"
	"github.com/jnewmano/public-notices/internal/processor"
	"github.com/jnewmano/public-notices/internal/server"
	"github.com/jnewmano/public-notices/internal/storage"
)

func main() {
	ctx := context.Background()

	httpAddr := os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		httpAddr = ":80"
	}

	url := os.Getenv("TARGET_URL")
	if url == "" {
		exit("missing target URL", fmt.Errorf("TARGET_URL not set"))
	}

	mapsAPIKey := os.Getenv("MAPS_API_KEY")
	if mapsAPIKey == "" {
		exit("missing maps credentials", fmt.Errorf("MAPS_API_KEY not set"))
	}

	addressSuffix := os.Getenv("ADDRESS_SUFFIX")
	if mapsAPIKey == "" {
		exit("missing address suffix", fmt.Errorf("ADDRESS_SUFFIX not set"))
	}

	storageBucket := os.Getenv("STORAGE_BUCKET")
	if storageBucket == "" {
		exit("missing storage bucket name", fmt.Errorf("STORAGE_BUCKET not set"))
	}

	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		exit("missing project id", fmt.Errorf("PROJECT_ID not set"))
	}

	entity := os.Getenv("TARGET_ENTITY")
	if entity == "" {
		exit("no entity name provided", fmt.Errorf("TARGET_ENTITY not set"))
	}

	fmt.Println("Configuring storage client")
	s, err := storage.New(ctx, storageBucket)
	if err != nil {
		exit("unable to setup storage client", err)
	}

	fmt.Println("Configuring datastore client")
	d, err := datastore.New(ctx, projectID)
	if err != nil {
		exit("unable to setup datastore client", err)
	}

	fmt.Println("Configuring location client")
	l, err := location.New(ctx, mapsAPIKey)
	if err != nil {
		exit("unable to setup location client", err)
	}

	fmt.Println("Configuring document processor")
	p, err := processor.New(addressSuffix, l, d, entity)
	if err != nil {
		exit("unable to setup processor", err)
	}

	fmt.Println(s)
	ch, err := checker.New(s.Write, p.ProcessDocument)
	if err != nil {
		exit("unable to setup web poller", err)
	}

	fmt.Println("loading future public meetings")
	pm, err := p.LoadFuturePublicMeetings(ctx, entity)
	if err != nil {
		exit("unable to load known public meetings", err)
	}

	fmt.Println(pm)

	if len(pm) == 1 {
		fmt.Println("Source:", pm[0].Source)
		fmt.Println("Version:", pm[0].Version)
		ch.SetLastTag(pm[0].Version)
	}

	// TODO: make sources more flexible (support both planning commission and city council)
	ch.SetURL(url)

	err = server.New(httpAddr, ch, p)
	if err != nil {
		exit("http server error", err)
	}
}

func exit(msg string, err error) {
	log.Println(msg)
	log.Println(err)
	time.Sleep(time.Second * 1)
	os.Exit(1)
}
