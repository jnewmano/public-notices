package processor

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/jnewmano/public-notices/internal/datastore"
	"github.com/jnewmano/public-notices/internal/location"
	"github.com/jnewmano/public-notices/internal/notice"
	"github.com/jnewmano/public-notices/internal/pdf"
	"github.com/jnewmano/public-notices/internal/tokenize"
)

const storeType = "PublicMeeting"

type Processor struct {
	addressSuffix string
	l             *location.LocationClient
	d             *datastore.DataStore
	entity        string

	meeting PublicMeeting
}

func New(addressSuffix string, l *location.LocationClient, d *datastore.DataStore, entity string) (*Processor, error) {

	p := Processor{
		addressSuffix: addressSuffix,
		l:             l,
		d:             d,
		entity:        entity,
	}

	return &p, nil

}

type PublicMeeting struct {
	Entity string
	Body   string

	Source  string
	Version string

	Date time.Time

	Notices []Notice `datastore:",noindex"`
}

type Notice struct {
	Notice   notice.Notice
	Location *location.Location
}

func (p *Processor) ProcessDocument(ctx context.Context, source string, version string, r io.Reader) error {

	entity := p.entity

	fmt.Println("Processing document", source, version)
	// extract text from the pdf
	fmt.Println("extracting text")
	txt, err := pdf.ExtractText(ctx, source, version, r)
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

	ns := make([]Notice, 0, len(notices))

	date := time.Time{}

	for _, v := range notices {

		if date.IsZero() {
			date = v.Date
		}

		// TODO: add address geo-location information
		if v.Address == nil || v.Address.Location == "" {
			continue
		}

		fmt.Println("checking notice", v.Date, v.Address)
		address := v.Address.Location + p.addressSuffix

		l, err := p.l.AddressLocation(ctx, address)
		if err != nil {
			return err
		}

		n := Notice{
			Notice:   v,
			Location: l,
		}

		ns = append(ns, n)

	}

	m := PublicMeeting{
		Entity: entity,

		Notices: ns,
		Date:    date,

		Source:  source,
		Version: version,
	}

	fmt.Println(m)

	// push the public meeting to offline storage
	key := entity + "_" + m.Date.Format("2006-01-02")
	fmt.Println("Saving public meeting data to:", key)

	// TODO: need a lock to update this
	p.meeting = m

	err = p.d.Put(ctx, storeType, key, &m)
	if err != nil {
		return err
	}

	return nil
}

func (p *Processor) LoadFuturePublicMeetings(ctx context.Context, entity string) ([]PublicMeeting, error) {

	dst := []PublicMeeting{}
	err := p.d.Future(ctx, storeType, entity, &dst)
	if err != nil {
		return nil, err
	}

	if len(dst) == 1 {
		p.meeting = dst[0]
	}

	return dst, nil

}

func (p *Processor) Meeting(entity string) PublicMeeting {
	// TODO: filter with the requested entity

	return p.meeting
}
