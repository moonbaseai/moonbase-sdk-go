// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// Manage your collections and items
//
// CollectionFieldService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectionFieldService] method instead.
type CollectionFieldService struct {
	Options []option.RequestOption
}

// NewCollectionFieldService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCollectionFieldService(opts ...option.RequestOption) (r CollectionFieldService) {
	r = CollectionFieldService{}
	r.Options = opts
	return
}

// Creates a new field in a collection.
func (r *CollectionFieldService) New(ctx context.Context, collectionID string, body CollectionFieldNewParams, opts ...option.RequestOption) (res *FieldUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if collectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("collections/%s/fields", collectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the details of a field in a collection.
func (r *CollectionFieldService) Get(ctx context.Context, id string, query CollectionFieldGetParams, opts ...option.RequestOption) (res *FieldUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.CollectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("collections/%s/fields/%s", query.CollectionID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing field in a collection.
func (r *CollectionFieldService) Update(ctx context.Context, id string, params CollectionFieldUpdateParams, opts ...option.RequestOption) (res *FieldUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CollectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("collections/%s/fields/%s", params.CollectionID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Permanently deletes a field from a collection.
func (r *CollectionFieldService) Delete(ctx context.Context, id string, body CollectionFieldDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CollectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("collections/%s/fields/%s", body.CollectionID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type CollectionFieldNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a single-line text field.
	OfFieldTextSingleLine *CollectionFieldNewParamsFieldFieldTextSingleLine `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a multi-line text field.
	OfFieldTextMultiLine *CollectionFieldNewParamsFieldFieldTextMultiLine `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfFieldIdentifier *CollectionFieldNewParamsFieldFieldIdentifier `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating an integer field.
	OfFieldNumberUnitlessInteger *CollectionFieldNewParamsFieldFieldNumberUnitlessInteger `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a decimal number field.
	OfFieldNumberUnitlessFloat *CollectionFieldNewParamsFieldFieldNumberUnitlessFloat `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a monetary field.
	OfFieldNumberMonetary *CollectionFieldNewParamsFieldFieldNumberMonetary `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a percentage field.
	OfFieldNumberPercentage *CollectionFieldNewParamsFieldFieldNumberPercentage `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a boolean field.
	OfFieldBoolean *CollectionFieldNewParamsFieldFieldBoolean `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating an email field.
	OfFieldEmail *CollectionFieldNewParamsFieldFieldEmail `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a URL field.
	OfFieldUriURL *CollectionFieldNewParamsFieldFieldUriURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a domain field.
	OfFieldUriDomain *CollectionFieldNewParamsFieldFieldUriDomain `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating an X (formerly Twitter) profile field.
	OfFieldUriSocialX *CollectionFieldNewParamsFieldFieldUriSocialX `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a LinkedIn profile field.
	OfFieldUriSocialLinkedIn *CollectionFieldNewParamsFieldFieldUriSocialLinkedIn `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a telephone number field.
	OfFieldTelephoneNumber *CollectionFieldNewParamsFieldFieldTelephoneNumber `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a geographic location field.
	OfFieldGeo *CollectionFieldNewParamsFieldFieldGeo `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a date field.
	OfFieldDate *CollectionFieldNewParamsFieldFieldDate `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a date and time field.
	OfFieldDatetime *CollectionFieldNewParamsFieldFieldDatetime `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a choice field with predefined options.
	OfFieldChoice *CollectionFieldNewParamsFieldFieldChoice `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a stage field.
	OfFieldStage *StageFieldCreateParams `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for creating a relation field that links items across collections.
	OfFieldRelation *CollectionFieldNewParamsFieldFieldRelation `json:",inline"`

	paramObj
}

func (u CollectionFieldNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFieldTextSingleLine,
		u.OfFieldTextMultiLine,
		u.OfFieldIdentifier,
		u.OfFieldNumberUnitlessInteger,
		u.OfFieldNumberUnitlessFloat,
		u.OfFieldNumberMonetary,
		u.OfFieldNumberPercentage,
		u.OfFieldBoolean,
		u.OfFieldEmail,
		u.OfFieldUriURL,
		u.OfFieldUriDomain,
		u.OfFieldUriSocialX,
		u.OfFieldUriSocialLinkedIn,
		u.OfFieldTelephoneNumber,
		u.OfFieldGeo,
		u.OfFieldDate,
		u.OfFieldDatetime,
		u.OfFieldChoice,
		u.OfFieldStage,
		u.OfFieldRelation)
}
func (r *CollectionFieldNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating a single-line text field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldTextSingleLine struct {
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
	Cardinality   string                     `json:"cardinality,omitzero"`
	DefaultValues []SingleLineTextValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/text/single_line`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/text/single_line".
	Type constant.FieldTextSingleLine `json:"type" default:"field/text/single_line"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldTextSingleLine) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldTextSingleLine
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldTextSingleLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldTextSingleLine](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a multi-line text field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldTextMultiLine struct {
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
	Cardinality   string                    `json:"cardinality,omitzero"`
	DefaultValues []MultiLineTextValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/text/multi_line`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/text/multi_line".
	Type constant.FieldTextMultiLine `json:"type" default:"field/text/multi_line"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldTextMultiLine) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldTextMultiLine
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldTextMultiLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldTextMultiLine](
		"cardinality", "one", "many",
	)
}

// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldIdentifier struct {
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Required    param.Opt[bool]   `json:"required,omitzero"`
	Unique      param.Opt[bool]   `json:"unique,omitzero"`
	// Any of "one", "many".
	Cardinality   string                 `json:"cardinality,omitzero"`
	DefaultValues []IdentifierValueParam `json:"default_values,omitzero"`
	// This field can be elided, and will marshal its zero value as "field/identifier".
	Type constant.FieldIdentifier `json:"type" default:"field/identifier"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldIdentifier) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldIdentifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldIdentifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldIdentifier](
		"cardinality", "one", "many",
	)
}

// Parameters for creating an integer field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldNumberUnitlessInteger struct {
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
	Cardinality   string              `json:"cardinality,omitzero"`
	DefaultValues []IntegerValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/number/unitless_integer`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/number/unitless_integer".
	Type constant.FieldNumberUnitlessInteger `json:"type" default:"field/number/unitless_integer"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldNumberUnitlessInteger) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldNumberUnitlessInteger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldNumberUnitlessInteger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldNumberUnitlessInteger](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a decimal number field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldNumberUnitlessFloat struct {
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
	Cardinality   string            `json:"cardinality,omitzero"`
	DefaultValues []FloatValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/number/unitless_float`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/number/unitless_float".
	Type constant.FieldNumberUnitlessFloat `json:"type" default:"field/number/unitless_float"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldNumberUnitlessFloat) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldNumberUnitlessFloat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldNumberUnitlessFloat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldNumberUnitlessFloat](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a monetary field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldNumberMonetary struct {
	// The human-readable name for the field.
	Name string `json:"name" api:"required"`
	// The default currency for the field, as a 3-letter ISO 4217 code (e.g., `USD`,
	// `EUR`, `GBP`).
	DefaultUnit param.Opt[string] `json:"default_unit,omitzero"`
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
	Cardinality   string               `json:"cardinality,omitzero"`
	DefaultValues []MonetaryValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/number/monetary`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/number/monetary".
	Type constant.FieldNumberMonetary `json:"type" default:"field/number/monetary"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldNumberMonetary) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldNumberMonetary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldNumberMonetary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldNumberMonetary](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a percentage field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldNumberPercentage struct {
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
	Cardinality   string                 `json:"cardinality,omitzero"`
	DefaultValues []PercentageValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/number/percentage`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/number/percentage".
	Type constant.FieldNumberPercentage `json:"type" default:"field/number/percentage"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldNumberPercentage) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldNumberPercentage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldNumberPercentage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldNumberPercentage](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a boolean field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldBoolean struct {
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
	Cardinality   string              `json:"cardinality,omitzero"`
	DefaultValues []BooleanValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/boolean`.
	//
	// This field can be elided, and will marshal its zero value as "field/boolean".
	Type constant.FieldBoolean `json:"type" default:"field/boolean"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldBoolean) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldBoolean
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldBoolean](
		"cardinality", "one", "many",
	)
}

// Parameters for creating an email field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldEmail struct {
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
	Cardinality   string            `json:"cardinality,omitzero"`
	DefaultValues []EmailValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/email`.
	//
	// This field can be elided, and will marshal its zero value as "field/email".
	Type constant.FieldEmail `json:"type" default:"field/email"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldEmail) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldEmail](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a URL field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldUriURL struct {
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
	Cardinality   string          `json:"cardinality,omitzero"`
	DefaultValues []URLValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/uri/url`.
	//
	// This field can be elided, and will marshal its zero value as "field/uri/url".
	Type constant.FieldUriURL `json:"type" default:"field/uri/url"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldUriURL) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldUriURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldUriURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldUriURL](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a domain field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldUriDomain struct {
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
	Cardinality   string             `json:"cardinality,omitzero"`
	DefaultValues []DomainValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/uri/domain`.
	//
	// This field can be elided, and will marshal its zero value as "field/uri/domain".
	Type constant.FieldUriDomain `json:"type" default:"field/uri/domain"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldUriDomain) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldUriDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldUriDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldUriDomain](
		"cardinality", "one", "many",
	)
}

// Parameters for creating an X (formerly Twitter) profile field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldUriSocialX struct {
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
	Cardinality   string              `json:"cardinality,omitzero"`
	DefaultValues []SocialXValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/uri/social_x`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/uri/social_x".
	Type constant.FieldUriSocialX `json:"type" default:"field/uri/social_x"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldUriSocialX) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldUriSocialX
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldUriSocialX) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldUriSocialX](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a LinkedIn profile field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldUriSocialLinkedIn struct {
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
	Cardinality   string                     `json:"cardinality,omitzero"`
	DefaultValues []SocialLinkedInValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/uri/social_linked_in`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/uri/social_linked_in".
	Type constant.FieldUriSocialLinkedIn `json:"type" default:"field/uri/social_linked_in"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldUriSocialLinkedIn) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldUriSocialLinkedIn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldUriSocialLinkedIn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldUriSocialLinkedIn](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a telephone number field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldTelephoneNumber struct {
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
	Cardinality   string                 `json:"cardinality,omitzero"`
	DefaultValues []TelephoneNumberParam `json:"default_values,omitzero"`
	// The field type. Must be `field/telephone_number`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/telephone_number".
	Type constant.FieldTelephoneNumber `json:"type" default:"field/telephone_number"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldTelephoneNumber) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldTelephoneNumber
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldTelephoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldTelephoneNumber](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a geographic location field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldGeo struct {
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
	Cardinality   string          `json:"cardinality,omitzero"`
	DefaultValues []GeoValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/geo`.
	//
	// This field can be elided, and will marshal its zero value as "field/geo".
	Type constant.FieldGeo `json:"type" default:"field/geo"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldGeo) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldGeo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldGeo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldGeo](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a date field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldDate struct {
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
	Cardinality   string                            `json:"cardinality,omitzero"`
	DefaultValues []DateFieldDefaultValueParamUnion `json:"default_values,omitzero"`
	// The field type. Must be `field/date`.
	//
	// This field can be elided, and will marshal its zero value as "field/date".
	Type constant.FieldDate `json:"type" default:"field/date"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldDate) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldDate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldDate](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a date and time field.
//
// The properties Name, Type are required.
type CollectionFieldNewParamsFieldFieldDatetime struct {
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
	Cardinality   string                                `json:"cardinality,omitzero"`
	DefaultValues []DatetimeFieldDefaultValueParamUnion `json:"default_values,omitzero"`
	// The field type. Must be `field/datetime`.
	//
	// This field can be elided, and will marshal its zero value as "field/datetime".
	Type constant.FieldDatetime `json:"type" default:"field/datetime"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldDatetime) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldDatetime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldDatetime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldDatetime](
		"cardinality", "one", "many",
	)
}

// Parameters for creating a choice field with predefined options.
//
// The properties Name, Options, Type are required.
type CollectionFieldNewParamsFieldFieldChoice struct {
	// The human-readable name for the field.
	Name string `json:"name" api:"required"`
	// A list of options to create for the field. Each option must have a `name`.
	Options []CollectionFieldNewParamsFieldFieldChoiceOption `json:"options,omitzero" api:"required"`
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
	Cardinality   string             `json:"cardinality,omitzero"`
	DefaultValues []ChoiceValueParam `json:"default_values,omitzero"`
	// The field type. Must be `field/choice`.
	//
	// This field can be elided, and will marshal its zero value as "field/choice".
	Type constant.FieldChoice `json:"type" default:"field/choice"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldChoice) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldChoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldChoice](
		"cardinality", "one", "many",
	)
}

// Parameters for defining an option in a choice field.
//
// The properties Color, Name are required.
type CollectionFieldNewParamsFieldFieldChoiceOption struct {
	// The color of the option.
	//
	// Any of "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime",
	// "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet",
	// "yellow".
	Color string `json:"color,omitzero" api:"required"`
	// The display name of the option.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldChoiceOption) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldChoiceOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldChoiceOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldChoiceOption](
		"color", "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime", "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet", "yellow",
	)
}

// Parameters for creating a relation field that links items across collections.
//
// The properties AllowedCollections, Name, RelationType, Type are required.
type CollectionFieldNewParamsFieldFieldRelation struct {
	// A list of collection IDs or `ref` values that are valid targets for this
	// relation.
	AllowedCollections []CollectionFieldNewParamsFieldFieldRelationAllowedCollection `json:"allowed_collections,omitzero" api:"required"`
	// The human-readable name for the field.
	Name string `json:"name" api:"required"`
	// The type of relationship: `one_way` for simple references, or `two_way` for
	// bidirectional relationships.
	//
	// Any of "one_way", "two_way".
	RelationType string `json:"relation_type,omitzero" api:"required"`
	// An optional description of the field's purpose.
	Description param.Opt[string] `json:"description,omitzero"`
	// If `true`, items must have a value for this field. Defaults to `false`.
	Required param.Opt[bool] `json:"required,omitzero"`
	// For `two_way` relations, the name of the reverse field created on the target
	// collection.
	ReverseFieldName param.Opt[string] `json:"reverse_field_name,omitzero"`
	// If `true`, values must be unique across all items. Defaults to `false`.
	Unique param.Opt[bool] `json:"unique,omitzero"`
	// Whether the field holds a single value (`one`) or multiple values (`many`).
	// Defaults to `one`.
	//
	// Any of "one", "many".
	Cardinality   string                                `json:"cardinality,omitzero"`
	DefaultValues []RelationFieldDefaultValueParamUnion `json:"default_values,omitzero"`
	// The field type. Must be `field/relation`.
	//
	// This field can be elided, and will marshal its zero value as "field/relation".
	Type constant.FieldRelation `json:"type" default:"field/relation"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldRelation) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldRelation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldRelation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldRelation](
		"relation_type", "one_way", "two_way",
	)
	apijson.RegisterFieldValidator[CollectionFieldNewParamsFieldFieldRelation](
		"cardinality", "one", "many",
	)
}

