package notice

import (
	"context"
	"fmt"
	"regexp"
	"strings"
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

	meetingDate := time.Time{}

	parsed := make([]Notice, 0, len(notices))

	for i, v := range notices {
		a, err := address.ExtractAddress(ctx, v)
		if err != nil {
			fmt.Printf("match error [%s] [%s]\n", err, v)
			if meetingDate.IsZero() == false {
				continue
			}

			meetingDate, _ = parseDate(ctx, v)
			continue
		}

		n := Notice{
			Date:    meetingDate,
			Address: a,
			Line:    i,
		}

		parsed = append(parsed, n)
	}

	return parsed, nil

}

var regexDate = regexp.MustCompile(`On ([A-Z][a-z]+ [0-9]{1,2}[a-z]{1,2}, [0-9]{4})`)

func parseDate(ctx context.Context, line string) (time.Time, error) {

	m := regexDate.FindAllStringSubmatch(line, -1)
	if len(m) == 0 {
		return time.Time{}, fmt.Errorf("no date matches")
	}

	// convert the string date to a time.Time
	d := m[0][1]

	r := strings.NewReplacer("st, ", " ", "nd, ", " ", "rd, ", " ", "th, ", " ")
	d = r.Replace(d)

	t, err := time.Parse("January 2 2006", d)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
