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
	"github.com/moonbaseai/moonbase-sdk-go/packages/respjson"
	"github.com/moonbaseai/moonbase-sdk-go/shared/constant"
)

// ViewService contains methods and other services that help with interacting with
// the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewViewService] method instead.
type ViewService struct {
	Options []option.RequestOption
	Items   ViewItemService
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

// Retrieves the details of an existing view.
func (r *ViewService) Get(ctx context.Context, id string, query ViewGetParams, opts ...option.RequestOption) (res *ViewGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("views/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A View represents a saved configuration for displaying items in a collection,
// including filters and sorting rules.
type ViewGetResponse struct {
	// Unique identifier for the object.
	ID    string               `json:"id,required"`
	Links ViewGetResponseLinks `json:"links,required"`
	// The name of the view.
	Name string `json:"name,required"`
	// String representing the object’s type. Always `view` for this object.
	Type constant.View `json:"type,required"`
	// The `Collection` this view belongs to.
	Collection Collection `json:"collection"`
	// The type of view, such as `table` or `board`.
	//
	// Any of "table", "board".
	ViewType ViewGetResponseViewType `json:"view_type"`
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
func (r ViewGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ViewGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewGetResponseLinks struct {
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
func (r ViewGetResponseLinks) RawJSON() string { return r.JSON.raw }
func (r *ViewGetResponseLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of view, such as `table` or `board`.
type ViewGetResponseViewType string

const (
	ViewGetResponseViewTypeTable ViewGetResponseViewType = "table"
	ViewGetResponseViewTypeBoard ViewGetResponseViewType = "board"
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
