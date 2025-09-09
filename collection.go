// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/apiquery"
	"github.com/moonbaseai/moonbase-sdk-go/internal/paramutil"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/pagination"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// CollectionService contains methods and other services that help with interacting
// with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectionService] method instead.
type CollectionService struct {
	Options []option.RequestOption
	Fields  CollectionFieldService
	Items   CollectionItemService
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

// Retrieves the details of an existing collection.
func (r *CollectionService) Get(ctx context.Context, id string, query CollectionGetParams, opts ...option.RequestOption) (res *Collection, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("collections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of your collections.
func (r *CollectionService) List(ctx context.Context, query CollectionListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Collection], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *CollectionService) ListAutoPaging(ctx context.Context, query CollectionListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Collection] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// A field that stores true or false values.
type BooleanField struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality BooleanFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Is Active").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `is_active`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/boolean` for this field.
	Type constant.FieldBoolean `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// True or false value
type BooleanValue struct {
	Data bool                  `json:"data,required"`
	Type constant.ValueBoolean `json:"type,required"`
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
	Data bool `json:"data,required"`
	// This field can be elided, and will marshal its zero value as "value/boolean".
	Type constant.ValueBoolean `json:"type,required"`
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
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality ChoiceFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Priority").
	Name string `json:"name,required"`
	// A list of `FieldOption` objects representing the available choices for this
	// field.
	Options []ChoiceFieldOption `json:"options,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `priority`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/choice` for this field.
	Type constant.FieldChoice `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Options     respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Represents a single selectable option within a choice field.
type ChoiceFieldOption struct {
	// Unique identifier for the option.
	ID string `json:"id,required"`
	// The human-readable text displayed for this option.
	Label string `json:"label,required"`
	// String representing the object’s type. Always `choice_field_option` for this
	// object.
	Type constant.ChoiceFieldOption `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
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

// ToParam converts this ChoiceFieldOption to a ChoiceFieldOptionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChoiceFieldOptionParam.Overrides()
func (r ChoiceFieldOption) ToParam() ChoiceFieldOptionParam {
	return param.Override[ChoiceFieldOptionParam](json.RawMessage(r.RawJSON()))
}

// Represents a single selectable option within a choice field.
//
// The properties ID, Label, Type are required.
type ChoiceFieldOptionParam struct {
	// Unique identifier for the option.
	ID string `json:"id,required"`
	// The human-readable text displayed for this option.
	Label string `json:"label,required"`
	// String representing the object’s type. Always `choice_field_option` for this
	// object.
	//
	// This field can be elided, and will marshal its zero value as
	// "choice_field_option".
	Type constant.ChoiceFieldOption `json:"type,required"`
	paramObj
}

func (r ChoiceFieldOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow ChoiceFieldOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChoiceFieldOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Selected choice option
type ChoiceValue struct {
	// An option that must match one of the predefined options for the field.
	Data ChoiceFieldOption    `json:"data,required"`
	Type constant.ValueChoice `json:"type,required"`
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
	Data ChoiceValueParamDataUnion `json:"data,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "value/choice".
	Type constant.ValueChoice `json:"type,required"`
	paramObj
}

func (r ChoiceValueParam) MarshalJSON() (data []byte, err error) {
	type shadow ChoiceValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChoiceValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChoiceValueParamDataUnion struct {
	OfChoiceFieldOption *ChoiceFieldOptionParam `json:",omitzero,inline"`
	OfPointer           *shared.PointerParam    `json:",omitzero,inline"`
	paramUnion
}

func (u ChoiceValueParamDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfChoiceFieldOption, u.OfPointer)
}
func (u *ChoiceValueParamDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChoiceValueParamDataUnion) asAny() any {
	if !param.IsOmitted(u.OfChoiceFieldOption) {
		return u.OfChoiceFieldOption
	} else if !param.IsOmitted(u.OfPointer) {
		return u.OfPointer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChoiceValueParamDataUnion) GetLabel() *string {
	if vt := u.OfChoiceFieldOption; vt != nil {
		return &vt.Label
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChoiceValueParamDataUnion) GetID() *string {
	if vt := u.OfChoiceFieldOption; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfPointer; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChoiceValueParamDataUnion) GetType() *string {
	if vt := u.OfChoiceFieldOption; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPointer; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A Collection is a container for structured data, similar to a database table or
// spreadsheet. It defines a schema using a set of `Fields` and holds the data as a
// list of `Items`.
type Collection struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A list of `Field` objects that define the schema for items in this collection.
	Fields []FieldUnion `json:"fields,required"`
	// The user-facing name of the collection (e.g., “Organizations”).
	Name string `json:"name,required"`
	// A unique, stable, machine-readable identifier for the collection. This reference
	// is used in API requests and does not change even if the `name` is updated.
	Ref string `json:"ref,required"`
	// String representing the object’s type. Always `collection` for this object.
	Type constant.Collection `json:"type,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the collection's purpose.
	Description string `json:"description"`
	// A list of saved `View` objects for presenting the collection's data.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Views []View `json:"views"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Fields      respjson.Field
		Name        respjson.Field
		Ref         respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
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

// A lightweight reference to a `Collection`, containing the minimal information
// needed to identify it.
type CollectionPointer struct {
	// Unique identifier of the collection.
	ID string `json:"id,required"`
	// The stable, machine-readable reference identifier of the collection.
	Ref string `json:"ref,required"`
	// String representing the object’s type. Always `collection` for this object.
	Type constant.Collection `json:"type,required"`
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

// ToParam converts this CollectionPointer to a CollectionPointerParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CollectionPointerParam.Overrides()
func (r CollectionPointer) ToParam() CollectionPointerParam {
	return param.Override[CollectionPointerParam](json.RawMessage(r.RawJSON()))
}

// A lightweight reference to a `Collection`, containing the minimal information
// needed to identify it.
//
// The properties ID, Ref, Type are required.
type CollectionPointerParam struct {
	// Unique identifier of the collection.
	ID string `json:"id,required"`
	// The stable, machine-readable reference identifier of the collection.
	Ref string `json:"ref,required"`
	// String representing the object’s type. Always `collection` for this object.
	//
	// This field can be elided, and will marshal its zero value as "collection".
	Type constant.Collection `json:"type,required"`
	paramObj
}

func (r CollectionPointerParam) MarshalJSON() (data []byte, err error) {
	type shadow CollectionPointerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionPointerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores dates without time information.
type DateField struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality DateFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Due Date").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `due_date`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/date` for this field.
	Type constant.FieldDate `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Date without time
type DateValue struct {
	Data time.Time          `json:"data,required" format:"date"`
	Type constant.ValueDate `json:"type,required"`
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
	Data time.Time `json:"data,required" format:"date"`
	// This field can be elided, and will marshal its zero value as "value/date".
	Type constant.ValueDate `json:"type,required"`
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
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality DatetimeFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Meeting Time").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `meeting_time`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/datetime` for this field.
	Type constant.FieldDatetime `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Date and time value
type DatetimeValue struct {
	Data time.Time              `json:"data,required" format:"date-time"`
	Type constant.ValueDatetime `json:"type,required"`
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
	Data time.Time `json:"data,required" format:"date-time"`
	// This field can be elided, and will marshal its zero value as "value/datetime".
	Type constant.ValueDatetime `json:"type,required"`
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
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality DomainFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Company Domain").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `company_domain`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/uri/domain` for this field.
	Type constant.FieldUriDomain `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Internet domain name
type DomainValue struct {
	// A valid internet domain name, without protocol (e.g., 'https://') or path.
	Data string                  `json:"data,required"`
	Type constant.ValueUriDomain `json:"type,required"`
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
	Data string `json:"data,required"`
	// This field can be elided, and will marshal its zero value as "value/uri/domain".
	Type constant.ValueUriDomain `json:"type,required"`
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
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality EmailFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Work Email").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `work_email`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/email` for this field.
	Type constant.FieldEmail `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Email address value
type EmailValue struct {
	// A valid email address.
	Data string              `json:"data,required" format:"email"`
	Type constant.ValueEmail `json:"type,required"`
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
	Data string `json:"data,required" format:"email"`
	// This field can be elided, and will marshal its zero value as "value/email".
	Type constant.ValueEmail `json:"type,required"`
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
// [SingleLineTextField], [MultiLineTextField], [IntegerField], [FloatField],
// [MonetaryField], [PercentageField], [BooleanField], [EmailField], [URLField],
// [DomainField], [SocialXField], [SocialLinkedInField], [TelephoneNumberField],
// [GeoField], [DateField], [DatetimeField], [ChoiceField], [StageField],
// [RelationField].
//
// Use the [FieldUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FieldUnion struct {
	ID          string    `json:"id"`
	Cardinality string    `json:"cardinality"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Readonly    bool      `json:"readonly"`
	Ref         string    `json:"ref"`
	Required    bool      `json:"required"`
	// Any of "field/text/single_line", "field/text/multi_line",
	// "field/number/unitless_integer", "field/number/unitless_float",
	// "field/number/monetary", "field/number/percentage", "field/boolean",
	// "field/email", "field/uri/url", "field/uri/domain", "field/uri/social_x",
	// "field/uri/social_linked_in", "field/telephone_number", "field/geo",
	// "field/date", "field/datetime", "field/choice", "field/stage", "field/relation".
	Type        string    `json:"type"`
	Unique      bool      `json:"unique"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	// This field is from variant [ChoiceField].
	Options []ChoiceFieldOption `json:"options"`
	// This field is from variant [StageField].
	Funnel Funnel `json:"funnel"`
	// This field is from variant [RelationField].
	RelationType RelationFieldRelationType `json:"relation_type"`
	JSON         struct {
		ID           respjson.Field
		Cardinality  respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Readonly     respjson.Field
		Ref          respjson.Field
		Required     respjson.Field
		Type         respjson.Field
		Unique       respjson.Field
		UpdatedAt    respjson.Field
		Description  respjson.Field
		Options      respjson.Field
		Funnel       respjson.Field
		RelationType respjson.Field
		raw          string
	} `json:"-"`
}

// anyField is implemented by each variant of [FieldUnion] to add type safety for
// the return type of [FieldUnion.AsAny]
type anyField interface {
	implFieldUnion()
}

func (SingleLineTextField) implFieldUnion()  {}
func (MultiLineTextField) implFieldUnion()   {}
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

// FieldValueUnion contains all possible properties and values from
// [SingleLineTextValue], [MultiLineTextValue], [IntegerValue], [FloatValue],
// [MonetaryValue], [PercentageValue], [BooleanValue], [EmailValue], [URLValue],
// [DomainValue], [SocialXValue], [SocialLinkedInValue], [TelephoneNumber],
// [GeoValue], [DateValue], [DatetimeValue], [ChoiceValue], [FunnelStepValue],
// [RelationValue], [[]ValueUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfArrayOfValues]
type FieldValueUnion struct {
	// This field will be present if the value is a [[]ValueUnion] instead of an
	// object.
	OfArrayOfValues []ValueUnion `json:",inline"`
	// This field is a union of [string], [string], [int64], [float64],
	// [MonetaryValueData], [float64], [bool], [string], [string], [string],
	// [SocialXValueData], [SocialLinkedInValueData], [string], [string], [time.Time],
	// [time.Time], [ChoiceFieldOption], [FunnelStep]
	Data FieldValueUnionData `json:"data"`
	Type string              `json:"type"`
	// This field is from variant [RelationValue].
	Item ItemPointer `json:"item"`
	JSON struct {
		OfArrayOfValues respjson.Field
		Data            respjson.Field
		Type            respjson.Field
		Item            respjson.Field
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
	// This field is from variant [ChoiceFieldOption].
	Label string `json:"label"`
	Type  string `json:"type"`
	// This field is from variant [FunnelStep].
	Name string `json:"name"`
	// This field is from variant [FunnelStep].
	StepType FunnelStepStepType `json:"step_type"`
	JSON     struct {
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
		Label        respjson.Field
		Type         respjson.Field
		Name         respjson.Field
		StepType     respjson.Field
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

func FieldValueParamOfX(data FieldValueParamXData) FieldValueParamUnion {
	var variant FieldValueParamX
	variant.Data = data
	return FieldValueParamUnion{OfX: &variant}
}

func FieldValueParamOfLinkedIn(data FieldValueParamLinkedInData) FieldValueParamUnion {
	var variant FieldValueParamLinkedIn
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

func FieldValueParamOfChoice[T ChoiceFieldOptionParam | shared.PointerParam](data T) FieldValueParamUnion {
	var variant ChoiceValueParam
	switch v := any(data).(type) {
	case ChoiceFieldOptionParam:
		variant.Data.OfChoiceFieldOption = &v
	case shared.PointerParam:
		variant.Data.OfPointer = &v
	}
	return FieldValueParamUnion{OfChoice: &variant}
}

func FieldValueParamOfFunnelStep[T FunnelStepParam | shared.PointerParam](data T) FieldValueParamUnion {
	var variant FunnelStepValueParam
	switch v := any(data).(type) {
	case FunnelStepParam:
		variant.Data.OfFunnelStep = &v
	case shared.PointerParam:
		variant.Data.OfPointer = &v
	}
	return FieldValueParamUnion{OfFunnelStep: &variant}
}

func FieldValueParamOfRelation[T ItemPointerParam | shared.PointerParam](item T) FieldValueParamUnion {
	var variant RelationValueParam
	switch v := any(item).(type) {
	case ItemPointerParam:
		variant.Item.OfItemPointer = &v
	case shared.PointerParam:
		variant.Item.OfPointer = &v
	}
	return FieldValueParamUnion{OfRelation: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FieldValueParamUnion struct {
	OfSingleLineText  *SingleLineTextValueParam `json:",omitzero,inline"`
	OfMultiLineText   *MultiLineTextValueParam  `json:",omitzero,inline"`
	OfInteger         *IntegerValueParam        `json:",omitzero,inline"`
	OfFloat           *FloatValueParam          `json:",omitzero,inline"`
	OfMonetary        *MonetaryValueParam       `json:",omitzero,inline"`
	OfPercentage      *PercentageValueParam     `json:",omitzero,inline"`
	OfBoolean         *BooleanValueParam        `json:",omitzero,inline"`
	OfEmail           *EmailValueParam          `json:",omitzero,inline"`
	OfURL             *URLValueParam            `json:",omitzero,inline"`
	OfDomain          *DomainValueParam         `json:",omitzero,inline"`
	OfX               *FieldValueParamX         `json:",omitzero,inline"`
	OfLinkedIn        *FieldValueParamLinkedIn  `json:",omitzero,inline"`
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
func (u FieldValueParamUnion) GetItem() *RelationValueParamItemUnion {
	if vt := u.OfRelation; vt != nil {
		return &vt.Item
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueParamUnion) GetType() *string {
	if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMultiLineText; vt != nil {
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
		res.any = vt.Data.asAny()
	} else if vt := u.OfFunnelStep; vt != nil {
		res.any = vt.Data.asAny()
	}
	return
}

// Can have the runtime types [*string], [*int64], [*float64],
// [*MonetaryValueDataParam], [*bool], [*FieldValueParamXData],
// [*FieldValueParamLinkedInData], [*time.Time], [*ChoiceFieldOptionParam],
// [*shared.PointerParam], [*FunnelStepParam]
type fieldValueParamUnionData struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *int64:
//	case *float64:
//	case *moonbase.MonetaryValueDataParam:
//	case *bool:
//	case *moonbase.FieldValueParamXData:
//	case *moonbase.FieldValueParamLinkedInData:
//	case *time.Time:
//	case *moonbase.ChoiceFieldOptionParam:
//	case *shared.PointerParam:
//	case *moonbase.FunnelStepParam:
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
func (u fieldValueParamUnionData) GetLabel() *string {
	switch vt := u.any.(type) {
	case *ChoiceValueParamDataUnion:
		return vt.GetLabel()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetName() *string {
	switch vt := u.any.(type) {
	case *FunnelStepValueParamDataUnion:
		return vt.GetName()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetStepType() *string {
	switch vt := u.any.(type) {
	case *FunnelStepValueParamDataUnion:
		return vt.GetStepType()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetURL() *string {
	switch vt := u.any.(type) {
	case *FieldValueParamXData:
		return paramutil.AddrIfPresent(vt.URL)
	case *FieldValueParamLinkedInData:
		return paramutil.AddrIfPresent(vt.URL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetUsername() *string {
	switch vt := u.any.(type) {
	case *FieldValueParamXData:
		return paramutil.AddrIfPresent(vt.Username)
	case *FieldValueParamLinkedInData:
		return paramutil.AddrIfPresent(vt.Username)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetID() *string {
	switch vt := u.any.(type) {
	case *ChoiceValueParamDataUnion:
		return vt.GetID()
	case *FunnelStepValueParamDataUnion:
		return vt.GetID()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueParamUnionData) GetType() *string {
	switch vt := u.any.(type) {
	case *ChoiceValueParamDataUnion:
		return vt.GetType()
	case *FunnelStepValueParamDataUnion:
		return vt.GetType()
	}
	return nil
}

// The social media profile for the X (formerly Twitter) platform
//
// The properties Data, Type are required.
type FieldValueParamX struct {
	// Social media profile information including both the full URL and extracted
	// username.
	Data FieldValueParamXData `json:"data,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/uri/social_x".
	Type constant.ValueUriSocialX `json:"type,required"`
	paramObj
}

func (r FieldValueParamX) MarshalJSON() (data []byte, err error) {
	type shadow FieldValueParamX
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FieldValueParamX) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Social media profile information including both the full URL and extracted
// username.
type FieldValueParamXData struct {
	// The full URL to the X profile, starting with 'https://x.com/'
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// The X username, up to 15 characters long, containing only lowercase letters
	// (a-z), uppercase letters (A-Z), numbers (0-9), and underscores (\_). Does not
	// include the '@' symbol prefix.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r FieldValueParamXData) MarshalJSON() (data []byte, err error) {
	type shadow FieldValueParamXData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FieldValueParamXData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The social media profile for the LinkedIn platform
//
// The properties Data, Type are required.
type FieldValueParamLinkedIn struct {
	// The social media profile for the LinkedIn platform
	Data FieldValueParamLinkedInData `json:"data,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/uri/social_linked_in".
	Type constant.ValueUriSocialLinkedIn `json:"type,required"`
	paramObj
}

func (r FieldValueParamLinkedIn) MarshalJSON() (data []byte, err error) {
	type shadow FieldValueParamLinkedIn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FieldValueParamLinkedIn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The social media profile for the LinkedIn platform
type FieldValueParamLinkedInData struct {
	// The full URL to the LinkedIn profile.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// The LinkedIn username, including the prefix 'company/' for company pages or
	// 'in/' for personal profiles.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r FieldValueParamLinkedInData) MarshalJSON() (data []byte, err error) {
	type shadow FieldValueParamLinkedInData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FieldValueParamLinkedInData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores decimal numbers with floating-point precision.
type FloatField struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality FloatFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Rating").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `rating`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/number/unitless_float` for this field.
	Type constant.FieldNumberUnitlessFloat `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Floating point number
type FloatValue struct {
	Data float64                           `json:"data,required"`
	Type constant.ValueNumberUnitlessFloat `json:"type,required"`
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
	Data float64 `json:"data,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/number/unitless_float".
	Type constant.ValueNumberUnitlessFloat `json:"type,required"`
	paramObj
}

func (r FloatValueParam) MarshalJSON() (data []byte, err error) {
	type shadow FloatValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FloatValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Funnel step value
type FunnelStepValue struct {
	// A specific funnel step, as configured on the Funnel
	Data FunnelStep               `json:"data,required"`
	Type constant.ValueFunnelStep `json:"type,required"`
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
	// A specific funnel step, as configured on the Funnel
	Data FunnelStepValueParamDataUnion `json:"data,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/funnel_step".
	Type constant.ValueFunnelStep `json:"type,required"`
	paramObj
}

func (r FunnelStepValueParam) MarshalJSON() (data []byte, err error) {
	type shadow FunnelStepValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelStepValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelStepValueParamDataUnion struct {
	OfFunnelStep *FunnelStepParam     `json:",omitzero,inline"`
	OfPointer    *shared.PointerParam `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelStepValueParamDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFunnelStep, u.OfPointer)
}
func (u *FunnelStepValueParamDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FunnelStepValueParamDataUnion) asAny() any {
	if !param.IsOmitted(u.OfFunnelStep) {
		return u.OfFunnelStep
	} else if !param.IsOmitted(u.OfPointer) {
		return u.OfPointer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunnelStepValueParamDataUnion) GetName() *string {
	if vt := u.OfFunnelStep; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunnelStepValueParamDataUnion) GetStepType() *string {
	if vt := u.OfFunnelStep; vt != nil {
		return (*string)(&vt.StepType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunnelStepValueParamDataUnion) GetID() *string {
	if vt := u.OfFunnelStep; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfPointer; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunnelStepValueParamDataUnion) GetType() *string {
	if vt := u.OfFunnelStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPointer; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A field that stores geographic coordinates or location data.
type GeoField struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality GeoFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Location").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `location`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/geo` for this field.
	Type constant.FieldGeo `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Geographic coordinate value
type GeoValue struct {
	// A string that represents some geographic location. The exact format may vary
	// based on context.
	Data string            `json:"data,required"`
	Type constant.ValueGeo `json:"type,required"`
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
	Data string `json:"data,required"`
	// This field can be elided, and will marshal its zero value as "value/geo".
	Type constant.ValueGeo `json:"type,required"`
	paramObj
}

func (r GeoValueParam) MarshalJSON() (data []byte, err error) {
	type shadow GeoValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeoValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores whole numbers without decimal places.
type IntegerField struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality IntegerFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Employee Count").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `employee_count`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/number/unitless_integer` for this
	// field.
	Type constant.FieldNumberUnitlessInteger `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Integer value without units
type IntegerValue struct {
	Data int64                               `json:"data,required"`
	Type constant.ValueNumberUnitlessInteger `json:"type,required"`
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
	Data int64 `json:"data,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/number/unitless_integer".
	Type constant.ValueNumberUnitlessInteger `json:"type,required"`
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
	ID string `json:"id,required"`
	// String representing the object’s type. Always `item` for this object.
	Type constant.Item `json:"type,required"`
	// A hash where keys are the `ref` of a `Field` and values are the data stored for
	// that field.
	Values map[string]FieldValueUnion `json:"values,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
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
	ID string `json:"id,required"`
	// A reference to the `Collection` containing this item.
	Collection CollectionPointer `json:"collection,required"`
	// String representing the object’s type. Always `item` for this object.
	Type constant.Item `json:"type,required"`
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

// ToParam converts this ItemPointer to a ItemPointerParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ItemPointerParam.Overrides()
func (r ItemPointer) ToParam() ItemPointerParam {
	return param.Override[ItemPointerParam](json.RawMessage(r.RawJSON()))
}

// A reference to an `Item` within a specific `Collection`, providing the context
// needed to locate the item.
//
// The properties ID, Collection, Type are required.
type ItemPointerParam struct {
	// Unique identifier of the item.
	ID string `json:"id,required"`
	// A reference to the `Collection` containing this item.
	Collection CollectionPointerParam `json:"collection,omitzero,required"`
	// String representing the object’s type. Always `item` for this object.
	//
	// This field can be elided, and will marshal its zero value as "item".
	Type constant.Item `json:"type,required"`
	paramObj
}

func (r ItemPointerParam) MarshalJSON() (data []byte, err error) {
	type shadow ItemPointerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemPointerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that stores monetary amounts with currency information.
type MonetaryField struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality MonetaryFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Deal Value").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `deal_value`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/number/monetary` for this field.
	Type constant.FieldNumberMonetary `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Monetary or currency value
type MonetaryValue struct {
	// A monetary amount is composed of the amount in the smallest unit of a currency
	// and an ISO currency code.
	Data MonetaryValueData            `json:"data,required"`
	Type constant.ValueNumberMonetary `json:"type,required"`
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
	Currency string `json:"currency,required"`
	// The amount in the minor units of the currency. For example, $10 (10 USD) would
	// be 1000. Minor units conversion depends on the currency.
	InMinorUnits int64 `json:"in_minor_units,required"`
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
	Data MonetaryValueDataParam `json:"data,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/number/monetary".
	Type constant.ValueNumberMonetary `json:"type,required"`
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
	Currency string `json:"currency,required"`
	// The amount in the minor units of the currency. For example, $10 (10 USD) would
	// be 1000. Minor units conversion depends on the currency.
	InMinorUnits int64 `json:"in_minor_units,required"`
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
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality MultiLineTextFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Description").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `description`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/text/multi_line` for this field.
	Type constant.FieldTextMultiLine `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Multiple lines of text
type MultiLineTextValue struct {
	// Text which may contain line breaks, can be up to 65,536 characters long. Do not
	// use markdown formatting, just plain text.
	Data string                      `json:"data,required"`
	Type constant.ValueTextMultiLine `json:"type,required"`
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
	Data string `json:"data,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/text/multi_line".
	Type constant.ValueTextMultiLine `json:"type,required"`
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
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality PercentageFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Win Probability").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `win_probability`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/number/percentage` for this field.
	Type constant.FieldNumberPercentage `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Percentage numeric value
type PercentageValue struct {
	// A floating-point number representing a percentage value, for example 50.21 for
	// 50.21% or -1000 for -1000% etc.
	Data float64                        `json:"data,required"`
	Type constant.ValueNumberPercentage `json:"type,required"`
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
	Data float64 `json:"data,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/number/percentage".
	Type constant.ValueNumberPercentage `json:"type,required"`
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
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality RelationFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Account").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `account`).
	Ref string `json:"ref,required"`
	// The type of relationship. Can be `one_way` for simple references or `two_way`
	// for bidirectional relationships.
	//
	// Any of "one_way", "two_way".
	RelationType RelationFieldRelationType `json:"relation_type,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/relation` for this field.
	Type constant.FieldRelation `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Cardinality  respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Readonly     respjson.Field
		Ref          respjson.Field
		RelationType respjson.Field
		Required     respjson.Field
		Type         respjson.Field
		Unique       respjson.Field
		UpdatedAt    respjson.Field
		Description  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
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

// The type of relationship. Can be `one_way` for simple references or `two_way`
// for bidirectional relationships.
type RelationFieldRelationType string

const (
	RelationFieldRelationTypeOneWay RelationFieldRelationType = "one_way"
	RelationFieldRelationTypeTwoWay RelationFieldRelationType = "two_way"
)

// Related item reference
type RelationValue struct {
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Item ItemPointer            `json:"item,required"`
	Type constant.ValueRelation `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Item        respjson.Field
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
// The properties Item, Type are required.
type RelationValueParam struct {
	// A reference to an `Item` within a specific `Collection`, providing the context
	// needed to locate the item.
	Item RelationValueParamItemUnion `json:"item,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "value/relation".
	Type constant.ValueRelation `json:"type,required"`
	paramObj
}

func (r RelationValueParam) MarshalJSON() (data []byte, err error) {
	type shadow RelationValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RelationValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RelationValueParamItemUnion struct {
	OfItemPointer *ItemPointerParam    `json:",omitzero,inline"`
	OfPointer     *shared.PointerParam `json:",omitzero,inline"`
	paramUnion
}

func (u RelationValueParamItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfItemPointer, u.OfPointer)
}
func (u *RelationValueParamItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RelationValueParamItemUnion) asAny() any {
	if !param.IsOmitted(u.OfItemPointer) {
		return u.OfItemPointer
	} else if !param.IsOmitted(u.OfPointer) {
		return u.OfPointer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RelationValueParamItemUnion) GetCollection() *CollectionPointerParam {
	if vt := u.OfItemPointer; vt != nil {
		return &vt.Collection
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RelationValueParamItemUnion) GetID() *string {
	if vt := u.OfItemPointer; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfPointer; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RelationValueParamItemUnion) GetType() *string {
	if vt := u.OfItemPointer; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPointer; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A field that stores a single line of text without line breaks.
type SingleLineTextField struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality SingleLineTextFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Company Name").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `company_name`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/text/single_line` for this field.
	Type constant.FieldTextSingleLine `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// A single line of text
type SingleLineTextValue struct {
	// A single line of text, up to 1024 characters long. It should not contain line
	// breaks.
	Data string                       `json:"data,required"`
	Type constant.ValueTextSingleLine `json:"type,required"`
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
	Data string `json:"data,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/text/single_line".
	Type constant.ValueTextSingleLine `json:"type,required"`
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
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality SocialLinkedInFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "LinkedIn Profile").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `linkedin_profile`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/uri/social_linked_in` for this field.
	Type constant.FieldUriSocialLinkedIn `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// The social media profile for the LinkedIn platform
type SocialLinkedInValue struct {
	// The social media profile for the LinkedIn platform
	Data SocialLinkedInValueData         `json:"data,required"`
	Type constant.ValueUriSocialLinkedIn `json:"type,required"`
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
	URL string `json:"url,required" format:"uri"`
	// The LinkedIn username, including the prefix 'company/' for company pages or
	// 'in/' for personal profiles.
	Username string `json:"username,required"`
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

// A field that stores X (formerly Twitter) profile information.
type SocialXField struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality SocialXFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "X Profile").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `x_profile`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/uri/social_x` for this field.
	Type constant.FieldUriSocialX `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// The social media profile for the X (formerly Twitter) platform
type SocialXValue struct {
	// Social media profile information including both the full URL and extracted
	// username.
	Data SocialXValueData         `json:"data,required"`
	Type constant.ValueUriSocialX `json:"type,required"`
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
	URL string `json:"url,required" format:"uri"`
	// The X username, up to 15 characters long, containing only lowercase letters
	// (a-z), uppercase letters (A-Z), numbers (0-9), and underscores (\_). Does not
	// include the '@' symbol prefix.
	Username string `json:"username,required"`
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

// A field that tracks an item's position in a funnel or pipeline workflow.
type StageField struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality StageFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The `Funnel` object that defines the available stages for this field.
	Funnel Funnel `json:"funnel,required"`
	// The human-readable name of the field (e.g., "Sales Stage").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `sales_stage`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/stage` for this field.
	Type constant.FieldStage `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Funnel      respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Telephone number value
type TelephoneNumber struct {
	// A telephone number in strictly formatted E.164 format. Do not include spaces,
	// dashes, or parentheses etc.
	Data string                        `json:"data,required"`
	Type constant.ValueTelephoneNumber `json:"type,required"`
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
	Data string `json:"data,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/telephone_number".
	Type constant.ValueTelephoneNumber `json:"type,required"`
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
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality TelephoneNumberFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Phone").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `phone`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/telephone_number` for this field.
	Type constant.FieldTelephoneNumber `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// A field that stores and validates web URLs.
type URLField struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality URLFieldCardinality `json:"cardinality,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The human-readable name of the field (e.g., "Website").
	Name string `json:"name,required"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `website`).
	Ref string `json:"ref,required"`
	// If `true`, this field must have a value.
	Required bool `json:"required,required"`
	// The data type of the field. Always `field/uri/url` for this field.
	Type constant.FieldUriURL `json:"type,required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cardinality respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Readonly    respjson.Field
		Ref         respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Unique      respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// URL or web address
type URLValue struct {
	// A valid URL, conforming to RFC 3986, up to 8,192 characters long. It should
	// include the protocol, for example 'https://' or 'mailto:support@moonbase.ai'
	// etc.
	Data string               `json:"data,required" format:"uri"`
	Type constant.ValueUriURL `json:"type,required"`
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
	Data string `json:"data,required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "value/uri/url".
	Type constant.ValueUriURL `json:"type,required"`
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
// [SingleLineTextValue], [MultiLineTextValue], [IntegerValue], [FloatValue],
// [MonetaryValue], [PercentageValue], [BooleanValue], [EmailValue], [URLValue],
// [DomainValue], [SocialXValue], [SocialLinkedInValue], [TelephoneNumber],
// [GeoValue], [DateValue], [DatetimeValue], [ChoiceValue], [FunnelStepValue],
// [RelationValue].
//
// Use the [ValueUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ValueUnion struct {
	// This field is a union of [string], [string], [int64], [float64],
	// [MonetaryValueData], [float64], [bool], [string], [string], [string],
	// [SocialXValueData], [SocialLinkedInValueData], [string], [string], [time.Time],
	// [time.Time], [ChoiceFieldOption], [FunnelStep]
	Data ValueUnionData `json:"data"`
	// Any of "value/text/single_line", "value/text/multi_line",
	// "value/number/unitless_integer", "value/number/unitless_float",
	// "value/number/monetary", "value/number/percentage", "value/boolean",
	// "value/email", "value/uri/url", "value/uri/domain", "value/uri/social_x",
	// "value/uri/social_linked_in", "value/telephone_number", "value/geo",
	// "value/date", "value/datetime", "value/choice", "value/funnel_step",
	// "value/relation".
	Type string `json:"type"`
	// This field is from variant [RelationValue].
	Item ItemPointer `json:"item"`
	JSON struct {
		Data respjson.Field
		Type respjson.Field
		Item respjson.Field
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
	// This field is from variant [ChoiceFieldOption].
	Label string `json:"label"`
	Type  string `json:"type"`
	// This field is from variant [FunnelStep].
	Name string `json:"name"`
	// This field is from variant [FunnelStep].
	StepType FunnelStepStepType `json:"step_type"`
	JSON     struct {
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
		Label        respjson.Field
		Type         respjson.Field
		Name         respjson.Field
		StepType     respjson.Field
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

func ValueParamOfValueUriSocialX(data ValueParamValueUriSocialXData) ValueParamUnion {
	var valueUriSocialX ValueParamValueUriSocialX
	valueUriSocialX.Data = data
	return ValueParamUnion{OfValueUriSocialX: &valueUriSocialX}
}

func ValueParamOfValueUriSocialLinkedIn(data ValueParamValueUriSocialLinkedInData) ValueParamUnion {
	var valueUriSocialLinkedIn ValueParamValueUriSocialLinkedIn
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

func ValueParamOfValueChoice[T ChoiceFieldOptionParam | shared.PointerParam](data T) ValueParamUnion {
	var valueChoice ChoiceValueParam
	switch v := any(data).(type) {
	case ChoiceFieldOptionParam:
		valueChoice.Data.OfChoiceFieldOption = &v
	case shared.PointerParam:
		valueChoice.Data.OfPointer = &v
	}
	return ValueParamUnion{OfValueChoice: &valueChoice}
}

func ValueParamOfValueFunnelStep[T FunnelStepParam | shared.PointerParam](data T) ValueParamUnion {
	var valueFunnelStep FunnelStepValueParam
	switch v := any(data).(type) {
	case FunnelStepParam:
		valueFunnelStep.Data.OfFunnelStep = &v
	case shared.PointerParam:
		valueFunnelStep.Data.OfPointer = &v
	}
	return ValueParamUnion{OfValueFunnelStep: &valueFunnelStep}
}

func ValueParamOfValueRelation[T ItemPointerParam | shared.PointerParam](item T) ValueParamUnion {
	var valueRelation RelationValueParam
	switch v := any(item).(type) {
	case ItemPointerParam:
		valueRelation.Item.OfItemPointer = &v
	case shared.PointerParam:
		valueRelation.Item.OfPointer = &v
	}
	return ValueParamUnion{OfValueRelation: &valueRelation}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ValueParamUnion struct {
	OfValueTextSingleLine        *SingleLineTextValueParam         `json:",omitzero,inline"`
	OfValueTextMultiLine         *MultiLineTextValueParam          `json:",omitzero,inline"`
	OfValueNumberUnitlessInteger *IntegerValueParam                `json:",omitzero,inline"`
	OfValueNumberUnitlessFloat   *FloatValueParam                  `json:",omitzero,inline"`
	OfValueNumberMonetary        *MonetaryValueParam               `json:",omitzero,inline"`
	OfValueNumberPercentage      *PercentageValueParam             `json:",omitzero,inline"`
	OfValueBoolean               *BooleanValueParam                `json:",omitzero,inline"`
	OfValueEmail                 *EmailValueParam                  `json:",omitzero,inline"`
	OfValueUriURL                *URLValueParam                    `json:",omitzero,inline"`
	OfValueUriDomain             *DomainValueParam                 `json:",omitzero,inline"`
	OfValueUriSocialX            *ValueParamValueUriSocialX        `json:",omitzero,inline"`
	OfValueUriSocialLinkedIn     *ValueParamValueUriSocialLinkedIn `json:",omitzero,inline"`
	OfValueTelephoneNumber       *TelephoneNumberParam             `json:",omitzero,inline"`
	OfValueGeo                   *GeoValueParam                    `json:",omitzero,inline"`
	OfValueDate                  *DateValueParam                   `json:",omitzero,inline"`
	OfValueDatetime              *DatetimeValueParam               `json:",omitzero,inline"`
	OfValueChoice                *ChoiceValueParam                 `json:",omitzero,inline"`
	OfValueFunnelStep            *FunnelStepValueParam             `json:",omitzero,inline"`
	OfValueRelation              *RelationValueParam               `json:",omitzero,inline"`
	paramUnion
}

func (u ValueParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfValueTextSingleLine,
		u.OfValueTextMultiLine,
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
func (u ValueParamUnion) GetItem() *RelationValueParamItemUnion {
	if vt := u.OfValueRelation; vt != nil {
		return &vt.Item
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueParamUnion) GetType() *string {
	if vt := u.OfValueTextSingleLine; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfValueTextMultiLine; vt != nil {
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
		res.any = vt.Data.asAny()
	} else if vt := u.OfValueFunnelStep; vt != nil {
		res.any = vt.Data.asAny()
	}
	return
}

// Can have the runtime types [*string], [*int64], [*float64],
// [*MonetaryValueDataParam], [*bool], [*ValueParamValueUriSocialXData],
// [*ValueParamValueUriSocialLinkedInData], [*time.Time],
// [*ChoiceFieldOptionParam], [*shared.PointerParam], [*FunnelStepParam]
type valueParamUnionData struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *int64:
//	case *float64:
//	case *moonbase.MonetaryValueDataParam:
//	case *bool:
//	case *moonbase.ValueParamValueUriSocialXData:
//	case *moonbase.ValueParamValueUriSocialLinkedInData:
//	case *time.Time:
//	case *moonbase.ChoiceFieldOptionParam:
//	case *shared.PointerParam:
//	case *moonbase.FunnelStepParam:
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
func (u valueParamUnionData) GetLabel() *string {
	switch vt := u.any.(type) {
	case *ChoiceValueParamDataUnion:
		return vt.GetLabel()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetName() *string {
	switch vt := u.any.(type) {
	case *FunnelStepValueParamDataUnion:
		return vt.GetName()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetStepType() *string {
	switch vt := u.any.(type) {
	case *FunnelStepValueParamDataUnion:
		return vt.GetStepType()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetURL() *string {
	switch vt := u.any.(type) {
	case *ValueParamValueUriSocialXData:
		return paramutil.AddrIfPresent(vt.URL)
	case *ValueParamValueUriSocialLinkedInData:
		return paramutil.AddrIfPresent(vt.URL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetUsername() *string {
	switch vt := u.any.(type) {
	case *ValueParamValueUriSocialXData:
		return paramutil.AddrIfPresent(vt.Username)
	case *ValueParamValueUriSocialLinkedInData:
		return paramutil.AddrIfPresent(vt.Username)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetID() *string {
	switch vt := u.any.(type) {
	case *ChoiceValueParamDataUnion:
		return vt.GetID()
	case *FunnelStepValueParamDataUnion:
		return vt.GetID()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueParamUnionData) GetType() *string {
	switch vt := u.any.(type) {
	case *ChoiceValueParamDataUnion:
		return vt.GetType()
	case *FunnelStepValueParamDataUnion:
		return vt.GetType()
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ValueParamUnion](
		"type",
		apijson.Discriminator[SingleLineTextValueParam]("value/text/single_line"),
		apijson.Discriminator[MultiLineTextValueParam]("value/text/multi_line"),
		apijson.Discriminator[IntegerValueParam]("value/number/unitless_integer"),
		apijson.Discriminator[FloatValueParam]("value/number/unitless_float"),
		apijson.Discriminator[MonetaryValueParam]("value/number/monetary"),
		apijson.Discriminator[PercentageValueParam]("value/number/percentage"),
		apijson.Discriminator[BooleanValueParam]("value/boolean"),
		apijson.Discriminator[EmailValueParam]("value/email"),
		apijson.Discriminator[URLValueParam]("value/uri/url"),
		apijson.Discriminator[DomainValueParam]("value/uri/domain"),
		apijson.Discriminator[ValueParamValueUriSocialX]("value/uri/social_x"),
		apijson.Discriminator[ValueParamValueUriSocialLinkedIn]("value/uri/social_linked_in"),
		apijson.Discriminator[TelephoneNumberParam]("value/telephone_number"),
		apijson.Discriminator[GeoValueParam]("value/geo"),
		apijson.Discriminator[DateValueParam]("value/date"),
		apijson.Discriminator[DatetimeValueParam]("value/datetime"),
		apijson.Discriminator[ChoiceValueParam]("value/choice"),
		apijson.Discriminator[FunnelStepValueParam]("value/funnel_step"),
		apijson.Discriminator[RelationValueParam]("value/relation"),
	)
}

// The social media profile for the X (formerly Twitter) platform
//
// The properties Data, Type are required.
type ValueParamValueUriSocialX struct {
	// Social media profile information including both the full URL and extracted
	// username.
	Data ValueParamValueUriSocialXData `json:"data,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/uri/social_x".
	Type constant.ValueUriSocialX `json:"type,required"`
	paramObj
}

func (r ValueParamValueUriSocialX) MarshalJSON() (data []byte, err error) {
	type shadow ValueParamValueUriSocialX
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ValueParamValueUriSocialX) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Social media profile information including both the full URL and extracted
// username.
type ValueParamValueUriSocialXData struct {
	// The full URL to the X profile, starting with 'https://x.com/'
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// The X username, up to 15 characters long, containing only lowercase letters
	// (a-z), uppercase letters (A-Z), numbers (0-9), and underscores (\_). Does not
	// include the '@' symbol prefix.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r ValueParamValueUriSocialXData) MarshalJSON() (data []byte, err error) {
	type shadow ValueParamValueUriSocialXData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ValueParamValueUriSocialXData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The social media profile for the LinkedIn platform
//
// The properties Data, Type are required.
type ValueParamValueUriSocialLinkedIn struct {
	// The social media profile for the LinkedIn platform
	Data ValueParamValueUriSocialLinkedInData `json:"data,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/uri/social_linked_in".
	Type constant.ValueUriSocialLinkedIn `json:"type,required"`
	paramObj
}

func (r ValueParamValueUriSocialLinkedIn) MarshalJSON() (data []byte, err error) {
	type shadow ValueParamValueUriSocialLinkedIn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ValueParamValueUriSocialLinkedIn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The social media profile for the LinkedIn platform
type ValueParamValueUriSocialLinkedInData struct {
	// The full URL to the LinkedIn profile.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// The LinkedIn username, including the prefix 'company/' for company pages or
	// 'in/' for personal profiles.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r ValueParamValueUriSocialLinkedInData) MarshalJSON() (data []byte, err error) {
	type shadow ValueParamValueUriSocialLinkedInData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ValueParamValueUriSocialLinkedInData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionGetParams struct {
	// Specifies which related objects to include in the response.
	//
	// Any of "views".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectionGetParams]'s query parameters as `url.Values`.
func (r CollectionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
