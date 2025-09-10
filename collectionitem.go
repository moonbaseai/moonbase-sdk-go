// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/moonbaseai/moonbase-sdk-go/internal/apijson"
	"github.com/moonbaseai/moonbase-sdk-go/internal/apiquery"
	"github.com/moonbaseai/moonbase-sdk-go/internal/requestconfig"
	"github.com/moonbaseai/moonbase-sdk-go/option"
	"github.com/moonbaseai/moonbase-sdk-go/packages/pagination"
	"github.com/moonbaseai/moonbase-sdk-go/packages/param"
)

// CollectionItemService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectionItemService] method instead.
type CollectionItemService struct {
	Options []option.RequestOption
}

// NewCollectionItemService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCollectionItemService(opts ...option.RequestOption) (r CollectionItemService) {
	r = CollectionItemService{}
	r.Options = opts
	return
}

// Creates a new item in a collection.
func (r *CollectionItemService) New(ctx context.Context, collectionID string, body CollectionItemNewParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = append(r.Options[:], opts...)
	if collectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return
	}
	path := fmt.Sprintf("collections/%s/items", collectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing item.
func (r *CollectionItemService) Get(ctx context.Context, id string, query CollectionItemGetParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = append(r.Options[:], opts...)
	if query.CollectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("collections/%s/items/%s", query.CollectionID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an item.
func (r *CollectionItemService) Update(ctx context.Context, id string, params CollectionItemUpdateParams, opts ...option.RequestOption) (res *Item, err error) {
	if !param.IsOmitted(params.UpdateManyStrategy) {
		opts = append(opts, option.WithHeader("update-many-strategy", fmt.Sprintf("%s", params.UpdateManyStrategy)))
	}
	if !param.IsOmitted(params.UpdateOneStrategy) {
		opts = append(opts, option.WithHeader("update-one-strategy", fmt.Sprintf("%s", params.UpdateOneStrategy)))
	}
	opts = append(r.Options[:], opts...)
	if params.CollectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("collections/%s/items/%s", params.CollectionID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of items that are part of the collection.
func (r *CollectionItemService) List(ctx context.Context, collectionID string, query CollectionItemListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Item], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if collectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return
	}
	path := fmt.Sprintf("collections/%s/items", collectionID)
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

// Returns a list of items that are part of the collection.
func (r *CollectionItemService) ListAutoPaging(ctx context.Context, collectionID string, query CollectionItemListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Item] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, collectionID, query, opts...))
}

// Permanently deletes an item.
func (r *CollectionItemService) Delete(ctx context.Context, id string, body CollectionItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.CollectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("collections/%s/items/%s", body.CollectionID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Find and update an existing item, or create a new one.
func (r *CollectionItemService) Upsert(ctx context.Context, collectionID string, params CollectionItemUpsertParams, opts ...option.RequestOption) (res *Item, err error) {
	if !param.IsOmitted(params.UpdateManyStrategy) {
		opts = append(opts, option.WithHeader("update-many-strategy", fmt.Sprintf("%s", params.UpdateManyStrategy)))
	}
	if !param.IsOmitted(params.UpdateOneStrategy) {
		opts = append(opts, option.WithHeader("update-one-strategy", fmt.Sprintf("%s", params.UpdateOneStrategy)))
	}
	opts = append(r.Options[:], opts...)
	if collectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return
	}
	path := fmt.Sprintf("collections/%s/items/upsert", collectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type CollectionItemNewParams struct {
	// A hash where keys are the `ref` of a `Field` and values are the data to be set.
	Values map[string]FieldValueParamUnion `json:"values,omitzero,required"`
	paramObj
}

func (r CollectionItemNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectionItemNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionItemNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionItemGetParams struct {
	CollectionID string `path:"collection_id,required" json:"-"`
	paramObj
}

type CollectionItemUpdateParams struct {
	CollectionID string `path:"collection_id,required" json:"-"`
	// A hash where keys are the `ref` of a `Field` and values are the new data to be
	// set.
	Values map[string]FieldValueParamUnion `json:"values,omitzero,required"`
	// Any of "replace", "preserve", "merge".
	UpdateManyStrategy CollectionItemUpdateParamsUpdateManyStrategy `header:"update-many-strategy,omitzero" json:"-"`
	// Any of "replace", "preserve".
	UpdateOneStrategy CollectionItemUpdateParamsUpdateOneStrategy `header:"update-one-strategy,omitzero" json:"-"`
	paramObj
}

func (r CollectionItemUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectionItemUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionItemUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionItemUpdateParamsUpdateManyStrategy string

const (
	CollectionItemUpdateParamsUpdateManyStrategyReplace  CollectionItemUpdateParamsUpdateManyStrategy = "replace"
	CollectionItemUpdateParamsUpdateManyStrategyPreserve CollectionItemUpdateParamsUpdateManyStrategy = "preserve"
	CollectionItemUpdateParamsUpdateManyStrategyMerge    CollectionItemUpdateParamsUpdateManyStrategy = "merge"
)

type CollectionItemUpdateParamsUpdateOneStrategy string

const (
	CollectionItemUpdateParamsUpdateOneStrategyReplace  CollectionItemUpdateParamsUpdateOneStrategy = "replace"
	CollectionItemUpdateParamsUpdateOneStrategyPreserve CollectionItemUpdateParamsUpdateOneStrategy = "preserve"
)

type CollectionItemListParams struct {
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

// URLQuery serializes [CollectionItemListParams]'s query parameters as
// `url.Values`.
func (r CollectionItemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectionItemDeleteParams struct {
	CollectionID string `path:"collection_id,required" json:"-"`
	paramObj
}

type CollectionItemUpsertParams struct {
	// A hash where keys are the `ref` of a `Field` and values are used to identify the
	// item to update. When multiple identifiers are provided, the update will find
	// items that match any of the identifiers.
	Identifiers map[string]FieldValueParamUnion `json:"identifiers,omitzero,required"`
	// A hash where keys are the `ref` of a `Field` and values are the data to be set.
	Values map[string]FieldValueParamUnion `json:"values,omitzero,required"`
	// Any of "replace", "preserve", "merge".
	UpdateManyStrategy CollectionItemUpsertParamsUpdateManyStrategy `header:"update-many-strategy,omitzero" json:"-"`
	// Any of "replace", "preserve".
	UpdateOneStrategy CollectionItemUpsertParamsUpdateOneStrategy `header:"update-one-strategy,omitzero" json:"-"`
	paramObj
}

func (r CollectionItemUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectionItemUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionItemUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionItemUpsertParamsUpdateManyStrategy string

const (
	CollectionItemUpsertParamsUpdateManyStrategyReplace  CollectionItemUpsertParamsUpdateManyStrategy = "replace"
	CollectionItemUpsertParamsUpdateManyStrategyPreserve CollectionItemUpsertParamsUpdateManyStrategy = "preserve"
	CollectionItemUpsertParamsUpdateManyStrategyMerge    CollectionItemUpsertParamsUpdateManyStrategy = "merge"
)

type CollectionItemUpsertParamsUpdateOneStrategy string

const (
	CollectionItemUpsertParamsUpdateOneStrategyReplace  CollectionItemUpsertParamsUpdateOneStrategy = "replace"
	CollectionItemUpsertParamsUpdateOneStrategyPreserve CollectionItemUpsertParamsUpdateOneStrategy = "preserve"
)
