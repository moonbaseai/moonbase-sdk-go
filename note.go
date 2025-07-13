// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moonbasesdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/moonbase-sdk-go/internal/apijson"
	"github.com/stainless-sdks/moonbase-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/moonbase-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/moonbase-sdk-go/option"
	"github.com/stainless-sdks/moonbase-sdk-go/packages/pagination"
	"github.com/stainless-sdks/moonbase-sdk-go/packages/param"
	"github.com/stainless-sdks/moonbase-sdk-go/packages/respjson"
	"github.com/stainless-sdks/moonbase-sdk-go/shared/constant"
)

// NoteService contains methods and other services that help with interacting with
// the Moonbase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNoteService] method instead.
type NoteService struct {
	Options []option.RequestOption
}

// NewNoteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNoteService(opts ...option.RequestOption) (r NoteService) {
	r = NoteService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing note.
func (r *NoteService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Note, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your notes.
func (r *NoteService) List(ctx context.Context, query NoteListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Note], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "notes"
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

// Returns a list of your notes.
func (r *NoteService) ListAutoPaging(ctx context.Context, query NoteListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Note] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// The Note object represents a block of text content, often used for meeting notes
// or summaries.
type Note struct {
	// Unique identifier for the object.
	ID    string    `json:"id,required"`
	Links NoteLinks `json:"links,required"`
	// String representing the object’s type. Always `note` for this object.
	Type constant.Note `json:"type,required"`
	// The main content of the note.
	Body string `json:"body"`
	// Time at which the object was created, as an RFC 3339 timestamp.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A short, system-generated summary of the note's content.
	Summary string `json:"summary"`
	// An optional title for the note.
	Title string `json:"title"`
	// Time at which the object was last updated, as an RFC 3339 timestamp.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		Type        respjson.Field
		Body        respjson.Field
		CreatedAt   respjson.Field
		Summary     respjson.Field
		Title       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Note) RawJSON() string { return r.JSON.raw }
func (r *Note) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteLinks struct {
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
func (r NoteLinks) RawJSON() string { return r.JSON.raw }
func (r *NoteLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NoteListParams struct {
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

// URLQuery serializes [NoteListParams]'s query parameters as `url.Values`.
func (r NoteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