// A reference to a `Collection` used in request bodies. Provide at least one of
// `id` or `ref` to identify the collection.
//
// The property Type is required.
type CollectionFieldNewParamsFieldFieldRelationAllowedCollection struct {
	// Unique identifier of the collection.
	ID param.Opt[string] `json:"id,omitzero"`
	// The stable, machine-readable reference identifier of the collection.
	Ref param.Opt[string] `json:"ref,omitzero"`
	// String representing the object’s type. Always `collection` for this object.
	//
	// This field can be elided, and will marshal its zero value as "collection".
	Type constant.Collection `json:"type" default:"collection"`
	paramObj
}

func (r CollectionFieldNewParamsFieldFieldRelationAllowedCollection) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldNewParamsFieldFieldRelationAllowedCollection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldNewParamsFieldFieldRelationAllowedCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionFieldGetParams struct {
	CollectionID string `path:"collection_id" api:"required" json:"-"`
	paramObj
}

type CollectionFieldUpdateParams struct {
	CollectionID string `path:"collection_id" api:"required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a single-line text field.
	OfFieldTextSingleLine *CollectionFieldUpdateParamsFieldFieldTextSingleLine `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a multi-line text field.
	OfFieldTextMultiLine *CollectionFieldUpdateParamsFieldFieldTextMultiLine `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfFieldIdentifier *CollectionFieldUpdateParamsFieldFieldIdentifier `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating an integer field.
	OfFieldNumberUnitlessInteger *CollectionFieldUpdateParamsFieldFieldNumberUnitlessInteger `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a decimal number field.
	OfFieldNumberUnitlessFloat *CollectionFieldUpdateParamsFieldFieldNumberUnitlessFloat `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a monetary field.
	OfFieldNumberMonetary *CollectionFieldUpdateParamsFieldFieldNumberMonetary `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a percentage field.
	OfFieldNumberPercentage *CollectionFieldUpdateParamsFieldFieldNumberPercentage `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a boolean field.
	OfFieldBoolean *CollectionFieldUpdateParamsFieldFieldBoolean `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating an email field.
	OfFieldEmail *CollectionFieldUpdateParamsFieldFieldEmail `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a URL field.
	OfFieldUriURL *CollectionFieldUpdateParamsFieldFieldUriURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a domain field.
	OfFieldUriDomain *CollectionFieldUpdateParamsFieldFieldUriDomain `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating an X (formerly Twitter) profile field.
	OfFieldUriSocialX *CollectionFieldUpdateParamsFieldFieldUriSocialX `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a LinkedIn profile field.
	OfFieldUriSocialLinkedIn *CollectionFieldUpdateParamsFieldFieldUriSocialLinkedIn `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a telephone number field.
	OfFieldTelephoneNumber *CollectionFieldUpdateParamsFieldFieldTelephoneNumber `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a geographic location field.
	OfFieldGeo *CollectionFieldUpdateParamsFieldFieldGeo `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a date field.
	OfFieldDate *CollectionFieldUpdateParamsFieldFieldDate `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a date and time field.
	OfFieldDatetime *CollectionFieldUpdateParamsFieldFieldDatetime `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a choice field.
	OfFieldChoice *CollectionFieldUpdateParamsFieldFieldChoice `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a stage field.
	OfFieldStage *StageFieldUpdateParams `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for updating a relation field.
	OfFieldRelation *CollectionFieldUpdateParamsFieldFieldRelation `json:",inline"`

	paramObj
}

func (u CollectionFieldUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFieldTextSingleLine,
		u.OfFieldTextMultiLine,
		u.OfFieldIdentifier,
		u.OfFieldNumberUnitlessInteger,
		u.OfFieldNumberUnitlessFloat,
		u.OfFieldNumberMonetary,
		u.OfFieldNumberPercentage,
		u.OfFieldBoolean,
		u.OfFieldEmail,
		u.OfFieldUriURL,
		u.OfFieldUriDomain,
		u.OfFieldUriSocialX,
		u.OfFieldUriSocialLinkedIn,
		u.OfFieldTelephoneNumber,
		u.OfFieldGeo,
		u.OfFieldDate,
		u.OfFieldDatetime,
		u.OfFieldChoice,
		u.OfFieldStage,
		u.OfFieldRelation)
}
func (r *CollectionFieldUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for updating a single-line text field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldTextSingleLine struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]            `json:"unique,omitzero"`
	DefaultValues []SingleLineTextValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/text/single_line`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/text/single_line".
	Type constant.FieldTextSingleLine `json:"type" default:"field/text/single_line"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldTextSingleLine) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldTextSingleLine
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldTextSingleLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldTextSingleLine](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a multi-line text field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldTextMultiLine struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]           `json:"unique,omitzero"`
	DefaultValues []MultiLineTextValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/text/multi_line`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/text/multi_line".
	Type constant.FieldTextMultiLine `json:"type" default:"field/text/multi_line"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldTextMultiLine) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldTextMultiLine
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldTextMultiLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldTextMultiLine](
		"cardinality", "one", "many",
	)
}

// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldIdentifier struct {
	Description   param.Opt[string]      `json:"description,omitzero"`
	Name          param.Opt[string]      `json:"name,omitzero"`
	Required      param.Opt[bool]        `json:"required,omitzero"`
	Unique        param.Opt[bool]        `json:"unique,omitzero"`
	DefaultValues []IdentifierValueParam `json:"default_values,omitzero"`
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// This field can be elided, and will marshal its zero value as "field/identifier".
	Type constant.FieldIdentifier `json:"type" default:"field/identifier"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldIdentifier) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldIdentifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldIdentifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldIdentifier](
		"cardinality", "one", "many",
	)
}

