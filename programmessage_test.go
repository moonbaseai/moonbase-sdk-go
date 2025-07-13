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

func TestProgramMessageSendWithOptionalParams(t *testing.T) {
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
	_, err := client.ProgramMessages.Send(context.TODO(), moonbasesdk.ProgramMessageSendParams{
		Person: moonbasesdk.ProgramMessageSendParamsPerson{
			Email: "person-71@example-71.com",
		},
		ProgramTemplateID: "1CR2QLhLQPX9WCiWH3cXCe",
		CustomVariables: map[string]any{
			"coupon_code": "bar",
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
