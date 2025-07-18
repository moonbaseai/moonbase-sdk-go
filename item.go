// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/paramutil"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// ItemService contains methods and other services that help with interacting with
// the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewItemService] method instead.
type ItemService struct {
	Options []option.RequestOption
}

// NewItemService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewItemService(opts ...option.RequestOption) (r ItemService) {
	r = ItemService{}
	r.Options = opts
	return
}

// Creates a new item in a collection.
func (r *ItemService) New(ctx context.Context, body ItemNewParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = append(r.Options[:], opts...)
	path := "items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing item.
func (r *ItemService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Item, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("items/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an item.
func (r *ItemService) Update(ctx context.Context, id string, params ItemUpdateParams, opts ...option.RequestOption) (res *Item, err error) {
	if !param.IsOmitted(params.UpdateManyStrategy) {
		opts = append(opts, option.WithHeader("update-many-strategy", fmt.Sprintf("%s", params.UpdateManyStrategy)))
	}
	if !param.IsOmitted(params.UpdateOneStrategy) {
		opts = append(opts, option.WithHeader("update-one-strategy", fmt.Sprintf("%s", params.UpdateOneStrategy)))
	}
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("items/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Permanently deletes an item.
func (r *ItemService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *Item, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("items/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Find and update an existing item, or create a new one.
func (r *ItemService) Upsert(ctx context.Context, params ItemUpsertParams, opts ...option.RequestOption) (res *Item, err error) {
	if !param.IsOmitted(params.UpdateManyStrategy) {
		opts = append(opts, option.WithHeader("update-many-strategy", fmt.Sprintf("%s", params.UpdateManyStrategy)))
	}
	if !param.IsOmitted(params.UpdateOneStrategy) {
		opts = append(opts, option.WithHeader("update-one-strategy", fmt.Sprintf("%s", params.UpdateOneStrategy)))
	}
	opts = append(r.Options[:], opts...)
	path := "items/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// True or false value
type BooleanValue struct {
	Boolean bool                  `json:"boolean,required"`
	Type    constant.ValueBoolean `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Boolean     respjson.Field
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
// The properties Boolean, Type are required.
type BooleanValueParam struct {
	Boolean bool `json:"boolean,required"`
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

// Selected choice option
type Choice struct {
	Option ChoiceOption         `json:"option,required"`
	Type   constant.ValueChoice `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Option      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Choice) RawJSON() string { return r.JSON.raw }
func (r *Choice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Choice to a ChoiceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChoiceParam.Overrides()
func (r Choice) ToParam() ChoiceParam {
	return param.Override[ChoiceParam](json.RawMessage(r.RawJSON()))
}

type ChoiceOption struct {
	ID    string               `json:"id,required"`
	Type  constant.FieldOption `json:"type,required"`
	Label string               `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChoiceOption) RawJSON() string { return r.JSON.raw }
func (r *ChoiceOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Selected choice option
//
// The properties Option, Type are required.
type ChoiceParam struct {
	Option ChoiceOptionParam `json:"option,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "value/choice".
	Type constant.ValueChoice `json:"type,required"`
	paramObj
}

func (r ChoiceParam) MarshalJSON() (data []byte, err error) {
	type shadow ChoiceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChoiceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Type are required.
type ChoiceOptionParam struct {
	ID    string            `json:"id,required"`
	Label param.Opt[string] `json:"label,omitzero"`
	// This field can be elided, and will marshal its zero value as "field_option".
	Type constant.FieldOption `json:"type,required"`
	paramObj
}

func (r ChoiceOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow ChoiceOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChoiceOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Date without time
type DateValue struct {
	Date time.Time          `json:"date,required" format:"date"`
	Type constant.ValueDate `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
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
// The properties Date, Type are required.
type DateValueParam struct {
	Date time.Time `json:"date,required" format:"date"`
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

// Date and time value
type DatetimeValue struct {
	Datetime time.Time              `json:"datetime,required" format:"date-time"`
	Type     constant.ValueDatetime `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Datetime    respjson.Field
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
// The properties Datetime, Type are required.
type DatetimeValueParam struct {
	Datetime time.Time `json:"datetime,required" format:"date-time"`
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

// Internet domain name
type DomainValue struct {
	Domain string                  `json:"domain,required"`
	Type   constant.ValueUriDomain `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
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
// The properties Domain, Type are required.
type DomainValueParam struct {
	Domain string `json:"domain,required"`
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

// Email address value
type EmailValue struct {
	Email string              `json:"email,required" format:"email"`
	Type  constant.ValueEmail `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
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
// The properties Email, Type are required.
type EmailValueParam struct {
	Email string `json:"email,required" format:"email"`
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

// FieldValueUnion contains all possible properties and values from
// [SingleLineTextValue], [MultiLineTextValue], [IntegerValue], [FloatValue],
// [MonetaryValue], [PercentageValue], [BooleanValue], [EmailValue], [URLValue],
// [DomainValue], [SocialXValue], [SocialLinkedInValue], [TelephoneNumber],
// [GeoValue], [DateValue], [DatetimeValue], [Choice], [FunnelStep],
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
	Text            string       `json:"text"`
	Type            string       `json:"type"`
	// This field is a union of [int64], [float64]
	Number FieldValueUnionNumber `json:"number"`
	// This field is from variant [MonetaryValue].
	Amount MonetaryValueAmount `json:"amount"`
	// This field is from variant [PercentageValue].
	Percentage float64 `json:"percentage"`
	// This field is from variant [BooleanValue].
	Boolean bool `json:"boolean"`
	// This field is from variant [EmailValue].
	Email string `json:"email"`
	// This field is from variant [URLValue].
	URL string `json:"url"`
	// This field is from variant [DomainValue].
	Domain string `json:"domain"`
	// This field is a union of [SocialXValueProfile], [SocialLinkedInValueProfile]
	Profile FieldValueUnionProfile `json:"profile"`
	// This field is from variant [TelephoneNumber].
	TelephoneNumber string `json:"telephone_number"`
	// This field is from variant [GeoValue].
	Geo string `json:"geo"`
	// This field is from variant [DateValue].
	Date time.Time `json:"date"`
	// This field is from variant [DatetimeValue].
	Datetime time.Time `json:"datetime"`
	// This field is from variant [Choice].
	Option ChoiceOption `json:"option"`
	// This field is from variant [FunnelStep].
	Step FunnelStepStep `json:"step"`
	// This field is from variant [RelationValue].
	Item Item `json:"item"`
	JSON struct {
		OfArrayOfValues respjson.Field
		Text            respjson.Field
		Type            respjson.Field
		Number          respjson.Field
		Amount          respjson.Field
		Percentage      respjson.Field
		Boolean         respjson.Field
		Email           respjson.Field
		URL             respjson.Field
		Domain          respjson.Field
		Profile         respjson.Field
		TelephoneNumber respjson.Field
		Geo             respjson.Field
		Date            respjson.Field
		Datetime        respjson.Field
		Option          respjson.Field
		Step            respjson.Field
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

func (u FieldValueUnion) AsChoice() (v Choice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FieldValueUnion) AsFunnelStep() (v FunnelStep) {
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

// FieldValueUnionNumber is an implicit subunion of [FieldValueUnion].
// FieldValueUnionNumber provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [FieldValueUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfFloat]
type FieldValueUnionNumber struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfInt   respjson.Field
		OfFloat respjson.Field
		raw     string
	} `json:"-"`
}

func (r *FieldValueUnionNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FieldValueUnionProfile is an implicit subunion of [FieldValueUnion].
// FieldValueUnionProfile provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [FieldValueUnion].
type FieldValueUnionProfile struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	JSON     struct {
		URL      respjson.Field
		Username respjson.Field
		raw      string
	} `json:"-"`
}

func (r *FieldValueUnionProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FieldValueUnion to a FieldValueUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FieldValueUnionParam.Overrides()
func (r FieldValueUnion) ToParam() FieldValueUnionParam {
	return param.Override[FieldValueUnionParam](json.RawMessage(r.RawJSON()))
}

func FieldValueParamOfSingleLineText(text string) FieldValueUnionParam {
	var variant SingleLineTextValueParam
	variant.Text = text
	return FieldValueUnionParam{OfSingleLineText: &variant}
}

func FieldValueParamOfMultiLineText(text string) FieldValueUnionParam {
	var variant MultiLineTextValueParam
	variant.Text = text
	return FieldValueUnionParam{OfMultiLineText: &variant}
}

func FieldValueParamOfInteger(number int64) FieldValueUnionParam {
	var variant IntegerValueParam
	variant.Number = number
	return FieldValueUnionParam{OfInteger: &variant}
}

func FieldValueParamOfFloat(number float64) FieldValueUnionParam {
	var variant FloatValueParam
	variant.Number = number
	return FieldValueUnionParam{OfFloat: &variant}
}

func FieldValueParamOfMonetary(amount MonetaryValueAmountParam) FieldValueUnionParam {
	var variant MonetaryValueParam
	variant.Amount = amount
	return FieldValueUnionParam{OfMonetary: &variant}
}

func FieldValueParamOfPercentage(percentage float64) FieldValueUnionParam {
	var variant PercentageValueParam
	variant.Percentage = percentage
	return FieldValueUnionParam{OfPercentage: &variant}
}

func FieldValueParamOfBoolean(boolean bool) FieldValueUnionParam {
	var variant BooleanValueParam
	variant.Boolean = boolean
	return FieldValueUnionParam{OfBoolean: &variant}
}

func FieldValueParamOfEmail(email string) FieldValueUnionParam {
	var variant EmailValueParam
	variant.Email = email
	return FieldValueUnionParam{OfEmail: &variant}
}

func FieldValueParamOfURL(url string) FieldValueUnionParam {
	var variant URLValueParam
	variant.URL = url
	return FieldValueUnionParam{OfURL: &variant}
}

func FieldValueParamOfDomain(domain string) FieldValueUnionParam {
	var variant DomainValueParam
	variant.Domain = domain
	return FieldValueUnionParam{OfDomain: &variant}
}

func FieldValueParamOfX(profile SocialXValueProfileParam) FieldValueUnionParam {
	var variant SocialXValueParam
	variant.Profile = profile
	return FieldValueUnionParam{OfX: &variant}
}

func FieldValueParamOfLinkedIn(profile SocialLinkedInValueProfileParam) FieldValueUnionParam {
	var variant SocialLinkedInValueParam
	variant.Profile = profile
	return FieldValueUnionParam{OfLinkedIn: &variant}
}

func FieldValueParamOfTelephoneNumber(telephoneNumber string) FieldValueUnionParam {
	var variant TelephoneNumberParam
	variant.TelephoneNumber = telephoneNumber
	return FieldValueUnionParam{OfTelephoneNumber: &variant}
}

func FieldValueParamOfGeo(geo string) FieldValueUnionParam {
	var variant GeoValueParam
	variant.Geo = geo
	return FieldValueUnionParam{OfGeo: &variant}
}

func FieldValueParamOfDate(date time.Time) FieldValueUnionParam {
	var variant DateValueParam
	variant.Date = date
	return FieldValueUnionParam{OfDate: &variant}
}

func FieldValueParamOfDateTime(datetime time.Time) FieldValueUnionParam {
	var variant DatetimeValueParam
	variant.Datetime = datetime
	return FieldValueUnionParam{OfDateTime: &variant}
}

func FieldValueParamOfChoice(option ChoiceOptionParam) FieldValueUnionParam {
	var variant ChoiceParam
	variant.Option = option
	return FieldValueUnionParam{OfChoice: &variant}
}

func FieldValueParamOfFunnelStep(step FunnelStepStepParam) FieldValueUnionParam {
	var variant FunnelStepParam
	variant.Step = step
	return FieldValueUnionParam{OfFunnelStep: &variant}
}

func FieldValueParamOfRelation(item ItemParam) FieldValueUnionParam {
	var variant RelationValueParam
	variant.Item = item
	return FieldValueUnionParam{OfRelation: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FieldValueUnionParam struct {
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
	OfX               *SocialXValueParam        `json:",omitzero,inline"`
	OfLinkedIn        *SocialLinkedInValueParam `json:",omitzero,inline"`
	OfTelephoneNumber *TelephoneNumberParam     `json:",omitzero,inline"`
	OfGeo             *GeoValueParam            `json:",omitzero,inline"`
	OfDate            *DateValueParam           `json:",omitzero,inline"`
	OfDateTime        *DatetimeValueParam       `json:",omitzero,inline"`
	OfChoice          *ChoiceParam              `json:",omitzero,inline"`
	OfFunnelStep      *FunnelStepParam          `json:",omitzero,inline"`
	OfRelation        *RelationValueParam       `json:",omitzero,inline"`
	OfArrayOfValues   []ValueUnionParam         `json:",omitzero,inline"`
	paramUnion
}

func (u FieldValueUnionParam) MarshalJSON() ([]byte, error) {
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
func (u *FieldValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FieldValueUnionParam) asAny() any {
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
func (u FieldValueUnionParam) GetAmount() *MonetaryValueAmountParam {
	if vt := u.OfMonetary; vt != nil {
		return &vt.Amount
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetPercentage() *float64 {
	if vt := u.OfPercentage; vt != nil {
		return &vt.Percentage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetBoolean() *bool {
	if vt := u.OfBoolean; vt != nil {
		return &vt.Boolean
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetEmail() *string {
	if vt := u.OfEmail; vt != nil {
		return &vt.Email
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetURL() *string {
	if vt := u.OfURL; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetDomain() *string {
	if vt := u.OfDomain; vt != nil {
		return &vt.Domain
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetTelephoneNumber() *string {
	if vt := u.OfTelephoneNumber; vt != nil {
		return &vt.TelephoneNumber
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetGeo() *string {
	if vt := u.OfGeo; vt != nil {
		return &vt.Geo
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetDate() *time.Time {
	if vt := u.OfDate; vt != nil {
		return &vt.Date
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetDatetime() *time.Time {
	if vt := u.OfDateTime; vt != nil {
		return &vt.Datetime
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetOption() *ChoiceOptionParam {
	if vt := u.OfChoice; vt != nil {
		return &vt.Option
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetStep() *FunnelStepStepParam {
	if vt := u.OfFunnelStep; vt != nil {
		return &vt.Step
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetItem() *ItemParam {
	if vt := u.OfRelation; vt != nil {
		return &vt.Item
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetText() *string {
	if vt := u.OfSingleLineText; vt != nil {
		return (*string)(&vt.Text)
	} else if vt := u.OfMultiLineText; vt != nil {
		return (*string)(&vt.Text)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FieldValueUnionParam) GetType() *string {
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
func (u FieldValueUnionParam) GetNumber() (res fieldValueUnionParamNumber) {
	if vt := u.OfInteger; vt != nil {
		res.any = &vt.Number
	} else if vt := u.OfFloat; vt != nil {
		res.any = &vt.Number
	}
	return
}

// Can have the runtime types [*int64], [*float64]
type fieldValueUnionParamNumber struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *int64:
//	case *float64:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u fieldValueUnionParamNumber) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FieldValueUnionParam) GetProfile() (res fieldValueUnionParamProfile) {
	if vt := u.OfX; vt != nil {
		res.any = &vt.Profile
	} else if vt := u.OfLinkedIn; vt != nil {
		res.any = &vt.Profile
	}
	return
}

// Can have the runtime types [*SocialXValueProfileParam],
// [*SocialLinkedInValueProfileParam]
type fieldValueUnionParamProfile struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *moonbase.SocialXValueProfileParam:
//	case *moonbase.SocialLinkedInValueProfileParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u fieldValueUnionParamProfile) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueUnionParamProfile) GetURL() *string {
	switch vt := u.any.(type) {
	case *SocialXValueProfileParam:
		return paramutil.AddrIfPresent(vt.URL)
	case *SocialLinkedInValueProfileParam:
		return paramutil.AddrIfPresent(vt.URL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u fieldValueUnionParamProfile) GetUsername() *string {
	switch vt := u.any.(type) {
	case *SocialXValueProfileParam:
		return paramutil.AddrIfPresent(vt.Username)
	case *SocialLinkedInValueProfileParam:
		return paramutil.AddrIfPresent(vt.Username)
	}
	return nil
}

// Floating point number
type FloatValue struct {
	Number float64                           `json:"number,required"`
	Type   constant.ValueNumberUnitlessFloat `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
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
// The properties Number, Type are required.
type FloatValueParam struct {
	Number float64 `json:"number,required"`
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
type FunnelStep struct {
	Step FunnelStepStep           `json:"step,required"`
	Type constant.ValueFunnelStep `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Step        respjson.Field
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

type FunnelStepStep struct {
	ID   string              `json:"id,required"`
	Type constant.FunnelStep `json:"type,required"`
	Name string              `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelStepStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelStepStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Funnel step value
//
// The properties Step, Type are required.
type FunnelStepParam struct {
	Step FunnelStepStepParam `json:"step,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/funnel_step".
	Type constant.ValueFunnelStep `json:"type,required"`
	paramObj
}

func (r FunnelStepParam) MarshalJSON() (data []byte, err error) {
	type shadow FunnelStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Type are required.
type FunnelStepStepParam struct {
	ID   string            `json:"id,required"`
	Name param.Opt[string] `json:"name,omitzero"`
	// This field can be elided, and will marshal its zero value as "funnel_step".
	Type constant.FunnelStep `json:"type,required"`
	paramObj
}

func (r FunnelStepStepParam) MarshalJSON() (data []byte, err error) {
	type shadow FunnelStepStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelStepStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Geographic coordinate value
type GeoValue struct {
	Geo  string            `json:"geo,required"`
	Type constant.ValueGeo `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geo         respjson.Field
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
// The properties Geo, Type are required.
type GeoValueParam struct {
	Geo string `json:"geo,required"`
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

// Integer value without units
type IntegerValue struct {
	Number int64                               `json:"number,required"`
	Type   constant.ValueNumberUnitlessInteger `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
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
// The properties Number, Type are required.
type IntegerValueParam struct {
	Number int64 `json:"number,required"`
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
	Type  constant.Item `json:"type,required"`
	Links ItemLinks     `json:"links"`
	// A hash where keys are the `ref` of a `Field` and values are the data stored for
	// that field.
	Values map[string]FieldValueUnion `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Links       respjson.Field
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

// ToParam converts this Item to a ItemParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ItemParam.Overrides()
func (r Item) ToParam() ItemParam {
	return param.Override[ItemParam](json.RawMessage(r.RawJSON()))
}

type ItemLinks struct {
	// A link to the `Collection` the item belongs to.
	Collection string `json:"collection" format:"uri"`
	// The canonical URL for this object.
	Self string `json:"self" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Collection  respjson.Field
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemLinks) RawJSON() string { return r.JSON.raw }
func (r *ItemLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Item represents a single record or row within a Collection. It holds a set of
// `values` corresponding to the Collection's `fields`.
//
// The properties ID, Type are required.
type ItemParam struct {
	// Unique identifier for the object.
	ID    string         `json:"id,required"`
	Links ItemLinksParam `json:"links,omitzero"`
	// A hash where keys are the `ref` of a `Field` and values are the data stored for
	// that field.
	Values map[string]FieldValueUnionParam `json:"values,omitzero"`
	// String representing the object’s type. Always `item` for this object.
	//
	// This field can be elided, and will marshal its zero value as "item".
	Type constant.Item `json:"type,required"`
	paramObj
}

func (r ItemParam) MarshalJSON() (data []byte, err error) {
	type shadow ItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ItemLinksParam struct {
	// A link to the `Collection` the item belongs to.
	Collection param.Opt[string] `json:"collection,omitzero" format:"uri"`
	// The canonical URL for this object.
	Self param.Opt[string] `json:"self,omitzero" format:"uri"`
	paramObj
}

func (r ItemLinksParam) MarshalJSON() (data []byte, err error) {
	type shadow ItemLinksParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemLinksParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monetary or currency value
type MonetaryValue struct {
	Amount MonetaryValueAmount          `json:"amount,required"`
	Type   constant.ValueNumberMonetary `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
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

type MonetaryValueAmount struct {
	AmountInMinorUnits int64             `json:"amount_in_minor_units,required"`
	Currency           string            `json:"currency,required"`
	Type               constant.Currency `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountInMinorUnits respjson.Field
		Currency           respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonetaryValueAmount) RawJSON() string { return r.JSON.raw }
func (r *MonetaryValueAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monetary or currency value
//
// The properties Amount, Type are required.
type MonetaryValueParam struct {
	Amount MonetaryValueAmountParam `json:"amount,omitzero,required"`
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

// The properties AmountInMinorUnits, Currency, Type are required.
type MonetaryValueAmountParam struct {
	AmountInMinorUnits int64  `json:"amount_in_minor_units,required"`
	Currency           string `json:"currency,required"`
	// This field can be elided, and will marshal its zero value as "currency".
	Type constant.Currency `json:"type,required"`
	paramObj
}

func (r MonetaryValueAmountParam) MarshalJSON() (data []byte, err error) {
	type shadow MonetaryValueAmountParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonetaryValueAmountParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Multiple lines of text
type MultiLineTextValue struct {
	Text string                      `json:"text,required"`
	Type constant.ValueTextMultiLine `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
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
// The properties Text, Type are required.
type MultiLineTextValueParam struct {
	Text string `json:"text,required"`
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

// Percentage numeric value
type PercentageValue struct {
	Percentage float64                        `json:"percentage,required"`
	Type       constant.ValueNumberPercentage `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
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
// The properties Percentage, Type are required.
type PercentageValueParam struct {
	Percentage float64 `json:"percentage,required"`
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

// Related item reference
type RelationValue struct {
	// An Item represents a single record or row within a Collection. It holds a set of
	// `values` corresponding to the Collection's `fields`.
	Item Item                   `json:"item,required"`
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

// ToParam converts this RelationValue to a RelationValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RelationValueParam.Overrides()
func (r RelationValue) ToParam() RelationValueParam {
	return param.Override[RelationValueParam](json.RawMessage(r.RawJSON()))
}

// Related item reference
//
// The properties Item, Type are required.
type RelationValueParam struct {
	// An Item represents a single record or row within a Collection. It holds a set of
	// `values` corresponding to the Collection's `fields`.
	Item ItemParam `json:"item,omitzero,required"`
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

// A single line of text
type SingleLineTextValue struct {
	Text string                       `json:"text,required"`
	Type constant.ValueTextSingleLine `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
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
// The properties Text, Type are required.
type SingleLineTextValueParam struct {
	Text string `json:"text,required"`
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

// LinkedIn profile link
type SocialLinkedInValue struct {
	Profile SocialLinkedInValueProfile      `json:"profile,required"`
	Type    constant.ValueUriSocialLinkedIn `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Profile     respjson.Field
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

// ToParam converts this SocialLinkedInValue to a SocialLinkedInValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SocialLinkedInValueParam.Overrides()
func (r SocialLinkedInValue) ToParam() SocialLinkedInValueParam {
	return param.Override[SocialLinkedInValueParam](json.RawMessage(r.RawJSON()))
}

type SocialLinkedInValueProfile struct {
	URL      string `json:"url" format:"uri"`
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialLinkedInValueProfile) RawJSON() string { return r.JSON.raw }
func (r *SocialLinkedInValueProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn profile link
//
// The properties Profile, Type are required.
type SocialLinkedInValueParam struct {
	Profile SocialLinkedInValueProfileParam `json:"profile,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/uri/social_linked_in".
	Type constant.ValueUriSocialLinkedIn `json:"type,required"`
	paramObj
}

func (r SocialLinkedInValueParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialLinkedInValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialLinkedInValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialLinkedInValueProfileParam struct {
	URL      param.Opt[string] `json:"url,omitzero" format:"uri"`
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r SocialLinkedInValueProfileParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialLinkedInValueProfileParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialLinkedInValueProfileParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X (formerly Twitter) username
type SocialXValue struct {
	Profile SocialXValueProfile      `json:"profile,required"`
	Type    constant.ValueUriSocialX `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Profile     respjson.Field
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

// ToParam converts this SocialXValue to a SocialXValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SocialXValueParam.Overrides()
func (r SocialXValue) ToParam() SocialXValueParam {
	return param.Override[SocialXValueParam](json.RawMessage(r.RawJSON()))
}

type SocialXValueProfile struct {
	URL      string `json:"url" format:"uri"`
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialXValueProfile) RawJSON() string { return r.JSON.raw }
func (r *SocialXValueProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X (formerly Twitter) username
//
// The properties Profile, Type are required.
type SocialXValueParam struct {
	Profile SocialXValueProfileParam `json:"profile,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "value/uri/social_x".
	Type constant.ValueUriSocialX `json:"type,required"`
	paramObj
}

func (r SocialXValueParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialXValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialXValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialXValueProfileParam struct {
	URL      param.Opt[string] `json:"url,omitzero" format:"uri"`
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r SocialXValueProfileParam) MarshalJSON() (data []byte, err error) {
	type shadow SocialXValueProfileParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialXValueProfileParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Telephone number value
type TelephoneNumber struct {
	TelephoneNumber string                        `json:"telephone_number,required"`
	Type            constant.ValueTelephoneNumber `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TelephoneNumber respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
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
// The properties TelephoneNumber, Type are required.
type TelephoneNumberParam struct {
	TelephoneNumber string `json:"telephone_number,required"`
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

// URL or web address
type URLValue struct {
	Type constant.ValueUriURL `json:"type,required"`
	URL  string               `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
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
// The properties Type, URL are required.
type URLValueParam struct {
	URL string `json:"url,required" format:"uri"`
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
// [GeoValue], [DateValue], [DatetimeValue], [Choice], [FunnelStep],
// [RelationValue].
//
// Use the [ValueUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ValueUnion struct {
	Text string `json:"text"`
	// Any of "value/text/single_line", "value/text/multi_line",
	// "value/number/unitless_integer", "value/number/unitless_float",
	// "value/number/monetary", "value/number/percentage", "value/boolean",
	// "value/email", "value/uri/url", "value/uri/domain", "value/uri/social_x",
	// "value/uri/social_linked_in", "value/telephone_number", "value/geo",
	// "value/date", "value/datetime", "value/choice", "value/funnel_step",
	// "value/relation".
	Type string `json:"type"`
	// This field is a union of [int64], [float64]
	Number ValueUnionNumber `json:"number"`
	// This field is from variant [MonetaryValue].
	Amount MonetaryValueAmount `json:"amount"`
	// This field is from variant [PercentageValue].
	Percentage float64 `json:"percentage"`
	// This field is from variant [BooleanValue].
	Boolean bool `json:"boolean"`
	// This field is from variant [EmailValue].
	Email string `json:"email"`
	// This field is from variant [URLValue].
	URL string `json:"url"`
	// This field is from variant [DomainValue].
	Domain string `json:"domain"`
	// This field is a union of [SocialXValueProfile], [SocialLinkedInValueProfile]
	Profile ValueUnionProfile `json:"profile"`
	// This field is from variant [TelephoneNumber].
	TelephoneNumber string `json:"telephone_number"`
	// This field is from variant [GeoValue].
	Geo string `json:"geo"`
	// This field is from variant [DateValue].
	Date time.Time `json:"date"`
	// This field is from variant [DatetimeValue].
	Datetime time.Time `json:"datetime"`
	// This field is from variant [Choice].
	Option ChoiceOption `json:"option"`
	// This field is from variant [FunnelStep].
	Step FunnelStepStep `json:"step"`
	// This field is from variant [RelationValue].
	Item Item `json:"item"`
	JSON struct {
		Text            respjson.Field
		Type            respjson.Field
		Number          respjson.Field
		Amount          respjson.Field
		Percentage      respjson.Field
		Boolean         respjson.Field
		Email           respjson.Field
		URL             respjson.Field
		Domain          respjson.Field
		Profile         respjson.Field
		TelephoneNumber respjson.Field
		Geo             respjson.Field
		Date            respjson.Field
		Datetime        respjson.Field
		Option          respjson.Field
		Step            respjson.Field
		Item            respjson.Field
		raw             string
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
func (Choice) implValueUnion()              {}
func (FunnelStep) implValueUnion()          {}
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
//	case moonbase.Choice:
//	case moonbase.FunnelStep:
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

func (u ValueUnion) AsValueChoice() (v Choice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValueUnion) AsValueFunnelStep() (v FunnelStep) {
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

// ValueUnionNumber is an implicit subunion of [ValueUnion]. ValueUnionNumber
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the [ValueUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfFloat]
type ValueUnionNumber struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfInt   respjson.Field
		OfFloat respjson.Field
		raw     string
	} `json:"-"`
}

func (r *ValueUnionNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ValueUnionProfile is an implicit subunion of [ValueUnion]. ValueUnionProfile
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the [ValueUnion].
type ValueUnionProfile struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	JSON     struct {
		URL      respjson.Field
		Username respjson.Field
		raw      string
	} `json:"-"`
}

func (r *ValueUnionProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ValueUnion to a ValueUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ValueUnionParam.Overrides()
func (r ValueUnion) ToParam() ValueUnionParam {
	return param.Override[ValueUnionParam](json.RawMessage(r.RawJSON()))
}

func ValueParamOfValueTextSingleLine(text string) ValueUnionParam {
	var valueTextSingleLine SingleLineTextValueParam
	valueTextSingleLine.Text = text
	return ValueUnionParam{OfValueTextSingleLine: &valueTextSingleLine}
}

func ValueParamOfValueTextMultiLine(text string) ValueUnionParam {
	var valueTextMultiLine MultiLineTextValueParam
	valueTextMultiLine.Text = text
	return ValueUnionParam{OfValueTextMultiLine: &valueTextMultiLine}
}

func ValueParamOfValueNumberUnitlessInteger(number int64) ValueUnionParam {
	var valueNumberUnitlessInteger IntegerValueParam
	valueNumberUnitlessInteger.Number = number
	return ValueUnionParam{OfValueNumberUnitlessInteger: &valueNumberUnitlessInteger}
}

func ValueParamOfValueNumberUnitlessFloat(number float64) ValueUnionParam {
	var valueNumberUnitlessFloat FloatValueParam
	valueNumberUnitlessFloat.Number = number
	return ValueUnionParam{OfValueNumberUnitlessFloat: &valueNumberUnitlessFloat}
}

func ValueParamOfValueNumberMonetary(amount MonetaryValueAmountParam) ValueUnionParam {
	var valueNumberMonetary MonetaryValueParam
	valueNumberMonetary.Amount = amount
	return ValueUnionParam{OfValueNumberMonetary: &valueNumberMonetary}
}

func ValueParamOfValueNumberPercentage(percentage float64) ValueUnionParam {
	var valueNumberPercentage PercentageValueParam
	valueNumberPercentage.Percentage = percentage
	return ValueUnionParam{OfValueNumberPercentage: &valueNumberPercentage}
}

func ValueParamOfValueBoolean(boolean bool) ValueUnionParam {
	var valueBoolean BooleanValueParam
	valueBoolean.Boolean = boolean
	return ValueUnionParam{OfValueBoolean: &valueBoolean}
}

func ValueParamOfValueEmail(email string) ValueUnionParam {
	var valueEmail EmailValueParam
	valueEmail.Email = email
	return ValueUnionParam{OfValueEmail: &valueEmail}
}

func ValueParamOfValueUriURL(url string) ValueUnionParam {
	var valueUriURL URLValueParam
	valueUriURL.URL = url
	return ValueUnionParam{OfValueUriURL: &valueUriURL}
}

func ValueParamOfValueUriDomain(domain string) ValueUnionParam {
	var valueUriDomain DomainValueParam
	valueUriDomain.Domain = domain
	return ValueUnionParam{OfValueUriDomain: &valueUriDomain}
}

func ValueParamOfValueUriSocialX(profile SocialXValueProfileParam) ValueUnionParam {
	var valueUriSocialX SocialXValueParam
	valueUriSocialX.Profile = profile
	return ValueUnionParam{OfValueUriSocialX: &valueUriSocialX}
}

func ValueParamOfValueUriSocialLinkedIn(profile SocialLinkedInValueProfileParam) ValueUnionParam {
	var valueUriSocialLinkedIn SocialLinkedInValueParam
	valueUriSocialLinkedIn.Profile = profile
	return ValueUnionParam{OfValueUriSocialLinkedIn: &valueUriSocialLinkedIn}
}

func ValueParamOfValueTelephoneNumber(telephoneNumber string) ValueUnionParam {
	var valueTelephoneNumber TelephoneNumberParam
	valueTelephoneNumber.TelephoneNumber = telephoneNumber
	return ValueUnionParam{OfValueTelephoneNumber: &valueTelephoneNumber}
}

func ValueParamOfValueGeo(geo string) ValueUnionParam {
	var valueGeo GeoValueParam
	valueGeo.Geo = geo
	return ValueUnionParam{OfValueGeo: &valueGeo}
}

func ValueParamOfValueDate(date time.Time) ValueUnionParam {
	var valueDate DateValueParam
	valueDate.Date = date
	return ValueUnionParam{OfValueDate: &valueDate}
}

func ValueParamOfValueDatetime(datetime time.Time) ValueUnionParam {
	var valueDatetime DatetimeValueParam
	valueDatetime.Datetime = datetime
	return ValueUnionParam{OfValueDatetime: &valueDatetime}
}

func ValueParamOfValueChoice(option ChoiceOptionParam) ValueUnionParam {
	var valueChoice ChoiceParam
	valueChoice.Option = option
	return ValueUnionParam{OfValueChoice: &valueChoice}
}

func ValueParamOfValueFunnelStep(step FunnelStepStepParam) ValueUnionParam {
	var valueFunnelStep FunnelStepParam
	valueFunnelStep.Step = step
	return ValueUnionParam{OfValueFunnelStep: &valueFunnelStep}
}

func ValueParamOfValueRelation(item ItemParam) ValueUnionParam {
	var valueRelation RelationValueParam
	valueRelation.Item = item
	return ValueUnionParam{OfValueRelation: &valueRelation}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ValueUnionParam struct {
	OfValueTextSingleLine        *SingleLineTextValueParam `json:",omitzero,inline"`
	OfValueTextMultiLine         *MultiLineTextValueParam  `json:",omitzero,inline"`
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
	OfValueChoice                *ChoiceParam              `json:",omitzero,inline"`
	OfValueFunnelStep            *FunnelStepParam          `json:",omitzero,inline"`
	OfValueRelation              *RelationValueParam       `json:",omitzero,inline"`
	paramUnion
}

func (u ValueUnionParam) MarshalJSON() ([]byte, error) {
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
func (u *ValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ValueUnionParam) asAny() any {
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
func (u ValueUnionParam) GetAmount() *MonetaryValueAmountParam {
	if vt := u.OfValueNumberMonetary; vt != nil {
		return &vt.Amount
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetPercentage() *float64 {
	if vt := u.OfValueNumberPercentage; vt != nil {
		return &vt.Percentage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetBoolean() *bool {
	if vt := u.OfValueBoolean; vt != nil {
		return &vt.Boolean
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetEmail() *string {
	if vt := u.OfValueEmail; vt != nil {
		return &vt.Email
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetURL() *string {
	if vt := u.OfValueUriURL; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetDomain() *string {
	if vt := u.OfValueUriDomain; vt != nil {
		return &vt.Domain
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetTelephoneNumber() *string {
	if vt := u.OfValueTelephoneNumber; vt != nil {
		return &vt.TelephoneNumber
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetGeo() *string {
	if vt := u.OfValueGeo; vt != nil {
		return &vt.Geo
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetDate() *time.Time {
	if vt := u.OfValueDate; vt != nil {
		return &vt.Date
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetDatetime() *time.Time {
	if vt := u.OfValueDatetime; vt != nil {
		return &vt.Datetime
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetOption() *ChoiceOptionParam {
	if vt := u.OfValueChoice; vt != nil {
		return &vt.Option
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetStep() *FunnelStepStepParam {
	if vt := u.OfValueFunnelStep; vt != nil {
		return &vt.Step
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetItem() *ItemParam {
	if vt := u.OfValueRelation; vt != nil {
		return &vt.Item
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetText() *string {
	if vt := u.OfValueTextSingleLine; vt != nil {
		return (*string)(&vt.Text)
	} else if vt := u.OfValueTextMultiLine; vt != nil {
		return (*string)(&vt.Text)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ValueUnionParam) GetType() *string {
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
func (u ValueUnionParam) GetNumber() (res valueUnionParamNumber) {
	if vt := u.OfValueNumberUnitlessInteger; vt != nil {
		res.any = &vt.Number
	} else if vt := u.OfValueNumberUnitlessFloat; vt != nil {
		res.any = &vt.Number
	}
	return
}

// Can have the runtime types [*int64], [*float64]
type valueUnionParamNumber struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *int64:
//	case *float64:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u valueUnionParamNumber) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ValueUnionParam) GetProfile() (res valueUnionParamProfile) {
	if vt := u.OfValueUriSocialX; vt != nil {
		res.any = &vt.Profile
	} else if vt := u.OfValueUriSocialLinkedIn; vt != nil {
		res.any = &vt.Profile
	}
	return
}

// Can have the runtime types [*SocialXValueProfileParam],
// [*SocialLinkedInValueProfileParam]
type valueUnionParamProfile struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *moonbase.SocialXValueProfileParam:
//	case *moonbase.SocialLinkedInValueProfileParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u valueUnionParamProfile) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u valueUnionParamProfile) GetURL() *string {
	switch vt := u.any.(type) {
	case *SocialXValueProfileParam:
		return paramutil.AddrIfPresent(vt.URL)
	case *SocialLinkedInValueProfileParam:
		return paramutil.AddrIfPresent(vt.URL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u valueUnionParamProfile) GetUsername() *string {
	switch vt := u.any.(type) {
	case *SocialXValueProfileParam:
		return paramutil.AddrIfPresent(vt.Username)
	case *SocialLinkedInValueProfileParam:
		return paramutil.AddrIfPresent(vt.Username)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ValueUnionParam](
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
		apijson.Discriminator[SocialXValueParam]("value/uri/social_x"),
		apijson.Discriminator[SocialLinkedInValueParam]("value/uri/social_linked_in"),
		apijson.Discriminator[TelephoneNumberParam]("value/telephone_number"),
		apijson.Discriminator[GeoValueParam]("value/geo"),
		apijson.Discriminator[DateValueParam]("value/date"),
		apijson.Discriminator[DatetimeValueParam]("value/datetime"),
		apijson.Discriminator[ChoiceParam]("value/choice"),
		apijson.Discriminator[FunnelStepParam]("value/funnel_step"),
		apijson.Discriminator[RelationValueParam]("value/relation"),
	)
}

type ItemNewParams struct {
	// The ID of the `Collection` to create the item in.
	CollectionID string `json:"collection_id,required"`
	// A hash where keys are the `ref` of a `Field` and values are the data to be set.
	Values map[string]FieldValueUnionParam `json:"values,omitzero,required"`
	paramObj
}

func (r ItemNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ItemNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ItemUpdateParams struct {
	// A hash where keys are the `ref` of a `Field` and values are the new data to be
	// set.
	Values map[string]FieldValueUnionParam `json:"values,omitzero,required"`
	// Any of "replace", "preserve", "merge".
	UpdateManyStrategy ItemUpdateParamsUpdateManyStrategy `header:"update-many-strategy,omitzero" json:"-"`
	// Any of "replace", "preserve".
	UpdateOneStrategy ItemUpdateParamsUpdateOneStrategy `header:"update-one-strategy,omitzero" json:"-"`
	paramObj
}

func (r ItemUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ItemUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ItemUpdateParamsUpdateManyStrategy string

const (
	ItemUpdateParamsUpdateManyStrategyReplace  ItemUpdateParamsUpdateManyStrategy = "replace"
	ItemUpdateParamsUpdateManyStrategyPreserve ItemUpdateParamsUpdateManyStrategy = "preserve"
	ItemUpdateParamsUpdateManyStrategyMerge    ItemUpdateParamsUpdateManyStrategy = "merge"
)

type ItemUpdateParamsUpdateOneStrategy string

const (
	ItemUpdateParamsUpdateOneStrategyReplace  ItemUpdateParamsUpdateOneStrategy = "replace"
	ItemUpdateParamsUpdateOneStrategyPreserve ItemUpdateParamsUpdateOneStrategy = "preserve"
)

type ItemUpsertParams struct {
	// The ID of the `Collection` to create the item in.
	CollectionID string `json:"collection_id,required"`
	// A hash where keys are the `ref` of a `Field` and values are used to identify the
	// item to update. When multiple identifiers are provided, the update will find
	// items that match any of the identifiers.
	Identifiers map[string]FieldValueUnionParam `json:"identifiers,omitzero,required"`
	// A hash where keys are the `ref` of a `Field` and values are the data to be set.
	Values map[string]FieldValueUnionParam `json:"values,omitzero,required"`
	// Any of "replace", "preserve", "merge".
	UpdateManyStrategy ItemUpsertParamsUpdateManyStrategy `header:"update-many-strategy,omitzero" json:"-"`
	// Any of "replace", "preserve".
	UpdateOneStrategy ItemUpsertParamsUpdateOneStrategy `header:"update-one-strategy,omitzero" json:"-"`
	paramObj
}

func (r ItemUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow ItemUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ItemUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ItemUpsertParamsUpdateManyStrategy string

const (
	ItemUpsertParamsUpdateManyStrategyReplace  ItemUpsertParamsUpdateManyStrategy = "replace"
	ItemUpsertParamsUpdateManyStrategyPreserve ItemUpsertParamsUpdateManyStrategy = "preserve"
	ItemUpsertParamsUpdateManyStrategyMerge    ItemUpsertParamsUpdateManyStrategy = "merge"
)

type ItemUpsertParamsUpdateOneStrategy string

const (
	ItemUpsertParamsUpdateOneStrategyReplace  ItemUpsertParamsUpdateOneStrategy = "replace"
	ItemUpsertParamsUpdateOneStrategyPreserve ItemUpsertParamsUpdateOneStrategy = "preserve"
)