// Parameters for updating an integer field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldNumberUnitlessInteger struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]     `json:"unique,omitzero"`
	DefaultValues []IntegerValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/number/unitless_integer`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/number/unitless_integer".
	Type constant.FieldNumberUnitlessInteger `json:"type" default:"field/number/unitless_integer"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldNumberUnitlessInteger) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldNumberUnitlessInteger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldNumberUnitlessInteger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldNumberUnitlessInteger](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a decimal number field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldNumberUnitlessFloat struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]   `json:"unique,omitzero"`
	DefaultValues []FloatValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/number/unitless_float`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/number/unitless_float".
	Type constant.FieldNumberUnitlessFloat `json:"type" default:"field/number/unitless_float"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldNumberUnitlessFloat) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldNumberUnitlessFloat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldNumberUnitlessFloat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldNumberUnitlessFloat](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a monetary field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldNumberMonetary struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The default currency for the field, as a 3-letter ISO 4217 code (e.g., `USD`,
	// `EUR`, `GBP`).
	DefaultUnit param.Opt[string] `json:"default_unit,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]      `json:"unique,omitzero"`
	DefaultValues []MonetaryValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/number/monetary`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/number/monetary".
	Type constant.FieldNumberMonetary `json:"type" default:"field/number/monetary"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldNumberMonetary) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldNumberMonetary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldNumberMonetary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldNumberMonetary](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a percentage field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldNumberPercentage struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]        `json:"unique,omitzero"`
	DefaultValues []PercentageValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/number/percentage`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/number/percentage".
	Type constant.FieldNumberPercentage `json:"type" default:"field/number/percentage"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldNumberPercentage) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldNumberPercentage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldNumberPercentage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldNumberPercentage](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a boolean field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldBoolean struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]     `json:"unique,omitzero"`
	DefaultValues []BooleanValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/boolean`.
	//
	// This field can be elided, and will marshal its zero value as "field/boolean".
	Type constant.FieldBoolean `json:"type" default:"field/boolean"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldBoolean) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldBoolean
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldBoolean](
		"cardinality", "one", "many",
	)
}

// Parameters for updating an email field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldEmail struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]   `json:"unique,omitzero"`
	DefaultValues []EmailValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/email`.
	//
	// This field can be elided, and will marshal its zero value as "field/email".
	Type constant.FieldEmail `json:"type" default:"field/email"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldEmail) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldEmail](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a URL field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldUriURL struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool] `json:"unique,omitzero"`
	DefaultValues []URLValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/uri/url`.
	//
	// This field can be elided, and will marshal its zero value as "field/uri/url".
	Type constant.FieldUriURL `json:"type" default:"field/uri/url"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldUriURL) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldUriURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldUriURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldUriURL](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a domain field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldUriDomain struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]    `json:"unique,omitzero"`
	DefaultValues []DomainValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/uri/domain`.
	//
	// This field can be elided, and will marshal its zero value as "field/uri/domain".
	Type constant.FieldUriDomain `json:"type" default:"field/uri/domain"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldUriDomain) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldUriDomain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldUriDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldUriDomain](
		"cardinality", "one", "many",
	)
}

// Parameters for updating an X (formerly Twitter) profile field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldUriSocialX struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]     `json:"unique,omitzero"`
	DefaultValues []SocialXValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/uri/social_x`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/uri/social_x".
	Type constant.FieldUriSocialX `json:"type" default:"field/uri/social_x"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldUriSocialX) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldUriSocialX
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldUriSocialX) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldUriSocialX](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a LinkedIn profile field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldUriSocialLinkedIn struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]            `json:"unique,omitzero"`
	DefaultValues []SocialLinkedInValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/uri/social_linked_in`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/uri/social_linked_in".
	Type constant.FieldUriSocialLinkedIn `json:"type" default:"field/uri/social_linked_in"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldUriSocialLinkedIn) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldUriSocialLinkedIn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldUriSocialLinkedIn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldUriSocialLinkedIn](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a telephone number field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldTelephoneNumber struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]        `json:"unique,omitzero"`
	DefaultValues []TelephoneNumberParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/telephone_number`.
	//
	// This field can be elided, and will marshal its zero value as
	// "field/telephone_number".
	Type constant.FieldTelephoneNumber `json:"type" default:"field/telephone_number"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldTelephoneNumber) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldTelephoneNumber
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldTelephoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldTelephoneNumber](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a geographic location field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldGeo struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool] `json:"unique,omitzero"`
	DefaultValues []GeoValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/geo`.
	//
	// This field can be elided, and will marshal its zero value as "field/geo".
	Type constant.FieldGeo `json:"type" default:"field/geo"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldGeo) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldGeo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldGeo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldGeo](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a date field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldDate struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]                   `json:"unique,omitzero"`
	DefaultValues []DateFieldDefaultValueParamUnion `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/date`.
	//
	// This field can be elided, and will marshal its zero value as "field/date".
	Type constant.FieldDate `json:"type" default:"field/date"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldDate) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldDate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldDate](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a date and time field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldDatetime struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]                       `json:"unique,omitzero"`
	DefaultValues []DatetimeFieldDefaultValueParamUnion `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The field type. Must be `field/datetime`.
	//
	// This field can be elided, and will marshal its zero value as "field/datetime".
	Type constant.FieldDatetime `json:"type" default:"field/datetime"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldDatetime) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldDatetime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldDatetime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldDatetime](
		"cardinality", "one", "many",
	)
}

