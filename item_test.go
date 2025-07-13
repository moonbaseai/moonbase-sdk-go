// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbasesdk_test

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
	client := moonbasesdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Items.New(context.TODO(), moonbasesdk.ItemNewParams{
		CollectionID: "1CR2QLsnhwrJX7Z33jnyGV",
		Values: map[string]moonbasesdk.FieldValueUnionParam{
			"name": {
				OfSingleLineText: &moonbasesdk.SingleLineTextValueParam{
					Text: "Aperture Science",
				},
			},
			"ceo": {
				OfRelation: &moonbasesdk.RelationValueParam{
					Item: moonbasesdk.ItemParam{
						ID: "1CR2QLtx9doK4wFiFB7VAS",
						Values: map[string]moonbasesdk.FieldValueUnionParam{
							"foo": {
								OfSingleLineText: &moonbasesdk.SingleLineTextValueParam{
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
		var apierr *moonbasesdk.Error
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
	client := moonbasesdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Items.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moonbasesdk.Error
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
	client := moonbasesdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Items.Update(
		context.TODO(),
		"id",
		moonbasesdk.ItemUpdateParams{
			Values: map[string]moonbasesdk.FieldValueUnionParam{
				"foo": {
					OfSingleLineText: &moonbasesdk.SingleLineTextValueParam{
						Text: "text",
					},
				},
			},
			UpdateManyStrategy: moonbasesdk.ItemUpdateParamsUpdateManyStrategyReplace,
			UpdateOneStrategy:  moonbasesdk.ItemUpdateParamsUpdateOneStrategyReplace,
		},
	)
	if err != nil {
		var apierr *moonbasesdk.Error
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
	client := moonbasesdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Items.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *moonbasesdk.Error
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
	client := moonbasesdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Items.Upsert(context.TODO(), moonbasesdk.ItemUpsertParams{
		CollectionID: "1CR2QLbeMAqKQ6PvQu39pZ",
		Identifiers: map[string]moonbasesdk.FieldValueUnionParam{
			"domain": {
				OfArrayOfValues: []moonbasesdk.ValueUnionParam{},
			},
		},
		Values: map[string]moonbasesdk.FieldValueUnionParam{
			"name": {
				OfSingleLineText: &moonbasesdk.SingleLineTextValueParam{
					Text: "Aperture Science",
				},
			},
			"domain": {
				OfArrayOfValues: []moonbasesdk.ValueUnionParam{},
			},
			"linked_in": {
				OfLinkedIn: &moonbasesdk.SocialLinkedInValueParam{
					Profile: moonbasesdk.SocialLinkedInValueProfileParam{
						URL:      moonbasesdk.String("https://linkedin.com/company/aperturescience"),
						Username: moonbasesdk.String("username"),
					},
				},
			},
		},
		UpdateManyStrategy: moonbasesdk.ItemUpsertParamsUpdateManyStrategyReplace,
		UpdateOneStrategy:  moonbasesdk.ItemUpsertParamsUpdateOneStrategyReplace,
	})
	if err != nil {
		var apierr *moonbasesdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
