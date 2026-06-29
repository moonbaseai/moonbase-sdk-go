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

func TestViewNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.New(context.TODO(), moonbase.ViewNewParams{
		Collection: moonbase.ViewNewParamsCollection{
			ID:  moonbase.String("id"),
			Ref: moonbase.String("people"),
		},
		Fields: []moonbase.ViewFieldParam{{
			Field:         "name",
			DisplayFields: []string{"string"},
			IsPinned:      moonbase.Bool(true),
			IsWrapped:     moonbase.Bool(true),
			Size: moonbase.ViewFieldSizeUnionParam{
				OfViewFieldSizeString: moonbase.String("fit"),
			},
		}, {
			Field:         "email",
			DisplayFields: []string{"string"},
			IsPinned:      moonbase.Bool(true),
			IsWrapped:     moonbase.Bool(true),
			Size: moonbase.ViewFieldSizeUnionParam{
				OfViewFieldSizeString: moonbase.String("fit"),
			},
		}},
		Name:     "Active leads",
		ViewType: moonbase.ViewNewParamsViewTypeTable,
		Aggregates: []moonbase.ViewAggregateUnionParam{{
			OfItemCount: &moonbase.ViewAggregateItemCountParam{
				Group: moonbase.String("group"),
			},
		}},
		Filter: moonbase.ItemsFilterUnionParam{
			OfItemsFilterValueMatches: &moonbase.ItemsFilterValueMatchesParam{
				Field: "name",
				Op:    moonbase.ItemsFilterValueMatchesOpContains,
				Value: moonbase.ItemsFilterValueMatchesValueUnionParam{
					OfString: moonbase.String("Acme"),
				},
			},
		},
		Groups: []string{"string"},
		RelationValueFilters: []moonbase.ViewRelationValueFilterParam{{
			Field: "field",
			Filter: moonbase.ItemsFilterUnionParam{
				OfItemsFilterValueMatches: &moonbase.ItemsFilterValueMatchesParam{
					Field: "field",
					Op:    moonbase.ItemsFilterValueMatchesOpStartsWith,
					Value: moonbase.ItemsFilterValueMatchesValueUnionParam{
						OfString: moonbase.String("string"),
					},
				},
			},
		}},
		Sort: []string{"-name"},
	})
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewGet(t *testing.T) {
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
	_, err := client.Views.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.Update(
		context.TODO(),
		"id",
		moonbase.ViewUpdateParams{
			Aggregates: []moonbase.ViewAggregateUnionParam{{
				OfItemCount: &moonbase.ViewAggregateItemCountParam{
					Group: moonbase.String("stage"),
				},
			}},
			Fields: []moonbase.ViewFieldParam{{
				Field:         "name",
				DisplayFields: []string{"string"},
				IsPinned:      moonbase.Bool(true),
				IsWrapped:     moonbase.Bool(true),
				Size: moonbase.ViewFieldSizeUnionParam{
					OfViewFieldSizeString: moonbase.String("fit"),
				},
			}, {
				Field:         "amount",
				DisplayFields: []string{"string"},
				IsPinned:      moonbase.Bool(true),
				IsWrapped:     moonbase.Bool(true),
				Size: moonbase.ViewFieldSizeUnionParam{
					OfViewFieldSizeString: moonbase.String("fit"),
				},
			}, {
				Field:         "stage",
				DisplayFields: []string{"string"},
				IsPinned:      moonbase.Bool(true),
				IsWrapped:     moonbase.Bool(true),
				Size: moonbase.ViewFieldSizeUnionParam{
					OfViewFieldSizeString: moonbase.String("fit"),
				},
			}, {
				Field:         "owner",
				DisplayFields: []string{"name", "email"},
				IsPinned:      moonbase.Bool(true),
				IsWrapped:     moonbase.Bool(true),
				Size: moonbase.ViewFieldSizeUnionParam{
					OfViewFieldSizeString: moonbase.String("fit"),
				},
			}, {
				Field:         "related_tasks",
				DisplayFields: []string{"string"},
				IsPinned:      moonbase.Bool(true),
				IsWrapped:     moonbase.Bool(true),
				Size: moonbase.ViewFieldSizeUnionParam{
					OfViewFieldSizeString: moonbase.String("fit"),
				},
			}},
			Filter: moonbase.ItemsFilterUnionParam{
				OfItemsFilterValueMatches: &moonbase.ItemsFilterValueMatchesParam{
					Field: "name",
					Op:    moonbase.ItemsFilterValueMatchesOpEq,
					Value: moonbase.ItemsFilterValueMatchesValueUnionParam{
						OfString: moonbase.String("Acme"),
					},
				},
			},
			Groups: []string{"stage"},
			Name:   moonbase.String("Active deals"),
			RelationValueFilters: []moonbase.ViewRelationValueFilterParam{{
				Field: "related_tasks",
				Filter: moonbase.ItemsFilterUnionParam{
					OfItemsFilterValueMatches: &moonbase.ItemsFilterValueMatchesParam{
						Field: "state",
						Op:    moonbase.ItemsFilterValueMatchesOpEq,
						Value: moonbase.ItemsFilterValueMatchesValueUnionParam{
							OfString: moonbase.String("Open"),
						},
					},
				},
			}},
			Sort:     []string{"-name"},
			ViewType: moonbase.ViewUpdateParamsViewTypeTable,
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

func TestViewListWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.List(context.TODO(), moonbase.ViewListParams{
		After:  moonbase.String("after"),
		Before: moonbase.String("before"),
		Limit:  moonbase.Int(1),
	})
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewDelete(t *testing.T) {
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
	err := client.Views.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
