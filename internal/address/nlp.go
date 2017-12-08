// +build none

package notice

/*
import (
	"encoding/json"
	"fmt"
	"log"

	// Imports the Google Cloud Natural Language API client package.
	language "cloud.google.com/go/language/apiv1"
	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {

	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx, option.WithServiceAccountFile("key.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to analyze.
	text := "Public hearing and recommendation of Patterson Construction’s request for a zone change on 0.32-acres of property located at 845 North 500 West changing the zoning from an A-1 to an R-2 zone."

	req := languagepb.AnalyzeSyntaxRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	}

	resp, err := client.AnalyzeSyntax(ctx, &req)
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	pretty.Println(resp)

	out, _ := json.Marshal(resp)
	fmt.Println(string(out))
}
*/
