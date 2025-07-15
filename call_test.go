// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go"
	"github.com/moonbaseai/moonbase-sdk-go/internal/testutil"
	"github.com/moonbaseai/moonbase-sdk-go/option"
)

func TestCallNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.New(context.TODO(), moonbase.CallNewParams{
		Direction: moonbase.CallNewParamsDirectionIncoming,
		Participants: []moonbase.CallNewParamsParticipant{{
			Phone: "+14155551212",
			Role:  "caller",
		}, {
			Phone: "+16505551212",
			Role:  "callee",
		}},
		Provider:   "openphone",
		ProviderID: "openphone_id_000000000002",
		StartAt:    time.Now(),
		Status:     moonbase.CallNewParamsStatusCompleted,
		AnsweredAt: moonbase.Time(time.Now()),
		EndAt:      moonbase.Time(time.Now()),
		ProviderMetadata: map[string]any{
			"answered_by":     "bar",
			"user_id":         "bar",
			"phone_number_id": "bar",
			"conversation_id": "bar",
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
