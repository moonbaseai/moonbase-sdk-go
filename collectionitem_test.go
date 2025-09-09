// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/moonbaseai/moonbase-sdk-go"
	"github.com/moonbaseai/moonbase-sdk-go/internal/testutil"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/shared"
)

func TestCollectionItemNew(t *testing.T) {
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
	_, err := client.Collections.Items.New(
		context.TODO(),
		"collection_id",
		moonbase.CollectionItemNewParams{
			Values: map[string]moonbase.FieldValueParamUnion{
				"name": {
					OfSingleLineText: &moonbase.SingleLineTextValueParam{
						Data: "Aperture Science",
					},
				},
				"ceo": {
					OfRelation: &moonbase.RelationValueParam{
						Item: moonbase.RelationValueParamItemUnion{
							OfPointer: &shared.PointerParam{
								ID:   "1CLJt2v2rARKGD4MLziBCw",
								Type: "item",
							},
						},
					},
				},
			},
		},
	)
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionItemGet(t *testing.T) {
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
	_, err := client.Collections.Items.Get(
		context.TODO(),
		"id",
		moonbase.CollectionItemGetParams{
			CollectionID: "collection_id",
		},
	)
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionItemUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Collections.Items.Update(
		context.TODO(),
		"id",
		moonbase.CollectionItemUpdateParams{
			CollectionID: "collection_id",
			Values: map[string]moonbase.FieldValueParamUnion{
				"name": {
					OfSingleLineText: &moonbase.SingleLineTextValueParam{
						Data: "Jony Appleseed",
					},
				},
			},
			UpdateManyStrategy: moonbase.CollectionItemUpdateParamsUpdateManyStrategyReplace,
			UpdateOneStrategy:  moonbase.CollectionItemUpdateParamsUpdateOneStrategyReplace,
		},
	)
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionItemListWithOptionalParams(t *testing.T) {
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
	_, err := client.Collections.Items.List(
		context.TODO(),
		"collection_id",
		moonbase.CollectionItemListParams{
			After:  moonbase.String("after"),
			Before: moonbase.String("before"),
			Limit:  moonbase.Int(1),
		},
	)
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionItemDelete(t *testing.T) {
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
	err := client.Collections.Items.Delete(
		context.TODO(),
		"id",
		moonbase.CollectionItemDeleteParams{
			CollectionID: "collection_id",
		},
	)
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionItemUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Collections.Items.Upsert(
		context.TODO(),
		"collection_id",
		moonbase.CollectionItemUpsertParams{
			Identifiers: map[string]moonbase.FieldValueParamUnion{
				"domain": {
					OfArrayOfValues: []moonbase.ValueParamUnion{{
						OfValueUriDomain: &moonbase.DomainValueParam{
							Data: "aperturescience.com",
						},
					}},
				},
			},
			Values: map[string]moonbase.FieldValueParamUnion{
				"name": {
					OfSingleLineText: &moonbase.SingleLineTextValueParam{
						Data: "Aperture Science",
					},
				},
				"domain": {
					OfArrayOfValues: []moonbase.ValueParamUnion{{
						OfValueUriDomain: &moonbase.DomainValueParam{
							Data: "aperturescience.com",
						},
					}},
				},
				"linked_in": {
					OfLinkedIn: &moonbase.FieldValueParamLinkedIn{
						Data: moonbase.FieldValueParamLinkedInData{
							URL:      moonbase.String("https://linkedin.com/company/aperturescience"),
							Username: moonbase.String("company/moonbaseai"),
						},
					},
				},
			},
			UpdateManyStrategy: moonbase.CollectionItemUpsertParamsUpdateManyStrategyReplace,
			UpdateOneStrategy:  moonbase.CollectionItemUpsertParamsUpdateOneStrategyReplace,
		},
	)
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
