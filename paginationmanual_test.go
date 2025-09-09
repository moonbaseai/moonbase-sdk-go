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

func TestManualPagination(t *testing.T) {
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
	people, err := client.Collections.Items.List(
		context.TODO(),
		"people",
		moonbase.CollectionItemListParams{
			Limit: moonbase.Int(5),
		},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, item := range people.Data {
		t.Logf("%+v\n", item.ID)
	}
	// Prism mock isn't going to give us real pagination
	people, err = people.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if people != nil {
		for _, item := range people.Data {
			t.Logf("%+v\n", item.ID)
		}
	}
}
