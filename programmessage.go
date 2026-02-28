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
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// Manage your marketing campaigns and forms
//
// ProgramMessageService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProgramMessageService] method instead.
type ProgramMessageService struct {
	Options []option.RequestOption
}

// NewProgramMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProgramMessageService(opts ...option.RequestOption) (r ProgramMessageService) {
	r = ProgramMessageService{}
	r.Options = opts
	return
}

// Sends a message using a program template.
func (r *ProgramMessageService) Send(ctx context.Context, body ProgramMessageSendParams, opts ...option.RequestOption) (res *ProgramMessage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "program_messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Represents a single message sent as part of a `Program`.
type ProgramMessage struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Time at which the message was created and enqueued for sending, as an ISO 8601
	// timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The `ProgramTemplate` used to generate this message.
	ProgramTemplate ProgramTemplate `json:"program_template" api:"required"`
	// String representing the object’s type. Always `program_message` for this object.
	Type constant.ProgramMessage `json:"type" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		ProgramTemplate respjson.Field
		Type            respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProgramMessage) RawJSON() string { return r.JSON.raw }
func (r *ProgramMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProgramMessageSendParams struct {
	// The person to send the message to.
	Person ProgramMessageSendParamsPerson `json:"person,omitzero" api:"required"`
	// The ID of the `ProgramTemplate` to use for sending the message.
	ProgramTemplateID string `json:"program_template_id" api:"required"`
	// Any custom Liquid variables to be interpolated into the message template.
	CustomVariables map[string]any `json:"custom_variables,omitzero"`
	paramObj
}

func (r ProgramMessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow ProgramMessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProgramMessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The person to send the message to.
//
// The property Email is required.
type ProgramMessageSendParamsPerson struct {
	Email string `json:"email" api:"required" format:"email"`
	paramObj
}

func (r ProgramMessageSendParamsPerson) MarshalJSON() (data []byte, err error) {
	type shadow ProgramMessageSendParamsPerson
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProgramMessageSendParamsPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
