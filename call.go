// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"net/http"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
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
	opts = append(r.Options[:], opts...)
	path := "calls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The Call object represents a phone call that has been logged in the system. It
// contains details about the participants, timing, and outcome of the call.
type Call struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The direction of the call, either `incoming` or `outgoing`.
	//
	// Any of "incoming", "outgoing".
	Direction CallDirection `json:"direction,required"`
	// The participants involved in the call.
	Participants []CallParticipant `json:"participants,required"`
	// The name of the phone provider that handled the call.
	Provider string `json:"provider,required"`
	// The unique identifier for the call from the provider's system.
	ProviderID string `json:"provider_id,required"`
	// The time the call started, as an RFC 3339 timestamp.
	StartAt time.Time `json:"start_at,required" format:"date-time"`
	// The current status of the call.
	//
	// Any of "queued", "initiated", "ringing", "in_progress", "completed", "busy",
	// "failed", "no_answer", "canceled", "missed", "answered", "forwarded",
	// "abandoned".
	Status CallStatus `json:"status,required"`
	// String representing the object’s type. Always `call` for this object.
	Type constant.Call `json:"type,required"`
	// The time the call was answered, if available, as an RFC 3339 timestamp.
	AnsweredAt time.Time `json:"answered_at" format:"date-time"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The time the call ended, if available, as an RFC 3339 timestamp.
	EndAt time.Time `json:"end_at" format:"date-time"`
	// A hash of additional metadata from the provider.
	ProviderMetadata map[string]any `json:"provider_metadata"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Direction        respjson.Field
		Participants     respjson.Field
		Provider         respjson.Field
		ProviderID       respjson.Field
		StartAt          respjson.Field
		Status           respjson.Field
		Type             respjson.Field
		AnsweredAt       respjson.Field
		CreatedAt        respjson.Field
		EndAt            respjson.Field
		ProviderMetadata respjson.Field
		UpdatedAt        respjson.Field
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
	ID string `json:"id,required"`
	// The E.164 formatted phone number of the participant.
	Phone string `json:"phone,required"`
	// The role of the participant in the call. Can be `caller`, `callee`, or `other`.
	//
	// Any of "caller", "callee", "other".
	Role string `json:"role,required"`
	// String representing the object’s type. Always `participant` for this object.
	Type constant.Participant `json:"type,required"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Phone       respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallParticipant) RawJSON() string { return r.JSON.raw }
func (r *CallParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the call.
type CallStatus string

const (
	CallStatusQueued     CallStatus = "queued"
	CallStatusInitiated  CallStatus = "initiated"
	CallStatusRinging    CallStatus = "ringing"
	CallStatusInProgress CallStatus = "in_progress"
	CallStatusCompleted  CallStatus = "completed"
	CallStatusBusy       CallStatus = "busy"
	CallStatusFailed     CallStatus = "failed"
	CallStatusNoAnswer   CallStatus = "no_answer"
	CallStatusCanceled   CallStatus = "canceled"
	CallStatusMissed     CallStatus = "missed"
	CallStatusAnswered   CallStatus = "answered"
	CallStatusForwarded  CallStatus = "forwarded"
	CallStatusAbandoned  CallStatus = "abandoned"
)

type CallNewParams struct {
	// The direction of the call, either `incoming` or `outgoing`.
	//
	// Any of "incoming", "outgoing".
	Direction CallNewParamsDirection `json:"direction,omitzero,required"`
	// An array of participants involved in the call.
	Participants []CallNewParamsParticipant `json:"participants,omitzero,required"`
	// The name of the phone provider that handled the call (e.g., `openphone`).
	Provider string `json:"provider,required"`
	// The unique identifier for the call from the provider's system.
	ProviderID string `json:"provider_id,required"`
	// The time the call started, as an RFC 3339 timestamp.
	StartAt time.Time `json:"start_at,required" format:"date-time"`
	// The status of the call.
	//
	// Any of "queued", "initiated", "ringing", "in_progress", "completed", "busy",
	// "failed", "no_answer", "canceled", "missed", "answered", "forwarded",
	// "abandoned".
	Status CallNewParamsStatus `json:"status,omitzero,required"`
	// The time the call was answered, as an RFC 3339 timestamp.
	AnsweredAt param.Opt[time.Time] `json:"answered_at,omitzero" format:"date-time"`
	// The time the call ended, as an RFC 3339 timestamp.
	EndAt param.Opt[time.Time] `json:"end_at,omitzero" format:"date-time"`
	// A hash of additional metadata from the provider.
	ProviderMetadata map[string]any `json:"provider_metadata,omitzero"`
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
	Phone string `json:"phone,required"`
	// The role of the participant in the call. Can be `caller`, `callee`, or `other`.
	//
	// Any of "caller", "callee", "other".
	Role string `json:"role,omitzero,required"`
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

// The status of the call.
type CallNewParamsStatus string

const (
	CallNewParamsStatusQueued     CallNewParamsStatus = "queued"
	CallNewParamsStatusInitiated  CallNewParamsStatus = "initiated"
	CallNewParamsStatusRinging    CallNewParamsStatus = "ringing"
	CallNewParamsStatusInProgress CallNewParamsStatus = "in_progress"
	CallNewParamsStatusCompleted  CallNewParamsStatus = "completed"
	CallNewParamsStatusBusy       CallNewParamsStatus = "busy"
	CallNewParamsStatusFailed     CallNewParamsStatus = "failed"
	CallNewParamsStatusNoAnswer   CallNewParamsStatus = "no_answer"
	CallNewParamsStatusCanceled   CallNewParamsStatus = "canceled"
	CallNewParamsStatusMissed     CallNewParamsStatus = "missed"
	CallNewParamsStatusAnswered   CallNewParamsStatus = "answered"
	CallNewParamsStatusForwarded  CallNewParamsStatus = "forwarded"
	CallNewParamsStatusAbandoned  CallNewParamsStatus = "abandoned"
)
