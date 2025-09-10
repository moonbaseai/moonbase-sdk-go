// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Structured content that can be rendered in multiple formats, currently
// supporting Markdown.
type FormattedText struct {
	// The content formatted as Markdown text.
	Markdown string `json:"markdown"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markdown    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormattedText) RawJSON() string { return r.JSON.raw }
func (r *FormattedText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A lightweight reference to another resource.
type Pointer struct {
	// Unique identifier for the referenced object.
	ID string `json:"id,required"`
	// String indicating the type of the referenced object.
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pointer) RawJSON() string { return r.JSON.raw }
func (r *Pointer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Pointer to a PointerParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PointerParam.Overrides()
func (r Pointer) ToParam() PointerParam {
	return param.Override[PointerParam](json.RawMessage(r.RawJSON()))
}

// A lightweight reference to another resource.
//
// The properties ID, Type are required.
type PointerParam struct {
	// Unique identifier for the referenced object.
	ID string `json:"id,required"`
	// String indicating the type of the referenced object.
	Type string `json:"type,required"`
	paramObj
}

func (r PointerParam) MarshalJSON() (data []byte, err error) {
	type shadow PointerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PointerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
