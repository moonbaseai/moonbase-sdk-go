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
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// Manage your collections and items
//
// ViewService contains methods and other services that help with interacting with
// the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewViewService] method instead.
type ViewService struct {
	Options []option.RequestOption
	// Manage your collections and items
	Items ViewItemService
}

// NewViewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewViewService(opts ...option.RequestOption) (r ViewService) {
	r = ViewService{}
	r.Options = opts
	r.Items = NewViewItemService(opts...)
	return
}

// Creates a new view in a collection.
func (r *ViewService) New(ctx context.Context, body ViewNewParams, opts ...option.RequestOption) (res *View, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "views"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the details of an existing view.
func (r *ViewService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *View, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("views/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a view. The change applies to the shared view that everyone in the
// workspace sees.
func (r *ViewService) Update(ctx context.Context, id string, body ViewUpdateParams, opts ...option.RequestOption) (res *View, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("views/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of views.
func (r *ViewService) List(ctx context.Context, query ViewListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ViewListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "views"
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

// Returns a list of views.
func (r *ViewService) ListAutoPaging(ctx context.Context, query ViewListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ViewListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a view. The default view of a collection cannot be deleted.
func (r *ViewService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("views/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A View represents a saved configuration for displaying items in a collection,
// including filters and sorting rules.
type View struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// The metrics computed over the view's items.
	Aggregates []ViewAggregateUnion `json:"aggregates" api:"required"`
	// The `Collection` this view belongs to.
	Collection CollectionPointer `json:"collection" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The view's columns, in display order.
	Fields []ViewField `json:"fields" api:"required"`
	// Return only items that match the filter conditions. Complex filters can be
	// created by nesting filters inside of `AND`, `OR`, and `NOT` filters.
	Filter ItemsFilterUnion `json:"filter" api:"required"`
	// Fields whose values group the view's items. Empty when the view is not grouped.
	Groups []string `json:"groups" api:"required"`
	// The name of the view.
	Name string `json:"name" api:"required"`
	// Filters limiting which related items the view's relation columns show.
	RelationValueFilters []ViewRelationValueFilter `json:"relation_value_filters" api:"required"`
	// Sort items returned by the specified fields. Empty when the view has no sort.
	Sort []string `json:"sort" api:"required"`
	// String representing the object’s type. Always `view` for this object.
	Type constant.View `json:"type" default:"view"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The type of view, such as `table` or `board`.
	//
	// Any of "table", "board".
	ViewType ViewViewType `json:"view_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Aggregates           respjson.Field
		Collection           respjson.Field
		CreatedAt            respjson.Field
		Fields               respjson.Field
		Filter               respjson.Field
		Groups               respjson.Field
		Name                 respjson.Field
		RelationValueFilters respjson.Field
		Sort                 respjson.Field
		Type                 respjson.Field
		UpdatedAt            respjson.Field
		ViewType             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r View) RawJSON() string { return r.JSON.raw }
func (r *View) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of view, such as `table` or `board`.
type ViewViewType string

const (
	ViewViewTypeTable ViewViewType = "table"
	ViewViewTypeBoard ViewViewType = "board"
)

// ViewAggregateUnion contains all possible properties and values from
// [ViewAggregateItemCount], [ViewAggregateFieldStatistic].
//
// Use the [ViewAggregateUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ViewAggregateUnion struct {
	// Any of "item_count", "field_statistic".
	Type  string `json:"type"`
	Group string `json:"group"`
	// This field is from variant [ViewAggregateFieldStatistic].
	Statistic ViewAggregateFieldStatisticStatistic `json:"statistic"`
	// This field is from variant [ViewAggregateFieldStatistic].
	Value string `json:"value"`
	// This field is from variant [ViewAggregateFieldStatistic].
	Weight string `json:"weight"`
	JSON   struct {
		Type      respjson.Field
		Group     respjson.Field
		Statistic respjson.Field
		Value     respjson.Field
		Weight    respjson.Field
		raw       string
	} `json:"-"`
}

// anyViewAggregate is implemented by each variant of [ViewAggregateUnion] to add
// type safety for the return type of [ViewAggregateUnion.AsAny]
type anyViewAggregate interface {
	implViewAggregateUnion()
}

func (ViewAggregateItemCount) implViewAggregateUnion()      {}
func (ViewAggregateFieldStatistic) implViewAggregateUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ViewAggregateUnion.AsAny().(type) {
//	case moonbase.ViewAggregateItemCount:
//	case moonbase.ViewAggregateFieldStatistic:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ViewAggregateUnion) AsAny() anyViewAggregate {
	switch u.Type {
	case "item_count":
		return u.AsItemCount()
	case "field_statistic":
		return u.AsFieldStatistic()
	}
	return nil
}

func (u ViewAggregateUnion) AsItemCount() (v ViewAggregateItemCount) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ViewAggregateUnion) AsFieldStatistic() (v ViewAggregateFieldStatistic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ViewAggregateUnion) RawJSON() string { return u.JSON.raw }

func (r *ViewAggregateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewAggregateUnion to a ViewAggregateUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewAggregateUnionParam.Overrides()
func (r ViewAggregateUnion) ToParam() ViewAggregateUnionParam {
	return param.Override[ViewAggregateUnionParam](json.RawMessage(r.RawJSON()))
}

func ViewAggregateParamOfFieldStatistic(statistic ViewAggregateFieldStatisticStatistic, value string) ViewAggregateUnionParam {
	var fieldStatistic ViewAggregateFieldStatisticParam
	fieldStatistic.Statistic = statistic
	fieldStatistic.Value = value
	return ViewAggregateUnionParam{OfFieldStatistic: &fieldStatistic}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ViewAggregateUnionParam struct {
	OfItemCount      *ViewAggregateItemCountParam      `json:",omitzero,inline"`
	OfFieldStatistic *ViewAggregateFieldStatisticParam `json:",omitzero,inline"`
	paramUnion
}

func (u ViewAggregateUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfItemCount, u.OfFieldStatistic)
}
func (u *ViewAggregateUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ViewAggregateUnionParam) asAny() any {
	if !param.IsOmitted(u.OfItemCount) {
		return u.OfItemCount
	} else if !param.IsOmitted(u.OfFieldStatistic) {
		return u.OfFieldStatistic
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ViewAggregateUnionParam) GetStatistic() *string {
	if vt := u.OfFieldStatistic; vt != nil {
		return (*string)(&vt.Statistic)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ViewAggregateUnionParam) GetValue() *string {
	if vt := u.OfFieldStatistic; vt != nil {
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ViewAggregateUnionParam) GetWeight() *string {
	if vt := u.OfFieldStatistic; vt != nil && vt.Weight.Valid() {
		return &vt.Weight.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ViewAggregateUnionParam) GetType() *string {
	if vt := u.OfItemCount; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFieldStatistic; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ViewAggregateUnionParam) GetGroup() *string {
	if vt := u.OfItemCount; vt != nil && vt.Group.Valid() {
		return &vt.Group.Value
	} else if vt := u.OfFieldStatistic; vt != nil && vt.Group.Valid() {
		return &vt.Group.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ViewAggregateUnionParam](
		"type",
		apijson.Discriminator[ViewAggregateItemCountParam]("item_count"),
		apijson.Discriminator[ViewAggregateFieldStatisticParam]("field_statistic"),
	)
}

// Computes a statistic over the values of a field.
type ViewAggregateFieldStatistic struct {
	// The statistic to compute. Scalar statistics (sum, mean, max, min) require a
	// number field as `value`.
	//
	// Any of "count", "sum", "mean", "max", "min", "filled_percentage".
	Statistic ViewAggregateFieldStatisticStatistic `json:"statistic" api:"required"`
	Type      constant.FieldStatistic              `json:"type" default:"field_statistic"`
	// The field whose values the statistic is computed over.
	Value string `json:"value" api:"required"`
	// An optional field whose values bucket the statistic.
	Group string `json:"group"`
	// An optional percentage field used to weight the statistic. Only supported for
	// scalar statistics.
	Weight string `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Statistic   respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		Group       respjson.Field
		Weight      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewAggregateFieldStatistic) RawJSON() string { return r.JSON.raw }
func (r *ViewAggregateFieldStatistic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewAggregateFieldStatistic to a
// ViewAggregateFieldStatisticParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewAggregateFieldStatisticParam.Overrides()
func (r ViewAggregateFieldStatistic) ToParam() ViewAggregateFieldStatisticParam {
	return param.Override[ViewAggregateFieldStatisticParam](json.RawMessage(r.RawJSON()))
}

// The statistic to compute. Scalar statistics (sum, mean, max, min) require a
// number field as `value`.
type ViewAggregateFieldStatisticStatistic string

const (
	ViewAggregateFieldStatisticStatisticCount            ViewAggregateFieldStatisticStatistic = "count"
	ViewAggregateFieldStatisticStatisticSum              ViewAggregateFieldStatisticStatistic = "sum"
	ViewAggregateFieldStatisticStatisticMean             ViewAggregateFieldStatisticStatistic = "mean"
	ViewAggregateFieldStatisticStatisticMax              ViewAggregateFieldStatisticStatistic = "max"
	ViewAggregateFieldStatisticStatisticMin              ViewAggregateFieldStatisticStatistic = "min"
	ViewAggregateFieldStatisticStatisticFilledPercentage ViewAggregateFieldStatisticStatistic = "filled_percentage"
)

// Computes a statistic over the values of a field.
//
// The properties Statistic, Type, Value are required.
type ViewAggregateFieldStatisticParam struct {
	// The statistic to compute. Scalar statistics (sum, mean, max, min) require a
	// number field as `value`.
	//
	// Any of "count", "sum", "mean", "max", "min", "filled_percentage".
	Statistic ViewAggregateFieldStatisticStatistic `json:"statistic,omitzero" api:"required"`
	// The field whose values the statistic is computed over.
	Value string `json:"value" api:"required"`
	// An optional field whose values bucket the statistic.
	Group param.Opt[string] `json:"group,omitzero"`
	// An optional percentage field used to weight the statistic. Only supported for
	// scalar statistics.
	Weight param.Opt[string] `json:"weight,omitzero"`
	// This field can be elided, and will marshal its zero value as "field_statistic".
	Type constant.FieldStatistic `json:"type" default:"field_statistic"`
	paramObj
}

func (r ViewAggregateFieldStatisticParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewAggregateFieldStatisticParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewAggregateFieldStatisticParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Counts the view's items.
type ViewAggregateItemCount struct {
	Type constant.ItemCount `json:"type" default:"item_count"`
	// An optional field whose values bucket the counts.
	Group string `json:"group"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Group       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewAggregateItemCount) RawJSON() string { return r.JSON.raw }
func (r *ViewAggregateItemCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewAggregateItemCount to a ViewAggregateItemCountParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewAggregateItemCountParam.Overrides()
func (r ViewAggregateItemCount) ToParam() ViewAggregateItemCountParam {
	return param.Override[ViewAggregateItemCountParam](json.RawMessage(r.RawJSON()))
}

// Counts the view's items.
//
// The property Type is required.
type ViewAggregateItemCountParam struct {
	// An optional field whose values bucket the counts.
	Group param.Opt[string] `json:"group,omitzero"`
	// This field can be elided, and will marshal its zero value as "item_count".
	Type constant.ItemCount `json:"type" default:"item_count"`
	paramObj
}

func (r ViewAggregateItemCountParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewAggregateItemCountParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewAggregateItemCountParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A column of the view.
type ViewField struct {
	// The field shown in this column.
	Field string `json:"field" api:"required"`
	// Which fields of the related item to show, relative to the related collection.
	// Omitted means the related collection's default display fields.
	DisplayFields []string `json:"display_fields"`
	// Whether the column is pinned.
	IsPinned bool `json:"is_pinned"`
	// Whether the column wraps its content.
	IsWrapped bool `json:"is_wrapped"`
	// The column width: a number of pixels, or `fit` (size to content), or `flex`
	// (fill available space).
	Size ViewFieldSizeUnion `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field         respjson.Field
		DisplayFields respjson.Field
		IsPinned      respjson.Field
		IsWrapped     respjson.Field
		Size          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewField) RawJSON() string { return r.JSON.raw }
func (r *ViewField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewField to a ViewFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewFieldParam.Overrides()
func (r ViewField) ToParam() ViewFieldParam {
	return param.Override[ViewFieldParam](json.RawMessage(r.RawJSON()))
}

// ViewFieldSizeUnion contains all possible properties and values from [float64],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfViewFieldSizeString]
type ViewFieldSizeUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfViewFieldSizeString string `json:",inline"`
	JSON                  struct {
		OfFloat               respjson.Field
		OfViewFieldSizeString respjson.Field
		raw                   string
	} `json:"-"`
}

func (u ViewFieldSizeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ViewFieldSizeUnion) AsViewFieldSizeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ViewFieldSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *ViewFieldSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewFieldSizeString string

const (
	ViewFieldSizeStringFit  ViewFieldSizeString = "fit"
	ViewFieldSizeStringFlex ViewFieldSizeString = "flex"
)

// A column of the view.
//
// The property Field is required.
type ViewFieldParam struct {
	// The field shown in this column.
	Field string `json:"field" api:"required"`
	// Whether the column is pinned.
	IsPinned param.Opt[bool] `json:"is_pinned,omitzero"`
	// Whether the column wraps its content.
	IsWrapped param.Opt[bool] `json:"is_wrapped,omitzero"`
	// Which fields of the related item to show, relative to the related collection.
	// Omitted means the related collection's default display fields.
	DisplayFields []string `json:"display_fields,omitzero"`
	// The column width: a number of pixels, or `fit` (size to content), or `flex`
	// (fill available space).
	Size ViewFieldSizeUnionParam `json:"size,omitzero"`
	paramObj
}

func (r ViewFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ViewFieldSizeUnionParam struct {
	OfFloat param.Opt[float64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfViewFieldSizeString)
	OfViewFieldSizeString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ViewFieldSizeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfViewFieldSizeString)
}
func (u *ViewFieldSizeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ViewFieldSizeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfViewFieldSizeString) {
		return &u.OfViewFieldSizeString
	}
	return nil
}

// Limits which related items a relation column shows: only related items matching
// `filter` appear.
type ViewRelationValueFilter struct {
	// The relation column whose related items are filtered.
	Field string `json:"field" api:"required"`
	// The filter the related items must match. Field paths are relative to the related
	// collection.
	Filter ItemsFilterUnion `json:"filter" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewRelationValueFilter) RawJSON() string { return r.JSON.raw }
func (r *ViewRelationValueFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewRelationValueFilter to a ViewRelationValueFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewRelationValueFilterParam.Overrides()
func (r ViewRelationValueFilter) ToParam() ViewRelationValueFilterParam {
	return param.Override[ViewRelationValueFilterParam](json.RawMessage(r.RawJSON()))
}

// Limits which related items a relation column shows: only related items matching
// `filter` appear.
//
// The properties Field, Filter are required.
type ViewRelationValueFilterParam struct {
	// The relation column whose related items are filtered.
	Field string `json:"field" api:"required"`
	// The filter the related items must match. Field paths are relative to the related
	// collection.
	Filter ItemsFilterUnionParam `json:"filter,omitzero" api:"required"`
	paramObj
}

func (r ViewRelationValueFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewRelationValueFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewRelationValueFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewListResponse struct {
	ID string `json:"id" api:"required"`
	// A lightweight reference to a `Collection`, containing the minimal information
	// needed to identify it.
	Collection CollectionPointer `json:"collection" api:"required"`
	CreatedAt  time.Time         `json:"created_at" api:"required" format:"date-time"`
	Name       string            `json:"name" api:"required"`
	Type       constant.View     `json:"type" default:"view"`
	UpdatedAt  time.Time         `json:"updated_at" api:"required" format:"date-time"`
	// Any of "table", "board".
	ViewType ViewListResponseViewType `json:"view_type" api:"required"`
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
func (r ViewListResponse) RawJSON() string { return r.JSON.raw }
func (r *ViewListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewListResponseViewType string

const (
	ViewListResponseViewTypeTable ViewListResponseViewType = "table"
	ViewListResponseViewTypeBoard ViewListResponseViewType = "board"
)

type ViewNewParams struct {
	// A pointer to the `Collection` the view belongs to.
	Collection ViewNewParamsCollection `json:"collection,omitzero" api:"required"`
	// The view's columns, in display order.
	Fields []ViewFieldParam `json:"fields,omitzero" api:"required"`
	// The name of the view.
	Name string `json:"name" api:"required"`
	// The type of view, `table` or `board`.
	//
	// Any of "table", "board".
	ViewType ViewNewParamsViewType `json:"view_type,omitzero" api:"required"`
	// The metrics computed over the view's items.
	Aggregates []ViewAggregateUnionParam `json:"aggregates,omitzero"`
	// The filter applied to the view's items.
	Filter ItemsFilterUnionParam `json:"filter,omitzero"`
	// Fields whose values group the view's items.
	Groups []string `json:"groups,omitzero"`
	// Filters limiting which related items the view's relation columns show.
	RelationValueFilters []ViewRelationValueFilterParam `json:"relation_value_filters,omitzero"`
	// Sort items returned by the specified fields.
	Sort []string `json:"sort,omitzero"`
	paramObj
}

func (r ViewNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ViewNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A pointer to the `Collection` the view belongs to.
//
// The property Type is required.
type ViewNewParamsCollection struct {
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

func (r ViewNewParamsCollection) MarshalJSON() (data []byte, err error) {
	type shadow ViewNewParamsCollection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewNewParamsCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of view, `table` or `board`.
type ViewNewParamsViewType string

const (
	ViewNewParamsViewTypeTable ViewNewParamsViewType = "table"
	ViewNewParamsViewTypeBoard ViewNewParamsViewType = "board"
)

type ViewUpdateParams struct {
	// The name of the view.
	Name param.Opt[string] `json:"name,omitzero"`
	// The metrics computed over the view's items. An empty array clears them.
	Aggregates []ViewAggregateUnionParam `json:"aggregates,omitzero"`
	// The view's columns, in display order. If provided, it must contain at least one
	// column.
	Fields []ViewFieldParam `json:"fields,omitzero"`
	// Return only items that match the filter conditions. Complex filters can be
	// created by nesting filters inside of `AND`, `OR`, and `NOT` filters.
	Filter ItemsFilterUnionParam `json:"filter,omitzero"`
	// Fields whose values group the view's items. An empty array clears the grouping.
	Groups []string `json:"groups,omitzero"`
	// Filters limiting which related items the view's relation columns show. An empty
	// array clears them.
	RelationValueFilters []ViewRelationValueFilterParam `json:"relation_value_filters,omitzero"`
	// Sort items returned by the specified fields. An empty array clears the sort.
	Sort []string `json:"sort,omitzero"`
	// The type of view, `table` or `board`.
	//
	// Any of "table", "board".
	ViewType ViewUpdateParamsViewType `json:"view_type,omitzero"`
	paramObj
}

func (r ViewUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ViewUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of view, `table` or `board`.
type ViewUpdateParamsViewType string

const (
	ViewUpdateParamsViewTypeTable ViewUpdateParamsViewType = "table"
	ViewUpdateParamsViewTypeBoard ViewUpdateParamsViewType = "board"
)

type ViewListParams struct {
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

// URLQuery serializes [ViewListParams]'s query parameters as `url.Values`.
func (r ViewListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
