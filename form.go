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

// Creates a new form with an auto-generated collection and default fields.
func (r *FormService) New(ctx context.Context, body FormNewParams, opts ...option.RequestOption) (res *Form, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "forms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the details of an existing form.
func (r *FormService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Form, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("forms/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing form.
func (r *FormService) Update(ctx context.Context, id string, body FormUpdateParams, opts ...option.RequestOption) (res *Form, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("forms/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of your forms.
func (r *FormService) List(ctx context.Context, query FormListParams, opts ...option.RequestOption) (res *pagination.CursorPage[FormListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *FormService) ListAutoPaging(ctx context.Context, query FormListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[FormListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a form. The backing collection is preserved.
func (r *FormService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("forms/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A Form provides a way to create `Items` in a `Collection`, often via a public
// URL for external users. Each form submission creates a new item.
type Form struct {
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// `true` if submissions require a business email address, blocking free and
	// disposable providers.
	BusinessEmailRequired bool `json:"business_email_required" api:"required"`
	// The `Collection` that submissions to this form are saved to.
	Collection CollectionPointer `json:"collection" api:"required"`
	// Time at which the object was created, as an ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The HTML snippet for embedding the form on your website.
	HTMLEmbed string `json:"html_embed" api:"required"`
	// The name of the form, used as the title on its public page.
	Name string `json:"name" api:"required"`
	// If `true`, a Moonbase Pages hosted page is enabled for this form, providing a
	// standalone public URL for sharing.
	PagesEnabled bool `json:"pages_enabled" api:"required"`
	// String representing the object’s type. Always `form` for this object.
	Type constant.Form `json:"type" default:"form"`
	// Time at which the object was last updated, as an ISO 8601 timestamp in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The public URL for the form, if `pages_enabled` is `true`.
	PagesURL string `json:"pages_url" format:"uri"`
	// Optional URL the user is redirected to after a successful submission. When
	// unset, no redirect occurs. Stored as a Liquid template; rendered at submission
	// time with form field values under `submission.<key>` (keyed by the field's
	// `key`) plus UTM params (`utm_source`, `utm_medium`, `utm_campaign`, `utm_term`,
	// `utm_content`) automatically appended. Use the `uri_encode` filter for URL-safe
	// values, e.g.
	// `https://example.com/thanks?email={{ submission.email | uri_encode }}`. The
	// rendered URL must parse as a valid URL or the submission errors.
	RedirectURL string `json:"redirect_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		BusinessEmailRequired respjson.Field
		Collection            respjson.Field
		CreatedAt             respjson.Field
		HTMLEmbed             respjson.Field
		Name                  respjson.Field
		PagesEnabled          respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		PagesURL              respjson.Field
		RedirectURL           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Form) RawJSON() string { return r.JSON.raw }
func (r *Form) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the most essential attributes of a Form (does not include the
// embed HTML).
type FormListResponse struct {
	ID                    string `json:"id" api:"required"`
	BusinessEmailRequired bool   `json:"business_email_required" api:"required"`
	// A lightweight reference to a `Collection`, containing the minimal information
	// needed to identify it.
	Collection   CollectionPointer `json:"collection" api:"required"`
	CreatedAt    time.Time         `json:"created_at" api:"required" format:"date-time"`
	Name         string            `json:"name" api:"required"`
	PagesEnabled bool              `json:"pages_enabled" api:"required"`
	Type         constant.Form     `json:"type" default:"form"`
	UpdatedAt    time.Time         `json:"updated_at" api:"required" format:"date-time"`
	PagesURL     string            `json:"pages_url" format:"uri"`
	RedirectURL  string            `json:"redirect_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		BusinessEmailRequired respjson.Field
		Collection            respjson.Field
		CreatedAt             respjson.Field
		Name                  respjson.Field
		PagesEnabled          respjson.Field
		Type                  respjson.Field
		UpdatedAt             respjson.Field
		PagesURL              respjson.Field
		RedirectURL           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormListResponse) RawJSON() string { return r.JSON.raw }
func (r *FormListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormNewParams struct {
	// The name of the form, used as the title on its public page.
	Name string `json:"name" api:"required"`
	// If `true`, submissions require a business email address. Defaults to `false`.
	BusinessEmailRequired param.Opt[bool] `json:"business_email_required,omitzero"`
	// If `true`, enables a Moonbase Pages hosted page for this form, providing a
	// standalone public URL for sharing. Defaults to `false`.
	PagesEnabled param.Opt[bool] `json:"pages_enabled,omitzero"`
	// Optional URL the user is redirected to after a successful submission. Omit to
	// leave submissions without a redirect. Stored as a Liquid template; rendered at
	// submission time with form field values under `submission.<key>` (keyed by the
	// field's `key`) plus UTM params (`utm_source`, `utm_medium`, `utm_campaign`,
	// `utm_term`, `utm_content`) automatically appended. Use the `uri_encode` filter
	// for URL-safe values, e.g.
	// `https://example.com/thanks?email={{ submission.email | uri_encode }}`. The
	// rendered URL must parse as a valid URL or the submission errors.
	RedirectURL param.Opt[string] `json:"redirect_url,omitzero"`
	paramObj
}

func (r FormNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FormNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormUpdateParams struct {
	// Updated redirect URL, or `null` to clear. Omit to leave the existing value
	// unchanged. Liquid template rendered at submission time with form field values
	// under `submission.<key>` (keyed by the field's `key`) plus UTM params
	// (`utm_source`, `utm_medium`, `utm_campaign`, `utm_term`, `utm_content`)
	// automatically appended. Use the `uri_encode` filter for URL-safe values. The
	// rendered URL must parse as a valid URL or the submission errors.
	RedirectURL param.Opt[string] `json:"redirect_url,omitzero"`
	// If `true`, submissions require a business email address.
	BusinessEmailRequired param.Opt[bool] `json:"business_email_required,omitzero"`
	// The new name for the form.
	Name param.Opt[string] `json:"name,omitzero"`
	// If `true`, a Moonbase Pages hosted page is enabled for this form, providing a
	// standalone public URL for sharing.
	PagesEnabled param.Opt[bool] `json:"pages_enabled,omitzero"`
	paramObj
}

func (r FormUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FormUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormUpdateParams) UnmarshalJSON(data []byte) error {
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
