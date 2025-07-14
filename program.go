// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/apiquery"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/pagination"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// ProgramService contains methods and other services that help with interacting
// with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProgramService] method instead.
type ProgramService struct {
	Options []option.RequestOption
}

// NewProgramService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProgramService(opts ...option.RequestOption) (r ProgramService) {
	r = ProgramService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing program.
func (r *ProgramService) Get(ctx context.Context, id string, query ProgramGetParams, opts ...option.RequestOption) (res *Program, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("programs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of your marketing programs.
func (r *ProgramService) List(ctx context.Context, query ProgramListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Program], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "programs"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of your marketing programs.
func (r *ProgramService) ListAutoPaging(ctx context.Context, query ProgramListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Program] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// The Program object represents an email campaign. It defines the sending behavior
// and tracks engagement metrics.
type Program struct {
	// Unique identifier for the object.
	ID    string       `json:"id,required"`
	Links ProgramLinks `json:"links,required"`
	// The current status of the program. Can be `draft`, `published`, `paused`, or
	// `archived`.
	//
	// Any of "draft", "published", "paused", "archived".
	Status ProgramStatus `json:"status,required"`
	// The sending trigger for the program. Can be `api` for transactional sends or
	// `broadcast` for scheduled sends.
	//
	// Any of "api", "broadcast".
	Trigger ProgramTrigger `json:"trigger,required"`
	// String representing the object’s type. Always `program` for this object.
	Type constant.Program `json:"type,required"`
	// A `ProgramActivityMetrics` object summarizing engagement for this program.
	ActivityMetrics ProgramActivityMetrics `json:"activity_metrics"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The user-facing name of the program.
	DisplayName string `json:"display_name"`
	// The `ProgramTemplate` used for messages in this program.
	ProgramTemplate ProgramTemplate `json:"program_template"`
	// For `broadcast` programs, the time the program is scheduled to send, as an RFC
	// 3339 timestamp.
	ScheduledAt time.Time `json:"scheduled_at" format:"date-time"`
	// `true` if link clicks are tracked for this program.
	TrackClicks bool `json:"track_clicks"`
	// `true` if email opens are tracked for this program.
	TrackOpens bool `json:"track_opens"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Links           respjson.Field
		Status          respjson.Field
		Trigger         respjson.Field
		Type            respjson.Field
		ActivityMetrics respjson.Field
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		ProgramTemplate respjson.Field
		ScheduledAt     respjson.Field
		TrackClicks     respjson.Field
		TrackOpens      respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Program) RawJSON() string { return r.JSON.raw }
func (r *Program) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProgramLinks struct {
	// A link to the `ProgramTemplate` for this program.
	ProgramTemplate string `json:"program_template,required" format:"uri"`
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProgramTemplate respjson.Field
		Self            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProgramLinks) RawJSON() string { return r.JSON.raw }
func (r *ProgramLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the program. Can be `draft`, `published`, `paused`, or
// `archived`.
type ProgramStatus string

const (
	ProgramStatusDraft     ProgramStatus = "draft"
	ProgramStatusPublished ProgramStatus = "published"
	ProgramStatusPaused    ProgramStatus = "paused"
	ProgramStatusArchived  ProgramStatus = "archived"
)

// The sending trigger for the program. Can be `api` for transactional sends or
// `broadcast` for scheduled sends.
type ProgramTrigger string

const (
	ProgramTriggerAPI       ProgramTrigger = "api"
	ProgramTriggerBroadcast ProgramTrigger = "broadcast"
)

// A `ProgramActivityMetrics` object summarizing engagement for this program.
type ProgramActivityMetrics struct {
	// The number of emails that could not be delivered.
	Bounced int64 `json:"bounced"`
	// The number of recipients who clicked at least one link.
	Clicked int64 `json:"clicked"`
	// The number of recipients who marked the email as spam.
	Complained int64 `json:"complained"`
	// The number of emails that failed to send due to a technical issue.
	Failed int64 `json:"failed"`
	// The number of recipients who opened the email.
	Opened int64 `json:"opened"`
	// The total number of emails successfully sent.
	Sent int64 `json:"sent"`
	// The number of emails blocked by delivery protection rules.
	Shielded int64 `json:"shielded"`
	// The number of recipients who unsubscribed.
	Unsubscribed int64 `json:"unsubscribed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced      respjson.Field
		Clicked      respjson.Field
		Complained   respjson.Field
		Failed       respjson.Field
		Opened       respjson.Field
		Sent         respjson.Field
		Shielded     respjson.Field
		Unsubscribed respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProgramActivityMetrics) RawJSON() string { return r.JSON.raw }
func (r *ProgramActivityMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProgramGetParams struct {
	// Specifies which related objects to include in the response. Valid options are
	// `activity_metrics` and `program_template`.
	//
	// Any of "activity_metrics", "program_template".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProgramGetParams]'s query parameters as `url.Values`.
func (r ProgramGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProgramListParams struct {
	// When specified, returns results starting immediately after the item identified
	// by this cursor. Use the cursor value from the previous response's metadata to
	// fetch the next page of results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// When specified, returns results starting immediately before the item identified
	// by this cursor. Use the cursor value from the response's metadata to fetch the
	// previous page of results.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return per page. Must be between 1 and 100. Defaults
	// to 20 if not specified.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProgramListParams]'s query parameters as `url.Values`.
func (r ProgramListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
