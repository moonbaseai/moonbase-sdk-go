// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/apiquery"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
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

// Returns items that match the search query.
func (r *ItemService) Search(ctx context.Context, query ItemSearchParams, opts ...option.RequestOption) (res *ItemSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "items/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ItemSearchResponse struct {
	Data []Item        `json:"data,required"`
	Type constant.List `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *ItemSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ItemSearchParams struct {
	// The search text to match against items.
	Query string `query:"query,required" json:"-"`
	// Filter results by one or more collection IDs or `ref` values.
	Filter ItemSearchParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemSearchParams]'s query parameters as `url.Values`.
func (r ItemSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter results by one or more collection IDs or `ref` values.
type ItemSearchParamsFilter struct {
	CollectionID ItemSearchParamsFilterCollectionID `query:"collection_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemSearchParamsFilter]'s query parameters as `url.Values`.
func (r ItemSearchParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ItemSearchParamsFilterCollectionID struct {
	In []string `query:"in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ItemSearchParamsFilterCollectionID]'s query parameters as
// `url.Values`.
func (r ItemSearchParamsFilterCollectionID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
