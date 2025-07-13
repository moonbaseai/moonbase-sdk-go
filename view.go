// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbasesdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/moonbase-sdk-go/internal/apijson"
	"github.com/stainless-sdks/moonbase-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/moonbase-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/moonbase-sdk-go/option"
	"github.com/stainless-sdks/moonbase-sdk-go/packages/pagination"
	"github.com/stainless-sdks/moonbase-sdk-go/packages/param"
	"github.com/stainless-sdks/moonbase-sdk-go/packages/respjson"
	"github.com/stainless-sdks/moonbase-sdk-go/shared/constant"
)

// ViewService contains methods and other services that help with interacting with
// the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewViewService] method instead.
type ViewService struct {
	Options []option.RequestOption
}

// NewViewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewViewService(opts ...option.RequestOption) (r ViewService) {
	r = ViewService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing view.
func (r *ViewService) Get(ctx context.Context, id string, query ViewGetParams, opts ...option.RequestOption) (res *View, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("views/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of items that are part of the specified view.
func (r *ViewService) ListItems(ctx context.Context, id string, query ViewListItemsParams, opts ...option.RequestOption) (res *pagination.CursorPage[Item], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("views/%s/items", id)
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

// Returns a list of items that are part of the specified view.
func (r *ViewService) ListItemsAutoPaging(ctx context.Context, id string, query ViewListItemsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Item] {
	return pagination.NewCursorPageAutoPager(r.ListItems(ctx, id, query, opts...))
}

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
	Collection Collection `json:"collection"`
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

type ViewGetParams struct {
	// Specifies which related objects to include in the response. Valid option is
	// `collection`.
	//
	// Any of "collection".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ViewGetParams]'s query parameters as `url.Values`.
func (r ViewGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ViewListItemsParams struct {
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

// URLQuery serializes [ViewListItemsParams]'s query parameters as `url.Values`.
func (r ViewListItemsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
