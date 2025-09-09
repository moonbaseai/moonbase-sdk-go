// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"encoding/json"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// FunnelService contains methods and other services that help with interacting
// with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunnelService] method instead.
type FunnelService struct {
	Options []option.RequestOption
}

// NewFunnelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFunnelService(opts ...option.RequestOption) (r FunnelService) {
	r = FunnelService{}
	r.Options = opts
	return
}

// A Funnel represents a series of steps used to track progression.
type Funnel struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// An ordered list of `FunnelStep` objects that make up the funnel.
	Steps []FunnelStep `json:"steps,required"`
	// String representing the object’s type. Always `funnel` for this object.
	Type constant.Funnel `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Steps       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Funnel) RawJSON() string { return r.JSON.raw }
func (r *Funnel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a single step within a `Funnel`.
type FunnelStep struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The name of the step.
	Name string `json:"name,required"`
	// The status of the step in the funnel flow.
	//
	// - `active`: represents an in progress state within the funnel
	// - `success`: completed successfully and exited the funnel
	// - `failure`: exited the funnel without conversion
	//
	// Any of "active", "success", "failure".
	StepType FunnelStepStepType `json:"step_type,required"`
	// String representing the object’s type. Always `funnel_step` for this object.
	Type constant.FunnelStep `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		StepType    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FunnelStep to a FunnelStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FunnelStepParam.Overrides()
func (r FunnelStep) ToParam() FunnelStepParam {
	return param.Override[FunnelStepParam](json.RawMessage(r.RawJSON()))
}

// The status of the step in the funnel flow.
//
// - `active`: represents an in progress state within the funnel
// - `success`: completed successfully and exited the funnel
// - `failure`: exited the funnel without conversion
type FunnelStepStepType string

const (
	FunnelStepStepTypeActive  FunnelStepStepType = "active"
	FunnelStepStepTypeSuccess FunnelStepStepType = "success"
	FunnelStepStepTypeFailure FunnelStepStepType = "failure"
)

// Represents a single step within a `Funnel`.
//
// The properties ID, Name, StepType, Type are required.
type FunnelStepParam struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The name of the step.
	Name string `json:"name,required"`
	// The status of the step in the funnel flow.
	//
	// - `active`: represents an in progress state within the funnel
	// - `success`: completed successfully and exited the funnel
	// - `failure`: exited the funnel without conversion
	//
	// Any of "active", "success", "failure".
	StepType FunnelStepStepType `json:"step_type,omitzero,required"`
	// String representing the object’s type. Always `funnel_step` for this object.
	//
	// This field can be elided, and will marshal its zero value as "funnel_step".
	Type constant.FunnelStep `json:"type,required"`
	paramObj
}

func (r FunnelStepParam) MarshalJSON() (data []byte, err error) {
	type shadow FunnelStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
