// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase_test

import (
	"context"
	"os"
	"testing"

	"github.com/moonbaseai/moonbase-sdk-go"
	"github.com/moonbaseai/moonbase-sdk-go/internal/testutil"
	"github.com/moonbaseai/moonbase-sdk-go/option"
)

func TestAutoPagination(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moonbase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	iter := client.Collections.Items.ListAutoPaging(
		context.TODO(),
		"people",
		moonbase.CollectionItemListParams{
			Limit: moonbase.Int(5),
		},
	)
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		item := iter.Current()
		t.Logf("%+v\n", item.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
