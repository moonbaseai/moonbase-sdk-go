// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/apiquery"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/pagination"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// CallService contains methods and other services that help with interacting with
// the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallService] method instead.
type CallService struct {
	Options []option.RequestOption
}

// NewCallService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCallService(opts ...option.RequestOption) (r CallService) {
	r = CallService{}
	r.Options = opts
	return
}

// Logs a phone call.
func (r *CallService) New(ctx context.Context, body CallNewParams, opts ...option.RequestOption) (res *Call, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "calls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing call.
func (r *CallService) Get(ctx context.Context, id string, query CallGetParams, opts ...option.RequestOption) (res *Call, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of calls.
func (r *CallService) List(ctx context.Context, query CallListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Call], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "calls"
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

// Returns a list of calls.
func (r *CallService) ListAutoPaging(ctx context.Context, query CallListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Call] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Find and update an existing phone call, or create a new one.
func (r *CallService) Upsert(ctx context.Context, body CallUpsertParams, opts ...option.RequestOption) (res *Call, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "calls/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The Call object represents a phone call that has been logged in the system. It
// contains details about the participants, timing, and outcome of the call.
type Call struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The direction of the call, either `incoming` or `outgoing`.
	//
	// Any of "incoming", "outgoing".
	Direction CallDirection `json:"direction" api:"required"`
	// The participants involved in the call.
	Participants []CallParticipant `json:"participants" api:"required"`
	// The name of the phone provider that handled the call.
	//
	// Any of "openphone", "user", "zoom_phone".
	Provider CallProvider `json:"provider" api:"required"`
	// The unique identifier for the call from the provider's system.
	ProviderID string `json:"provider_id" api:"required"`
	// The current status of the call.
	ProviderStatus string `json:"provider_status" api:"required"`
	// The time the call started, as an ISO 8601 timestamp in UTC.
	StartAt time.Time `json:"start_at" api:"required" format:"date-time"`
	// String representing the object’s type. Always `call` for this object.
	Type constant.Call `json:"type" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The time the call was answered, if available, as an ISO 8601 timestamp in UTC.
	AnsweredAt time.Time `json:"answered_at" format:"date-time"`
	// The time the call ended, if available, as an ISO 8601 timestamp in UTC.
	EndAt time.Time `json:"end_at" format:"date-time"`
	// The Note object represents a block of text content, often used for meeting notes
	// or summaries.
	Note Note `json:"note"`
	// A hash of additional metadata from the provider.
	ProviderMetadata map[string]any `json:"provider_metadata"`
	// The Note object represents a block of text content, often used for meeting notes
	// or summaries.
	Summary    Note           `json:"summary"`
	Transcript CallTranscript `json:"transcript" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Direction        respjson.Field
		Participants     respjson.Field
		Provider         respjson.Field
		ProviderID       respjson.Field
		ProviderStatus   respjson.Field
		StartAt          respjson.Field
		Type             respjson.Field
		UpdatedAt        respjson.Field
		AnsweredAt       respjson.Field
		EndAt            respjson.Field
		Note             respjson.Field
		ProviderMetadata respjson.Field
		Summary          respjson.Field
		Transcript       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Call) RawJSON() string { return r.JSON.raw }
func (r *Call) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the call, either `incoming` or `outgoing`.
type CallDirection string

const (
	CallDirectionIncoming CallDirection = "incoming"
	CallDirectionOutgoing CallDirection = "outgoing"
)

// Represents a participant in a call.
type CallParticipant struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// The E.164 formatted phone number of the participant.
	Phone string `json:"phone" api:"required"`
	// The role of the participant in the call. Can be `caller`, `callee`, or `other`.
	//
	// Any of "caller", "callee", "other".
	Role string `json:"role" api:"required"`
	// String representing the object’s type. Always `call_participant` for this
	// object.
	Type constant.CallParticipant `json:"type" api:"required"`
	// A lightweight reference to another resource.
	Organization shared.Pointer `json:"organization"`
	// A lightweight reference to another resource.
	Person shared.Pointer `json:"person"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Phone        respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		Organization respjson.Field
		Person       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallParticipant) RawJSON() string { return r.JSON.raw }
func (r *CallParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The name of the phone provider that handled the call.
type CallProvider string

const (
	CallProviderOpenphone CallProvider = "openphone"
	CallProviderUser      CallProvider = "user"
	CallProviderZoomPhone CallProvider = "zoom_phone"
)

type CallTranscript struct {
	Cues []CallTranscriptCue `json:"cues" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cues        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallTranscript) RawJSON() string { return r.JSON.raw }
func (r *CallTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallTranscriptCue struct {
	From    float64                  `json:"from" api:"required"`
	Speaker CallTranscriptCueSpeaker `json:"speaker" api:"required"`
	Text    string                   `json:"text" api:"required"`
	To      float64                  `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Speaker     respjson.Field
		Text        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallTranscriptCue) RawJSON() string { return r.JSON.raw }
func (r *CallTranscriptCue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallTranscriptCueSpeaker struct {
	AttendeeID string `json:"attendee_id"`
	Label      string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttendeeID  respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallTranscriptCueSpeaker) RawJSON() string { return r.JSON.raw }
func (r *CallTranscriptCueSpeaker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallNewParams struct {
	// The direction of the call, either `incoming` or `outgoing`.
	//
	// Any of "incoming", "outgoing".
	Direction CallNewParamsDirection `json:"direction,omitzero" api:"required"`
	// An array of participants involved in the call.
	Participants []CallNewParamsParticipant `json:"participants,omitzero" api:"required"`
	// The name of the phone provider that handled the call (e.g., `openphone`).
	//
	// Any of "openphone", "user", "zoom_phone".
	Provider CallNewParamsProvider `json:"provider,omitzero" api:"required"`
	// The unique identifier for the call from the provider's system.
	ProviderID string `json:"provider_id" api:"required"`
	// The status of the call.
	ProviderStatus string `json:"provider_status" api:"required"`
	// The time the call started, as an ISO 8601 timestamp in UTC.
	StartAt time.Time `json:"start_at" api:"required" format:"date-time"`
	// The time the call was answered, as an ISO 8601 timestamp in UTC.
	AnsweredAt param.Opt[time.Time] `json:"answered_at,omitzero" format:"date-time"`
	// The time the call ended, as an ISO 8601 timestamp in UTC.
	EndAt param.Opt[time.Time] `json:"end_at,omitzero" format:"date-time"`
	// A hash of additional metadata from the provider.
	ProviderMetadata map[string]any `json:"provider_metadata,omitzero"`
	// Any recordings associated with the call.
	Recordings []CallNewParamsRecording `json:"recordings,omitzero"`
	// A transcript of the call.
	Transcript CallNewParamsTranscript `json:"transcript,omitzero"`
	paramObj
}

func (r CallNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CallNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the call, either `incoming` or `outgoing`.
type CallNewParamsDirection string

const (
	CallNewParamsDirectionIncoming CallNewParamsDirection = "incoming"
	CallNewParamsDirectionOutgoing CallNewParamsDirection = "outgoing"
)

// Parameters for creating a `Participant` object.
//
// The properties Phone, Role are required.
type CallNewParamsParticipant struct {
	// The E.164 formatted phone number of the participant.
	Phone string `json:"phone" api:"required"`
	// The role of the participant in the call. Can be `caller`, `callee`, or `other`.
	//
	// Any of "caller", "callee", "other".
	Role string `json:"role,omitzero" api:"required"`
	paramObj
}

func (r CallNewParamsParticipant) MarshalJSON() (data []byte, err error) {
	type shadow CallNewParamsParticipant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallNewParamsParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallNewParamsParticipant](
		"role", "caller", "callee", "other",
	)
}

// The name of the phone provider that handled the call (e.g., `openphone`).
type CallNewParamsProvider string

const (
	CallNewParamsProviderOpenphone CallNewParamsProvider = "openphone"
	CallNewParamsProviderUser      CallNewParamsProvider = "user"
	CallNewParamsProviderZoomPhone CallNewParamsProvider = "zoom_phone"
)

// Parameters for creating a `CallRecording` object.
//
// The properties ContentType, ProviderID, URL are required.
type CallNewParamsRecording struct {
	// The content type of the recording. Note that only `audio/mpeg` is supported at
	// this time.
	//
	// Any of "audio/mpeg".
	ContentType string `json:"content_type,omitzero" api:"required"`
	// The unique identifier for the recording from the provider's system.
	ProviderID string `json:"provider_id" api:"required"`
	// The URL pointing to the recording.
	URL string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r CallNewParamsRecording) MarshalJSON() (data []byte, err error) {
	type shadow CallNewParamsRecording
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallNewParamsRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallNewParamsRecording](
		"content_type", "audio/mpeg",
	)
}

// A transcript of the call.
//
// The property Cues is required.
type CallNewParamsTranscript struct {
	// A list of cues that identify the text spoken in specific time slices of the
	// call.
	Cues []CallNewParamsTranscriptCue `json:"cues,omitzero" api:"required"`
	paramObj
}

func (r CallNewParamsTranscript) MarshalJSON() (data []byte, err error) {
	type shadow CallNewParamsTranscript
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallNewParamsTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating a `CallTranscriptCue` object to capture the text spoken
// in a specific time slice.
//
// The properties From, Speaker, Text, To are required.
type CallNewParamsTranscriptCue struct {
	// The start time of the slice, in fractional seconds from the start of the call.
	From float64 `json:"from" api:"required"`
	// The E.164 formatted phone number of the speaker.
	Speaker string `json:"speaker" api:"required"`
	// The text spoken during the slice.
	Text string `json:"text" api:"required"`
	// The end time of the slice, in fractional seconds from the start of the call.
	To float64 `json:"to" api:"required"`
	paramObj
}

func (r CallNewParamsTranscriptCue) MarshalJSON() (data []byte, err error) {
	type shadow CallNewParamsTranscriptCue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallNewParamsTranscriptCue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetParams struct {
	// Specifies which related objects to include in the response. Valid options are
	// `transcript`, `note`, and `summary`.
	//
	// Any of "transcript", "note", "summary".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallGetParams]'s query parameters as `url.Values`.
func (r CallGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CallListParams struct {
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

// URLQuery serializes [CallListParams]'s query parameters as `url.Values`.
func (r CallListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CallUpsertParams struct {
	// The direction of the call, either `incoming` or `outgoing`.
	//
	// Any of "incoming", "outgoing".
	Direction CallUpsertParamsDirection `json:"direction,omitzero" api:"required"`
	// An array of participants involved in the call.
	Participants []CallUpsertParamsParticipant `json:"participants,omitzero" api:"required"`
	// The name of the phone provider that handled the call (e.g., `openphone`).
	//
	// Any of "openphone", "user", "zoom_phone".
	Provider CallUpsertParamsProvider `json:"provider,omitzero" api:"required"`
	// The unique identifier for the call from the provider's system.
	ProviderID string `json:"provider_id" api:"required"`
	// The status of the call.
	ProviderStatus string `json:"provider_status" api:"required"`
	// The time the call started, as an ISO 8601 timestamp in UTC.
	StartAt time.Time `json:"start_at" api:"required" format:"date-time"`
	// The time the call was answered, as an ISO 8601 timestamp in UTC.
	AnsweredAt param.Opt[time.Time] `json:"answered_at,omitzero" format:"date-time"`
	// The time the call ended, as an ISO 8601 timestamp in UTC.
	EndAt param.Opt[time.Time] `json:"end_at,omitzero" format:"date-time"`
	// A hash of additional metadata from the provider.
	ProviderMetadata map[string]any `json:"provider_metadata,omitzero"`
	// Any recordings associated with the call.
	Recordings []CallUpsertParamsRecording `json:"recordings,omitzero"`
	// A transcript of the call.
	Transcript CallUpsertParamsTranscript `json:"transcript,omitzero"`
	paramObj
}

func (r CallUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow CallUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the call, either `incoming` or `outgoing`.
type CallUpsertParamsDirection string

const (
	CallUpsertParamsDirectionIncoming CallUpsertParamsDirection = "incoming"
	CallUpsertParamsDirectionOutgoing CallUpsertParamsDirection = "outgoing"
)

// Parameters for creating a `Participant` object.
//
// The properties Phone, Role are required.
type CallUpsertParamsParticipant struct {
	// The E.164 formatted phone number of the participant.
	Phone string `json:"phone" api:"required"`
	// The role of the participant in the call. Can be `caller`, `callee`, or `other`.
	//
	// Any of "caller", "callee", "other".
	Role string `json:"role,omitzero" api:"required"`
	paramObj
}

func (r CallUpsertParamsParticipant) MarshalJSON() (data []byte, err error) {
	type shadow CallUpsertParamsParticipant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallUpsertParamsParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallUpsertParamsParticipant](
		"role", "caller", "callee", "other",
	)
}

// The name of the phone provider that handled the call (e.g., `openphone`).
type CallUpsertParamsProvider string

const (
	CallUpsertParamsProviderOpenphone CallUpsertParamsProvider = "openphone"
	CallUpsertParamsProviderUser      CallUpsertParamsProvider = "user"
	CallUpsertParamsProviderZoomPhone CallUpsertParamsProvider = "zoom_phone"
)

// Parameters for creating a `CallRecording` object.
//
// The properties ContentType, ProviderID, URL are required.
type CallUpsertParamsRecording struct {
	// The content type of the recording. Note that only `audio/mpeg` is supported at
	// this time.
	//
	// Any of "audio/mpeg".
	ContentType string `json:"content_type,omitzero" api:"required"`
	// The unique identifier for the recording from the provider's system.
	ProviderID string `json:"provider_id" api:"required"`
	// The URL pointing to the recording.
	URL string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r CallUpsertParamsRecording) MarshalJSON() (data []byte, err error) {
	type shadow CallUpsertParamsRecording
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallUpsertParamsRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallUpsertParamsRecording](
		"content_type", "audio/mpeg",
	)
}

// A transcript of the call.
//
// The property Cues is required.
type CallUpsertParamsTranscript struct {
	// A list of cues that identify the text spoken in specific time slices of the
	// call.
	Cues []CallUpsertParamsTranscriptCue `json:"cues,omitzero" api:"required"`
	paramObj
}

func (r CallUpsertParamsTranscript) MarshalJSON() (data []byte, err error) {
	type shadow CallUpsertParamsTranscript
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallUpsertParamsTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating a `CallTranscriptCue` object to capture the text spoken
// in a specific time slice.
//
// The properties From, Speaker, Text, To are required.
type CallUpsertParamsTranscriptCue struct {
	// The start time of the slice, in fractional seconds from the start of the call.
	From float64 `json:"from" api:"required"`
	// The E.164 formatted phone number of the speaker.
	Speaker string `json:"speaker" api:"required"`
	// The text spoken during the slice.
	Text string `json:"text" api:"required"`
	// The end time of the slice, in fractional seconds from the start of the call.
	To float64 `json:"to" api:"required"`
	paramObj
}

func (r CallUpsertParamsTranscriptCue) MarshalJSON() (data []byte, err error) {
	type shadow CallUpsertParamsTranscriptCue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallUpsertParamsTranscriptCue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
