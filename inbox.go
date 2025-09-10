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

// InboxService contains methods and other services that help with interacting with
// the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxService] method instead.
type InboxService struct {
	Options []option.RequestOption
}

// NewInboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInboxService(opts ...option.RequestOption) (r InboxService) {
	r = InboxService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing inbox.
func (r *InboxService) Get(ctx context.Context, id string, query InboxGetParams, opts ...option.RequestOption) (res *Inbox, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("inboxes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of shared inboxes.
func (r *InboxService) List(ctx context.Context, query InboxListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Inbox], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inboxes"
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

// Returns a list of shared inboxes.
func (r *InboxService) ListAutoPaging(ctx context.Context, query InboxListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Inbox] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// The Inbox object represents a shared inbox for receiving and sending messages.
type Inbox struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The display name of the inbox.
	Name string `json:"name,required"`
	// String representing the object’s type. Always `inbox` for this object.
	Type constant.Inbox `json:"type,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The `Tagset` associated with this inbox, which defines the tags available for
	// its conversations.
	//
	// **Note:** Only present when requested using the `include` query parameter.
	Tagset Tagset `json:"tagset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Tagset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Inbox) RawJSON() string { return r.JSON.raw }
func (r *Inbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxGetParams struct {
	// Specifies which related objects to include in the response. Valid option is
	// `tagset`.
	//
	// Any of "tagset".
	Include InboxGetParamsInclude `query:"include[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxGetParams]'s query parameters as `url.Values`.
func (r InboxGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies which related objects to include in the response. Valid option is
// `tagset`.
type InboxGetParamsInclude string

const (
	InboxGetParamsIncludeTagset InboxGetParamsInclude = "tagset"
)

type InboxListParams struct {
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
	// Any of "tagset".
	Include InboxListParamsInclude `query:"include[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxListParams]'s query parameters as `url.Values`.
func (r InboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxListParamsInclude string

const (
	InboxListParamsIncludeTagset InboxListParamsInclude = "tagset"
)
