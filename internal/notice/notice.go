package notice

import (
	"context"
	"fmt"
	"time"

	"github.com/jnewmano/public-notices/internal/address"
)

type Notice struct {
	Date    time.Time
	Action  string
	Address *address.Address

	Line int
}

func ProcessNotices(ctx context.Context, notices []string) ([]Notice, error) {

	n, err := parseNotices(ctx, notices)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func parseNotices(ctx context.Context, notices []string) ([]Notice, error) {

	parsed := make([]Notice, 0, len(notices))

	for i, v := range notices {
		a, err := address.ExtractAddress(ctx, v)
		if err != nil {
			fmt.Printf("match error [%s] [%s]\n", err, v)
			continue
		}

		n := Notice{
			Address: a,
			Line:    i,
		}

		parsed = append(parsed, n)
	}

	return parsed, nil

}
