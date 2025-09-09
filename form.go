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

// FormService contains methods and other services that help with interacting with
// the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFormService] method instead.
type FormService struct {
	Options []option.RequestOption
}

// NewFormService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFormService(opts ...option.RequestOption) (r FormService) {
	r = FormService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing form.
func (r *FormService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Form, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("forms/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your forms.
func (r *FormService) List(ctx context.Context, query FormListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Form], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "forms"
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

// Returns a list of your forms.
func (r *FormService) ListAutoPaging(ctx context.Context, query FormListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Form] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// A Form provides a way to create `Items` in a `Collection`, often via a public
// URL for external users. Each form submission creates a new item.
type Form struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The `Collection` that submissions to this form are saved to.
	Collection Collection `json:"collection,required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the form, used as the title on its public page.
	Name string `json:"name,required"`
	// `true` if the form is available at a public URL.
	PagesEnabled bool `json:"pages_enabled,required"`
	// String representing the object’s type. Always `form` for this object.
	Type constant.Form `json:"type,required"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The public URL for the form, if `pages_enabled` is `true`.
	PagesURL string `json:"pages_url" format:"uri"`
	// An optional URL to redirect users to after a successful submission.
	RedirectURL string `json:"redirect_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Collection   respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		PagesEnabled respjson.Field
		Type         respjson.Field
		UpdatedAt    respjson.Field
		PagesURL     respjson.Field
		RedirectURL  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Form) RawJSON() string { return r.JSON.raw }
func (r *Form) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormListParams struct {
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

// URLQuery serializes [FormListParams]'s query parameters as `url.Values`.
func (r FormListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
