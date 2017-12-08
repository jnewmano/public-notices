package pdf

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	ctx := context.Background()

	f, err := os.Open("./sample.pdf")
	if err != nil {
		t.Fatalf("unable to open sample pdf: %s\n", err)
	}
	defer f.Close()

	s, err := ExtractText(ctx, "sample.pdf", f)
	if err != nil {
		t.Fatalf("unable to extract text: %s\n", err)
	}

	fmt.Println(s)

}
