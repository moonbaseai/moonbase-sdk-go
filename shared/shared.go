// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
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

// ToParam converts this FormattedText to a FormattedTextParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FormattedTextParam.Overrides()
func (r FormattedText) ToParam() FormattedTextParam {
	return param.Override[FormattedTextParam](json.RawMessage(r.RawJSON()))
}

// Structured content that can be rendered in multiple formats, currently
// supporting Markdown.
type FormattedTextParam struct {
	// The content formatted as Markdown text.
	Markdown param.Opt[string] `json:"markdown,omitzero"`
	paramObj
}

func (r FormattedTextParam) MarshalJSON() (data []byte, err error) {
	type shadow FormattedTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormattedTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Tag is a label that can be applied to supported resources (such as
// conversations, calls, and meetings) for organization and filtering.
type Tag struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// The color for the tag.
	//
	// Any of "amber", "blue", "cyan", "emerald", "fuchsia", "green", "indigo", "lime",
	// "lunar", "orange", "pink", "purple", "red", "rose", "sky", "teal", "violet",
	// "yellow".
	Color TagColor `json:"color" api:"required"`
	// The name of the tag.
	Name string `json:"name" api:"required"`
	// String representing the object’s type. Always `tag` for this object.
	Type constant.Tag `json:"type" default:"tag"`
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
func (r Tag) RawJSON() string { return r.JSON.raw }
func (r *Tag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The color for the tag.
type TagColor string

const (
	TagColorAmber   TagColor = "amber"
	TagColorBlue    TagColor = "blue"
	TagColorCyan    TagColor = "cyan"
	TagColorEmerald TagColor = "emerald"
	TagColorFuchsia TagColor = "fuchsia"
	TagColorGreen   TagColor = "green"
	TagColorIndigo  TagColor = "indigo"
	TagColorLime    TagColor = "lime"
	TagColorLunar   TagColor = "lunar"
	TagColorOrange  TagColor = "orange"
	TagColorPink    TagColor = "pink"
	TagColorPurple  TagColor = "purple"
	TagColorRed     TagColor = "red"
	TagColorRose    TagColor = "rose"
	TagColorSky     TagColor = "sky"
	TagColorTeal    TagColor = "teal"
	TagColorViolet  TagColor = "violet"
	TagColorYellow  TagColor = "yellow"
)

// A lightweight reference to a `Tag` used in request bodies.
//
// The properties ID, Type are required.
type TagPointerParam struct {
	// Unique identifier of the tag.
	ID string `json:"id" api:"required"`
	// String representing the object’s type. Always `tag` for this object.
	//
	// This field can be elided, and will marshal its zero value as "tag".
	Type constant.Tag `json:"type" default:"tag"`
	paramObj
}

func (r TagPointerParam) MarshalJSON() (data []byte, err error) {
	type shadow TagPointerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagPointerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
