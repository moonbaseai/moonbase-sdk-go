// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/apiquery"
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// A list of search results.
type SearchResponse struct {
	Data []SearchResponseData `json:"data" api:"required"`
	Type constant.List        `json:"type" default:"list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A search result entry.
type SearchResponseData struct {
	// An Item represents a single record or row within a Collection. It holds a set of
	// `values` corresponding to the Collection's `fields`.
	Data SearchResponseDataDataUnion `json:"data" api:"required"`
	Type constant.SearchResult       `json:"type" default:"search_result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SearchResponseDataDataUnion contains all possible properties and values from
// [Item], [MoonbaseFile].
//
// Use the [SearchResponseDataDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SearchResponseDataDataUnion struct {
	ID string `json:"id"`
	// This field is from variant [Item].
	Collection CollectionPointer `json:"collection"`
	// Any of "item", "file".
	Type string `json:"type"`
	// This field is from variant [Item].
	Values map[string]FieldValueUnion `json:"values"`
	// This field is from variant [MoonbaseFile].
	Associations []ItemPointer `json:"associations"`
	// This field is from variant [MoonbaseFile].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [MoonbaseFile].
	DownloadURL string `json:"download_url"`
	// This field is from variant [MoonbaseFile].
	Filename string `json:"filename"`
	// This field is from variant [MoonbaseFile].
	Name string `json:"name"`
	// This field is from variant [MoonbaseFile].
	Size float64 `json:"size"`
	// This field is from variant [MoonbaseFile].
	UpdatedAt time.Time `json:"updated_at"`
	JSON      struct {
		ID           respjson.Field
		Collection   respjson.Field
		Type         respjson.Field
		Values       respjson.Field
		Associations respjson.Field
		CreatedAt    respjson.Field
		DownloadURL  respjson.Field
		Filename     respjson.Field
		Name         respjson.Field
		Size         respjson.Field
		UpdatedAt    respjson.Field
		raw          string
	} `json:"-"`
}

// anySearchResponseDataData is implemented by each variant of
// [SearchResponseDataDataUnion] to add type safety for the return type of
// [SearchResponseDataDataUnion.AsAny]
type anySearchResponseDataData interface {
	implSearchResponseDataDataUnion()
}

func (Item) implSearchResponseDataDataUnion()         {}
func (MoonbaseFile) implSearchResponseDataDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := SearchResponseDataDataUnion.AsAny().(type) {
//	case moonbase.Item:
//	case moonbase.MoonbaseFile:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u SearchResponseDataDataUnion) AsAny() anySearchResponseDataData {
	switch u.Type {
	case "item":
		return u.AsItem()
	case "file":
		return u.AsFile()
	}
	return nil
}

func (u SearchResponseDataDataUnion) AsItem() (v Item) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SearchResponseDataDataUnion) AsFile() (v MoonbaseFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SearchResponseDataDataUnion) RawJSON() string { return u.JSON.raw }

func (r *SearchResponseDataDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchParams struct {
	// The search text to match against items and files.
	Query string `query:"query" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [SearchParams]'s query parameters as `url.Values`.
func (r SearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
