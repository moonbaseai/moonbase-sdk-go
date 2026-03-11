// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
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

// Manage your marketing campaigns and forms
//
// UnsubscribeService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUnsubscribeService] method instead.
type UnsubscribeService struct {
	Options []option.RequestOption
}

// NewUnsubscribeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUnsubscribeService(opts ...option.RequestOption) (r UnsubscribeService) {
	r = UnsubscribeService{}
	r.Options = opts
	return
}

// Create a new unsubscribe.
func (r *UnsubscribeService) New(ctx context.Context, body UnsubscribeNewParams, opts ...option.RequestOption) (res *Unsubscribe, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "unsubscribes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a list of unsubscribes.
func (r *UnsubscribeService) List(ctx context.Context, query UnsubscribeListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Unsubscribe], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "unsubscribes"
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

// Returns a list of unsubscribes.
func (r *UnsubscribeService) ListAutoPaging(ctx context.Context, query UnsubscribeListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Unsubscribe] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes an unsubscribe by email address.
func (r *UnsubscribeService) Delete(ctx context.Context, email string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if email == "" {
		err = errors.New("missing required email parameter")
		return err
	}
	path := fmt.Sprintf("unsubscribes/%s", email)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type Unsubscribe struct {
	CreatedAt time.Time            `json:"created_at" api:"required" format:"date-time"`
	Email     string               `json:"email" api:"required"`
	Type      constant.Unsubscribe `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Email       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Unsubscribe) RawJSON() string { return r.JSON.raw }
func (r *Unsubscribe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnsubscribeNewParams struct {
	Email string `json:"email" api:"required"`
	paramObj
}

func (r UnsubscribeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UnsubscribeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnsubscribeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnsubscribeListParams struct {
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

// URLQuery serializes [UnsubscribeListParams]'s query parameters as `url.Values`.
func (r UnsubscribeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
