// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

// TagsetService contains methods and other services that help with interacting
// with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagsetService] method instead.
type TagsetService struct {
	Options []option.RequestOption
}

// NewTagsetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTagsetService(opts ...option.RequestOption) (r TagsetService) {
	r = TagsetService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing tagset.
func (r *TagsetService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Tagset, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("tagsets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your tagsets.
func (r *TagsetService) List(ctx context.Context, query TagsetListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Tagset], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "tagsets"
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

// Returns a list of your tagsets.
func (r *TagsetService) ListAutoPaging(ctx context.Context, query TagsetListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Tagset] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// A Tagset is a collection of `Tag` objects that can be applied within a specific
// `Inbox`.
type Tagset struct {
	// Unique identifier for the object.
	ID    string      `json:"id,required"`
	Links TagsetLinks `json:"links,required"`
	// The name of the tagset.
	Name string `json:"name,required"`
	// String representing the object’s type. Always `tagset` for this object.
	Type constant.Tagset `json:"type,required"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An optional description of the tagset's purpose.
	Description string `json:"description"`
	// A list of `Tag` objects belonging to this tagset.
	Tags []TagsetTag `json:"tags"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Tags        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tagset) RawJSON() string { return r.JSON.raw }
func (r *Tagset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagsetLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagsetLinks) RawJSON() string { return r.JSON.raw }
func (r *TagsetLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Tag is a label that can be applied to `Conversation` objects for organization
// and filtering.
type TagsetTag struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The name of the tag.
	Name string `json:"name,required"`
	// String representing the object’s type. Always `tag` for this object.
	Type constant.Tag `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagsetTag) RawJSON() string { return r.JSON.raw }
func (r *TagsetTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagsetListParams struct {
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

// URLQuery serializes [TagsetListParams]'s query parameters as `url.Values`.
func (r TagsetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
