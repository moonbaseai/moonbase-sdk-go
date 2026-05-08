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
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// Manage your collections and items
//
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

// Creates a new funnel.
func (r *FunnelService) New(ctx context.Context, body FunnelNewParams, opts ...option.RequestOption) (res *Funnel, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "funnels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the details of an existing funnel.
func (r *FunnelService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Funnel, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("funnels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a funnel.
func (r *FunnelService) Update(ctx context.Context, id string, body FunnelUpdateParams, opts ...option.RequestOption) (res *Funnel, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("funnels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of funnels.
func (r *FunnelService) List(ctx context.Context, query FunnelListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Funnel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "funnels"
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

// Returns a list of funnels.
func (r *FunnelService) ListAutoPaging(ctx context.Context, query FunnelListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Funnel] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a funnel.
func (r *FunnelService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("funnels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A Funnel represents a series of steps used to track progression.
type Funnel struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the funnel.
	Name string `json:"name" api:"required"`
	// An ordered list of `FunnelStep` objects that make up the funnel.
	Steps []FunnelStep `json:"steps" api:"required"`
	// String representing the object’s type. Always `funnel` for this object.
	Type constant.Funnel `json:"type" default:"funnel"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Steps       respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
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
	ID string `json:"id" api:"required"`
	// The display color of the step.
	//
	// Any of "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime",
	// "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet",
	// "yellow".
	Color FunnelStepColor `json:"color" api:"required"`
	// The name of the step.
	Name string `json:"name" api:"required"`
	// The status of the step in the funnel flow.
	//
	// - `active`: represents an in progress state within the funnel
	// - `success`: completed successfully and exited the funnel
	// - `failure`: exited the funnel without conversion
	//
	// Any of "active", "success", "failure".
	StepType FunnelStepStepType `json:"step_type" api:"required"`
	// String representing the object’s type. Always `funnel_step` for this object.
	Type constant.FunnelStep `json:"type" default:"funnel_step"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Color       respjson.Field
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

// The display color of the step.
type FunnelStepColor string

const (
	FunnelStepColorAmber   FunnelStepColor = "amber"
	FunnelStepColorBlue    FunnelStepColor = "blue"
	FunnelStepColorCyan    FunnelStepColor = "cyan"
	FunnelStepColorEmerald FunnelStepColor = "emerald"
	FunnelStepColorFuchsia FunnelStepColor = "fuchsia"
	FunnelStepColorGreen   FunnelStepColor = "green"
	FunnelStepColorIndigo  FunnelStepColor = "indigo"
	FunnelStepColorLime    FunnelStepColor = "lime"
	FunnelStepColorLunar   FunnelStepColor = "lunar"
	FunnelStepColorOrange  FunnelStepColor = "orange"
	FunnelStepColorPink    FunnelStepColor = "pink"
	FunnelStepColorPurple  FunnelStepColor = "purple"
	FunnelStepColorRed     FunnelStepColor = "red"
	FunnelStepColorRose    FunnelStepColor = "rose"
	FunnelStepColorSky     FunnelStepColor = "sky"
	FunnelStepColorTeal    FunnelStepColor = "teal"
	FunnelStepColorViolet  FunnelStepColor = "violet"
	FunnelStepColorYellow  FunnelStepColor = "yellow"
)

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

// The properties ID, Type are required.
type FunnelStepPointerParam struct {
	ID string `json:"id" api:"required"`
	// This field can be elided, and will marshal its zero value as "funnel_step".
	Type constant.FunnelStep `json:"type" default:"funnel_step"`
	paramObj
}

func (r FunnelStepPointerParam) MarshalJSON() (data []byte, err error) {
	type shadow FunnelStepPointerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelStepPointerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewParams struct {
	// The name of the funnel.
	Name string `json:"name" api:"required"`
	// An ordered list of steps to create. Array order determines step order.
	Steps []FunnelNewParamsStep `json:"steps,omitzero"`
	paramObj
}

func (r FunnelNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating a funnel step.
//
// The properties Color, Name, StepType are required.
type FunnelNewParamsStep struct {
	// The display color of the step.
	//
	// Any of "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime",
	// "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet",
	// "yellow".
	Color string `json:"color,omitzero" api:"required"`
	// The name of the step.
	Name string `json:"name" api:"required"`
	// The status of the step in the funnel flow.
	//
	// - `active`: represents an in progress state within the funnel
	// - `success`: completed successfully and exited the funnel
	// - `failure`: exited the funnel without conversion
	//
	// Any of "active", "success", "failure".
	StepType string `json:"step_type,omitzero" api:"required"`
	paramObj
}

func (r FunnelNewParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsStep](
		"color", "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime", "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet", "yellow",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsStep](
		"step_type", "active", "success", "failure",
	)
}

type FunnelUpdateParams struct {
	// The name of the funnel.
	Name param.Opt[string] `json:"name,omitzero"`
	// An ordered list of steps. Providing this replaces all existing steps. Omitting
	// preserves existing steps.
	Steps []FunnelUpdateParamsStep `json:"steps,omitzero"`
	paramObj
}

func (r FunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for updating a funnel step. Include `id` to update an existing step,
// or omit `id` to create a new one. Steps not included are removed.
//
// The properties Color, Name, StepType are required.
type FunnelUpdateParamsStep struct {
	// The display color of the step.
	//
	// Any of "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime",
	// "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet",
	// "yellow".
	Color string `json:"color,omitzero" api:"required"`
	// The name of the step.
	Name string `json:"name" api:"required"`
	// The status of the step in the funnel flow.
	//
	// - `active`: represents an in progress state within the funnel
	// - `success`: completed successfully and exited the funnel
	// - `failure`: exited the funnel without conversion
	//
	// Any of "active", "success", "failure".
	StepType string `json:"step_type,omitzero" api:"required"`
	// The ID of an existing step to update. Omit to create a new step.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r FunnelUpdateParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsStep](
		"color", "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime", "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet", "yellow",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsStep](
		"step_type", "active", "success", "failure",
	)
}

type FunnelListParams struct {
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

// URLQuery serializes [FunnelListParams]'s query parameters as `url.Values`.
func (r FunnelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
