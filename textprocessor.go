package main

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/jnewmano/public-notices/internal/location"
	"github.com/jnewmano/public-notices/internal/notice"
	"github.com/jnewmano/public-notices/internal/pdf"
	"github.com/jnewmano/public-notices/internal/tokenize"
)

type Processor struct {
	addressSuffix string
	l             *location.LocationClient
}

func New(addressSuffix string, l *location.LocationClient) (*Processor, error) {

	p := Processor{
		addressSuffix: addressSuffix,
		l:             l,
	}

	return &p, nil

}
func (p *Processor) ProcessDocument(ctx context.Context, name string, r io.Reader) error {

	// extract text from the pdf
	fmt.Println("extracting text")
	txt, err := pdf.ExtractText(ctx, name, r)
	if err != nil {
		return err
	}

	r = bytes.NewBuffer([]byte(txt))

	// tokenize the text
	fmt.Println("tokenizing")
	tokens, err := tokenize.Tokenize(ctx, r)
	if err != nil {
		return err
	}

	// extract notices from the text
	fmt.Println("processing notices")
	notices, err := notice.ProcessNotices(ctx, tokens)
	if err != nil {
		return err
	}

	for _, v := range notices {

		// TODO: add address geo-location information
		if v.Address == nil || v.Address.Location == "" {
			continue
		}

		fmt.Println("checking notice", v.Address)
		address := v.Address.Location + p.addressSuffix

		l, err := p.l.AddressLocation(ctx, address)
		if err != nil {
			return err
		}

		fmt.Println(l)
	}

	return nil
}
