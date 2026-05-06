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

// Manage your meetings, files, and notes
//
// TagsetService contains methods and other services that help with interacting
// with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagsetService] method instead.
type TagsetService struct {
	Options []option.RequestOption
}

// NewTagsetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTagsetService(opts ...option.RequestOption) (r TagsetService) {
	r = TagsetService{}
	r.Options = opts
	return
}

// Create a new tagset.
func (r *TagsetService) New(ctx context.Context, body TagsetNewParams, opts ...option.RequestOption) (res *Tagset, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tagsets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the details of an existing tagset.
func (r *TagsetService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Tagset, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tagsets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing tagset.
func (r *TagsetService) Update(ctx context.Context, id string, body TagsetUpdateParams, opts ...option.RequestOption) (res *Tagset, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tagsets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of your tagsets.
func (r *TagsetService) List(ctx context.Context, query TagsetListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Tagset], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "tagsets"
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

// Returns a list of your tagsets.
func (r *TagsetService) ListAutoPaging(ctx context.Context, query TagsetListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Tagset] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a tagset.
func (r *TagsetService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("tagsets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A Tagset is a collection of `Tag` objects whose tags can be applied to
// conversations, calls, and meetings.
type Tagset struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the tagset.
	Name string `json:"name" api:"required"`
	// A list of `Tag` objects belonging to this tagset.
	Tags []shared.Tag `json:"tags" api:"required"`
	// String representing the object’s type. Always `tagset` for this object.
	Type constant.Tagset `json:"type" default:"tagset"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional description of the tagset's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Tags        respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tagset) RawJSON() string { return r.JSON.raw }
func (r *Tagset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagsetPointer struct {
	ID   string          `json:"id" api:"required"`
	Type constant.Tagset `json:"type" default:"tagset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagsetPointer) RawJSON() string { return r.JSON.raw }
func (r *TagsetPointer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagsetNewParams struct {
	// The name of the tagset.
	Name string `json:"name" api:"required"`
	// An optional description of the tagset's purpose.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional list of tags to create with this tagset. Tags are ordered by their
	// position in the list.
	Tags []TagsetNewParamsTag `json:"tags,omitzero"`
	paramObj
}

func (r TagsetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagsetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagsetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating or updating a tag within a tagset.
//
// The properties Color, Name are required.
type TagsetNewParamsTag struct {
	// The color for the tag.
	//
	// Any of "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime",
	// "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet",
	// "yellow".
	Color string `json:"color,omitzero" api:"required"`
	// The name of the tag.
	Name string `json:"name" api:"required"`
	// Existing tag identifier. Include to update an existing tag, omit to create a new
	// tag.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r TagsetNewParamsTag) MarshalJSON() (data []byte, err error) {
	type shadow TagsetNewParamsTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagsetNewParamsTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TagsetNewParamsTag](
		"color", "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime", "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet", "yellow",
	)
}

type TagsetUpdateParams struct {
	// An updated description of the tagset.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name of the tagset.
	Name param.Opt[string] `json:"name,omitzero"`
	// Optional full list of tags for this tagset. If provided, tags are ordered by
	// array position.
	Tags []TagsetUpdateParamsTag `json:"tags,omitzero"`
	paramObj
}

func (r TagsetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TagsetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagsetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating or updating a tag within a tagset.
//
// The properties Color, Name are required.
type TagsetUpdateParamsTag struct {
	// The color for the tag.
	//
	// Any of "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime",
	// "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet",
	// "yellow".
	Color string `json:"color,omitzero" api:"required"`
	// The name of the tag.
	Name string `json:"name" api:"required"`
	// Existing tag identifier. Include to update an existing tag, omit to create a new
	// tag.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r TagsetUpdateParamsTag) MarshalJSON() (data []byte, err error) {
	type shadow TagsetUpdateParamsTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagsetUpdateParamsTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TagsetUpdateParamsTag](
		"color", "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime", "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet", "yellow",
	)
}

type TagsetListParams struct {
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

// URLQuery serializes [TagsetListParams]'s query parameters as `url.Values`.
func (r TagsetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
