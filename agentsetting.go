// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// AgentSettingService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentSettingService] method instead.
type AgentSettingService struct {
	Options []option.RequestOption
}

// NewAgentSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentSettingService(opts ...option.RequestOption) (r AgentSettingService) {
	r = AgentSettingService{}
	r.Options = opts
	return
}

func (r *AgentSettingService) Get(ctx context.Context, opts ...option.RequestOption) (res *AgentSettingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent_settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentSettingGetResponse struct {
	CreatedAt             time.Time              `json:"created_at" api:"required" format:"date-time"`
	Type                  constant.AgentSettings `json:"type" api:"required"`
	UpdatedAt             time.Time              `json:"updated_at" api:"required" format:"date-time"`
	DealSummaryModel      string                 `json:"deal_summary_model"`
	DealSummaryPrompt     string                 `json:"deal_summary_prompt"`
	MeetingAgentModel     string                 `json:"meeting_agent_model"`
	MeetingPrebriefPrompt string                 `json:"meeting_prebrief_prompt"`
	MeetingSummaryPrompt  string                 `json:"meeting_summary_prompt"`
	MeetingWebSearch      bool                   `json:"meeting_web_search"`
	OrganizationInfo      string                 `json:"organization_info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt             respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		DealSummaryModel      respjson.Field
		DealSummaryPrompt     respjson.Field
		MeetingAgentModel     respjson.Field
		MeetingPrebriefPrompt respjson.Field
		MeetingSummaryPrompt  respjson.Field
		MeetingWebSearch      respjson.Field
		OrganizationInfo      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentSettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentSettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
