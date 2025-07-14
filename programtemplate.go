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

// ProgramTemplateService contains methods and other services that help with
// interacting with the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProgramTemplateService] method instead.
type ProgramTemplateService struct {
	Options []option.RequestOption
}

// NewProgramTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProgramTemplateService(opts ...option.RequestOption) (r ProgramTemplateService) {
	r = ProgramTemplateService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing program template.
func (r *ProgramTemplateService) Get(ctx context.Context, id string, query ProgramTemplateGetParams, opts ...option.RequestOption) (res *ProgramTemplate, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("program_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of your program templates.
func (r *ProgramTemplateService) List(ctx context.Context, query ProgramTemplateListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ProgramTemplate], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "program_templates"
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

// Returns a list of your program templates.
func (r *ProgramTemplateService) ListAutoPaging(ctx context.Context, query ProgramTemplateListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ProgramTemplate] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// The ProgramTemplate object defines the content of a message sent by a `Program`,
// including support for Liquid templating.
type ProgramTemplate struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The body content of the email, which can include Liquid variables.
	Body  string               `json:"body,required"`
	Links ProgramTemplateLinks `json:"links,required"`
	// The subject line of the email, which can include Liquid variables.
	Subject string `json:"subject,required"`
	// String representing the object’s type. Always `program_template` for this
	// object.
	Type constant.ProgramTemplate `json:"type,required"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The `Program` that uses this template.
	Program Program `json:"program"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Body        respjson.Field
		Links       respjson.Field
		Subject     respjson.Field
		Type        respjson.Field
		CreatedAt   respjson.Field
		Program     respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProgramTemplate) RawJSON() string { return r.JSON.raw }
func (r *ProgramTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProgramTemplateLinks struct {
	// The canonical URL for this object.
	Self string `json:"self,required" format:"uri"`
	// A link to the `Program` using this template.
	Program string `json:"program" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Self        respjson.Field
		Program     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProgramTemplateLinks) RawJSON() string { return r.JSON.raw }
func (r *ProgramTemplateLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProgramTemplateGetParams struct {
	// Specifies which related objects to include in the response. Valid option is
	// `program`.
	//
	// Any of "program".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProgramTemplateGetParams]'s query parameters as
// `url.Values`.
func (r ProgramTemplateGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProgramTemplateListParams struct {
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
	// Specifies which related objects to include in the response. Valid option is
	// `program`.
	//
	// Any of "program".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProgramTemplateListParams]'s query parameters as
// `url.Values`.
func (r ProgramTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
