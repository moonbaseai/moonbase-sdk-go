// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"encoding/json"
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
	// Where a tagset is available (`calls`, `meetings`, or `inbox` with an inbox ID).
	Associations []TagsetAssociationUnion `json:"associations" api:"required"`
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
		ID           respjson.Field
		Associations respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Tags         respjson.Field
		Type         respjson.Field
		UpdatedAt    respjson.Field
		Description  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tagset) RawJSON() string { return r.JSON.raw }
func (r *Tagset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TagsetAssociationUnion contains all possible properties and values from
// [TagsetAssociationCalls], [TagsetAssociationMeetings], [TagsetAssociationInbox].
//
// Use the [TagsetAssociationUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TagsetAssociationUnion struct {
	// Any of "calls", "meetings", "inbox".
	Type string `json:"type"`
	// This field is from variant [TagsetAssociationInbox].
	ID   string `json:"id"`
	JSON struct {
		Type respjson.Field
		ID   respjson.Field
		raw  string
	} `json:"-"`
}

// anyTagsetAssociation is implemented by each variant of [TagsetAssociationUnion]
// to add type safety for the return type of [TagsetAssociationUnion.AsAny]
type anyTagsetAssociation interface {
	implTagsetAssociationUnion()
}

func (TagsetAssociationCalls) implTagsetAssociationUnion()    {}
func (TagsetAssociationMeetings) implTagsetAssociationUnion() {}
func (TagsetAssociationInbox) implTagsetAssociationUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TagsetAssociationUnion.AsAny().(type) {
//	case moonbase.TagsetAssociationCalls:
//	case moonbase.TagsetAssociationMeetings:
//	case moonbase.TagsetAssociationInbox:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TagsetAssociationUnion) AsAny() anyTagsetAssociation {
	switch u.Type {
	case "calls":
		return u.AsCalls()
	case "meetings":
		return u.AsMeetings()
	case "inbox":
		return u.AsInbox()
	}
	return nil
}

func (u TagsetAssociationUnion) AsCalls() (v TagsetAssociationCalls) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagsetAssociationUnion) AsMeetings() (v TagsetAssociationMeetings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagsetAssociationUnion) AsInbox() (v TagsetAssociationInbox) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TagsetAssociationUnion) RawJSON() string { return u.JSON.raw }

func (r *TagsetAssociationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TagsetAssociationUnion to a TagsetAssociationUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TagsetAssociationUnionParam.Overrides()
func (r TagsetAssociationUnion) ToParam() TagsetAssociationUnionParam {
	return param.Override[TagsetAssociationUnionParam](json.RawMessage(r.RawJSON()))
}

// Makes this tagset available for calls.
type TagsetAssociationCalls struct {
	// String representing the association type. Always `calls` for call tagset
	// associations.
	Type constant.Calls `json:"type" default:"calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagsetAssociationCalls) RawJSON() string { return r.JSON.raw }
func (r *TagsetAssociationCalls) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Makes this tagset available for meetings.
type TagsetAssociationMeetings struct {
	// String representing the association type. Always `meetings` for meeting tagset
	// associations.
	Type constant.Meetings `json:"type" default:"meetings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagsetAssociationMeetings) RawJSON() string { return r.JSON.raw }
func (r *TagsetAssociationMeetings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Makes this tagset available in an inbox.
type TagsetAssociationInbox struct {
	// Unique identifier of the inbox this tagset is assigned to.
	ID string `json:"id" api:"required"`
	// String representing the association type. Always `inbox` for inbox tagset
	// associations.
	Type constant.Inbox `json:"type" default:"inbox"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagsetAssociationInbox) RawJSON() string { return r.JSON.raw }
func (r *TagsetAssociationInbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func TagsetAssociationParamOfInbox(id string) TagsetAssociationUnionParam {
	var inbox TagsetAssociationInboxParam
	inbox.ID = id
	return TagsetAssociationUnionParam{OfInbox: &inbox}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TagsetAssociationUnionParam struct {
	OfCalls    *TagsetAssociationCallsParam    `json:",omitzero,inline"`
	OfMeetings *TagsetAssociationMeetingsParam `json:",omitzero,inline"`
	OfInbox    *TagsetAssociationInboxParam    `json:",omitzero,inline"`
	paramUnion
}

func (u TagsetAssociationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCalls, u.OfMeetings, u.OfInbox)
}
func (u *TagsetAssociationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TagsetAssociationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfCalls) {
		return u.OfCalls
	} else if !param.IsOmitted(u.OfMeetings) {
		return u.OfMeetings
	} else if !param.IsOmitted(u.OfInbox) {
		return u.OfInbox
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TagsetAssociationUnionParam) GetID() *string {
	if vt := u.OfInbox; vt != nil {
		return &vt.ID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TagsetAssociationUnionParam) GetType() *string {
	if vt := u.OfCalls; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMeetings; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInbox; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[TagsetAssociationUnionParam](
		"type",
		apijson.Discriminator[TagsetAssociationCallsParam]("calls"),
		apijson.Discriminator[TagsetAssociationMeetingsParam]("meetings"),
		apijson.Discriminator[TagsetAssociationInboxParam]("inbox"),
	)
}

func NewTagsetAssociationCallsParam() TagsetAssociationCallsParam {
	return TagsetAssociationCallsParam{
		Type: "calls",
	}
}

// Makes this tagset available for calls.
//
// This struct has a constant value, construct it with
// [NewTagsetAssociationCallsParam].
type TagsetAssociationCallsParam struct {
	// String representing the association type. Always `calls` for call tagset
	// associations.
	Type constant.Calls `json:"type" default:"calls"`
	paramObj
}

func (r TagsetAssociationCallsParam) MarshalJSON() (data []byte, err error) {
	type shadow TagsetAssociationCallsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagsetAssociationCallsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewTagsetAssociationMeetingsParam() TagsetAssociationMeetingsParam {
	return TagsetAssociationMeetingsParam{
		Type: "meetings",
	}
}

// Makes this tagset available for meetings.
//
// This struct has a constant value, construct it with
// [NewTagsetAssociationMeetingsParam].
type TagsetAssociationMeetingsParam struct {
	// String representing the association type. Always `meetings` for meeting tagset
	// associations.
	Type constant.Meetings `json:"type" default:"meetings"`
	paramObj
}

func (r TagsetAssociationMeetingsParam) MarshalJSON() (data []byte, err error) {
	type shadow TagsetAssociationMeetingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagsetAssociationMeetingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Makes this tagset available in an inbox.
//
// The properties ID, Type are required.
type TagsetAssociationInboxParam struct {
	// Unique identifier of the inbox this tagset is assigned to.
	ID string `json:"id" api:"required"`
	// String representing the association type. Always `inbox` for inbox tagset
	// associations.
	//
	// This field can be elided, and will marshal its zero value as "inbox".
	Type constant.Inbox `json:"type" default:"inbox"`
	paramObj
}

func (r TagsetAssociationInboxParam) MarshalJSON() (data []byte, err error) {
	type shadow TagsetAssociationInboxParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagsetAssociationInboxParam) UnmarshalJSON(data []byte) error {
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
	// Optional list of associations for this tagset. Include `{type: "calls"}`,
	// `{type: "meetings"}`, or `{type: "inbox", id}`.
	Associations []TagsetAssociationUnionParam `json:"associations,omitzero"`
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
	// Optional full list of associations for this tagset. If provided, it replaces all
	// existing associations. An empty array clears all associations, and omitting it
	// preserves existing associations.
	Associations []TagsetAssociationUnionParam `json:"associations,omitzero"`
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
