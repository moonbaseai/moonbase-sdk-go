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

func TestMeetingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Meetings.Get(
		context.TODO(),
		"id",
		moonbase.MeetingGetParams{
			Include: []string{"organizer"},
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

func TestMeetingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Meetings.Update(
		context.TODO(),
		"id",
		moonbase.MeetingUpdateParams{
			Recording: moonbase.MeetingUpdateParamsRecording{
				ContentType: "video/mp4",
				ProviderID:  "abc123",
				URL:         "https://example.com/recording.mp4",
			},
			Tags: []shared.TagPointerParam{{
				ID: "1CLJt2vJy3SZLhqYW8rQoN",
			}},
			Transcript: moonbase.MeetingUpdateParamsTranscript{
				Cues: []moonbase.MeetingUpdateParamsTranscriptCue{{
					From:    0.71999997,
					Speaker: "Jony Appleseed",
					Text:    "Hello.",
					To:      1.22,
				}, {
					From:    1.52,
					Speaker: "Jane Doe",
					Text:    "Hey! It's been too long.",
					To:      3.22,
				}},
				Provider:   "example",
				ProviderID: "def456",
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

func TestMeetingListWithOptionalParams(t *testing.T) {
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
	_, err := client.Meetings.List(context.TODO(), moonbase.MeetingListParams{
		After:  moonbase.String("after"),
		Before: moonbase.String("before"),
		ICalUid: moonbase.MeetingListParamsICalUid{
			Eq: moonbase.String("eq"),
		},
		Limit: moonbase.Int(1),
	})
	if err != nil {
		var apierr *moonbase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
