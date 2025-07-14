// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbasesdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

// CollectionService contains methods and other services that help with interacting
// with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectionService] method instead.
type CollectionService struct {
	Options []option.RequestOption
	Fields  CollectionFieldService
}

// NewCollectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCollectionService(opts ...option.RequestOption) (r CollectionService) {
	r = CollectionService{}
	r.Options = opts
	r.Fields = NewCollectionFieldService(opts...)
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

// A Collection is a container for structured data, similar to a database table or
// spreadsheet. It defines a schema using a set of `Fields` and holds the data as a
// list of `Items`.
type Collection struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// A hash of related links.
	Links CollectionLinks `json:"links,required"`
	// The user-facing name of the collection (e.g., “Organizations”).
	Name string `json:"name,required"`
	// A unique, stable, machine-readable identifier for the collection. This reference
	// is used in API requests and does not change even if the `name` is updated.
	Ref string `json:"ref,required"`
	// String representing the object’s type. Always `collection` for this object.
	Type constant.Collection `json:"type,required"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An optional, longer-form description of the collection's purpose.
	Description string `json:"description"`
	// A list of `Field` objects that define the schema for items in this collection.
	Fields []Field `json:"fields"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// A list of saved `View` objects for presenting the collection's data.
	Views []shared.View `json:"views"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		Name        respjson.Field
		Ref         respjson.Field
		Type        respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Fields      respjson.Field
		UpdatedAt   respjson.Field
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

// A hash of related links.
type CollectionLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionLinks) RawJSON() string { return r.JSON.raw }
func (r *CollectionLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Field defines a single column in a Collection's schema. It specifies the data
// type, validation rules, and other properties for the values stored in that
// column.
type Field struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Specifies whether the field can hold a single value (`one`) or multiple values
	// (`many`).
	//
	// Any of "one", "many".
	Cardinality FieldCardinality `json:"cardinality,required"`
	Links       FieldLinks       `json:"links,required"`
	// The human-readable name of the field (e.g., “Due Date”).
	Name string `json:"name,required"`
	// A unique, stable, machine-readable identifier for the field within its
	// collection (e.g., `due_date`).
	Ref string `json:"ref,required"`
	// The data type of the field. This determines how values are stored, validated,
	// and rendered.
	//
	// Any of "field/text/single_line", "field/text/multi_line",
	// "field/number/unitless_integer", "field/number/unitless_float",
	// "field/number/monetary", "field/number/percentage", "field/boolean",
	// "field/email", "field/uri/url", "field/uri/domain", "field/uri/social_x",
	// "field/uri/social_linked_in", "field/telephone_number", "field/geo",
	// "field/date", "field/datetime", "field/choice", "field/stage",
	// "field/relation/one_way", "field/relation/two_way".
	Type FieldType `json:"type,required"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An optional, longer-form description of the field's purpose.
	Description string `json:"description"`
	// If the field `type` is `field/choice`, this is the list of available
	// `FieldOption` objects.
	FieldOptions []FieldFieldOption `json:"field_options"`
	// If the field `type` is `field/stage`, this is the associated `Funnel` object.
	Funnel FieldFunnel `json:"funnel"`
	// If `true`, the value of this field is system-managed and cannot be updated via
	// the API.
	Readonly bool `json:"readonly"`
	// If `true`, this field must have a value.
	Required bool `json:"required"`
	// If `true`, values for this field must be unique across all items in the
	// collection.
	Unique bool `json:"unique"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Cardinality  respjson.Field
		Links        respjson.Field
		Name         respjson.Field
		Ref          respjson.Field
		Type         respjson.Field
		CreatedAt    respjson.Field
		Description  respjson.Field
		FieldOptions respjson.Field
		Funnel       respjson.Field
		Readonly     respjson.Field
		Required     respjson.Field
		Unique       respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Field) RawJSON() string { return r.JSON.raw }
func (r *Field) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies whether the field can hold a single value (`one`) or multiple values
// (`many`).
type FieldCardinality string

const (
	FieldCardinalityOne  FieldCardinality = "one"
	FieldCardinalityMany FieldCardinality = "many"
)

type FieldLinks struct {
	// The `Collection` this field belongs to.
	Collection string `json:"collection,required" format:"uri"`
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Collection  respjson.Field
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FieldLinks) RawJSON() string { return r.JSON.raw }
func (r *FieldLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The data type of the field. This determines how values are stored, validated,
// and rendered.
type FieldType string

const (
	FieldTypeFieldTextSingleLine        FieldType = "field/text/single_line"
	FieldTypeFieldTextMultiLine         FieldType = "field/text/multi_line"
	FieldTypeFieldNumberUnitlessInteger FieldType = "field/number/unitless_integer"
	FieldTypeFieldNumberUnitlessFloat   FieldType = "field/number/unitless_float"
	FieldTypeFieldNumberMonetary        FieldType = "field/number/monetary"
	FieldTypeFieldNumberPercentage      FieldType = "field/number/percentage"
	FieldTypeFieldBoolean               FieldType = "field/boolean"
	FieldTypeFieldEmail                 FieldType = "field/email"
	FieldTypeFieldUriURL                FieldType = "field/uri/url"
	FieldTypeFieldUriDomain             FieldType = "field/uri/domain"
	FieldTypeFieldUriSocialX            FieldType = "field/uri/social_x"
	FieldTypeFieldUriSocialLinkedIn     FieldType = "field/uri/social_linked_in"
	FieldTypeFieldTelephoneNumber       FieldType = "field/telephone_number"
	FieldTypeFieldGeo                   FieldType = "field/geo"
	FieldTypeFieldDate                  FieldType = "field/date"
	FieldTypeFieldDatetime              FieldType = "field/datetime"
	FieldTypeFieldChoice                FieldType = "field/choice"
	FieldTypeFieldStage                 FieldType = "field/stage"
	FieldTypeFieldRelationOneWay        FieldType = "field/relation/one_way"
	FieldTypeFieldRelationTwoWay        FieldType = "field/relation/two_way"
)

// Represents a selectable option within a `choice` field.
type FieldFieldOption struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The human-readable text for the option.
	Label string `json:"label,required"`
	// String representing the object’s type. Always `field_option` for this object.
	Type constant.FieldOption `json:"type,required"`
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
func (r FieldFieldOption) RawJSON() string { return r.JSON.raw }
func (r *FieldFieldOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If the field `type` is `field/stage`, this is the associated `Funnel` object.
type FieldFunnel struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// String representing the object’s type. Always `funnel` for this object.
	Type constant.Funnel `json:"type,required"`
	// An ordered list of `FunnelStep` objects that make up the funnel.
	Steps []FieldFunnelStep `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FieldFunnel) RawJSON() string { return r.JSON.raw }
func (r *FieldFunnel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a single step within a `Funnel`.
type FieldFunnelStep struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The name of the step.
	Name string `json:"name,required"`
	// The type of step, which can be `funnel_step/active`, `funnel_step/success`, or
	// `funnel_step/failure`.
	//
	// Any of "funnel_step/active", "funnel_step/success", "funnel_step/failure".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FieldFunnelStep) RawJSON() string { return r.JSON.raw }
func (r *FieldFunnelStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionGetParams struct {
	// Specifies which related objects to include in the response. Valid options are
	// `fields` and `views`.
	//
	// Any of "fields", "views".
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
