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

func TestInboxMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.InboxMessages.New(context.TODO(), moonbase.InboxMessageNewParams{
		Body: shared.FormattedTextParam{
			Markdown: moonbase.String("This is the body of the message. It supports [markdown](https://en.wikipedia.org/wiki/Markdown)."),
		},
		InboxID: "1CLJt2v6KXDyzDuM57pQqo",
		Bcc: []moonbase.InboxMessageNewParamsBcc{{
			Email: "steve@example.com",
			Name:  moonbase.String("Steve"),
		}},
		Cc: []moonbase.InboxMessageNewParamsCc{{
			Email: "joe@example.com",
			Name:  moonbase.String("Joe"),
		}},
		ConversationID: moonbase.String("conversation_id"),
		Subject:        moonbase.String("Test Subject"),
		To: []moonbase.InboxMessageNewParamsTo{{
			Email: "bob@example.com",
			Name:  moonbase.String("Bob"),
		}, {
			Email: "jack@example.com",
			Name:  moonbase.String("name"),
		}},
	})
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboxMessageGetWithOptionalParams(t *testing.T) {
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
	_, err := client.InboxMessages.Get(
		context.TODO(),
		"id",
		moonbase.InboxMessageGetParams{
			Include: []string{"addresses"},
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

func TestInboxMessageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.InboxMessages.Update(
		context.TODO(),
		"id",
		moonbase.InboxMessageUpdateParams{
			LockVersion: 0,
			Bcc: []moonbase.InboxMessageUpdateParamsBcc{{
				Email: "steve@example.com",
				Name:  moonbase.String("Steve"),
			}},
			Body: shared.FormattedTextParam{
				Markdown: moonbase.String("This is the body of the message. It supports [markdown](https://en.wikipedia.org/wiki/Markdown)."),
			},
			Cc: []moonbase.InboxMessageUpdateParamsCc{{
				Email: "joe@example.com",
				Name:  moonbase.String("Joe"),
			}},
			Subject: moonbase.String("Test Subject"),
			To: []moonbase.InboxMessageUpdateParamsTo{{
				Email: "bob@example.com",
				Name:  moonbase.String("Bob"),
			}, {
				Email: "jack@example.com",
				Name:  moonbase.String("name"),
			}},
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

func TestInboxMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.InboxMessages.List(context.TODO(), moonbase.InboxMessageListParams{
		After:  moonbase.String("after"),
		Before: moonbase.String("before"),
		Filter: moonbase.InboxMessageListParamsFilter{
			ConversationID: moonbase.InboxMessageListParamsFilterConversationID{
				Eq: moonbase.String("eq"),
			},
			InboxID: moonbase.InboxMessageListParamsFilterInboxID{
				Eq: moonbase.String("eq"),
			},
		},
		Include: []string{"addresses"},
		Limit:   moonbase.Int(1),
	})
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboxMessageDelete(t *testing.T) {
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
	err := client.InboxMessages.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
