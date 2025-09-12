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
		Provider:       "openphone",
		ProviderID:     "openphone_id_000000000003",
		ProviderStatus: "completed",
		StartAt:        time.Now(),
		AnsweredAt:     moonbase.Time(time.Now()),
		EndAt:          moonbase.Time(time.Now()),
		ProviderMetadata: map[string]any{
			"answered_by":     "bar",
			"user_id":         "bar",
			"phone_number_id": "bar",
			"conversation_id": "bar",
		},
		Recordings: []moonbase.CallNewParamsRecording{{
			ContentType: "content_type",
			ProviderID:  "provider_id",
			URL:         "https://example.com",
		}},
		Transcript: moonbase.CallNewParamsTranscript{
			Cues: []moonbase.CallNewParamsTranscriptCue{{
				From:    0,
				Speaker: "speaker",
				Text:    "text",
				To:      0,
			}},
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

func TestCallUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Upsert(context.TODO(), moonbase.CallUpsertParams{
		Direction: moonbase.CallUpsertParamsDirectionIncoming,
		Participants: []moonbase.CallUpsertParamsParticipant{{
			Phone: "+14155551212",
			Role:  "caller",
		}, {
			Phone: "+16505551212",
			Role:  "callee",
		}},
		Provider:       "openphone",
		ProviderID:     "openphone_id_000000000006",
		ProviderStatus: "completed",
		StartAt:        time.Now(),
		AnsweredAt:     moonbase.Time(time.Now()),
		EndAt:          moonbase.Time(time.Now()),
		ProviderMetadata: map[string]any{
			"answered_by":     "bar",
			"user_id":         "bar",
			"phone_number_id": "bar",
			"conversation_id": "bar",
		},
		Recordings: []moonbase.CallUpsertParamsRecording{{
			ContentType: "content_type",
			ProviderID:  "provider_id",
			URL:         "https://example.com",
		}},
		Transcript: moonbase.CallUpsertParamsTranscript{
			Cues: []moonbase.CallUpsertParamsTranscriptCue{{
				From:    0,
				Speaker: "speaker",
				Text:    "text",
				To:      0,
			}},
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
