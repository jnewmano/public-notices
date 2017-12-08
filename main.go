package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jnewmano/public-notices/internal/checker"
	"github.com/jnewmano/public-notices/internal/location"
	// "github.com/jnewmano/public-notices/internal/storage"
)

func main() {
	ctx := context.Background()

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

	/*
		fmt.Println("Configuring storage client")
		s, err := storage.New(ctx, "")
		if err != nil {
			exit("unable to setup storage client", err)
		}
	*/

	fmt.Println("Configuring location client")
	l, err := location.New(ctx, mapsAPIKey)
	if err != nil {
		exit("unable to setup location client", err)
	}

	fmt.Println("Configuring document processor")
	p, err := New(addressSuffix, l)
	if err != nil {
		exit("unable to setup processor", err)
	}

	//ch, err := checker.New(s.Write, p.ProcessDocument)
	ch, err := checker.New(p.ProcessDocument)
	if err != nil {
		exit("unable to setup web poller", err)
	}

	// TODO: track last tag and continuously poll
	fmt.Println("Running document check")
	ctx, done := context.WithTimeout(ctx, time.Second*30)
	defer done()

	initialTag := "<>"
	tag, err := ch.Do(ctx, url, initialTag)
	if err != nil {
		exit("unable to check document", err)
	}

	fmt.Println("Last tag was:", tag)

}

func exit(msg string, err error) {
	log.Println(msg)
	log.Println(err)
	time.Sleep(time.Second * 1)
	os.Exit(1)
}
