// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/moonbaseai/moonbase-sdk-go"
	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// A View represents a saved configuration for displaying items in a collection,
// including filters and sorting rules.
type View struct {
	// Unique identifier for the object.
	ID    string    `json:"id,required"`
	Links ViewLinks `json:"links,required"`
	// The name of the view.
	Name string `json:"name,required"`
	// String representing the object’s type. Always `view` for this object.
	Type constant.View `json:"type,required"`
	// The `Collection` this view belongs to.
	Collection moonbase.Collection `json:"collection"`
	// The type of view, such as `table` or `board`.
	//
	// Any of "table", "board".
	ViewType ViewViewType `json:"view_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Collection  respjson.Field
		ViewType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r View) RawJSON() string { return r.JSON.raw }
func (r *View) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewLinks struct {
	// A link to the `Collection` this view belongs to.
	Collection string `json:"collection,required" format:"uri"`
	// A link to the list of `Item` objects that are visible in this view.
	Items string `json:"items,required" format:"uri"`
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Collection  respjson.Field
		Items       respjson.Field
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewLinks) RawJSON() string { return r.JSON.raw }
func (r *ViewLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of view, such as `table` or `board`.
type ViewViewType string

const (
	ViewViewTypeTable ViewViewType = "table"
	ViewViewTypeBoard ViewViewType = "board"
)
