package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jnewmano/public-notices/internal/checker"
	"github.com/jnewmano/public-notices/internal/pdf"
	"github.com/jnewmano/public-notices/internal/storage"
)

func main() {
	ctx := context.Background()

	s, err := storage.New(ctx, "")
	if err != nil {
		exit(err)
	}

	textProcessor := func(ctx context.Context, name string, r io.Reader) error {
		txt, err := pdf.ExtractText(ctx, name, r)
		if err != nil {
			return err
		}

		fmt.Println(txt)

		// TODO: analyze the text body
		//       pull out some descriptive information

		return nil
	}

	ch, err := checker.New(s.Write, textProcessor)
	if err != nil {
		exit(err)
	}

	fmt.Println(ch)
}

func exit(err error) {
	log.Println(err)
	time.Sleep(time.Second * 1)
	os.Exit(1)
}