// Parameters for updating a choice field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldChoice struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]    `json:"unique,omitzero"`
	DefaultValues []ChoiceValueParam `json:"default_values,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The complete set of options for this field. Omit to leave unchanged. Array order
	// determines display order.
	Options []CollectionFieldUpdateParamsFieldFieldChoiceOption `json:"options,omitzero"`
	// The field type. Must be `field/choice`.
	//
	// This field can be elided, and will marshal its zero value as "field/choice".
	Type constant.FieldChoice `json:"type" default:"field/choice"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldChoice) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldChoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldChoice](
		"cardinality", "one", "many",
	)
}

// A choice field option. Items with an `id` update existing options; items without
// an `id` are added as new options.
//
// The properties Color, Name are required.
type CollectionFieldUpdateParamsFieldFieldChoiceOption struct {
	// The color of the option.
	//
	// Any of "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime",
	// "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet",
	// "yellow".
	Color string `json:"color,omitzero" api:"required"`
	// The display name of the option.
	Name string `json:"name" api:"required"`
	// The ID of an existing option to update. When absent, a new option is created.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldChoiceOption) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldChoiceOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldChoiceOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldChoiceOption](
		"color", "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime", "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet", "yellow",
	)
}

// Parameters for updating a relation field.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldRelation struct {
	// An updated description, or `null` to clear it.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the field.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, items must have a value for this field.
	Required param.Opt[bool] `json:"required,omitzero"`
	// If `true`, values must be unique across all items.
	Unique        param.Opt[bool]                       `json:"unique,omitzero"`
	DefaultValues []RelationFieldDefaultValueParamUnion `json:"default_values,omitzero"`
	// The complete set of allowed collections. Omit to leave unchanged. Array replaces
	// the current set.
	AllowedCollections []CollectionFieldUpdateParamsFieldFieldRelationAllowedCollection `json:"allowed_collections,omitzero"`
	// Updated cardinality: `one` or `many`.
	//
	// Any of "one", "many".
	Cardinality string `json:"cardinality,omitzero"`
	// The type of relationship: `one_way` for simple references, or `two_way` for
	// bidirectional relationships.
	//
	// Any of "one_way", "two_way".
	RelationType string `json:"relation_type,omitzero"`
	// The field type. Must be `field/relation`.
	//
	// This field can be elided, and will marshal its zero value as "field/relation".
	Type constant.FieldRelation `json:"type" default:"field/relation"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldRelation) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldRelation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldRelation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldRelation](
		"cardinality", "one", "many",
	)
	apijson.RegisterFieldValidator[CollectionFieldUpdateParamsFieldFieldRelation](
		"relation_type", "one_way", "two_way",
	)
}

// A reference to a `Collection` used in request bodies. Provide at least one of
// `id` or `ref` to identify the collection.
//
// The property Type is required.
type CollectionFieldUpdateParamsFieldFieldRelationAllowedCollection struct {
	// Unique identifier of the collection.
	ID param.Opt[string] `json:"id,omitzero"`
	// The stable, machine-readable reference identifier of the collection.
	Ref param.Opt[string] `json:"ref,omitzero"`
	// String representing the object’s type. Always `collection` for this object.
	//
	// This field can be elided, and will marshal its zero value as "collection".
	Type constant.Collection `json:"type" default:"collection"`
	paramObj
}

func (r CollectionFieldUpdateParamsFieldFieldRelationAllowedCollection) MarshalJSON() (data []byte, err error) {
	type shadow CollectionFieldUpdateParamsFieldFieldRelationAllowedCollection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionFieldUpdateParamsFieldFieldRelationAllowedCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionFieldDeleteParams struct {
	CollectionID string `path:"collection_id" api:"required" json:"-"`
	paramObj
}
