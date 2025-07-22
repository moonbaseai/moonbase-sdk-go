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
)

func TestItemNew(t *testing.T) {
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
	_, err := client.Items.New(context.TODO(), moonbase.ItemNewParams{
		CollectionID: "1CRMPEPBnmyMYXMvKr9hg7",
		Values: map[string]moonbase.FieldValueUnionParam{
			"name": {
				OfSingleLineText: &moonbase.SingleLineTextValueParam{
					Text: "Aperture Science",
				},
			},
			"ceo": {
				OfRelation: &moonbase.RelationValueParam{
					Item: moonbase.ItemParam{
						ID: "1CRMPEQMz9G7AbnjP4j1mr",
						Links: moonbase.ItemLinksParam{
							Collection: moonbase.String("https://example.com"),
							Self:       moonbase.String("https://example.com"),
						},
						Values: map[string]moonbase.FieldValueUnionParam{
							"foo": {
								OfSingleLineText: &moonbase.SingleLineTextValueParam{
									Text: "text",
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestItemGet(t *testing.T) {
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
	_, err := client.Items.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestItemUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Items.Update(
		context.TODO(),
		"id",
		moonbase.ItemUpdateParams{
			Values: map[string]moonbase.FieldValueUnionParam{
				"name": {
					OfSingleLineText: &moonbase.SingleLineTextValueParam{
						Text: "Jony Appleseed",
					},
				},
			},
			UpdateManyStrategy: moonbase.ItemUpdateParamsUpdateManyStrategyReplace,
			UpdateOneStrategy:  moonbase.ItemUpdateParamsUpdateOneStrategyReplace,
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

func TestItemDelete(t *testing.T) {
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
	_, err := client.Items.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestItemUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Items.Upsert(context.TODO(), moonbase.ItemUpsertParams{
		CollectionID: "1CRMPDfqNedouo8m9fNbc3",
		Identifiers: map[string]moonbase.FieldValueUnionParam{
			"domain": {
				OfArrayOfValues: []moonbase.ValueUnionParam{},
			},
		},
		Values: map[string]moonbase.FieldValueUnionParam{
			"name": {
				OfSingleLineText: &moonbase.SingleLineTextValueParam{
					Text: "Aperture Science",
				},
			},
			"domain": {
				OfArrayOfValues: []moonbase.ValueUnionParam{},
			},
			"linked_in": {
				OfLinkedIn: &moonbase.SocialLinkedInValueParam{
					Profile: moonbase.SocialLinkedInValueProfileParam{
						URL:      moonbase.String("https://linkedin.com/company/aperturescience"),
						Username: moonbase.String("username"),
					},
				},
			},
		},
		UpdateManyStrategy: moonbase.ItemUpsertParamsUpdateManyStrategyReplace,
		UpdateOneStrategy:  moonbase.ItemUpsertParamsUpdateOneStrategyReplace,
	})
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
