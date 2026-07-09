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
	"github.com/moonbaseai/moonbase-sdk-go/internal/paramutil"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/pagination"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// Manage your collections and items
//
// CollectionService contains methods and other services that help with interacting
// with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectionService] method instead.
type CollectionService struct {
	Options []option.RequestOption
	// Manage your collections and items
	Fields CollectionFieldService
	// Manage your collections and items
	Items CollectionItemService
}

// NewCollectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCollectionService(opts ...option.RequestOption) (r CollectionService) {
	r = CollectionService{}
	r.Options = opts
	r.Fields = NewCollectionFieldService(opts...)
	r.Items = NewCollectionItemService(opts...)
	return
}

// Creates a new collection with default fields (name, created_at, updated_at) and
// a default view.
func (r *CollectionService) New(ctx context.Context, body CollectionNewParams, opts ...option.RequestOption) (res *Collection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "collections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the details of an existing collection.
func (r *CollectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Collection, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("collections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing collection.
func (r *CollectionService) Update(ctx context.Context, id string, body CollectionUpdateParams, opts ...option.RequestOption) (res *Collection, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("collections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of your collections.
func (r *CollectionService) List(ctx context.Context, query CollectionListParams, opts ...option.RequestOption) (res *pagination.CursorPage[CollectionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "collections"
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

// Returns a list of your collections.
func (r *CollectionService) ListAutoPaging(ctx context.Context, query CollectionListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[CollectionListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a collection.
func (r *CollectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("collections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A field that stores true or false values.
type BooleanField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality BooleanFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind BooleanFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Is Active").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `is_active`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/boolean` for this field.
	Type constant.FieldBoolean `json:"type" default:"field/boolean"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BooleanField) RawJSON() string { return r.JSON.raw }
func (r *BooleanField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type BooleanFieldCardinality string

const (
	BooleanFieldCardinalityOne  BooleanFieldCardinality = "one"
	BooleanFieldCardinalityMany BooleanFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type BooleanFieldKind string

const (
	BooleanFieldKindSystem  BooleanFieldKind = "system"
	BooleanFieldKindInverse BooleanFieldKind = "inverse"
	BooleanFieldKindCustom  BooleanFieldKind = "custom"
)

// True or false value
type BooleanValue struct {
	Data bool                  `json:"data" api:"required"`
	Type constant.ValueBoolean `json:"type" default:"value/boolean"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BooleanValue) RawJSON() string { return r.JSON.raw }
func (r *BooleanValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BooleanValue to a BooleanValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BooleanValueParam.Overrides()
func (r BooleanValue) ToParam() BooleanValueParam {
	return param.Override[BooleanValueParam](json.RawMessage(r.RawJSON()))
}

// True or false value
//
// The properties Data, Type are required.
type BooleanValueParam struct {
	Data bool `json:"data" api:"required"`
	// This field can be elided, and will marshal its zero value as "value/boolean".
	Type constant.ValueBoolean `json:"type" default:"value/boolean"`
	paramObj
}

func (r BooleanValueParam) MarshalJSON() (data []byte, err error) {
	type shadow BooleanValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BooleanValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores one or more predefined options from a list of choices.
type ChoiceField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality ChoiceFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind ChoiceFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Priority").
	Name string `json:"name" api:"required"`
	// A list of `FieldOption` objects representing the available choices for this
	// field.
	Options []ChoiceFieldOption `json:"options" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `priority`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/choice` for this field.
	Type constant.FieldChoice `json:"type" default:"field/choice"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Options       respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChoiceField) RawJSON() string { return r.JSON.raw }
func (r *ChoiceField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type ChoiceFieldCardinality string

const (
	ChoiceFieldCardinalityOne  ChoiceFieldCardinality = "one"
	ChoiceFieldCardinalityMany ChoiceFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type ChoiceFieldKind string

const (
	ChoiceFieldKindSystem  ChoiceFieldKind = "system"
	ChoiceFieldKindInverse ChoiceFieldKind = "inverse"
	ChoiceFieldKindCustom  ChoiceFieldKind = "custom"
)

// Represents a single selectable option within a choice field.
type ChoiceFieldOption struct {
	// Unique identifier for the option.
	ID string `json:"id" api:"required"`
	// The color of the option.
	//
	// Any of "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime",
	// "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet",
	// "yellow".
	Color ChoiceFieldOptionColor `json:"color" api:"required"`
	// The human-readable text displayed for this option.
	Name string `json:"name" api:"required"`
	// String representing the object’s type. Always `choice_field_option` for this
	// object.
	Type constant.ChoiceFieldOption `json:"type" default:"choice_field_option"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Color       respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChoiceFieldOption) RawJSON() string { return r.JSON.raw }
func (r *ChoiceFieldOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The color of the option.
type ChoiceFieldOptionColor string

const (
	ChoiceFieldOptionColorAmber   ChoiceFieldOptionColor = "amber"
	ChoiceFieldOptionColorBlue    ChoiceFieldOptionColor = "blue"
	ChoiceFieldOptionColorCyan    ChoiceFieldOptionColor = "cyan"
	ChoiceFieldOptionColorEmerald ChoiceFieldOptionColor = "emerald"
	ChoiceFieldOptionColorFuchsia ChoiceFieldOptionColor = "fuchsia"
	ChoiceFieldOptionColorGreen   ChoiceFieldOptionColor = "green"
	ChoiceFieldOptionColorIndigo  ChoiceFieldOptionColor = "indigo"
	ChoiceFieldOptionColorLime    ChoiceFieldOptionColor = "lime"
	ChoiceFieldOptionColorLunar   ChoiceFieldOptionColor = "lunar"
	ChoiceFieldOptionColorOrange  ChoiceFieldOptionColor = "orange"
	ChoiceFieldOptionColorPink    ChoiceFieldOptionColor = "pink"
	ChoiceFieldOptionColorPurple  ChoiceFieldOptionColor = "purple"
	ChoiceFieldOptionColorRed     ChoiceFieldOptionColor = "red"
	ChoiceFieldOptionColorRose    ChoiceFieldOptionColor = "rose"
	ChoiceFieldOptionColorSky     ChoiceFieldOptionColor = "sky"
	ChoiceFieldOptionColorTeal    ChoiceFieldOptionColor = "teal"
	ChoiceFieldOptionColorViolet  ChoiceFieldOptionColor = "violet"
	ChoiceFieldOptionColorYellow  ChoiceFieldOptionColor = "yellow"
)

// The properties ID, Type are required.
type ChoiceFieldOptionPointerParam struct {
	ID string `json:"id" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "choice_field_option".
	Type constant.ChoiceFieldOption `json:"type" default:"choice_field_option"`
	paramObj
}

func (r ChoiceFieldOptionPointerParam) MarshalJSON() (data []byte, err error) {
	type shadow ChoiceFieldOptionPointerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChoiceFieldOptionPointerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Selected choice option
type ChoiceValue struct {
	// An option that must match one of the predefined options for the field.
	Data ChoiceFieldOption    `json:"data" api:"required"`
	Type constant.ValueChoice `json:"type" default:"value/choice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChoiceValue) RawJSON() string { return r.JSON.raw }
func (r *ChoiceValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Selected choice option
//
// The properties Data, Type are required.
type ChoiceValueParam struct {
	// An option that must match one of the predefined options for the field.
	Data ChoiceFieldOptionPointerParam `json:"data,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "value/choice".
	Type constant.ValueChoice `json:"type" default:"value/choice"`
	paramObj
}

func (r ChoiceValueParam) MarshalJSON() (data []byte, err error) {
	type shadow ChoiceValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChoiceValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Collection is a container for structured data, similar to a database table or
// spreadsheet. It defines a schema using a set of `Fields` and holds the data as a
// list of `Items`.
type Collection struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A list of `Field` objects that define the schema for items in this collection.
	Fields []FieldUnion `json:"fields" api:"required"`
	// `system` collections are managed by Moonbase (e.g., People, Organizations),
	// `form` collections back a Form, and `custom` collections are user-created.
	//
	// Any of "system", "form", "custom".
	Kind CollectionKind `json:"kind" api:"required"`
	// The user-facing name of the collection (e.g., “Organizations”).
	Name string `json:"name" api:"required"`
	// A unique, stable, machine-readable identifier for the collection. This reference
	// is used in API requests and does not change even if the `name` is updated.
	Ref string `json:"ref" api:"required"`
	// String representing the object’s type. Always `collection` for this object.
	Type constant.Collection `json:"type" default:"collection"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the collection's purpose.
	Description string `json:"description"`
	// The collection's icon, as a Phosphor icon name in kebab-case (e.g. `users`,
	// `chart-bar`). Only present when an icon is set.
	IconName string `json:"icon_name"`
	// A list of saved `View` objects for presenting the collection's data.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Views []CollectionView `json:"views"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Fields      respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		Ref         respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		IconName    respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Collection) RawJSON() string { return r.JSON.raw }
func (r *Collection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `system` collections are managed by Moonbase (e.g., People, Organizations),
// `form` collections back a Form, and `custom` collections are user-created.
type CollectionKind string

const (
	CollectionKindSystem CollectionKind = "system"
	CollectionKindForm   CollectionKind = "form"
	CollectionKindCustom CollectionKind = "custom"
)

type CollectionView struct {
	ID string `json:"id" api:"required"`
	// A lightweight reference to a `Collection`, containing the minimal information
	// needed to identify it.
	Collection CollectionPointer `json:"collection" api:"required"`
	CreatedAt  time.Time         `json:"created_at" api:"required" format:"date-time"`
	Name       string            `json:"name" api:"required"`
	Type       constant.View     `json:"type" default:"view"`
	UpdatedAt  time.Time         `json:"updated_at" api:"required" format:"date-time"`
	// Any of "table", "board".
	ViewType string `json:"view_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Collection  respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ViewType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionView) RawJSON() string { return r.JSON.raw }
func (r *CollectionView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A lightweight reference to a `Collection`, containing the minimal information
// needed to identify it.
type CollectionPointer struct {
	// Unique identifier of the collection.
	ID string `json:"id" api:"required"`
	// The stable, machine-readable reference identifier of the collection.
	Ref string `json:"ref" api:"required"`
	// String representing the object’s type. Always `collection` for this object.
	Type constant.Collection `json:"type" default:"collection"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Ref         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionPointer) RawJSON() string { return r.JSON.raw }
func (r *CollectionPointer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolves to today's date at the time the record is created.
type CurrentDate struct {
	Type constant.CurrentDate `json:"type" default:"current_date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrentDate) RawJSON() string { return r.JSON.raw }
func (r *CurrentDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CurrentDate to a CurrentDateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CurrentDateParam.Overrides()
func (r CurrentDate) ToParam() CurrentDateParam {
	return param.Override[CurrentDateParam](json.RawMessage(r.RawJSON()))
}

func NewCurrentDateParam() CurrentDateParam {
	return CurrentDateParam{
		Type: "current_date",
	}
}

// Resolves to today's date at the time the record is created.
//
// This struct has a constant value, construct it with [NewCurrentDateParam].
type CurrentDateParam struct {
	Type constant.CurrentDate `json:"type" default:"current_date"`
	paramObj
}

func (r CurrentDateParam) MarshalJSON() (data []byte, err error) {
	type shadow CurrentDateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CurrentDateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolves to the current date and time at the time the record is created.
type CurrentDatetime struct {
	Type constant.CurrentDatetime `json:"type" default:"current_datetime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrentDatetime) RawJSON() string { return r.JSON.raw }
func (r *CurrentDatetime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CurrentDatetime to a CurrentDatetimeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CurrentDatetimeParam.Overrides()
func (r CurrentDatetime) ToParam() CurrentDatetimeParam {
	return param.Override[CurrentDatetimeParam](json.RawMessage(r.RawJSON()))
}

func NewCurrentDatetimeParam() CurrentDatetimeParam {
	return CurrentDatetimeParam{
		Type: "current_datetime",
	}
}

// Resolves to the current date and time at the time the record is created.
//
// This struct has a constant value, construct it with [NewCurrentDatetimeParam].
type CurrentDatetimeParam struct {
	Type constant.CurrentDatetime `json:"type" default:"current_datetime"`
	paramObj
}

func (r CurrentDatetimeParam) MarshalJSON() (data []byte, err error) {
	type shadow CurrentDatetimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CurrentDatetimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolves to the team member who creates the record.
type CurrentMember struct {
	Type constant.CurrentMember `json:"type" default:"current_member"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrentMember) RawJSON() string { return r.JSON.raw }
func (r *CurrentMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CurrentMember to a CurrentMemberParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CurrentMemberParam.Overrides()
func (r CurrentMember) ToParam() CurrentMemberParam {
	return param.Override[CurrentMemberParam](json.RawMessage(r.RawJSON()))
}

func NewCurrentMemberParam() CurrentMemberParam {
	return CurrentMemberParam{
		Type: "current_member",
	}
}

// Resolves to the team member who creates the record.
//
// This struct has a constant value, construct it with [NewCurrentMemberParam].
type CurrentMemberParam struct {
	Type constant.CurrentMember `json:"type" default:"current_member"`
	paramObj
}

func (r CurrentMemberParam) MarshalJSON() (data []byte, err error) {
	type shadow CurrentMemberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CurrentMemberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores dates without time information.
type DateField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality DateFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind DateFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Due Date").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `due_date`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/date` for this field.
	Type constant.FieldDate `json:"type" default:"field/date"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DateField) RawJSON() string { return r.JSON.raw }
func (r *DateField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type DateFieldCardinality string

const (
	DateFieldCardinalityOne  DateFieldCardinality = "one"
	DateFieldCardinalityMany DateFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type DateFieldKind string

const (
	DateFieldKindSystem  DateFieldKind = "system"
	DateFieldKindInverse DateFieldKind = "inverse"
	DateFieldKindCustom  DateFieldKind = "custom"
)

func DateFieldDefaultValueParamOfValueDate(data time.Time) DateFieldDefaultValueParamUnion {
	var valueDate DateValueParam
	valueDate.Data = data
	return DateFieldDefaultValueParamUnion{OfValueDate: &valueDate}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DateFieldDefaultValueParamUnion struct {
	OfValueDate   *DateValueParam   `json:",omitzero,inline"`
	OfCurrentDate *CurrentDateParam `json:",omitzero,inline"`
	paramUnion
}

func (u DateFieldDefaultValueParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfValueDate, u.OfCurrentDate)
}
func (u *DateFieldDefaultValueParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DateFieldDefaultValueParamUnion) asAny() any {
	if !param.IsOmitted(u.OfValueDate) {
		return u.OfValueDate
	} else if !param.IsOmitted(u.OfCurrentDate) {
		return u.OfCurrentDate
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DateFieldDefaultValueParamUnion) GetData() *time.Time {
	if vt := u.OfValueDate; vt != nil {
		return &vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DateFieldDefaultValueParamUnion) GetType() *string {
	if vt := u.OfValueDate; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCurrentDate; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[DateFieldDefaultValueParamUnion](
		"type",
		apijson.Discriminator[DateValueParam]("value/date"),
		apijson.Discriminator[CurrentDateParam]("current_date"),
	)
}

// Date without time
type DateValue struct {
	Data time.Time          `json:"data" api:"required" format:"date"`
	Type constant.ValueDate `json:"type" default:"value/date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DateValue) RawJSON() string { return r.JSON.raw }
func (r *DateValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DateValue to a DateValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DateValueParam.Overrides()
func (r DateValue) ToParam() DateValueParam {
	return param.Override[DateValueParam](json.RawMessage(r.RawJSON()))
}

// Date without time
//
// The properties Data, Type are required.
type DateValueParam struct {
	Data time.Time `json:"data" api:"required" format:"date"`
	// This field can be elided, and will marshal its zero value as "value/date".
	Type constant.ValueDate `json:"type" default:"value/date"`
	paramObj
}

func (r DateValueParam) MarshalJSON() (data []byte, err error) {
	type shadow DateValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DateValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores dates with time information.
type DatetimeField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality DatetimeFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind DatetimeFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Meeting Time").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `meeting_time`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/datetime` for this field.
	Type constant.FieldDatetime `json:"type" default:"field/datetime"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatetimeField) RawJSON() string { return r.JSON.raw }
func (r *DatetimeField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type DatetimeFieldCardinality string

const (
	DatetimeFieldCardinalityOne  DatetimeFieldCardinality = "one"
	DatetimeFieldCardinalityMany DatetimeFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type DatetimeFieldKind string

const (
	DatetimeFieldKindSystem  DatetimeFieldKind = "system"
	DatetimeFieldKindInverse DatetimeFieldKind = "inverse"
	DatetimeFieldKindCustom  DatetimeFieldKind = "custom"
)

func DatetimeFieldDefaultValueParamOfValueDatetime(data time.Time) DatetimeFieldDefaultValueParamUnion {
	var valueDatetime DatetimeValueParam
	valueDatetime.Data = data
	return DatetimeFieldDefaultValueParamUnion{OfValueDatetime: &valueDatetime}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DatetimeFieldDefaultValueParamUnion struct {
	OfValueDatetime   *DatetimeValueParam   `json:",omitzero,inline"`
	OfCurrentDatetime *CurrentDatetimeParam `json:",omitzero,inline"`
	paramUnion
}

func (u DatetimeFieldDefaultValueParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfValueDatetime, u.OfCurrentDatetime)
}
func (u *DatetimeFieldDefaultValueParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DatetimeFieldDefaultValueParamUnion) asAny() any {
	if !param.IsOmitted(u.OfValueDatetime) {
		return u.OfValueDatetime
	} else if !param.IsOmitted(u.OfCurrentDatetime) {
		return u.OfCurrentDatetime
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DatetimeFieldDefaultValueParamUnion) GetData() *time.Time {
	if vt := u.OfValueDatetime; vt != nil {
		return &vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DatetimeFieldDefaultValueParamUnion) GetType() *string {
	if vt := u.OfValueDatetime; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCurrentDatetime; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[DatetimeFieldDefaultValueParamUnion](
		"type",
		apijson.Discriminator[DatetimeValueParam]("value/datetime"),
		apijson.Discriminator[CurrentDatetimeParam]("current_datetime"),
	)
}

// Date and time value
type DatetimeValue struct {
	Data time.Time              `json:"data" api:"required" format:"date-time"`
	Type constant.ValueDatetime `json:"type" default:"value/datetime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatetimeValue) RawJSON() string { return r.JSON.raw }
func (r *DatetimeValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DatetimeValue to a DatetimeValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DatetimeValueParam.Overrides()
func (r DatetimeValue) ToParam() DatetimeValueParam {
	return param.Override[DatetimeValueParam](json.RawMessage(r.RawJSON()))
}

// Date and time value
//
// The properties Data, Type are required.
type DatetimeValueParam struct {
	Data time.Time `json:"data" api:"required" format:"date-time"`
	// This field can be elided, and will marshal its zero value as "value/datetime".
	Type constant.ValueDatetime `json:"type" default:"value/datetime"`
	paramObj
}

func (r DatetimeValueParam) MarshalJSON() (data []byte, err error) {
	type shadow DatetimeValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatetimeValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores internet domain names.
type DomainField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality DomainFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind DomainFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Company Domain").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `company_domain`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/uri/domain` for this field.
	Type constant.FieldUriDomain `json:"type" default:"field/uri/domain"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainField) RawJSON() string { return r.JSON.raw }
func (r *DomainField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type DomainFieldCardinality string

const (
	DomainFieldCardinalityOne  DomainFieldCardinality = "one"
	DomainFieldCardinalityMany DomainFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type DomainFieldKind string

const (
	DomainFieldKindSystem  DomainFieldKind = "system"
	DomainFieldKindInverse DomainFieldKind = "inverse"
	DomainFieldKindCustom  DomainFieldKind = "custom"
)

// Internet domain name
type DomainValue struct {
	// A valid internet domain name, without protocol (e.g., 'https://') or path.
	Data string                  `json:"data" api:"required"`
	Type constant.ValueUriDomain `json:"type" default:"value/uri/domain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainValue) RawJSON() string { return r.JSON.raw }
func (r *DomainValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DomainValue to a DomainValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DomainValueParam.Overrides()
func (r DomainValue) ToParam() DomainValueParam {
	return param.Override[DomainValueParam](json.RawMessage(r.RawJSON()))
}

// Internet domain name
//
// The properties Data, Type are required.
type DomainValueParam struct {
	// A valid internet domain name, without protocol (e.g., 'https://') or path.
	Data string `json:"data" api:"required"`
	// This field can be elided, and will marshal its zero value as "value/uri/domain".
	Type constant.ValueUriDomain `json:"type" default:"value/uri/domain"`
	paramObj
}

func (r DomainValueParam) MarshalJSON() (data []byte, err error) {
	type shadow DomainValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores and validates email addresses.
type EmailField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality EmailFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind EmailFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Work Email").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `work_email`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/email` for this field.
	Type constant.FieldEmail `json:"type" default:"field/email"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailField) RawJSON() string { return r.JSON.raw }
func (r *EmailField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type EmailFieldCardinality string

const (
	EmailFieldCardinalityOne  EmailFieldCardinality = "one"
	EmailFieldCardinalityMany EmailFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type EmailFieldKind string

const (
	EmailFieldKindSystem  EmailFieldKind = "system"
	EmailFieldKindInverse EmailFieldKind = "inverse"
	EmailFieldKindCustom  EmailFieldKind = "custom"
)

// Email address value
type EmailValue struct {
	// A valid email address.
	Data string              `json:"data" api:"required" format:"email"`
	Type constant.ValueEmail `json:"type" default:"value/email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailValue) RawJSON() string { return r.JSON.raw }
func (r *EmailValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailValue to a EmailValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailValueParam.Overrides()
func (r EmailValue) ToParam() EmailValueParam {
	return param.Override[EmailValueParam](json.RawMessage(r.RawJSON()))
}

// Email address value
//
// The properties Data, Type are required.
type EmailValueParam struct {
	// A valid email address.
	Data string `json:"data" api:"required" format:"email"`
	// This field can be elided, and will marshal its zero value as "value/email".
	Type constant.ValueEmail `json:"type" default:"value/email"`
	paramObj
}

func (r EmailValueParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FieldUnion contains all possible properties and values from
// [SingleLineTextField], [MultiLineTextField], [IdentifierField], [IntegerField],
// [FloatField], [MonetaryField], [PercentageField], [BooleanField], [EmailField],
// [URLField], [DomainField], [SocialXField], [SocialLinkedInField],
// [TelephoneNumberField], [GeoField], [DateField], [DatetimeField], [ChoiceField],
// [StageField], [RelationField].
//
// Use the [FieldUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FieldUnion struct {
	ID            string                   `json:"id"`
	Cardinality   string                   `json:"cardinality"`
	CreatedAt     time.Time                `json:"created_at"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values"`
	Kind          string                   `json:"kind"`
	Name          string                   `json:"name"`
	Readonly      bool                     `json:"readonly"`
	Ref           string                   `json:"ref"`
	Required      bool                     `json:"required"`
	// Any of "field/text/single_line", "field/text/multi_line", "field/identifier",
	// "field/number/unitless_integer", "field/number/unitless_float",
	// "field/number/monetary", "field/number/percentage", "field/boolean",
	// "field/email", "field/uri/url", "field/uri/domain", "field/uri/social_x",
	// "field/uri/social_linked_in", "field/telephone_number", "field/geo",
	// "field/date", "field/datetime", "field/choice", "field/stage", "field/relation".
	Type        string    `json:"type"`
	Unique      bool      `json:"unique"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	// This field is from variant [MonetaryField].
	DefaultUnit string `json:"default_unit"`
	// This field is from variant [ChoiceField].
	Options []ChoiceFieldOption `json:"options"`
	// This field is from variant [StageField].
	Funnel Funnel `json:"funnel"`
	// This field is from variant [RelationField].
	AllowedCollections []CollectionPointer `json:"allowed_collections"`
	// This field is from variant [RelationField].
	RelationType RelationFieldRelationType `json:"relation_type"`
	// This field is from variant [RelationField].
	ReverseFieldName string `json:"reverse_field_name"`
	// This field is from variant [RelationField].
	ReverseFields []FieldPointer `json:"reverse_fields"`
	// This field is from variant [RelationField].
	SourceField FieldPointer `json:"source_field"`
	JSON        struct {
		ID                 respjson.Field
		Cardinality        respjson.Field
		CreatedAt          respjson.Field
		DefaultValues      respjson.Field
		Kind               respjson.Field
		Name               respjson.Field
		Readonly           respjson.Field
		Ref                respjson.Field
		Required           respjson.Field
		Type               respjson.Field
		Unique             respjson.Field
		UpdatedAt          respjson.Field
		Description        respjson.Field
		DefaultUnit        respjson.Field
		Options            respjson.Field
		Funnel             respjson.Field
		AllowedCollections respjson.Field
		RelationType       respjson.Field
		ReverseFieldName   respjson.Field
		ReverseFields      respjson.Field
		SourceField        respjson.Field
		raw                string
	} `json:"-"`
}

// anyField is implemented by each variant of [FieldUnion] to add type safety for
// the return type of [FieldUnion.AsAny]
type anyField interface {
	implFieldUnion()
}

func (SingleLineTextField) implFieldUnion()  {}
func (MultiLineTextField) implFieldUnion()   {}
func (IdentifierField) implFieldUnion()      {}
func (IntegerField) implFieldUnion()         {}
func (FloatField) implFieldUnion()           {}
func (MonetaryField) implFieldUnion()        {}
func (PercentageField) implFieldUnion()      {}
func (BooleanField) implFieldUnion()         {}
func (EmailField) implFieldUnion()           {}
func (URLField) implFieldUnion()             {}
func (DomainField) implFieldUnion()          {}
func (SocialXField) implFieldUnion()         {}
func (SocialLinkedInField) implFieldUnion()  {}
func (TelephoneNumberField) implFieldUnion() {}
func (GeoField) implFieldUnion()             {}
func (DateField) implFieldUnion()            {}
func (DatetimeField) implFieldUnion()        {}
func (ChoiceField) implFieldUnion()          {}
func (StageField) implFieldUnion()           {}
func (RelationField) implFieldUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FieldUnion.AsAny().(type) {
//	case moonbase.SingleLineTextField:
//	case moonbase.MultiLineTextField:
//	case moonbase.IdentifierField:
//	case moonbase.IntegerField:
//	case moonbase.FloatField:
//	case moonbase.MonetaryField:
//	case moonbase.PercentageField:
//	case moonbase.BooleanField:
//	case moonbase.EmailField:
//	case moonbase.URLField:
//	case moonbase.DomainField:
//	case moonbase.SocialXField:
//	case moonbase.SocialLinkedInField:
//	case moonbase.TelephoneNumberField:
//	case moonbase.GeoField:
//	case moonbase.DateField:
//	case moonbase.DatetimeField:
//	case moonbase.ChoiceField:
//	case moonbase.StageField:
//	case moonbase.RelationField:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FieldUnion) AsAny() anyField {
	switch u.Type {
	case "field/text/single_line":
		return u.AsFieldTextSingleLine()
	case "field/text/multi_line":
		return u.AsFieldTextMultiLine()
	case "field/identifier":
		return u.AsFieldIdentifier()
	case "field/number/unitless_integer":
		return u.AsFieldNumberUnitlessInteger()
	case "field/number/unitless_float":
		return u.AsFieldNumberUnitlessFloat()
	case "field/number/monetary":
		return u.AsFieldNumberMonetary()
	case "field/number/percentage":
		return u.AsFieldNumberPercentage()
	case "field/boolean":
		return u.AsFieldBoolean()
	case "field/email":
		return u.AsFieldEmail()
	case "field/uri/url":
		return u.AsFieldUriURL()
	case "field/uri/domain":
		return u.AsFieldUriDomain()
	case "field/uri/social_x":
		return u.AsFieldUriSocialX()
	case "field/uri/social_linked_in":
		return u.AsFieldUriSocialLinkedIn()
	case "field/telephone_number":
		return u.AsFieldTelephoneNumber()
	case "field/geo":
		return u.AsFieldGeo()
	case "field/date":
		return u.AsFieldDate()
	case "field/datetime":
		return u.AsFieldDatetime()
	case "field/choice":
		return u.AsFieldChoice()
	case "field/stage":
		return u.AsFieldStage()
	case "field/relation":
		return u.AsFieldRelation()
	}
	return nil
}

func (u FieldUnion) AsFieldTextSingleLine() (v SingleLineTextField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldTextMultiLine() (v MultiLineTextField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldIdentifier() (v IdentifierField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldNumberUnitlessInteger() (v IntegerField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldNumberUnitlessFloat() (v FloatField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldNumberMonetary() (v MonetaryField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldNumberPercentage() (v PercentageField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldBoolean() (v BooleanField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldEmail() (v EmailField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldUriURL() (v URLField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldUriDomain() (v DomainField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldUriSocialX() (v SocialXField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldUriSocialLinkedIn() (v SocialLinkedInField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldTelephoneNumber() (v TelephoneNumberField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldGeo() (v GeoField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldDate() (v DateField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldDatetime() (v DatetimeField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldChoice() (v ChoiceField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldStage() (v StageField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldUnion) AsFieldRelation() (v RelationField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FieldUnion) RawJSON() string { return u.JSON.raw }

func (r *FieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FieldDefaultValueUnion contains all possible properties and values from
// [SingleLineTextValue], [MultiLineTextValue], [IdentifierValue], [IntegerValue],
// [FloatValue], [MonetaryValue], [PercentageValue], [BooleanValue], [EmailValue],
// [URLValue], [DomainValue], [SocialXValue], [SocialLinkedInValue],
// [TelephoneNumber], [GeoValue], [DateValue], [CurrentDate], [DatetimeValue],
// [CurrentDatetime], [ChoiceValue], [FunnelStepValue], [RelationValue],
// [CurrentMember].
//
// Use the [FieldDefaultValueUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FieldDefaultValueUnion struct {
	// This field is a union of [string], [string], [string], [int64], [float64],
	// [MonetaryValueData], [float64], [bool], [string], [string], [string],
	// [SocialXValueData], [SocialLinkedInValueData], [string], [string], [time.Time],
	// [time.Time], [ChoiceFieldOption], [FunnelStep], [ItemPointer]
	Data FieldDefaultValueUnionData `json:"data"`
	// Any of "value/text/single_line", "value/text/multi_line", "value/identifier",
	// "value/number/unitless_integer", "value/number/unitless_float",
	// "value/number/monetary", "value/number/percentage", "value/boolean",
	// "value/email", "value/uri/url", "value/uri/domain", "value/uri/social_x",
	// "value/uri/social_linked_in", "value/telephone_number", "value/geo",
	// "value/date", "current_date", "value/datetime", "current_datetime",
	// "value/choice", "value/funnel_step", "value/relation", "current_member".
	Type string `json:"type"`
	JSON struct {
		Data respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

// anyFieldDefaultValue is implemented by each variant of [FieldDefaultValueUnion]
// to add type safety for the return type of [FieldDefaultValueUnion.AsAny]
type anyFieldDefaultValue interface {
	implFieldDefaultValueUnion()
}

func (SingleLineTextValue) implFieldDefaultValueUnion() {}
func (MultiLineTextValue) implFieldDefaultValueUnion()  {}
func (IdentifierValue) implFieldDefaultValueUnion()     {}
func (IntegerValue) implFieldDefaultValueUnion()        {}
func (FloatValue) implFieldDefaultValueUnion()          {}
func (MonetaryValue) implFieldDefaultValueUnion()       {}
func (PercentageValue) implFieldDefaultValueUnion()     {}
func (BooleanValue) implFieldDefaultValueUnion()        {}
func (EmailValue) implFieldDefaultValueUnion()          {}
func (URLValue) implFieldDefaultValueUnion()            {}
func (DomainValue) implFieldDefaultValueUnion()         {}
func (SocialXValue) implFieldDefaultValueUnion()        {}
func (SocialLinkedInValue) implFieldDefaultValueUnion() {}
func (TelephoneNumber) implFieldDefaultValueUnion()     {}
func (GeoValue) implFieldDefaultValueUnion()            {}
func (DateValue) implFieldDefaultValueUnion()           {}
func (CurrentDate) implFieldDefaultValueUnion()         {}
func (DatetimeValue) implFieldDefaultValueUnion()       {}
func (CurrentDatetime) implFieldDefaultValueUnion()     {}
func (ChoiceValue) implFieldDefaultValueUnion()         {}
func (FunnelStepValue) implFieldDefaultValueUnion()     {}
func (RelationValue) implFieldDefaultValueUnion()       {}
func (CurrentMember) implFieldDefaultValueUnion()       {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FieldDefaultValueUnion.AsAny().(type) {
//	case moonbase.SingleLineTextValue:
//	case moonbase.MultiLineTextValue:
//	case moonbase.IdentifierValue:
//	case moonbase.IntegerValue:
//	case moonbase.FloatValue:
//	case moonbase.MonetaryValue:
//	case moonbase.PercentageValue:
//	case moonbase.BooleanValue:
//	case moonbase.EmailValue:
//	case moonbase.URLValue:
//	case moonbase.DomainValue:
//	case moonbase.SocialXValue:
//	case moonbase.SocialLinkedInValue:
//	case moonbase.TelephoneNumber:
//	case moonbase.GeoValue:
//	case moonbase.DateValue:
//	case moonbase.CurrentDate:
//	case moonbase.DatetimeValue:
//	case moonbase.CurrentDatetime:
//	case moonbase.ChoiceValue:
//	case moonbase.FunnelStepValue:
//	case moonbase.RelationValue:
//	case moonbase.CurrentMember:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FieldDefaultValueUnion) AsAny() anyFieldDefaultValue {
	switch u.Type {
	case "value/text/single_line":
		return u.AsValueTextSingleLine()
	case "value/text/multi_line":
		return u.AsValueTextMultiLine()
	case "value/identifier":
		return u.AsValueIdentifier()
	case "value/number/unitless_integer":
		return u.AsValueNumberUnitlessInteger()
	case "value/number/unitless_float":
		return u.AsValueNumberUnitlessFloat()
	case "value/number/monetary":
		return u.AsValueNumberMonetary()
	case "value/number/percentage":
		return u.AsValueNumberPercentage()
	case "value/boolean":
		return u.AsValueBoolean()
	case "value/email":
		return u.AsValueEmail()
	case "value/uri/url":
		return u.AsValueUriURL()
	case "value/uri/domain":
		return u.AsValueUriDomain()
	case "value/uri/social_x":
		return u.AsValueUriSocialX()
	case "value/uri/social_linked_in":
		return u.AsValueUriSocialLinkedIn()
	case "value/telephone_number":
		return u.AsValueTelephoneNumber()
	case "value/geo":
		return u.AsValueGeo()
	case "value/date":
		return u.AsValueDate()
	case "current_date":
		return u.AsCurrentDate()
	case "value/datetime":
		return u.AsValueDatetime()
	case "current_datetime":
		return u.AsCurrentDatetime()
	case "value/choice":
		return u.AsValueChoice()
	case "value/funnel_step":
		return u.AsValueFunnelStep()
	case "value/relation":
		return u.AsValueRelation()
	case "current_member":
		return u.AsCurrentMember()
	}
	return nil
}

func (u FieldDefaultValueUnion) AsValueTextSingleLine() (v SingleLineTextValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueTextMultiLine() (v MultiLineTextValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueIdentifier() (v IdentifierValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueNumberUnitlessInteger() (v IntegerValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueNumberUnitlessFloat() (v FloatValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueNumberMonetary() (v MonetaryValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueNumberPercentage() (v PercentageValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueBoolean() (v BooleanValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueEmail() (v EmailValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueUriURL() (v URLValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueUriDomain() (v DomainValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueUriSocialX() (v SocialXValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueUriSocialLinkedIn() (v SocialLinkedInValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueTelephoneNumber() (v TelephoneNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueGeo() (v GeoValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueDate() (v DateValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsCurrentDate() (v CurrentDate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueDatetime() (v DatetimeValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsCurrentDatetime() (v CurrentDatetime) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueChoice() (v ChoiceValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueFunnelStep() (v FunnelStepValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsValueRelation() (v RelationValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldDefaultValueUnion) AsCurrentMember() (v CurrentMember) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FieldDefaultValueUnion) RawJSON() string { return u.JSON.raw }

func (r *FieldDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FieldDefaultValueUnionData is an implicit subunion of [FieldDefaultValueUnion].
// FieldDefaultValueUnionData provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [FieldDefaultValueUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfFloat OfBool OfTime]
type FieldDefaultValueUnionData struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [time.Time] instead of an object.
	OfTime time.Time `json:",inline"`
	// This field is from variant [MonetaryValueData].
	Currency string `json:"currency"`
	// This field is from variant [MonetaryValueData].
	InMinorUnits int64  `json:"in_minor_units"`
	URL          string `json:"url"`
	Username     string `json:"username"`
	ID           string `json:"id"`
	Color        string `json:"color"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	// This field is from variant [FunnelStep].
	StepType FunnelStepStepType `json:"step_type"`
	// This field is from variant [ItemPointer].
	Collection CollectionPointer `json:"collection"`
	JSON       struct {
		OfString     respjson.Field
		OfInt        respjson.Field
		OfFloat      respjson.Field
		OfBool       respjson.Field
		OfTime       respjson.Field
		Currency     respjson.Field
		InMinorUnits respjson.Field
		URL          respjson.Field
		Username     respjson.Field
		ID           respjson.Field
		Color        respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		StepType     respjson.Field
		Collection   respjson.Field
		raw          string
	} `json:"-"`
}

func (r *FieldDefaultValueUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A lightweight reference to a `Field`, containing the minimal information needed
// to identify it.
type FieldPointer struct {
	// Unique identifier of the field.
	ID string `json:"id" api:"required"`
	// A reference to the `Collection` containing this field.
	Collection CollectionPointer `json:"collection" api:"required"`
	// The stable, machine-readable reference identifier of the field.
	Ref string `json:"ref" api:"required"`
	// String representing the object’s type. Always `field` for this object.
	Type constant.Field `json:"type" default:"field"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Collection  respjson.Field
		Ref         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FieldPointer) RawJSON() string { return r.JSON.raw }
func (r *FieldPointer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FieldValueUnion contains all possible properties and values from
// [SingleLineTextValue], [MultiLineTextValue], [IdentifierValue], [IntegerValue],
// [FloatValue], [MonetaryValue], [PercentageValue], [BooleanValue], [EmailValue],
// [URLValue], [DomainValue], [SocialXValue], [SocialLinkedInValue],
// [TelephoneNumber], [GeoValue], [DateValue], [DatetimeValue], [ChoiceValue],
// [FunnelStepValue], [RelationValue], [[]ValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfArrayOfValues]
type FieldValueUnion struct {
	// This field will be present if the value is a [[]ValueUnion] instead of an
	// object.
	OfArrayOfValues []ValueUnion `json:",inline"`
	// This field is a union of [string], [string], [string], [int64], [float64],
	// [MonetaryValueData], [float64], [bool], [string], [string], [string],
	// [SocialXValueData], [SocialLinkedInValueData], [string], [string], [time.Time],
	// [time.Time], [ChoiceFieldOption], [FunnelStep], [ItemPointer]
	Data FieldValueUnionData `json:"data"`
	Type string              `json:"type"`
	JSON struct {
		OfArrayOfValues respjson.Field
		Data            respjson.Field
		Type            respjson.Field
		raw             string
	} `json:"-"`
}

func (u FieldValueUnion) AsSingleLineText() (v SingleLineTextValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsMultiLineText() (v MultiLineTextValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsIdentifier() (v IdentifierValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsInteger() (v IntegerValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsFloat() (v FloatValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsMonetary() (v MonetaryValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsPercentage() (v PercentageValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsBoolean() (v BooleanValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsEmail() (v EmailValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsURL() (v URLValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsDomain() (v DomainValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsX() (v SocialXValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsLinkedIn() (v SocialLinkedInValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsTelephoneNumber() (v TelephoneNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsGeo() (v GeoValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsDate() (v DateValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsDateTime() (v DatetimeValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsChoice() (v ChoiceValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsFunnelStep() (v FunnelStepValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsRelation() (v RelationValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsArrayOfValues() (v []ValueUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *FieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FieldValueUnionData is an implicit subunion of [FieldValueUnion].
// FieldValueUnionData provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [FieldValueUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfFloat OfBool OfTime]
type FieldValueUnionData struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [time.Time] instead of an object.
	OfTime time.Time `json:",inline"`
	// This field is from variant [MonetaryValueData].
	Currency string `json:"currency"`
	// This field is from variant [MonetaryValueData].
	InMinorUnits int64  `json:"in_minor_units"`
	URL          string `json:"url"`
	Username     string `json:"username"`
	ID           string `json:"id"`
	Color        string `json:"color"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	// This field is from variant [FunnelStep].
	StepType FunnelStepStepType `json:"step_type"`
	// This field is from variant [ItemPointer].
	Collection CollectionPointer `json:"collection"`
	JSON       struct {
		OfString     respjson.Field
		OfInt        respjson.Field
		OfFloat      respjson.Field
		OfBool       respjson.Field
		OfTime       respjson.Field
		Currency     respjson.Field
		InMinorUnits respjson.Field
		URL          respjson.Field
		Username     respjson.Field
		ID           respjson.Field
		Color        respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		StepType     respjson.Field
		Collection   respjson.Field
		raw          string
	} `json:"-"`
}

func (r *FieldValueUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func FieldValueParamOfSingleLineText(data string) FieldValueParamUnion {
	var variant SingleLineTextValueParam
	variant.Data = data
	return FieldValueParamUnion{OfSingleLineText: &variant}
}

func FieldValueParamOfMultiLineText(data string) FieldValueParamUnion {
	var variant MultiLineTextValueParam
	variant.Data = data
	return FieldValueParamUnion{OfMultiLineText: &variant}
}

func FieldValueParamOfIdentifier(data string) FieldValueParamUnion {
	var variant IdentifierValueParam
	variant.Data = data
	return FieldValueParamUnion{OfIdentifier: &variant}
}

func FieldValueParamOfInteger(data int64) FieldValueParamUnion {
	var variant IntegerValueParam
	variant.Data = data
	return FieldValueParamUnion{OfInteger: &variant}
}

func FieldValueParamOfFloat(data float64) FieldValueParamUnion {
	var variant FloatValueParam
	variant.Data = data
	return FieldValueParamUnion{OfFloat: &variant}
}

func FieldValueParamOfMonetary(data MonetaryValueDataParam) FieldValueParamUnion {
	var variant MonetaryValueParam
	variant.Data = data
	return FieldValueParamUnion{OfMonetary: &variant}
}

func FieldValueParamOfPercentage(data float64) FieldValueParamUnion {
	var variant PercentageValueParam
	variant.Data = data
	return FieldValueParamUnion{OfPercentage: &variant}
}

func FieldValueParamOfBoolean(data bool) FieldValueParamUnion {
	var variant BooleanValueParam
	variant.Data = data
	return FieldValueParamUnion{OfBoolean: &variant}
}

func FieldValueParamOfEmail(data string) FieldValueParamUnion {
	var variant EmailValueParam
	variant.Data = data
	return FieldValueParamUnion{OfEmail: &variant}
}

func FieldValueParamOfURL(data string) FieldValueParamUnion {
	var variant URLValueParam
	variant.Data = data
	return FieldValueParamUnion{OfURL: &variant}
}

func FieldValueParamOfDomain(data string) FieldValueParamUnion {
	var variant DomainValueParam
	variant.Data = data
	return FieldValueParamUnion{OfDomain: &variant}
}

func FieldValueParamOfX(data SocialProfileXParam) FieldValueParamUnion {
	var variant SocialXValueParam
	variant.Data = data
	return FieldValueParamUnion{OfX: &variant}
}

func FieldValueParamOfLinkedIn(data SocialProfileLinkedInParam) FieldValueParamUnion {
	var variant SocialLinkedInValueParam
	variant.Data = data
	return FieldValueParamUnion{OfLinkedIn: &variant}
}

func FieldValueParamOfTelephoneNumber(data string) FieldValueParamUnion {
	var variant TelephoneNumberParam
	variant.Data = data
	return FieldValueParamUnion{OfTelephoneNumber: &variant}
}

func FieldValueParamOfGeo(data string) FieldValueParamUnion {
	var variant GeoValueParam
	variant.Data = data
	return FieldValueParamUnion{OfGeo: &variant}
}

func FieldValueParamOfDate(data time.Time) FieldValueParamUnion {
	var variant DateValueParam
	variant.Data = data
	return FieldValueParamUnion{OfDate: &variant}
}

func FieldValueParamOfDateTime(data time.Time) FieldValueParamUnion {
	var variant DatetimeValueParam
	variant.Data = data
	return FieldValueParamUnion{OfDateTime: &variant}
}

func FieldValueParamOfChoice(data ChoiceFieldOptionPointerParam) FieldValueParamUnion {
	var variant ChoiceValueParam
	variant.Data = data
	return FieldValueParamUnion{OfChoice: &variant}
}

func FieldValueParamOfFunnelStep(data FunnelStepPointerParam) FieldValueParamUnion {
	var variant FunnelStepValueParam
	variant.Data = data
	return FieldValueParamUnion{OfFunnelStep: &variant}
}

func FieldValueParamOfRelation(data ItemPointerParam) FieldValueParamUnion {
	var variant RelationValueParam
	variant.Data = data
	return FieldValueParamUnion{OfRelation: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FieldValueParamUnion struct {
	OfSingleLineText  *SingleLineTextValueParam `json:",omitzero,inline"`
	OfMultiLineText   *MultiLineTextValueParam  `json:",omitzero,inline"`
	OfIdentifier      *IdentifierValueParam     `json:",omitzero,inline"`
	OfInteger         *IntegerValueParam        `json:",omitzero,inline"`
	OfFloat           *FloatValueParam          `json:",omitzero,inline"`
	OfMonetary        *MonetaryValueParam       `json:",omitzero,inline"`
	OfPercentage      *PercentageValueParam     `json:",omitzero,inline"`
	OfBoolean         *BooleanValueParam        `json:",omitzero,inline"`
	OfEmail           *EmailValueParam          `json:",omitzero,inline"`
	OfURL             *URLValueParam            `json:",omitzero,inline"`
	OfDomain          *DomainValueParam         `json:",omitzero,inline"`
	OfX               *SocialXValueParam        `json:",omitzero,inline"`
	OfLinkedIn        *SocialLinkedInValueParam `json:",omitzero,inline"`
	OfTelephoneNumber *TelephoneNumberParam     `json:",omitzero,inline"`
	OfGeo             *GeoValueParam            `json:",omitzero,inline"`
	OfDate            *DateValueParam           `json:",omitzero,inline"`
	OfDateTime        *DatetimeValueParam       `json:",omitzero,inline"`
	OfChoice          *ChoiceValueParam         `json:",omitzero,inline"`
	OfFunnelStep      *FunnelStepValueParam     `json:",omitzero,inline"`
	OfRelation        *RelationValueParam       `json:",omitzero,inline"`
	OfArrayOfValues   []ValueParamUnion         `json:",omitzero,inline"`
	paramUnion
}

func (u FieldValueParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSingleLineText,
		u.OfMultiLineText,
		u.OfIdentifier,
		u.OfInteger,
		u.OfFloat,
		u.OfMonetary,
		u.OfPercentage,
		u.OfBoolean,
		u.OfEmail,
		u.OfURL,
		u.OfDomain,
		u.OfX,
		u.OfLinkedIn,
		u.OfTelephoneNumber,
		u.OfGeo,
		u.OfDate,
		u.OfDateTime,
		u.OfChoice,
		u.OfFunnelStep,
		u.OfRelation,
		u.OfArrayOfValues)
}
func (u *FieldValueParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FieldValueParamUnion) asAny() any {
	if !param.IsOmitted(u.OfSingleLineText) {
		return u.OfSingleLineText
	} else if !param.IsOmitted(u.OfMultiLineText) {
		return u.OfMultiLineText
	} else if !param.IsOmitted(u.OfIdentifier) {
		return u.OfIdentifier
	} else if !param.IsOmitted(u.OfInteger) {
		return u.OfInteger
	} else if !param.IsOmitted(u.OfFloat) {
		return u.OfFloat
	} else if !param.IsOmitted(u.OfMonetary) {
		return u.OfMonetary
	} else if !param.IsOmitted(u.OfPercentage) {
		return u.OfPercentage
	} else if !param.IsOmitted(u.OfBoolean) {
		return u.OfBoolean
	} else if !param.IsOmitted(u.OfEmail) {
		return u.OfEmail
	} else if !param.IsOmitted(u.OfURL) {
		return u.OfURL
	} else if !param.IsOmitted(u.OfDomain) {
		return u.OfDomain
	} else if !param.IsOmitted(u.OfX) {
		return u.OfX
	} else if !param.IsOmitted(u.OfLinkedIn) {
		return u.OfLinkedIn
	} else if !param.IsOmitted(u.OfTelephoneNumber) {
		return u.OfTelephoneNumber
	} else if !param.IsOmitted(u.OfGeo) {
		return u.OfGeo
	} else if !param.IsOmitted(u.OfDate) {
		return u.OfDate
	} else if !param.IsOmitted(u.OfDateTime) {
		return u.OfDateTime
	} else if !param.IsOmitted(u.OfChoice) {
		return u.OfChoice
	} else if !param.IsOmitted(u.OfFunnelStep) {
		return u.OfFunnelStep
	} else if !param.IsOmitted(u.OfRelation) {
		return u.OfRelation
	} else if !param.IsOmitted(u.OfArrayOfValues) {
		return &u.OfArrayOfValues
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueParamUnion) GetType() *string {
	if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfIdentifier; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInteger; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFloat; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMonetary; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPercentage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBoolean; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEmail; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfURL; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDomain; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfX; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLinkedIn; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTelephoneNumber; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGeo; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDate; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDateTime; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfChoice; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunnelStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRelation; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FieldValueParamUnion) GetData() (res fieldValueParamUnionData) {
	if vt := u.OfSingleLineText; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfMultiLineText; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfIdentifier; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfInteger; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfFloat; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfMonetary; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfPercentage; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfBoolean; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfEmail; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfURL; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfDomain; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfX; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfLinkedIn; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfTelephoneNumber; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfGeo; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfDate; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfDateTime; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfChoice; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfFunnelStep; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfRelation; vt != nil {
		res.any = &vt.Data
	}
	return
}

// Can have the runtime types [*string], [*int64], [*float64],
// [*MonetaryValueDataParam], [*bool], [*SocialProfileXParam],
// [*SocialProfileLinkedInParam], [*time.Time], [*ChoiceFieldOptionPointerParam],
// [*FunnelStepPointerParam], [*ItemPointerParam]
type fieldValueParamUnionData struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *int64:
//	case *float64:
//	case *moonbase.MonetaryValueDataParam:
//	case *bool:
//	case *moonbase.SocialProfileXParam:
//	case *moonbase.SocialProfileLinkedInParam:
//	case *time.Time:
//	case *moonbase.ChoiceFieldOptionPointerParam:
//	case *moonbase.FunnelStepPointerParam:
//	case *moonbase.ItemPointerParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u fieldValueParamUnionData) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetCurrency() *string {
	switch vt := u.any.(type) {
	case *MonetaryValueDataParam:
		return &vt.Currency
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetInMinorUnits() *int64 {
	switch vt := u.any.(type) {
	case *MonetaryValueDataParam:
		return &vt.InMinorUnits
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetURL() *string {
	switch vt := u.any.(type) {
	case *SocialProfileXParam:
		return paramutil.AddrIfPresent(vt.URL)
	case *SocialProfileLinkedInParam:
		return paramutil.AddrIfPresent(vt.URL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetUsername() *string {
	switch vt := u.any.(type) {
	case *SocialProfileXParam:
		return paramutil.AddrIfPresent(vt.Username)
	case *SocialProfileLinkedInParam:
		return paramutil.AddrIfPresent(vt.Username)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetID() *string {
	switch vt := u.any.(type) {
	case *ChoiceFieldOptionPointerParam:
		return (*string)(&vt.ID)
	case *FunnelStepPointerParam:
		return (*string)(&vt.ID)
	case *ItemPointerParam:
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetType() *string {
	switch vt := u.any.(type) {
	case *ChoiceFieldOptionPointerParam:
		return (*string)(&vt.Type)
	case *FunnelStepPointerParam:
		return (*string)(&vt.Type)
	case *ItemPointerParam:
		return (*string)(&vt.Type)
	}
	return nil
}

// A field that stores decimal numbers with floating-point precision.
type FloatField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality FloatFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind FloatFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Rating").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `rating`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/number/unitless_float` for this field.
	Type constant.FieldNumberUnitlessFloat `json:"type" default:"field/number/unitless_float"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatField) RawJSON() string { return r.JSON.raw }
func (r *FloatField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type FloatFieldCardinality string

const (
	FloatFieldCardinalityOne  FloatFieldCardinality = "one"
	FloatFieldCardinalityMany FloatFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type FloatFieldKind string

const (
	FloatFieldKindSystem  FloatFieldKind = "system"
	FloatFieldKindInverse FloatFieldKind = "inverse"
	FloatFieldKindCustom  FloatFieldKind = "custom"
)

// Floating point number
type FloatValue struct {
	Data float64                           `json:"data" api:"required"`
	Type constant.ValueNumberUnitlessFloat `json:"type" default:"value/number/unitless_float"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatValue) RawJSON() string { return r.JSON.raw }
func (r *FloatValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FloatValue to a FloatValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FloatValueParam.Overrides()
func (r FloatValue) ToParam() FloatValueParam {
	return param.Override[FloatValueParam](json.RawMessage(r.RawJSON()))
}

// Floating point number
//
// The properties Data, Type are required.
type FloatValueParam struct {
	Data float64 `json:"data" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "value/number/unitless_float".
	Type constant.ValueNumberUnitlessFloat `json:"type" default:"value/number/unitless_float"`
	paramObj
}

func (r FloatValueParam) MarshalJSON() (data []byte, err error) {
	type shadow FloatValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FloatValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A pointer to a Funnel, used as a parameter.
//
// The properties ID, Type are required.
type FunnelPointerParam struct {
	// The ID of the funnel.
	ID string `json:"id" api:"required"`
	// String representing the object's type. Always `funnel` for this parameter.
	//
	// This field can be elided, and will marshal its zero value as "funnel".
	Type constant.Funnel `json:"type" default:"funnel"`
	paramObj
}

func (r FunnelPointerParam) MarshalJSON() (data []byte, err error) {
	type shadow FunnelPointerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelPointerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Funnel step value
type FunnelStepValue struct {
	// A specific funnel step, as configured on the Funnel.
	Data FunnelStep               `json:"data" api:"required"`
	Type constant.ValueFunnelStep `json:"type" default:"value/funnel_step"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelStepValue) RawJSON() string { return r.JSON.raw }
func (r *FunnelStepValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Funnel step value
//
// The properties Data, Type are required.
type FunnelStepValueParam struct {
	// A specific funnel step, as configured on the Funnel.
	Data FunnelStepPointerParam `json:"data,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "value/funnel_step".
	Type constant.ValueFunnelStep `json:"type" default:"value/funnel_step"`
	paramObj
}

func (r FunnelStepValueParam) MarshalJSON() (data []byte, err error) {
	type shadow FunnelStepValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelStepValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores geographic coordinates or location data.
type GeoField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality GeoFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind GeoFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Location").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `location`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/geo` for this field.
	Type constant.FieldGeo `json:"type" default:"field/geo"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeoField) RawJSON() string { return r.JSON.raw }
func (r *GeoField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type GeoFieldCardinality string

const (
	GeoFieldCardinalityOne  GeoFieldCardinality = "one"
	GeoFieldCardinalityMany GeoFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type GeoFieldKind string

const (
	GeoFieldKindSystem  GeoFieldKind = "system"
	GeoFieldKindInverse GeoFieldKind = "inverse"
	GeoFieldKindCustom  GeoFieldKind = "custom"
)

// Geographic coordinate value
type GeoValue struct {
	// A string that represents some geographic location. The exact format may vary
	// based on context.
	Data string            `json:"data" api:"required"`
	Type constant.ValueGeo `json:"type" default:"value/geo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeoValue) RawJSON() string { return r.JSON.raw }
func (r *GeoValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GeoValue to a GeoValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GeoValueParam.Overrides()
func (r GeoValue) ToParam() GeoValueParam {
	return param.Override[GeoValueParam](json.RawMessage(r.RawJSON()))
}

// Geographic coordinate value
//
// The properties Data, Type are required.
type GeoValueParam struct {
	// A string that represents some geographic location. The exact format may vary
	// based on context.
	Data string `json:"data" api:"required"`
	// This field can be elided, and will marshal its zero value as "value/geo".
	Type constant.ValueGeo `json:"type" default:"value/geo"`
	paramObj
}

func (r GeoValueParam) MarshalJSON() (data []byte, err error) {
	type shadow GeoValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeoValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores opaque external identifiers verbatim.
type IdentifierField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality IdentifierFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// Any of "system", "inverse", "custom".
	Kind IdentifierFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Stripe Id").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `stripe_id`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/identifier` for this field.
	Type constant.FieldIdentifier `json:"type" default:"field/identifier"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentifierField) RawJSON() string { return r.JSON.raw }
func (r *IdentifierField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type IdentifierFieldCardinality string

const (
	IdentifierFieldCardinalityOne  IdentifierFieldCardinality = "one"
	IdentifierFieldCardinalityMany IdentifierFieldCardinality = "many"
)

type IdentifierFieldKind string

const (
	IdentifierFieldKindSystem  IdentifierFieldKind = "system"
	IdentifierFieldKindInverse IdentifierFieldKind = "inverse"
	IdentifierFieldKindCustom  IdentifierFieldKind = "custom"
)

// Identifier string
type IdentifierValue struct {
	// An external identifier as text, uo to 255 characters in length.
	Data string                   `json:"data" api:"required"`
	Type constant.ValueIdentifier `json:"type" default:"value/identifier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentifierValue) RawJSON() string { return r.JSON.raw }
func (r *IdentifierValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this IdentifierValue to a IdentifierValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IdentifierValueParam.Overrides()
func (r IdentifierValue) ToParam() IdentifierValueParam {
	return param.Override[IdentifierValueParam](json.RawMessage(r.RawJSON()))
}

// Identifier string
//
// The properties Data, Type are required.
type IdentifierValueParam struct {
	// An external identifier as text, uo to 255 characters in length.
	Data string `json:"data" api:"required"`
	// This field can be elided, and will marshal its zero value as "value/identifier".
	Type constant.ValueIdentifier `json:"type" default:"value/identifier"`
	paramObj
}

func (r IdentifierValueParam) MarshalJSON() (data []byte, err error) {
	type shadow IdentifierValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IdentifierValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores whole numbers without decimal places.
type IntegerField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality IntegerFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind IntegerFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Employee Count").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `employee_count`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/number/unitless_integer` for this
	// field.
	Type constant.FieldNumberUnitlessInteger `json:"type" default:"field/number/unitless_integer"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegerField) RawJSON() string { return r.JSON.raw }
func (r *IntegerField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type IntegerFieldCardinality string

const (
	IntegerFieldCardinalityOne  IntegerFieldCardinality = "one"
	IntegerFieldCardinalityMany IntegerFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type IntegerFieldKind string

const (
	IntegerFieldKindSystem  IntegerFieldKind = "system"
	IntegerFieldKindInverse IntegerFieldKind = "inverse"
	IntegerFieldKindCustom  IntegerFieldKind = "custom"
)

// Integer value without units
type IntegerValue struct {
	Data int64                               `json:"data" api:"required"`
	Type constant.ValueNumberUnitlessInteger `json:"type" default:"value/number/unitless_integer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegerValue) RawJSON() string { return r.JSON.raw }
func (r *IntegerValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this IntegerValue to a IntegerValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IntegerValueParam.Overrides()
func (r IntegerValue) ToParam() IntegerValueParam {
	return param.Override[IntegerValueParam](json.RawMessage(r.RawJSON()))
}

// Integer value without units
//
// The properties Data, Type are required.
type IntegerValueParam struct {
	Data int64 `json:"data" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "value/number/unitless_integer".
	Type constant.ValueNumberUnitlessInteger `json:"type" default:"value/number/unitless_integer"`
	paramObj
}

func (r IntegerValueParam) MarshalJSON() (data []byte, err error) {
	type shadow IntegerValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegerValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Item represents a single record or row within a Collection. It holds a set of
// `values` corresponding to the Collection's `fields`.
type Item struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// A lightweight reference to a `Collection`, containing the minimal information
	// needed to identify it.
	Collection CollectionPointer `json:"collection" api:"required"`
	// String representing the object’s type. Always `item` for this object.
	Type constant.Item `json:"type" default:"item"`
	// A hash where keys are the `ref` of a `Field` and values are the data stored for
	// that field.
	Values map[string]FieldValueUnion `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Collection  respjson.Field
		Type        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Item) RawJSON() string { return r.JSON.raw }
func (r *Item) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to an `Item` within a specific `Collection`, providing the context
// needed to locate the item.
type ItemPointer struct {
	// Unique identifier of the item.
	ID string `json:"id" api:"required"`
	// A reference to the `Collection` containing this item.
	Collection CollectionPointer `json:"collection" api:"required"`
	// String representing the object’s type. Always `item` for this object.
	Type constant.Item `json:"type" default:"item"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Collection  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemPointer) RawJSON() string { return r.JSON.raw }
func (r *ItemPointer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A lightweight reference to an `Item` used in request bodies.
//
// The properties ID, Type are required.
type ItemPointerParam struct {
	// Unique identifier of the item.
	ID string `json:"id" api:"required"`
	// String representing the object’s type. Always `item` for this object.
	//
	// This field can be elided, and will marshal its zero value as "item".
	Type constant.Item `json:"type" default:"item"`
	paramObj
}

func (r ItemPointerParam) MarshalJSON() (data []byte, err error) {
	type shadow ItemPointerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemPointerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ItemsFilterUnion contains all possible properties and values from
// [ItemsFilterValueMatches], [ItemsFilterValueExists], [ItemsFilterAndGroup],
// [ItemsFilterOrGroup], [ItemsFilterNotGroup].
//
// Use the [ItemsFilterUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ItemsFilterUnion struct {
	Field string `json:"field"`
	// Any of nil, "exists", "and", "or", "not".
	Op string `json:"op"`
	// This field is from variant [ItemsFilterValueMatches].
	Value   ItemsFilterValueMatchesValueUnion `json:"value"`
	Filters []ItemsFilterUnion                `json:"filters"`
	// This field is from variant [ItemsFilterNotGroup].
	Filter ItemsFilterUnion `json:"filter"`
	JSON   struct {
		Field   respjson.Field
		Op      respjson.Field
		Value   respjson.Field
		Filters respjson.Field
		Filter  respjson.Field
		raw     string
	} `json:"-"`
}

func (u ItemsFilterUnion) AsItemsFilterValueMatches() (v ItemsFilterValueMatches) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ItemsFilterUnion) AsExists() (v ItemsFilterValueExists) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ItemsFilterUnion) AsAnd() (v ItemsFilterAndGroup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ItemsFilterUnion) AsOr() (v ItemsFilterOrGroup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ItemsFilterUnion) AsNot() (v ItemsFilterNotGroup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ItemsFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *ItemsFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ItemsFilterUnion to a ItemsFilterUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ItemsFilterUnionParam.Overrides()
func (r ItemsFilterUnion) ToParam() ItemsFilterUnionParam {
	return param.Override[ItemsFilterUnionParam](json.RawMessage(r.RawJSON()))
}

func ItemsFilterParamOfItemsFilterValueMatches[T string | float64 | bool](field string, op ItemsFilterValueMatchesOp, value T) ItemsFilterUnionParam {
	var variant ItemsFilterValueMatchesParam
	variant.Field = field
	variant.Op = op
	switch v := any(value).(type) {
	case string:
		variant.Value.OfString = param.NewOpt(v)
	case float64:
		variant.Value.OfFloat = param.NewOpt(v)
	case bool:
		variant.Value.OfBool = param.NewOpt(v)
	}
	return ItemsFilterUnionParam{OfItemsFilterValueMatches: &variant}
}

func ItemsFilterParamOfExists(field string) ItemsFilterUnionParam {
	var exists ItemsFilterValueExistsParam
	exists.Field = field
	return ItemsFilterUnionParam{OfExists: &exists}
}

func ItemsFilterParamOfAnd(filters []ItemsFilterUnionParam) ItemsFilterUnionParam {
	var and ItemsFilterAndGroupParam
	and.Filters = filters
	return ItemsFilterUnionParam{OfAnd: &and}
}

func ItemsFilterParamOfOr(filters []ItemsFilterUnionParam) ItemsFilterUnionParam {
	var or ItemsFilterOrGroupParam
	or.Filters = filters
	return ItemsFilterUnionParam{OfOr: &or}
}

func ItemsFilterParamOfNot[
	T ItemsFilterValueMatchesParam | ItemsFilterValueExistsParam | ItemsFilterAndGroupParam | ItemsFilterOrGroupParam | ItemsFilterNotGroupParam,
](filter T) ItemsFilterUnionParam {
	var not ItemsFilterNotGroupParam
	switch v := any(filter).(type) {
	case ItemsFilterValueMatchesParam:
		not.Filter.OfItemsFilterValueMatches = &v
	case ItemsFilterValueExistsParam:
		not.Filter.OfExists = &v
	case ItemsFilterAndGroupParam:
		not.Filter.OfAnd = &v
	case ItemsFilterOrGroupParam:
		not.Filter.OfOr = &v
	case ItemsFilterNotGroupParam:
		not.Filter.OfNot = &v
	}
	return ItemsFilterUnionParam{OfNot: &not}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ItemsFilterUnionParam struct {
	OfItemsFilterValueMatches *ItemsFilterValueMatchesParam `json:",omitzero,inline"`
	OfExists                  *ItemsFilterValueExistsParam  `json:",omitzero,inline"`
	OfAnd                     *ItemsFilterAndGroupParam     `json:",omitzero,inline"`
	OfOr                      *ItemsFilterOrGroupParam      `json:",omitzero,inline"`
	OfNot                     *ItemsFilterNotGroupParam     `json:",omitzero,inline"`
	paramUnion
}

func (u ItemsFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfItemsFilterValueMatches,
		u.OfExists,
		u.OfAnd,
		u.OfOr,
		u.OfNot)
}
func (u *ItemsFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ItemsFilterUnionParam) asAny() any {
	if !param.IsOmitted(u.OfItemsFilterValueMatches) {
		return u.OfItemsFilterValueMatches
	} else if !param.IsOmitted(u.OfExists) {
		return u.OfExists
	} else if !param.IsOmitted(u.OfAnd) {
		return u.OfAnd
	} else if !param.IsOmitted(u.OfOr) {
		return u.OfOr
	} else if !param.IsOmitted(u.OfNot) {
		return u.OfNot
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ItemsFilterUnionParam) GetValue() *ItemsFilterValueMatchesValueUnionParam {
	if vt := u.OfItemsFilterValueMatches; vt != nil {
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ItemsFilterUnionParam) GetFilter() *ItemsFilterUnionParam {
	if vt := u.OfNot; vt != nil {
		return &vt.Filter
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ItemsFilterUnionParam) GetField() *string {
	if vt := u.OfItemsFilterValueMatches; vt != nil {
		return (*string)(&vt.Field)
	} else if vt := u.OfExists; vt != nil {
		return (*string)(&vt.Field)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ItemsFilterUnionParam) GetOp() *string {
	if vt := u.OfItemsFilterValueMatches; vt != nil {
		return (*string)(&vt.Op)
	} else if vt := u.OfExists; vt != nil {
		return (*string)(&vt.Op)
	} else if vt := u.OfAnd; vt != nil {
		return (*string)(&vt.Op)
	} else if vt := u.OfOr; vt != nil {
		return (*string)(&vt.Op)
	} else if vt := u.OfNot; vt != nil {
		return (*string)(&vt.Op)
	}
	return nil
}

// Returns a pointer to the underlying variant's Filters property, if present.
func (u ItemsFilterUnionParam) GetFilters() []ItemsFilterUnionParam {
	if vt := u.OfAnd; vt != nil {
		return vt.Filters
	} else if vt := u.OfOr; vt != nil {
		return vt.Filters
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ItemsFilterUnionParam](
		"op",
		apijson.Discriminator[ItemsFilterValueMatchesParam]("starts_with"),
		apijson.Discriminator[ItemsFilterValueMatchesParam]("ends_with"),
		apijson.Discriminator[ItemsFilterValueMatchesParam]("contains"),
		apijson.Discriminator[ItemsFilterValueMatchesParam]("not_contains"),
		apijson.Discriminator[ItemsFilterValueMatchesParam]("eq"),
		apijson.Discriminator[ItemsFilterValueMatchesParam]("not_eq"),
		apijson.Discriminator[ItemsFilterValueMatchesParam]("gt"),
		apijson.Discriminator[ItemsFilterValueMatchesParam]("lt"),
		apijson.Discriminator[ItemsFilterValueMatchesParam]("gte"),
		apijson.Discriminator[ItemsFilterValueMatchesParam]("lte"),
		apijson.Discriminator[ItemsFilterValueExistsParam]("exists"),
		apijson.Discriminator[ItemsFilterAndGroupParam]("and"),
		apijson.Discriminator[ItemsFilterOrGroupParam]("or"),
		apijson.Discriminator[ItemsFilterNotGroupParam]("not"),
	)
}

// Include only items that match ALL of the filters in `filters`.
type ItemsFilterAndGroup struct {
	// An array of filters, ALL of which must be satisfied for this `and` filter to
	// match.
	Filters []ItemsFilterUnion `json:"filters" api:"required"`
	Op      constant.And       `json:"op" default:"and"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Op          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemsFilterAndGroup) RawJSON() string { return r.JSON.raw }
func (r *ItemsFilterAndGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ItemsFilterAndGroup to a ItemsFilterAndGroupParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ItemsFilterAndGroupParam.Overrides()
func (r ItemsFilterAndGroup) ToParam() ItemsFilterAndGroupParam {
	return param.Override[ItemsFilterAndGroupParam](json.RawMessage(r.RawJSON()))
}

// Include only items that match ALL of the filters in `filters`.
//
// The properties Filters, Op are required.
type ItemsFilterAndGroupParam struct {
	// An array of filters, ALL of which must be satisfied for this `and` filter to
	// match.
	Filters []ItemsFilterUnionParam `json:"filters,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "and".
	Op constant.And `json:"op" default:"and"`
	paramObj
}

func (r ItemsFilterAndGroupParam) MarshalJSON() (data []byte, err error) {
	type shadow ItemsFilterAndGroupParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemsFilterAndGroupParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Include only items that do NOT match the nested `filter`.
type ItemsFilterNotGroup struct {
	// A nested filter which must NOT match in order for this `not` filter to match.
	Filter ItemsFilterUnion `json:"filter" api:"required"`
	Op     constant.Not     `json:"op" default:"not"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Op          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemsFilterNotGroup) RawJSON() string { return r.JSON.raw }
func (r *ItemsFilterNotGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ItemsFilterNotGroup to a ItemsFilterNotGroupParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ItemsFilterNotGroupParam.Overrides()
func (r ItemsFilterNotGroup) ToParam() ItemsFilterNotGroupParam {
	return param.Override[ItemsFilterNotGroupParam](json.RawMessage(r.RawJSON()))
}

// Include only items that do NOT match the nested `filter`.
//
// The properties Filter, Op are required.
type ItemsFilterNotGroupParam struct {
	// A nested filter which must NOT match in order for this `not` filter to match.
	Filter ItemsFilterUnionParam `json:"filter,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "not".
	Op constant.Not `json:"op" default:"not"`
	paramObj
}

func (r ItemsFilterNotGroupParam) MarshalJSON() (data []byte, err error) {
	type shadow ItemsFilterNotGroupParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemsFilterNotGroupParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Include only items that match ANY of the filters in `filters`.
type ItemsFilterOrGroup struct {
	// An array of filters, ANY of which must be satisfied for this `or` filter to
	// match.
	Filters []ItemsFilterUnion `json:"filters" api:"required"`
	Op      constant.Or        `json:"op" default:"or"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Op          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemsFilterOrGroup) RawJSON() string { return r.JSON.raw }
func (r *ItemsFilterOrGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ItemsFilterOrGroup to a ItemsFilterOrGroupParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ItemsFilterOrGroupParam.Overrides()
func (r ItemsFilterOrGroup) ToParam() ItemsFilterOrGroupParam {
	return param.Override[ItemsFilterOrGroupParam](json.RawMessage(r.RawJSON()))
}

// Include only items that match ANY of the filters in `filters`.
//
// The properties Filters, Op are required.
type ItemsFilterOrGroupParam struct {
	// An array of filters, ANY of which must be satisfied for this `or` filter to
	// match.
	Filters []ItemsFilterUnionParam `json:"filters,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "or".
	Op constant.Or `json:"op" default:"or"`
	paramObj
}

func (r ItemsFilterOrGroupParam) MarshalJSON() (data []byte, err error) {
	type shadow ItemsFilterOrGroupParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemsFilterOrGroupParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Include only items that have a value in the given `field`.
type ItemsFilterValueExists struct {
	// The id or key of the field for which a value must exist, or a path to the field
	// for which a value must exist.
	Field string          `json:"field" api:"required"`
	Op    constant.Exists `json:"op" default:"exists"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Op          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemsFilterValueExists) RawJSON() string { return r.JSON.raw }
func (r *ItemsFilterValueExists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ItemsFilterValueExists to a ItemsFilterValueExistsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ItemsFilterValueExistsParam.Overrides()
func (r ItemsFilterValueExists) ToParam() ItemsFilterValueExistsParam {
	return param.Override[ItemsFilterValueExistsParam](json.RawMessage(r.RawJSON()))
}

// Include only items that have a value in the given `field`.
//
// The properties Field, Op are required.
type ItemsFilterValueExistsParam struct {
	// The id or key of the field for which a value must exist, or a path to the field
	// for which a value must exist.
	Field string `json:"field" api:"required"`
	// This field can be elided, and will marshal its zero value as "exists".
	Op constant.Exists `json:"op" default:"exists"`
	paramObj
}

func (r ItemsFilterValueExistsParam) MarshalJSON() (data []byte, err error) {
	type shadow ItemsFilterValueExistsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemsFilterValueExistsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Include only items with a value in the given `field` that satisfies the `op`
// condition.
type ItemsFilterValueMatches struct {
	// The id or key of the field in which values are matched, or a path to the field
	// in which values are matched.
	Field string `json:"field" api:"required"`
	// The matching operator for this filter.
	//
	// Any of "starts_with", "ends_with", "contains", "not_contains", "eq", "not_eq",
	// "gt", "lt", "gte", "lte".
	Op ItemsFilterValueMatchesOp `json:"op" api:"required"`
	// The value to match against. Use ISO8601 format for dates and datetime fields.
	// For date fields, the time portion of the date-time will be ignored. For currency
	// fields, the amount should be in the smallest unit of currency (eg: cents for
	// USD).
	Value ItemsFilterValueMatchesValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Op          respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemsFilterValueMatches) RawJSON() string { return r.JSON.raw }
func (r *ItemsFilterValueMatches) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ItemsFilterValueMatches to a ItemsFilterValueMatchesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ItemsFilterValueMatchesParam.Overrides()
func (r ItemsFilterValueMatches) ToParam() ItemsFilterValueMatchesParam {
	return param.Override[ItemsFilterValueMatchesParam](json.RawMessage(r.RawJSON()))
}

// The matching operator for this filter.
type ItemsFilterValueMatchesOp string

const (
	ItemsFilterValueMatchesOpStartsWith  ItemsFilterValueMatchesOp = "starts_with"
	ItemsFilterValueMatchesOpEndsWith    ItemsFilterValueMatchesOp = "ends_with"
	ItemsFilterValueMatchesOpContains    ItemsFilterValueMatchesOp = "contains"
	ItemsFilterValueMatchesOpNotContains ItemsFilterValueMatchesOp = "not_contains"
	ItemsFilterValueMatchesOpEq          ItemsFilterValueMatchesOp = "eq"
	ItemsFilterValueMatchesOpNotEq       ItemsFilterValueMatchesOp = "not_eq"
	ItemsFilterValueMatchesOpGt          ItemsFilterValueMatchesOp = "gt"
	ItemsFilterValueMatchesOpLt          ItemsFilterValueMatchesOp = "lt"
	ItemsFilterValueMatchesOpGte         ItemsFilterValueMatchesOp = "gte"
	ItemsFilterValueMatchesOpLte         ItemsFilterValueMatchesOp = "lte"
)

// ItemsFilterValueMatchesValueUnion contains all possible properties and values
// from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ItemsFilterValueMatchesValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ItemsFilterValueMatchesValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ItemsFilterValueMatchesValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ItemsFilterValueMatchesValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ItemsFilterValueMatchesValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ItemsFilterValueMatchesValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Include only items with a value in the given `field` that satisfies the `op`
// condition.
//
// The properties Field, Op, Value are required.
type ItemsFilterValueMatchesParam struct {
	// The id or key of the field in which values are matched, or a path to the field
	// in which values are matched.
	Field string `json:"field" api:"required"`
	// The matching operator for this filter.
	//
	// Any of "starts_with", "ends_with", "contains", "not_contains", "eq", "not_eq",
	// "gt", "lt", "gte", "lte".
	Op ItemsFilterValueMatchesOp `json:"op,omitzero" api:"required"`
	// The value to match against. Use ISO8601 format for dates and datetime fields.
	// For date fields, the time portion of the date-time will be ignored. For currency
	// fields, the amount should be in the smallest unit of currency (eg: cents for
	// USD).
	Value ItemsFilterValueMatchesValueUnionParam `json:"value,omitzero" api:"required"`
	paramObj
}

func (r ItemsFilterValueMatchesParam) MarshalJSON() (data []byte, err error) {
	type shadow ItemsFilterValueMatchesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemsFilterValueMatchesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ItemsFilterValueMatchesValueUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ItemsFilterValueMatchesValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ItemsFilterValueMatchesValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ItemsFilterValueMatchesValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// A field that stores monetary amounts with currency information.
type MonetaryField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality MonetaryFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The default currency for the field, as a 3-letter ISO 4217 code (e.g., `USD`,
	// `EUR`, `GBP`).
	DefaultUnit   string                   `json:"default_unit" api:"required"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind MonetaryFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Deal Value").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `deal_value`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/number/monetary` for this field.
	Type constant.FieldNumberMonetary `json:"type" default:"field/number/monetary"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultUnit   respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonetaryField) RawJSON() string { return r.JSON.raw }
func (r *MonetaryField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type MonetaryFieldCardinality string

const (
	MonetaryFieldCardinalityOne  MonetaryFieldCardinality = "one"
	MonetaryFieldCardinalityMany MonetaryFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type MonetaryFieldKind string

const (
	MonetaryFieldKindSystem  MonetaryFieldKind = "system"
	MonetaryFieldKindInverse MonetaryFieldKind = "inverse"
	MonetaryFieldKindCustom  MonetaryFieldKind = "custom"
)

// Monetary or currency value
type MonetaryValue struct {
	// A monetary amount is composed of the amount in the smallest unit of a currency
	// and an ISO currency code.
	Data MonetaryValueData            `json:"data" api:"required"`
	Type constant.ValueNumberMonetary `json:"type" default:"value/number/monetary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonetaryValue) RawJSON() string { return r.JSON.raw }
func (r *MonetaryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MonetaryValue to a MonetaryValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MonetaryValueParam.Overrides()
func (r MonetaryValue) ToParam() MonetaryValueParam {
	return param.Override[MonetaryValueParam](json.RawMessage(r.RawJSON()))
}

// A monetary amount is composed of the amount in the smallest unit of a currency
// and an ISO currency code.
type MonetaryValueData struct {
	// The 3-letter ISO 4217 currency code
	Currency string `json:"currency" api:"required"`
	// The amount in the minor units of the currency. For example, $10 (10 USD) would
	// be 1000. Minor units conversion depends on the currency.
	InMinorUnits int64 `json:"in_minor_units" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency     respjson.Field
		InMinorUnits respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonetaryValueData) RawJSON() string { return r.JSON.raw }
func (r *MonetaryValueData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monetary or currency value
//
// The properties Data, Type are required.
type MonetaryValueParam struct {
	// A monetary amount is composed of the amount in the smallest unit of a currency
	// and an ISO currency code.
	Data MonetaryValueDataParam `json:"data,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "value/number/monetary".
	Type constant.ValueNumberMonetary `json:"type" default:"value/number/monetary"`
	paramObj
}

func (r MonetaryValueParam) MarshalJSON() (data []byte, err error) {
	type shadow MonetaryValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonetaryValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A monetary amount is composed of the amount in the smallest unit of a currency
// and an ISO currency code.
//
// The properties Currency, InMinorUnits are required.
type MonetaryValueDataParam struct {
	// The 3-letter ISO 4217 currency code
	Currency string `json:"currency" api:"required"`
	// The amount in the minor units of the currency. For example, $10 (10 USD) would
	// be 1000. Minor units conversion depends on the currency.
	InMinorUnits int64 `json:"in_minor_units" api:"required"`
	paramObj
}

func (r MonetaryValueDataParam) MarshalJSON() (data []byte, err error) {
	type shadow MonetaryValueDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonetaryValueDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores multiple lines of text with line breaks preserved.
type MultiLineTextField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality MultiLineTextFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind MultiLineTextFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Description").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `description`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/text/multi_line` for this field.
	Type constant.FieldTextMultiLine `json:"type" default:"field/text/multi_line"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultiLineTextField) RawJSON() string { return r.JSON.raw }
func (r *MultiLineTextField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type MultiLineTextFieldCardinality string

const (
	MultiLineTextFieldCardinalityOne  MultiLineTextFieldCardinality = "one"
	MultiLineTextFieldCardinalityMany MultiLineTextFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type MultiLineTextFieldKind string

const (
	MultiLineTextFieldKindSystem  MultiLineTextFieldKind = "system"
	MultiLineTextFieldKindInverse MultiLineTextFieldKind = "inverse"
	MultiLineTextFieldKindCustom  MultiLineTextFieldKind = "custom"
)

// Multiple lines of text
type MultiLineTextValue struct {
	// Text which may contain line breaks, can be up to 65,536 characters long. Do not
	// use markdown formatting, just plain text.
	Data string                      `json:"data" api:"required"`
	Type constant.ValueTextMultiLine `json:"type" default:"value/text/multi_line"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultiLineTextValue) RawJSON() string { return r.JSON.raw }
func (r *MultiLineTextValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MultiLineTextValue to a MultiLineTextValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MultiLineTextValueParam.Overrides()
func (r MultiLineTextValue) ToParam() MultiLineTextValueParam {
	return param.Override[MultiLineTextValueParam](json.RawMessage(r.RawJSON()))
}

// Multiple lines of text
//
// The properties Data, Type are required.
type MultiLineTextValueParam struct {
	// Text which may contain line breaks, can be up to 65,536 characters long. Do not
	// use markdown formatting, just plain text.
	Data string `json:"data" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "value/text/multi_line".
	Type constant.ValueTextMultiLine `json:"type" default:"value/text/multi_line"`
	paramObj
}

func (r MultiLineTextValueParam) MarshalJSON() (data []byte, err error) {
	type shadow MultiLineTextValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MultiLineTextValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores percentage values as decimal numbers.
type PercentageField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality PercentageFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind PercentageFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Win Probability").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `win_probability`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/number/percentage` for this field.
	Type constant.FieldNumberPercentage `json:"type" default:"field/number/percentage"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PercentageField) RawJSON() string { return r.JSON.raw }
func (r *PercentageField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type PercentageFieldCardinality string

const (
	PercentageFieldCardinalityOne  PercentageFieldCardinality = "one"
	PercentageFieldCardinalityMany PercentageFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type PercentageFieldKind string

const (
	PercentageFieldKindSystem  PercentageFieldKind = "system"
	PercentageFieldKindInverse PercentageFieldKind = "inverse"
	PercentageFieldKindCustom  PercentageFieldKind = "custom"
)

// Percentage numeric value
type PercentageValue struct {
	// A floating-point number representing a percentage value, for example 50.21 for
	// 50.21% or -1000 for -1000% etc.
	Data float64                        `json:"data" api:"required"`
	Type constant.ValueNumberPercentage `json:"type" default:"value/number/percentage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PercentageValue) RawJSON() string { return r.JSON.raw }
func (r *PercentageValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PercentageValue to a PercentageValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PercentageValueParam.Overrides()
func (r PercentageValue) ToParam() PercentageValueParam {
	return param.Override[PercentageValueParam](json.RawMessage(r.RawJSON()))
}

// Percentage numeric value
//
// The properties Data, Type are required.
type PercentageValueParam struct {
	// A floating-point number representing a percentage value, for example 50.21 for
	// 50.21% or -1000 for -1000% etc.
	Data float64 `json:"data" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "value/number/percentage".
	Type constant.ValueNumberPercentage `json:"type" default:"value/number/percentage"`
	paramObj
}

func (r PercentageValueParam) MarshalJSON() (data []byte, err error) {
	type shadow PercentageValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PercentageValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that creates a link between items in different collections, enabling
// cross-collection relationships.
type RelationField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// The set of collections that are valid targets for this relation.
	AllowedCollections []CollectionPointer `json:"allowed_collections" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality RelationFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind RelationFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Account").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `account`).
	Ref string `json:"ref" api:"required"`
	// The type of relationship. Can be `one_way` for simple references or `two_way`
	// for bidirectional relationships.
	//
	// Any of "one_way", "two_way".
	RelationType RelationFieldRelationType `json:"relation_type" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/relation` for this field.
	Type constant.FieldRelation `json:"type" default:"field/relation"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// The name given to auto-created reverse fields on target collections. Only
	// present on `two_way` source fields.
	ReverseFieldName string `json:"reverse_field_name"`
	// A list of reverse fields created on each target collection. Only present on
	// `two_way` source fields.
	ReverseFields []FieldPointer `json:"reverse_fields"`
	// A reference to the source field that manages this reverse field. Only present on
	// reverse (contingent) fields.
	SourceField FieldPointer `json:"source_field"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AllowedCollections respjson.Field
		Cardinality        respjson.Field
		CreatedAt          respjson.Field
		DefaultValues      respjson.Field
		Kind               respjson.Field
		Name               respjson.Field
		Readonly           respjson.Field
		Ref                respjson.Field
		RelationType       respjson.Field
		Required           respjson.Field
		Type               respjson.Field
		Unique             respjson.Field
		UpdatedAt          respjson.Field
		Description        respjson.Field
		ReverseFieldName   respjson.Field
		ReverseFields      respjson.Field
		SourceField        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RelationField) RawJSON() string { return r.JSON.raw }
func (r *RelationField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type RelationFieldCardinality string

const (
	RelationFieldCardinalityOne  RelationFieldCardinality = "one"
	RelationFieldCardinalityMany RelationFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type RelationFieldKind string

const (
	RelationFieldKindSystem  RelationFieldKind = "system"
	RelationFieldKindInverse RelationFieldKind = "inverse"
	RelationFieldKindCustom  RelationFieldKind = "custom"
)

// The type of relationship. Can be `one_way` for simple references or `two_way`
// for bidirectional relationships.
type RelationFieldRelationType string

const (
	RelationFieldRelationTypeOneWay RelationFieldRelationType = "one_way"
	RelationFieldRelationTypeTwoWay RelationFieldRelationType = "two_way"
)

func RelationFieldDefaultValueParamOfValueRelation(data ItemPointerParam) RelationFieldDefaultValueParamUnion {
	var valueRelation RelationValueParam
	valueRelation.Data = data
	return RelationFieldDefaultValueParamUnion{OfValueRelation: &valueRelation}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RelationFieldDefaultValueParamUnion struct {
	OfValueRelation *RelationValueParam `json:",omitzero,inline"`
	OfCurrentMember *CurrentMemberParam `json:",omitzero,inline"`
	paramUnion
}

func (u RelationFieldDefaultValueParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfValueRelation, u.OfCurrentMember)
}
func (u *RelationFieldDefaultValueParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RelationFieldDefaultValueParamUnion) asAny() any {
	if !param.IsOmitted(u.OfValueRelation) {
		return u.OfValueRelation
	} else if !param.IsOmitted(u.OfCurrentMember) {
		return u.OfCurrentMember
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RelationFieldDefaultValueParamUnion) GetData() *ItemPointerParam {
	if vt := u.OfValueRelation; vt != nil {
		return &vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RelationFieldDefaultValueParamUnion) GetType() *string {
	if vt := u.OfValueRelation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCurrentMember; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[RelationFieldDefaultValueParamUnion](
		"type",
		apijson.Discriminator[RelationValueParam]("value/relation"),
		apijson.Discriminator[CurrentMemberParam]("current_member"),
	)
}

// Related item reference
type RelationValue struct {
	// A reference to another Moonbase item.
	Data ItemPointer            `json:"data" api:"required"`
	Type constant.ValueRelation `json:"type" default:"value/relation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RelationValue) RawJSON() string { return r.JSON.raw }
func (r *RelationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Related item reference
//
// The properties Data, Type are required.
type RelationValueParam struct {
	// A reference to another Moonbase item.
	Data ItemPointerParam `json:"data,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "value/relation".
	Type constant.ValueRelation `json:"type" default:"value/relation"`
	paramObj
}

func (r RelationValueParam) MarshalJSON() (data []byte, err error) {
	type shadow RelationValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RelationValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores a single line of text without line breaks.
type SingleLineTextField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality SingleLineTextFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind SingleLineTextFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Company Name").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `company_name`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/text/single_line` for this field.
	Type constant.FieldTextSingleLine `json:"type" default:"field/text/single_line"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SingleLineTextField) RawJSON() string { return r.JSON.raw }
func (r *SingleLineTextField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type SingleLineTextFieldCardinality string

const (
	SingleLineTextFieldCardinalityOne  SingleLineTextFieldCardinality = "one"
	SingleLineTextFieldCardinalityMany SingleLineTextFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type SingleLineTextFieldKind string

const (
	SingleLineTextFieldKindSystem  SingleLineTextFieldKind = "system"
	SingleLineTextFieldKindInverse SingleLineTextFieldKind = "inverse"
	SingleLineTextFieldKindCustom  SingleLineTextFieldKind = "custom"
)

// A single line of text
type SingleLineTextValue struct {
	// A single line of text, up to 1024 characters long. It should not contain line
	// breaks.
	Data string                       `json:"data" api:"required"`
	Type constant.ValueTextSingleLine `json:"type" default:"value/text/single_line"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SingleLineTextValue) RawJSON() string { return r.JSON.raw }
func (r *SingleLineTextValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SingleLineTextValue to a SingleLineTextValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SingleLineTextValueParam.Overrides()
func (r SingleLineTextValue) ToParam() SingleLineTextValueParam {
	return param.Override[SingleLineTextValueParam](json.RawMessage(r.RawJSON()))
}

// A single line of text
//
// The properties Data, Type are required.
type SingleLineTextValueParam struct {
	// A single line of text, up to 1024 characters long. It should not contain line
	// breaks.
	Data string `json:"data" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "value/text/single_line".
	Type constant.ValueTextSingleLine `json:"type" default:"value/text/single_line"`
	paramObj
}

func (r SingleLineTextValueParam) MarshalJSON() (data []byte, err error) {
	type shadow SingleLineTextValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SingleLineTextValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores LinkedIn profile information.
type SocialLinkedInField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality SocialLinkedInFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind SocialLinkedInFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "LinkedIn Profile").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `linkedin_profile`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/uri/social_linked_in` for this field.
	Type constant.FieldUriSocialLinkedIn `json:"type" default:"field/uri/social_linked_in"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialLinkedInField) RawJSON() string { return r.JSON.raw }
func (r *SocialLinkedInField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type SocialLinkedInFieldCardinality string

const (
	SocialLinkedInFieldCardinalityOne  SocialLinkedInFieldCardinality = "one"
	SocialLinkedInFieldCardinalityMany SocialLinkedInFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type SocialLinkedInFieldKind string

const (
	SocialLinkedInFieldKindSystem  SocialLinkedInFieldKind = "system"
	SocialLinkedInFieldKindInverse SocialLinkedInFieldKind = "inverse"
	SocialLinkedInFieldKindCustom  SocialLinkedInFieldKind = "custom"
)

// The social media profile for the LinkedIn platform
type SocialLinkedInValue struct {
	// The social media profile for the LinkedIn platform
	Data SocialLinkedInValueData         `json:"data" api:"required"`
	Type constant.ValueUriSocialLinkedIn `json:"type" default:"value/uri/social_linked_in"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialLinkedInValue) RawJSON() string { return r.JSON.raw }
func (r *SocialLinkedInValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The social media profile for the LinkedIn platform
type SocialLinkedInValueData struct {
	// The full URL to the LinkedIn profile.
	URL string `json:"url" api:"required" format:"uri"`
	// The LinkedIn username, including the prefix 'company/' for company pages or
	// 'in/' for personal profiles.
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialLinkedInValueData) RawJSON() string { return r.JSON.raw }
func (r *SocialLinkedInValueData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The social media profile for the LinkedIn platform
//
// The properties Data, Type are required.
type SocialLinkedInValueParam struct {
	// The social media profile for the LinkedIn platform
	Data SocialProfileLinkedInParam `json:"data,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "value/uri/social_linked_in".
	Type constant.ValueUriSocialLinkedIn `json:"type" default:"value/uri/social_linked_in"`
	paramObj
}

func (r SocialLinkedInValueParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialLinkedInValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialLinkedInValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Social media profile information including both the full URL and extracted
// username.
type SocialProfileLinkedInParam struct {
	// The full URL to the LinkedIn profile.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// The LinkedIn username, including the prefix 'company/' for company pages or
	// 'in/' for personal profiles.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r SocialProfileLinkedInParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialProfileLinkedInParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialProfileLinkedInParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Social media profile information including both the full URL and extracted
// username.
type SocialProfileXParam struct {
	// The full URL to the X profile, starting with 'https://x.com/'
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// The X username, up to 15 characters long, containing only lowercase letters
	// (a-z), uppercase letters (A-Z), numbers (0-9), and underscores (\_). Does not
	// include the '@' symbol prefix.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r SocialProfileXParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialProfileXParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialProfileXParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores X (formerly Twitter) profile information.
type SocialXField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality SocialXFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind SocialXFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "X Profile").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `x_profile`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/uri/social_x` for this field.
	Type constant.FieldUriSocialX `json:"type" default:"field/uri/social_x"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialXField) RawJSON() string { return r.JSON.raw }
func (r *SocialXField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type SocialXFieldCardinality string

const (
	SocialXFieldCardinalityOne  SocialXFieldCardinality = "one"
	SocialXFieldCardinalityMany SocialXFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type SocialXFieldKind string

const (
	SocialXFieldKindSystem  SocialXFieldKind = "system"
	SocialXFieldKindInverse SocialXFieldKind = "inverse"
	SocialXFieldKindCustom  SocialXFieldKind = "custom"
)

// The social media profile for the X (formerly Twitter) platform
type SocialXValue struct {
	// Social media profile information including both the full URL and extracted
	// username.
	Data SocialXValueData         `json:"data" api:"required"`
	Type constant.ValueUriSocialX `json:"type" default:"value/uri/social_x"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialXValue) RawJSON() string { return r.JSON.raw }
func (r *SocialXValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Social media profile information including both the full URL and extracted
// username.
type SocialXValueData struct {
	// The full URL to the X profile, starting with 'https://x.com/'
	URL string `json:"url" api:"required" format:"uri"`
	// The X username, up to 15 characters long, containing only lowercase letters
	// (a-z), uppercase letters (A-Z), numbers (0-9), and underscores (\_). Does not
	// include the '@' symbol prefix.
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialXValueData) RawJSON() string { return r.JSON.raw }
func (r *SocialXValueData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The social media profile for the X (formerly Twitter) platform
//
// The properties Data, Type are required.
type SocialXValueParam struct {
	// Social media profile information including both the full URL and extracted
	// username.
	Data SocialProfileXParam `json:"data,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "value/uri/social_x".
	Type constant.ValueUriSocialX `json:"type" default:"value/uri/social_x"`
	paramObj
}

func (r SocialXValueParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialXValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialXValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that tracks an item's position in a funnel or pipeline workflow.
type StageField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality StageFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// The `Funnel` object that defines the available stages for this field.
	Funnel Funnel `json:"funnel" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind StageFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Sales Stage").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `sales_stage`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/stage` for this field.
	Type constant.FieldStage `json:"type" default:"field/stage"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Funnel        respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StageField) RawJSON() string { return r.JSON.raw }
func (r *StageField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type StageFieldCardinality string

const (
	StageFieldCardinalityOne  StageFieldCardinality = "one"
	StageFieldCardinalityMany StageFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type StageFieldKind string

const (
	StageFieldKindSystem  StageFieldKind = "system"
	StageFieldKindInverse StageFieldKind = "inverse"
	StageFieldKindCustom  StageFieldKind = "custom"
)

// Parameters for creating a stage field.
//
// The properties Funnel, Name, Type are required.
type StageFieldCreateParams struct {
	// The funnel that defines the available stages for this field.
	Funnel FunnelPointerParam `json:"funnel,omitzero" api:"required"`
	// The human-readable name for the field.
	Name string `json:"name" api:"required"`
	// An optional description of the field's purpose.
	Description param.Opt[string] `json:"description,omitzero"`
	// If `true`, items must have a value for this field. Defaults to `false`.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items. Defaults to `false`.
	Unique param.Opt[bool] `json:"unique,omitzero"`
	// Whether the field holds a single value (`one`) or multiple values (`many`).
	// Defaults to `one`.
	//
	// Any of "one", "many".
	Cardinality   StageFieldCreateParamsCardinality `json:"cardinality,omitzero"`
	DefaultValues []FunnelStepValueParam            `json:"default_values,omitzero"`
	// The field type. Must be `field/stage`.
	//
	// This field can be elided, and will marshal its zero value as "field/stage".
	Type constant.FieldStage `json:"type" default:"field/stage"`
	paramObj
}

func (r StageFieldCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow StageFieldCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StageFieldCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the field holds a single value (`one`) or multiple values (`many`).
// Defaults to `one`.
type StageFieldCreateParamsCardinality string

const (
	StageFieldCreateParamsCardinalityOne  StageFieldCreateParamsCardinality = "one"
	StageFieldCreateParamsCardinalityMany StageFieldCreateParamsCardinality = "many"
)

// Parameters for updating a stage field.
//
// The property Type is required.
type StageFieldUpdateParams struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]        `json:"unique,omitzero"`
	DefaultValues []FunnelStepValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality StageFieldUpdateParamsCardinality `json:"cardinality,omitzero"`
	// A new funnel to use for this field, or omit to keep the current funnel.
	Funnel FunnelPointerParam `json:"funnel,omitzero"`
	// The field type. Must be `field/stage`.
	//
	// This field can be elided, and will marshal its zero value as "field/stage".
	Type constant.FieldStage `json:"type" default:"field/stage"`
	paramObj
}

func (r StageFieldUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StageFieldUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StageFieldUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Updated cardinality: `one` or `many`.
type StageFieldUpdateParamsCardinality string

const (
	StageFieldUpdateParamsCardinalityOne  StageFieldUpdateParamsCardinality = "one"
	StageFieldUpdateParamsCardinalityMany StageFieldUpdateParamsCardinality = "many"
)

// Telephone number value
type TelephoneNumber struct {
	// A telephone number in strictly formatted E.164 format. Do not include spaces,
	// dashes, or parentheses etc.
	Data string                        `json:"data" api:"required"`
	Type constant.ValueTelephoneNumber `json:"type" default:"value/telephone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephoneNumber) RawJSON() string { return r.JSON.raw }
func (r *TelephoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TelephoneNumber to a TelephoneNumberParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TelephoneNumberParam.Overrides()
func (r TelephoneNumber) ToParam() TelephoneNumberParam {
	return param.Override[TelephoneNumberParam](json.RawMessage(r.RawJSON()))
}

// Telephone number value
//
// The properties Data, Type are required.
type TelephoneNumberParam struct {
	// A telephone number in strictly formatted E.164 format. Do not include spaces,
	// dashes, or parentheses etc.
	Data string `json:"data" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "value/telephone_number".
	Type constant.ValueTelephoneNumber `json:"type" default:"value/telephone_number"`
	paramObj
}

func (r TelephoneNumberParam) MarshalJSON() (data []byte, err error) {
	type shadow TelephoneNumberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelephoneNumberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores phone numbers in E.164 format.
type TelephoneNumberField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality TelephoneNumberFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind TelephoneNumberFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Phone").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `phone`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/telephone_number` for this field.
	Type constant.FieldTelephoneNumber `json:"type" default:"field/telephone_number"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephoneNumberField) RawJSON() string { return r.JSON.raw }
func (r *TelephoneNumberField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type TelephoneNumberFieldCardinality string

const (
	TelephoneNumberFieldCardinalityOne  TelephoneNumberFieldCardinality = "one"
	TelephoneNumberFieldCardinalityMany TelephoneNumberFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type TelephoneNumberFieldKind string

const (
	TelephoneNumberFieldKindSystem  TelephoneNumberFieldKind = "system"
	TelephoneNumberFieldKindInverse TelephoneNumberFieldKind = "inverse"
	TelephoneNumberFieldKindCustom  TelephoneNumberFieldKind = "custom"
)

// A field that stores and validates web URLs.
type URLField struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality URLFieldCardinality `json:"cardinality" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt     time.Time                `json:"created_at" api:"required" format:"date-time"`
	DefaultValues []FieldDefaultValueUnion `json:"default_values" api:"required"`
	// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
	// of a two-way relation, and `custom` fields are user-created.
	//
	// Any of "system", "inverse", "custom".
	Kind URLFieldKind `json:"kind" api:"required"`
	// The human-readable name of the field (e.g., "Website").
	Name string `json:"name" api:"required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly" api:"required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `website`).
	Ref string `json:"ref" api:"required"`
	// If `true`, this field must have a value.
	Required bool `json:"required" api:"required"`
	// The data type of the field. Always `field/uri/url` for this field.
	Type constant.FieldUriURL `json:"type" default:"field/uri/url"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique" api:"required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Cardinality   respjson.Field
		CreatedAt     respjson.Field
		DefaultValues respjson.Field
		Kind          respjson.Field
		Name          respjson.Field
		Readonly      respjson.Field
		Ref           respjson.Field
		Required      respjson.Field
		Type          respjson.Field
		Unique        respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLField) RawJSON() string { return r.JSON.raw }
func (r *URLField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type URLFieldCardinality string

const (
	URLFieldCardinalityOne  URLFieldCardinality = "one"
	URLFieldCardinalityMany URLFieldCardinality = "many"
)

// `system` fields are managed by Moonbase, `inverse` fields are the reverse side
// of a two-way relation, and `custom` fields are user-created.
type URLFieldKind string

const (
	URLFieldKindSystem  URLFieldKind = "system"
	URLFieldKindInverse URLFieldKind = "inverse"
	URLFieldKindCustom  URLFieldKind = "custom"
)

// URL or web address
type URLValue struct {
	// A valid URL, conforming to RFC 3986, up to 8,192 characters long. It should
	// include the protocol, for example 'https://' or 'mailto:support@moonbase.ai'
	// etc.
	Data string               `json:"data" api:"required" format:"uri"`
	Type constant.ValueUriURL `json:"type" default:"value/uri/url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLValue) RawJSON() string { return r.JSON.raw }
func (r *URLValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this URLValue to a URLValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// URLValueParam.Overrides()
func (r URLValue) ToParam() URLValueParam {
	return param.Override[URLValueParam](json.RawMessage(r.RawJSON()))
}

// URL or web address
//
// The properties Data, Type are required.
type URLValueParam struct {
	// A valid URL, conforming to RFC 3986, up to 8,192 characters long. It should
	// include the protocol, for example 'https://' or 'mailto:support@moonbase.ai'
	// etc.
	Data string `json:"data" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "value/uri/url".
	Type constant.ValueUriURL `json:"type" default:"value/uri/url"`
	paramObj
}

func (r URLValueParam) MarshalJSON() (data []byte, err error) {
	type shadow URLValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ValueUnion contains all possible properties and values from
// [SingleLineTextValue], [MultiLineTextValue], [IdentifierValue], [IntegerValue],
// [FloatValue], [MonetaryValue], [PercentageValue], [BooleanValue], [EmailValue],
// [URLValue], [DomainValue], [SocialXValue], [SocialLinkedInValue],
// [TelephoneNumber], [GeoValue], [DateValue], [DatetimeValue], [ChoiceValue],
// [FunnelStepValue], [RelationValue].
//
// Use the [ValueUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ValueUnion struct {
	// This field is a union of [string], [string], [string], [int64], [float64],
	// [MonetaryValueData], [float64], [bool], [string], [string], [string],
	// [SocialXValueData], [SocialLinkedInValueData], [string], [string], [time.Time],
	// [time.Time], [ChoiceFieldOption], [FunnelStep], [ItemPointer]
	Data ValueUnionData `json:"data"`
	// Any of "value/text/single_line", "value/text/multi_line", "value/identifier",
	// "value/number/unitless_integer", "value/number/unitless_float",
	// "value/number/monetary", "value/number/percentage", "value/boolean",
	// "value/email", "value/uri/url", "value/uri/domain", "value/uri/social_x",
	// "value/uri/social_linked_in", "value/telephone_number", "value/geo",
	// "value/date", "value/datetime", "value/choice", "value/funnel_step",
	// "value/relation".
	Type string `json:"type"`
	JSON struct {
		Data respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

// anyValue is implemented by each variant of [ValueUnion] to add type safety for
// the return type of [ValueUnion.AsAny]
type anyValue interface {
	implValueUnion()
}

func (SingleLineTextValue) implValueUnion() {}
func (MultiLineTextValue) implValueUnion()  {}
func (IdentifierValue) implValueUnion()     {}
func (IntegerValue) implValueUnion()        {}
func (FloatValue) implValueUnion()          {}
func (MonetaryValue) implValueUnion()       {}
func (PercentageValue) implValueUnion()     {}
func (BooleanValue) implValueUnion()        {}
func (EmailValue) implValueUnion()          {}
func (URLValue) implValueUnion()            {}
func (DomainValue) implValueUnion()         {}
func (SocialXValue) implValueUnion()        {}
func (SocialLinkedInValue) implValueUnion() {}
func (TelephoneNumber) implValueUnion()     {}
func (GeoValue) implValueUnion()            {}
func (DateValue) implValueUnion()           {}
func (DatetimeValue) implValueUnion()       {}
func (ChoiceValue) implValueUnion()         {}
func (FunnelStepValue) implValueUnion()     {}
func (RelationValue) implValueUnion()       {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ValueUnion.AsAny().(type) {
//	case moonbase.SingleLineTextValue:
//	case moonbase.MultiLineTextValue:
//	case moonbase.IdentifierValue:
//	case moonbase.IntegerValue:
//	case moonbase.FloatValue:
//	case moonbase.MonetaryValue:
//	case moonbase.PercentageValue:
//	case moonbase.BooleanValue:
//	case moonbase.EmailValue:
//	case moonbase.URLValue:
//	case moonbase.DomainValue:
//	case moonbase.SocialXValue:
//	case moonbase.SocialLinkedInValue:
//	case moonbase.TelephoneNumber:
//	case moonbase.GeoValue:
//	case moonbase.DateValue:
//	case moonbase.DatetimeValue:
//	case moonbase.ChoiceValue:
//	case moonbase.FunnelStepValue:
//	case moonbase.RelationValue:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ValueUnion) AsAny() anyValue {
	switch u.Type {
	case "value/text/single_line":
		return u.AsValueTextSingleLine()
	case "value/text/multi_line":
		return u.AsValueTextMultiLine()
	case "value/identifier":
		return u.AsValueIdentifier()
	case "value/number/unitless_integer":
		return u.AsValueNumberUnitlessInteger()
	case "value/number/unitless_float":
		return u.AsValueNumberUnitlessFloat()
	case "value/number/monetary":
		return u.AsValueNumberMonetary()
	case "value/number/percentage":
		return u.AsValueNumberPercentage()
	case "value/boolean":
		return u.AsValueBoolean()
	case "value/email":
		return u.AsValueEmail()
	case "value/uri/url":
		return u.AsValueUriURL()
	case "value/uri/domain":
		return u.AsValueUriDomain()
	case "value/uri/social_x":
		return u.AsValueUriSocialX()
	case "value/uri/social_linked_in":
		return u.AsValueUriSocialLinkedIn()
	case "value/telephone_number":
		return u.AsValueTelephoneNumber()
	case "value/geo":
		return u.AsValueGeo()
	case "value/date":
		return u.AsValueDate()
	case "value/datetime":
		return u.AsValueDatetime()
	case "value/choice":
		return u.AsValueChoice()
	case "value/funnel_step":
		return u.AsValueFunnelStep()
	case "value/relation":
		return u.AsValueRelation()
	}
	return nil
}

func (u ValueUnion) AsValueTextSingleLine() (v SingleLineTextValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueTextMultiLine() (v MultiLineTextValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueIdentifier() (v IdentifierValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueNumberUnitlessInteger() (v IntegerValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueNumberUnitlessFloat() (v FloatValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueNumberMonetary() (v MonetaryValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueNumberPercentage() (v PercentageValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueBoolean() (v BooleanValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueEmail() (v EmailValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueUriURL() (v URLValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueUriDomain() (v DomainValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueUriSocialX() (v SocialXValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueUriSocialLinkedIn() (v SocialLinkedInValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueTelephoneNumber() (v TelephoneNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueGeo() (v GeoValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueDate() (v DateValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueDatetime() (v DatetimeValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueChoice() (v ChoiceValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueFunnelStep() (v FunnelStepValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueRelation() (v RelationValue) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ValueUnionData is an implicit subunion of [ValueUnion]. ValueUnionData provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the [ValueUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfFloat OfBool OfTime]
type ValueUnionData struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [time.Time] instead of an object.
	OfTime time.Time `json:",inline"`
	// This field is from variant [MonetaryValueData].
	Currency string `json:"currency"`
	// This field is from variant [MonetaryValueData].
	InMinorUnits int64  `json:"in_minor_units"`
	URL          string `json:"url"`
	Username     string `json:"username"`
	ID           string `json:"id"`
	Color        string `json:"color"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	// This field is from variant [FunnelStep].
	StepType FunnelStepStepType `json:"step_type"`
	// This field is from variant [ItemPointer].
	Collection CollectionPointer `json:"collection"`
	JSON       struct {
		OfString     respjson.Field
		OfInt        respjson.Field
		OfFloat      respjson.Field
		OfBool       respjson.Field
		OfTime       respjson.Field
		Currency     respjson.Field
		InMinorUnits respjson.Field
		URL          respjson.Field
		Username     respjson.Field
		ID           respjson.Field
		Color        respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		StepType     respjson.Field
		Collection   respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ValueUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ValueParamOfValueTextSingleLine(data string) ValueParamUnion {
	var valueTextSingleLine SingleLineTextValueParam
	valueTextSingleLine.Data = data
	return ValueParamUnion{OfValueTextSingleLine: &valueTextSingleLine}
}

func ValueParamOfValueTextMultiLine(data string) ValueParamUnion {
	var valueTextMultiLine MultiLineTextValueParam
	valueTextMultiLine.Data = data
	return ValueParamUnion{OfValueTextMultiLine: &valueTextMultiLine}
}

func ValueParamOfValueIdentifier(data string) ValueParamUnion {
	var valueIdentifier IdentifierValueParam
	valueIdentifier.Data = data
	return ValueParamUnion{OfValueIdentifier: &valueIdentifier}
}

func ValueParamOfValueNumberUnitlessInteger(data int64) ValueParamUnion {
	var valueNumberUnitlessInteger IntegerValueParam
	valueNumberUnitlessInteger.Data = data
	return ValueParamUnion{OfValueNumberUnitlessInteger: &valueNumberUnitlessInteger}
}

func ValueParamOfValueNumberUnitlessFloat(data float64) ValueParamUnion {
	var valueNumberUnitlessFloat FloatValueParam
	valueNumberUnitlessFloat.Data = data
	return ValueParamUnion{OfValueNumberUnitlessFloat: &valueNumberUnitlessFloat}
}

func ValueParamOfValueNumberMonetary(data MonetaryValueDataParam) ValueParamUnion {
	var valueNumberMonetary MonetaryValueParam
	valueNumberMonetary.Data = data
	return ValueParamUnion{OfValueNumberMonetary: &valueNumberMonetary}
}

func ValueParamOfValueNumberPercentage(data float64) ValueParamUnion {
	var valueNumberPercentage PercentageValueParam
	valueNumberPercentage.Data = data
	return ValueParamUnion{OfValueNumberPercentage: &valueNumberPercentage}
}

func ValueParamOfValueBoolean(data bool) ValueParamUnion {
	var valueBoolean BooleanValueParam
	valueBoolean.Data = data
	return ValueParamUnion{OfValueBoolean: &valueBoolean}
}

func ValueParamOfValueEmail(data string) ValueParamUnion {
	var valueEmail EmailValueParam
	valueEmail.Data = data
	return ValueParamUnion{OfValueEmail: &valueEmail}
}

func ValueParamOfValueUriURL(data string) ValueParamUnion {
	var valueUriURL URLValueParam
	valueUriURL.Data = data
	return ValueParamUnion{OfValueUriURL: &valueUriURL}
}

func ValueParamOfValueUriDomain(data string) ValueParamUnion {
	var valueUriDomain DomainValueParam
	valueUriDomain.Data = data
	return ValueParamUnion{OfValueUriDomain: &valueUriDomain}
}

func ValueParamOfValueUriSocialX(data SocialProfileXParam) ValueParamUnion {
	var valueUriSocialX SocialXValueParam
	valueUriSocialX.Data = data
	return ValueParamUnion{OfValueUriSocialX: &valueUriSocialX}
}

func ValueParamOfValueUriSocialLinkedIn(data SocialProfileLinkedInParam) ValueParamUnion {
	var valueUriSocialLinkedIn SocialLinkedInValueParam
	valueUriSocialLinkedIn.Data = data
	return ValueParamUnion{OfValueUriSocialLinkedIn: &valueUriSocialLinkedIn}
}

func ValueParamOfValueTelephoneNumber(data string) ValueParamUnion {
	var valueTelephoneNumber TelephoneNumberParam
	valueTelephoneNumber.Data = data
	return ValueParamUnion{OfValueTelephoneNumber: &valueTelephoneNumber}
}

func ValueParamOfValueGeo(data string) ValueParamUnion {
	var valueGeo GeoValueParam
	valueGeo.Data = data
	return ValueParamUnion{OfValueGeo: &valueGeo}
}

func ValueParamOfValueDate(data time.Time) ValueParamUnion {
	var valueDate DateValueParam
	valueDate.Data = data
	return ValueParamUnion{OfValueDate: &valueDate}
}

func ValueParamOfValueDatetime(data time.Time) ValueParamUnion {
	var valueDatetime DatetimeValueParam
	valueDatetime.Data = data
	return ValueParamUnion{OfValueDatetime: &valueDatetime}
}

func ValueParamOfValueChoice(data ChoiceFieldOptionPointerParam) ValueParamUnion {
	var valueChoice ChoiceValueParam
	valueChoice.Data = data
	return ValueParamUnion{OfValueChoice: &valueChoice}
}

func ValueParamOfValueFunnelStep(data FunnelStepPointerParam) ValueParamUnion {
	var valueFunnelStep FunnelStepValueParam
	valueFunnelStep.Data = data
	return ValueParamUnion{OfValueFunnelStep: &valueFunnelStep}
}

func ValueParamOfValueRelation(data ItemPointerParam) ValueParamUnion {
	var valueRelation RelationValueParam
	valueRelation.Data = data
	return ValueParamUnion{OfValueRelation: &valueRelation}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ValueParamUnion struct {
	OfValueTextSingleLine        *SingleLineTextValueParam `json:",omitzero,inline"`
	OfValueTextMultiLine         *MultiLineTextValueParam  `json:",omitzero,inline"`
	OfValueIdentifier            *IdentifierValueParam     `json:",omitzero,inline"`
	OfValueNumberUnitlessInteger *IntegerValueParam        `json:",omitzero,inline"`
	OfValueNumberUnitlessFloat   *FloatValueParam          `json:",omitzero,inline"`
	OfValueNumberMonetary        *MonetaryValueParam       `json:",omitzero,inline"`
	OfValueNumberPercentage      *PercentageValueParam     `json:",omitzero,inline"`
	OfValueBoolean               *BooleanValueParam        `json:",omitzero,inline"`
	OfValueEmail                 *EmailValueParam          `json:",omitzero,inline"`
	OfValueUriURL                *URLValueParam            `json:",omitzero,inline"`
	OfValueUriDomain             *DomainValueParam         `json:",omitzero,inline"`
	OfValueUriSocialX            *SocialXValueParam        `json:",omitzero,inline"`
	OfValueUriSocialLinkedIn     *SocialLinkedInValueParam `json:",omitzero,inline"`
	OfValueTelephoneNumber       *TelephoneNumberParam     `json:",omitzero,inline"`
	OfValueGeo                   *GeoValueParam            `json:",omitzero,inline"`
	OfValueDate                  *DateValueParam           `json:",omitzero,inline"`
	OfValueDatetime              *DatetimeValueParam       `json:",omitzero,inline"`
	OfValueChoice                *ChoiceValueParam         `json:",omitzero,inline"`
	OfValueFunnelStep            *FunnelStepValueParam     `json:",omitzero,inline"`
	OfValueRelation              *RelationValueParam       `json:",omitzero,inline"`
	paramUnion
}

func (u ValueParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfValueTextSingleLine,
		u.OfValueTextMultiLine,
		u.OfValueIdentifier,
		u.OfValueNumberUnitlessInteger,
		u.OfValueNumberUnitlessFloat,
		u.OfValueNumberMonetary,
		u.OfValueNumberPercentage,
		u.OfValueBoolean,
		u.OfValueEmail,
		u.OfValueUriURL,
		u.OfValueUriDomain,
		u.OfValueUriSocialX,
		u.OfValueUriSocialLinkedIn,
		u.OfValueTelephoneNumber,
		u.OfValueGeo,
		u.OfValueDate,
		u.OfValueDatetime,
		u.OfValueChoice,
		u.OfValueFunnelStep,
		u.OfValueRelation)
}
func (u *ValueParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ValueParamUnion) asAny() any {
	if !param.IsOmitted(u.OfValueTextSingleLine) {
		return u.OfValueTextSingleLine
	} else if !param.IsOmitted(u.OfValueTextMultiLine) {
		return u.OfValueTextMultiLine
	} else if !param.IsOmitted(u.OfValueIdentifier) {
		return u.OfValueIdentifier
	} else if !param.IsOmitted(u.OfValueNumberUnitlessInteger) {
		return u.OfValueNumberUnitlessInteger
	} else if !param.IsOmitted(u.OfValueNumberUnitlessFloat) {
		return u.OfValueNumberUnitlessFloat
	} else if !param.IsOmitted(u.OfValueNumberMonetary) {
		return u.OfValueNumberMonetary
	} else if !param.IsOmitted(u.OfValueNumberPercentage) {
		return u.OfValueNumberPercentage
	} else if !param.IsOmitted(u.OfValueBoolean) {
		return u.OfValueBoolean
	} else if !param.IsOmitted(u.OfValueEmail) {
		return u.OfValueEmail
	} else if !param.IsOmitted(u.OfValueUriURL) {
		return u.OfValueUriURL
	} else if !param.IsOmitted(u.OfValueUriDomain) {
		return u.OfValueUriDomain
	} else if !param.IsOmitted(u.OfValueUriSocialX) {
		return u.OfValueUriSocialX
	} else if !param.IsOmitted(u.OfValueUriSocialLinkedIn) {
		return u.OfValueUriSocialLinkedIn
	} else if !param.IsOmitted(u.OfValueTelephoneNumber) {
		return u.OfValueTelephoneNumber
	} else if !param.IsOmitted(u.OfValueGeo) {
		return u.OfValueGeo
	} else if !param.IsOmitted(u.OfValueDate) {
		return u.OfValueDate
	} else if !param.IsOmitted(u.OfValueDatetime) {
		return u.OfValueDatetime
	} else if !param.IsOmitted(u.OfValueChoice) {
		return u.OfValueChoice
	} else if !param.IsOmitted(u.OfValueFunnelStep) {
		return u.OfValueFunnelStep
	} else if !param.IsOmitted(u.OfValueRelation) {
		return u.OfValueRelation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueParamUnion) GetType() *string {
	if vt := u.OfValueTextSingleLine; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueTextMultiLine; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueIdentifier; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueNumberUnitlessInteger; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueNumberUnitlessFloat; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueNumberMonetary; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueNumberPercentage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueBoolean; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueEmail; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueUriURL; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueUriDomain; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueUriSocialX; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueUriSocialLinkedIn; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueTelephoneNumber; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueGeo; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueDate; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueDatetime; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueChoice; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueFunnelStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueRelation; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ValueParamUnion) GetData() (res valueParamUnionData) {
	if vt := u.OfValueTextSingleLine; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueTextMultiLine; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueIdentifier; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueNumberUnitlessInteger; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueNumberUnitlessFloat; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueNumberMonetary; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueNumberPercentage; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueBoolean; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueEmail; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueUriURL; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueUriDomain; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueUriSocialX; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueUriSocialLinkedIn; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueTelephoneNumber; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueGeo; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueDate; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueDatetime; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueChoice; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueFunnelStep; vt != nil {
		res.any = &vt.Data
	} else if vt := u.OfValueRelation; vt != nil {
		res.any = &vt.Data
	}
	return
}

// Can have the runtime types [*string], [*int64], [*float64],
// [*MonetaryValueDataParam], [*bool], [*SocialProfileXParam],
// [*SocialProfileLinkedInParam], [*time.Time], [*ChoiceFieldOptionPointerParam],
// [*FunnelStepPointerParam], [*ItemPointerParam]
type valueParamUnionData struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *int64:
//	case *float64:
//	case *moonbase.MonetaryValueDataParam:
//	case *bool:
//	case *moonbase.SocialProfileXParam:
//	case *moonbase.SocialProfileLinkedInParam:
//	case *time.Time:
//	case *moonbase.ChoiceFieldOptionPointerParam:
//	case *moonbase.FunnelStepPointerParam:
//	case *moonbase.ItemPointerParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u valueParamUnionData) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetCurrency() *string {
	switch vt := u.any.(type) {
	case *MonetaryValueDataParam:
		return &vt.Currency
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetInMinorUnits() *int64 {
	switch vt := u.any.(type) {
	case *MonetaryValueDataParam:
		return &vt.InMinorUnits
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetURL() *string {
	switch vt := u.any.(type) {
	case *SocialProfileXParam:
		return paramutil.AddrIfPresent(vt.URL)
	case *SocialProfileLinkedInParam:
		return paramutil.AddrIfPresent(vt.URL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetUsername() *string {
	switch vt := u.any.(type) {
	case *SocialProfileXParam:
		return paramutil.AddrIfPresent(vt.Username)
	case *SocialProfileLinkedInParam:
		return paramutil.AddrIfPresent(vt.Username)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetID() *string {
	switch vt := u.any.(type) {
	case *ChoiceFieldOptionPointerParam:
		return (*string)(&vt.ID)
	case *FunnelStepPointerParam:
		return (*string)(&vt.ID)
	case *ItemPointerParam:
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetType() *string {
	switch vt := u.any.(type) {
	case *ChoiceFieldOptionPointerParam:
		return (*string)(&vt.Type)
	case *FunnelStepPointerParam:
		return (*string)(&vt.Type)
	case *ItemPointerParam:
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ValueParamUnion](
		"type",
		apijson.Discriminator[SingleLineTextValueParam]("value/text/single_line"),
		apijson.Discriminator[MultiLineTextValueParam]("value/text/multi_line"),
		apijson.Discriminator[IdentifierValueParam]("value/identifier"),
		apijson.Discriminator[IntegerValueParam]("value/number/unitless_integer"),
		apijson.Discriminator[FloatValueParam]("value/number/unitless_float"),
		apijson.Discriminator[MonetaryValueParam]("value/number/monetary"),
		apijson.Discriminator[PercentageValueParam]("value/number/percentage"),
		apijson.Discriminator[BooleanValueParam]("value/boolean"),
		apijson.Discriminator[EmailValueParam]("value/email"),
		apijson.Discriminator[URLValueParam]("value/uri/url"),
		apijson.Discriminator[DomainValueParam]("value/uri/domain"),
		apijson.Discriminator[SocialXValueParam]("value/uri/social_x"),
		apijson.Discriminator[SocialLinkedInValueParam]("value/uri/social_linked_in"),
		apijson.Discriminator[TelephoneNumberParam]("value/telephone_number"),
		apijson.Discriminator[GeoValueParam]("value/geo"),
		apijson.Discriminator[DateValueParam]("value/date"),
		apijson.Discriminator[DatetimeValueParam]("value/datetime"),
		apijson.Discriminator[ChoiceValueParam]("value/choice"),
		apijson.Discriminator[FunnelStepValueParam]("value/funnel_step"),
		apijson.Discriminator[RelationValueParam]("value/relation"),
	)
}

// Information about the most essential attributes of a Collection (does not
// include the collection's field definitions).
type CollectionListResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "system", "form", "custom".
	Kind        CollectionListResponseKind `json:"kind" api:"required"`
	Name        string                     `json:"name" api:"required"`
	Ref         string                     `json:"ref" api:"required"`
	Type        constant.Collection        `json:"type" default:"collection"`
	UpdatedAt   time.Time                  `json:"updated_at" api:"required" format:"date-time"`
	Description string                     `json:"description"`
	IconName    string                     `json:"icon_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		Ref         respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		IconName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionListResponseKind string

const (
	CollectionListResponseKindSystem CollectionListResponseKind = "system"
	CollectionListResponseKindForm   CollectionListResponseKind = "form"
	CollectionListResponseKindCustom CollectionListResponseKind = "custom"
)

type CollectionNewParams struct {
	// The user-facing name of the collection (e.g., "Leads"). A `ref` is automatically
	// derived from the name.
	Name string `json:"name" api:"required"`
	// An optional, longer-form description of the collection's purpose.
	Description param.Opt[string] `json:"description,omitzero"`
	// An optional icon for the collection, as a Phosphor icon name in kebab-case (e.g.
	// `users`, `chart-bar`).
	IconName param.Opt[string] `json:"icon_name,omitzero"`
	paramObj
}

func (r CollectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionUpdateParams struct {
	// The collection's icon, as a Phosphor icon name in kebab-case (e.g. `users`,
	// `chart-bar`), or `null` to clear it.
	IconName param.Opt[string] `json:"icon_name,omitzero"`
	// An optional, longer-form description of the collection's purpose.
	Description param.Opt[string] `json:"description,omitzero"`
	// The user-facing name of the collection.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r CollectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionListParams struct {
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

// URLQuery serializes [CollectionListParams]'s query parameters as `url.Values`.
func (r CollectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
